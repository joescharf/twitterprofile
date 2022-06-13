// https://blog.radwell.codes/2022/01/go-program-for-a-unique-twitter-profile-banner/
// https://pkg.go.dev/github.com/dghubble/oauth1

package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
	"golang.org/x/sync/errgroup"

	"github.com/dghubble/go-twitter/twitter"
	twitterlogin "github.com/dghubble/gologin/v2/twitter"
	"github.com/dghubble/oauth1"
	twitteroa1 "github.com/dghubble/oauth1/twitter"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"

	"github.com/joescharf/twitterprofile/v2/ent"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type TP struct {
	APIKey       string
	APISecret    string
	AccessToken  string
	AccessSecret string
	ClientID     string
	SessionKey   string
}

type App struct {
	Tp              *TP
	Server          *http.Server
	Store           *sessions.CookieStore
	Oauth2Config    *oauth2.Config
	Oauth1Config    *oauth1.Config
	CodeVerifier    string
	Token           *oauth2.Token
	HttpClient1     *http.Client
	UserDescription string
}

var app *App

func init() {
	gob.Register(&twitter.User{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// **** INITIALIZE
	app = &App{}
	app.Tp = &TP{
		APIKey:       os.Getenv("TP_API_KEY"),
		APISecret:    os.Getenv("TP_API_KEY_SECRET"),
		AccessToken:  os.Getenv("TP_ACCESS_TOKEN"),
		AccessSecret: os.Getenv("TP_ACCESS_TOKEN_SECRET"),
		ClientID:     os.Getenv("TP_CLIENT_ID"),
		SessionKey:   os.Getenv("TP_SESSION_KEY"),
	}
	app.Oauth1Config = &oauth1.Config{
		ConsumerKey:    app.Tp.APIKey,
		ConsumerSecret: app.Tp.APISecret,
		CallbackURL:    "http://localhost:3000/auth/callback",
		Endpoint:       twitteroa1.AuthorizeEndpoint,
	}

	// Initialize Session Store
	// fmt.Println(hex.EncodeToString(securecookie.GenerateRandomKey(32)))
	app.Store = sessions.NewCookieStore([]byte(app.Tp.SessionKey))

	// DB
	dbConnStr := "postgres://postgres:postgres@localhost:15432/twitterprofile?sslmode=disable"
	dbClient, err := ent.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer dbClient.Close()
	// Run the auto migration tool.
	if err := dbClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// WEBSERVER and ROUTES

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Static locations:
	workDir, _ := os.Getwd()
	cssDir := http.Dir(filepath.Join(workDir, "dist/css"))
	FileServer(r, "/css", cssDir)

	// Routes
	r.Get("/", indexHandler)
	// r.Get("/login", loginHandler)
	// r.Get("/auth/callback", HandleOAuth1Callback)
	r.Get("/login", twitterlogin.LoginHandler(app.Oauth1Config, nil).ServeHTTP)
	r.Get("/auth/callback", twitterlogin.CallbackHandler(app.Oauth1Config, http.HandlerFunc(loginSuccessHandler), nil).ServeHTTP)

	r.Get("/profile", getProfileHandler)
	r.Post("/profile", updateProfileHandler)

	r.Get("/hc", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// START SERVERS and GOROUTINES
	g := errgroup.Group{}
	g.Go(func() error {
		return http.ListenAndServe(":3000", r)
	})

	// Final wait group
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
