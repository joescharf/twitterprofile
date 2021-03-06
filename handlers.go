package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/dghubble/go-twitter/twitter"
	gologinOauth1 "github.com/dghubble/gologin/v2/oauth1"
	gologinTwitter "github.com/dghubble/gologin/v2/twitter"
	"github.com/dghubble/oauth1"
	"github.com/gorilla/sessions"
	"github.com/joescharf/twitterprofile/v2/ent"
	"github.com/joescharf/twitterprofile/v2/templates"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func FlashError(w http.ResponseWriter, r *http.Request, session *sessions.Session, title, message, uri string, status int) {
	flash := &templates.Flash{
		Title:   title,
		Message: message,
	}
	session.AddFlash(flash)
	session.Save(r, w)
	http.Redirect(w, r, uri, status)
}

// indexHandler
// flash handling: https://github.com/gorilla/sessions/issues/57
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Get the flashes and check to see if there are any
	val := r.Context().Value("flashes")
	flashesInt, ok := val.([]interface{})
	flashes := make([]*templates.Flash, len(flashesInt))
	if ok {
		for i, v := range flashesInt {
			flashes[i] = v.(*templates.Flash)
		}
	}

	p := templates.HomeParams{
		Flashes: flashes,
	}
	templates.Home(w, p)
}

func loginSuccessHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session, _ := app.Store.Get(r, "twitterprofile")

	// Get authentication information and create
	// Twitter client from the oauth1.Config
	// https://github.com/dghubble/go-twitter/blob/master/examples/direct_messages.go
	accessToken, accessSecret, _ := gologinOauth1.AccessTokenFromContext(ctx)
	token := oauth1.NewToken(accessToken, accessSecret)
	// Save the auth'd httpClient to app struct:
	app.HttpClient1 = app.Oauth1Config.Client(oauth1.NoContext, token)
	app.TwitterClient = twitter.NewClient(app.HttpClient1)

	// Get User info from context
	twitterUser, err := gologinTwitter.UserFromContext(ctx)
	if err != nil {
		FlashError(w, r, session, "Error", "Error Logging In", "/", http.StatusSeeOther)
	}

	// Create twitterInfo struct:
	twitterInfo := &TwitterInfo{
		UserID: twitterUser.ID,
	}

	// SAVE SESSION
	session.Values["twitterInfo"] = twitterInfo
	err = session.Save(r, w)
	if err != nil {
		FlashError(w, r, session, "Error", "Error saving session", "/", http.StatusSeeOther)
		return
	}

	// Upsert User in DB
	err = app.DB.User.Create().
		SetTwitterUserID(twitterInfo.UserID).
		SetScreenName(twitterUser.ScreenName).
		SetDescription(twitterUser.Description).
		SetToken(accessToken).
		SetTokenSecret(accessSecret).
		SetTwitterProfileImageURL(twitterUser.ProfileImageURLHttps).
		SetUpdatedAt(time.Now()).
		OnConflictColumns("twitter_user_id").
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		FlashError(w, r, session, "Error", "Error saving user to database: "+err.Error(), "/", http.StatusSeeOther)
		return
	}

	// Set the template layout parameters:
	templates.SetLayoutParams(templates.LayoutParams{ProfileImageURL: twitterUser.ProfileImageURLHttps})

	http.Redirect(w, r, "/profile", http.StatusFound)
}

func getProfileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session, err := app.Store.Get(r, "twitterprofile")

	// Check we have a valid TwitterClient
	if app.TwitterClient == nil {
		FlashError(w, r, session, "Error", "Could not find your Twitter session, Please Login", "/", http.StatusSeeOther)
	}

	// Get user from context
	user := r.Context().Value("user").(*ent.User)

	// Get latest description from twitter
	twUser, _, err := app.TwitterClient.Users.Show(&twitter.UserShowParams{
		ScreenName: user.ScreenName,
	})
	if err != nil {
		FlashError(w, r, session, "Error", "Error getting user from twitter", "/", http.StatusSeeOther)
	}
	spew.Dump(twUser.ProfileImageURLHttps)

	// Update user description in DB
	user, err = user.Update().
		SetDescription(twUser.Description).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		FlashError(w, r, session, "Error", "Error updating user in DB", "/", http.StatusSeeOther)
	}

	templates.SetLayoutParams(templates.LayoutParams{ProfileImageURL: twUser.ProfileImageURLHttps})

	p := templates.ProfileParams{
		UserInfo: templates.UserInfo{
			ID:          user.ID,
			TokenSecret: user.TokenSecret,
		},
		Min:         user.Min,
		Max:         user.Max,
		ScreenName:  user.ScreenName,
		Description: user.Description,
	}
	templates.Profile(w, p)
}

// updateProfileHandler updates the twitter profile
func updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user := r.Context().Value("user").(*ent.User)

	newDesc := user.Description + "\nHello World."
	err := app.UpdateProfileDesc(app.HttpClient1, newDesc)
	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error updating profile: %v", err)))
	}
	w.Write([]byte(fmt.Sprintf("Profile Updated Successfully")))
}

type ProfileRequest struct {
	*ent.User
}

func (p *ProfileRequest) Bind(r *http.Request) error {
	return nil
}

// updateAPIProfileHandler updates the twitter profile
func updateAPIProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user := r.Context().Value("user").(*ent.User)

	// Get data from request:
	data := &ProfileRequest{User: user}
	if err := render.Bind(r, data); err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("API error updating profile: %v", err)))
		return
	}

	// Set min and max
	user.Min = data.Min
	user.Max = data.Max

	w.Write([]byte(fmt.Sprintf("User ID: %d, Min: %d, Max: %d", user.ID, user.Min, user.Max)))
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
