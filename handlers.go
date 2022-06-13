package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
	oauth1login "github.com/dghubble/gologin/v2/oauth1"
	twitterlogin "github.com/dghubble/gologin/v2/twitter"
	"github.com/joescharf/twitterprofile/v2/templates"

	"github.com/go-chi/chi/v5"
)

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
			flash := &templates.Flash{
				Title:   "Error",
				Message: "Error getting session",
			}
			session.AddFlash(flash)
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Get the values from the session
		val := session.Values["twitterInfo"]
		var twitterInfo = &TwitterInfo{}
		twitterInfo, ok := val.(*TwitterInfo)
		if !ok {
			flash := &templates.Flash{
				Title:   "Error",
				Message: "Error retrieving twitterInfo",
			}
			session.AddFlash(flash)
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Set the request contexts:
		ctx := context.WithValue(r.Context(), "twitterInfo", twitterInfo)

		// call the next handler in the chain, passing the response writer and
		// the updated request object with the new context value.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

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
		AccessToken:  accessToken,
		AccessSecret: accessSecret,
		ScreenName:   twitterUser.ScreenName,
		UserID:       twitterUser.ID,
		Description:  twitterUser.Description,
	}

	// SAVE SESSION
	session, _ := app.Store.Get(r, "twitterprofile")
	session.Values["twitterInfo"] = twitterInfo
	err = session.Save(r, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error storing session : %v", err)))
	}

	spew.Dump(accessToken, accessSecret, twitterUser.Name, twitterUser.IDStr, twitterUser.ScreenName, err)
	http.Redirect(w, r, "/profile", http.StatusFound)
}

func getProfileHandler(w http.ResponseWriter, r *http.Request) {

	twitterInfo := r.Context().Value("twitterInfo").(*TwitterInfo)

	p := templates.ProfileParams{
		ScreenName:  twitterInfo.ScreenName,
		Description: twitterInfo.Description,
	}
	templates.Profile(w, p)
}
func updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Update the profile
	twitterInfo := r.Context().Value("twitterInfo").(*TwitterInfo)

	newDesc := twitterInfo.Description + "\nHello World."
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
