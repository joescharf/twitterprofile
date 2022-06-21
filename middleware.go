package main

import (
	"context"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	gologinOauth1 "github.com/dghubble/gologin/v2/oauth1"
	"github.com/dghubble/oauth1"
	"github.com/joescharf/twitterprofile/v2/ent/user"
	"github.com/joescharf/twitterprofile/v2/templates"
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

// AuthMiddleware - see https://go-chi.io/#/pages/middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
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

		// Set the APP HttpClient and TwitterClient if not already set
		if app.HttpClient1 == nil {
			accessToken, accessSecret, _ := gologinOauth1.AccessTokenFromContext(ctx)
			token := oauth1.NewToken(accessToken, accessSecret)
			// Save the auth'd httpClient to app struct:
			app.HttpClient1 = app.Oauth1Config.Client(oauth1.NoContext, token)
			app.TwitterClient = twitter.NewClient(app.HttpClient1)
		}

		// Set the template layout parameters:
		templates.SetLayoutParams(templates.LayoutParams{ProfileImageURL: user.TwitterProfileImageURL})

		// Set the request contexts:
		// ctx := context.WithValue(r.Context(), "twitterInfo", twitterInfo)
		ctx = context.WithValue(ctx, "user", user)

		// call the next handler in the chain, passing the response writer and
		// the updated request object with the new context value.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
