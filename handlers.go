package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	oauth1login "github.com/dghubble/gologin/v2/oauth1"
	twitterlogin "github.com/dghubble/gologin/v2/twitter"
	"github.com/gorilla/sessions"
	"github.com/joescharf/twitterprofile/v2/ent"
	"github.com/joescharf/twitterprofile/v2/ent/user"
	"github.com/joescharf/twitterprofile/v2/templates"

	"github.com/go-chi/chi/v5"
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

func GetFlashMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Get the session from the request
		session, err := app.Store.Get(r, "twitterprofile")
		if err == nil {
			// Get the previous flashes, if any.
			if flashes := session.Flashes(); len(flashes) > 0 {
				ctx = context.WithValue(r.Context(), "flashes", flashes)
				session.Save(r, w)
			}
		}
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

// GetCookieMiddleware - see https://go-chi.io/#/pages/middleware
func GetCookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the session from the request
		session, err := app.Store.Get(r, "twitterprofile")
		if err != nil {
			FlashError(w, r, session, "Error", "Error getting session", "/", http.StatusSeeOther)
			return
		}

		// Get the values from the session
		val := session.Values["twitterInfo"]
		var twitterInfo = &TwitterInfo{}
		twitterInfo, ok := val.(*TwitterInfo)
		if !ok {
			FlashError(w, r, session, "Error", "Error getting twitterInfo from session", "/", http.StatusSeeOther)
			return
		}

		// Get the user from the database
		user, err := app.DB.User.Query().
			Where(user.TwitterUserIDEQ(twitterInfo.UserID)).
			Only(r.Context())
		if err != nil {
			FlashError(w, r, session, "Error", "Error getting user from database", "/", http.StatusInternalServerError)
			return
		}
		spew.Dump("DB USER:", user)

		// Set the request contexts:
		// ctx := context.WithValue(r.Context(), "twitterInfo", twitterInfo)
		ctx := context.WithValue(r.Context(), "user", user)

		// call the next handler in the chain, passing the response writer and
		// the updated request object with the new context value.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
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

	// Get authentication information
	accessToken, accessSecret, _ := oauth1login.AccessTokenFromContext(ctx)
	twitterUser, err := twitterlogin.UserFromContext(ctx)
	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error logging in : %v", err)))
	}

	// Create twitterInfo struct:
	twitterInfo := &TwitterInfo{
		UserID: twitterUser.ID,
		// AccessToken:  accessToken,
		// AccessSecret: accessSecret,
		// ScreenName:   twitterUser.ScreenName,
		// Description:  twitterUser.Description,
	}

	// SAVE SESSION
	session, _ := app.Store.Get(r, "twitterprofile")
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
		SetUpdatedAt(time.Now()).
		OnConflictColumns("twitter_user_id").
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		FlashError(w, r, session, "Error", "Error saving user to database: "+err.Error(), "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/profile", http.StatusFound)
}

func getProfileHandler(w http.ResponseWriter, r *http.Request) {

	// Get user from context
	user := r.Context().Value("user").(*ent.User)

	p := templates.ProfileParams{
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
