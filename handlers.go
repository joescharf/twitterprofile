package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/go-chi/chi/v5"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	err := app.AuthTwitter1()
	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error logging in: %v", err)))
	}
}
func getProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Get the profile:
	user, _, err := app.TwitterClient.Users.Show(&twitter.UserShowParams{
		ScreenName: "joescharf",
	})
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf("error getting profile: %v", err)))
	}
	app.UserDescription = user.Description
	w.Write([]byte(fmt.Sprintf("User Profile Description:\n%s", user.Description)))
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
