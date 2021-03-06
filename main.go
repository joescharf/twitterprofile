// https://bapp.Logger.radwell.codes/2022/01/go-program-for-a-unique-twitter-profile-banner/
// https://pkg.go.dev/github.com/dghubble/oauth1

package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"entgo.io/ent/dialect/sql/schema"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	gologinTwitter "github.com/dghubble/gologin/v2/twitter"
	"github.com/dghubble/oauth1"
	gologinOauth1 "github.com/dghubble/oauth1/twitter"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"

	"github.com/joescharf/twitterprofile/v2/ent"
	"github.com/joescharf/twitterprofile/v2/ent/migrate"
	"github.com/joescharf/twitterprofile/v2/templates"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var app *App

func init() {
	// https://pkg.go.dev/github.com/gorilla/sessions#pkg-overview
	gob.Register(&TwitterInfo{})
	gob.Register(&templates.Flash{})
}

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		app.Logger.Fatal("Error loading .env file")
	}

	// **** INITIALIZE
	app = &App{
		Env:        os.Getenv("TP_ENV"),
		SessionKey: os.Getenv("TP_SESSION_KEY"),
		DBURL:      os.Getenv("TP_DB_URL"),
	}
	app.TwitterAPIConfig = &TwitterAPIConfig{
		APIKey:       os.Getenv("TP_API_KEY"),
		APISecret:    os.Getenv("TP_API_KEY_SECRET"),
		AccessToken:  os.Getenv("TP_ACCESS_TOKEN"),
		AccessSecret: os.Getenv("TP_ACCESS_TOKEN_SECRET"),
	}
	app.Oauth1Config = &oauth1.Config{
		ConsumerKey:    app.TwitterAPIConfig.APIKey,
		ConsumerSecret: app.TwitterAPIConfig.APISecret,
		CallbackURL:    "http://localhost:3000/auth/callback",
		Endpoint:       gologinOauth1.AuthorizeEndpoint,
	}
	// Logger
	var logger *zap.Logger
	if app.Env == "production" {
		logger, _ = zap.NewProduction()
		app.Logger = logger.Sugar()
	} else {
		logger, _ = zap.NewDevelopment()
		app.Logger = logger.Sugar()
	}
	defer logger.Sync() // flushes buffer, if any

	// DB
	dbClient, err := ent.Open("postgres", app.DBURL)
	if err != nil {
		app.Logger.Fatalf("failed opening connection to postgres: %v", err)
	}
	app.DB = dbClient
	defer app.DB.Close()
	// Run the auto migration tool.
	if err := app.DB.Schema.Create(ctx, migrate.WithGlobalUniqueID(true), schema.WithAtlas(true)); err != nil {
		app.Logger.Fatalf("failed creating schema resources: %v", err)
	}

	// WEBSERVER and ROUTES
	// https://go-chi.io/#/pages/routing?id=routing-groups

	// Initialize Session Store
	app.Store = sessions.NewCookieStore([]byte(app.SessionKey))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(GetFlashMiddleware)

	// Static locations:
	workDir, _ := os.Getwd()
	cssDir := http.Dir(filepath.Join(workDir, "dist/css"))
	FileServer(r, "/css", cssDir)

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/", indexHandler)

		// https://github.com/dghubble/gologin#twitter-oauth1
		r.Get("/login", gologinTwitter.LoginHandler(app.Oauth1Config, nil).ServeHTTP)
		r.Get("/auth/callback", gologinTwitter.CallbackHandler(app.Oauth1Config, http.HandlerFunc(loginSuccessHandler), nil).ServeHTTP)

		r.Get("/hc", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
	})

	// Private Routes
	r.Group(func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Get("/profile", getProfileHandler)
		r.Post("/profile", updateProfileHandler)
	})

	// API Routes
	r.Group(func(r chi.Router) {
		r.Use(APIAuthMiddleware)
		r.Post("/api/v1/profile", updateAPIProfileHandler)
	})
	// START SERVERS and GOROUTINES
	g := errgroup.Group{}
	g.Go(func() error {
		return http.ListenAndServe(":3000", r)
	})

	app.Logger.Infow("Startup complete", "time", time.Now().String(), "APP", "twitterprofile")
	// Final wait group
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
