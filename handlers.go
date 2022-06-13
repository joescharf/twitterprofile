package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
	oauth1login "github.com/dghubble/gologin/v2/oauth1"
	twitterlogin "github.com/dghubble/gologin/v2/twitter"
	"github.com/joescharf/twitterprofile/v2/templates"

	"github.com/go-chi/chi/v5"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := templates.HomeParams{}
	templates.Home(w, p)
}

func loginSuccessHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	accessToken, accessSecret, _ := oauth1login.AccessTokenFromContext(ctx)
	twitterUser, err := twitterlogin.UserFromContext(ctx)
	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error logging in : %v", err)))
	}
	// SAVE SESSION
	session, _ := app.Store.Get(r, "twitterprofile")
	session.Values["accessToken"] = accessToken
	session.Values["accessSecret"] = accessSecret
	session.Values["twitterUsername"] = twitterUser.Name
	session.Values["twitterDescription"] = twitterUser.Description
	err = session.Save(r, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error storing session : %v", err)))
	}

	spew.Dump(accessToken, accessSecret, twitterUser, err)
	http.Redirect(w, r, "/profile", http.StatusFound)
}

func getProfileHandler(w http.ResponseWriter, r *http.Request) {
	// GET SESSION
	session, err := app.Store.Get(r, "twitterprofile")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error retrieving session : %v", err)))
	}
	twitterDesc := session.Values["twitterDescription"].(string)
	app.UserDescription = twitterDesc

	p := templates.ProfileParams{
		Description: twitterDesc,
	}
	templates.Profile(w, p)
}
func updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Update the profile
	newDesc := app.UserDescription + "\nHello World."
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
