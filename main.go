// https://blog.radwell.codes/2022/01/go-program-for-a-unique-twitter-profile-banner/
// https://pkg.go.dev/github.com/dghubble/oauth1

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
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
}

type App struct {
	Tp               *TP
	Server           *http.Server
	Oauth2Config     *oauth2.Config
	Oauth1Config     *oauth1.Config
	OA1RequestSecret string
	CodeVerifier     string
	Token            *oauth2.Token
	Token1           *oauth1.Token
}

var app *App

func main() {
	ctx := context.Background()
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
	}

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

	// AUTHENTICATE TWITTER
	err = app.AuthTwitter1()
	if err != nil {
		fmt.Println("AuthTwitter1 Error: ", err)
		os.Exit(1)
	}
	// Setup twitter client
	httpClient := app.Oauth1Config.Client(ctx, app.Token1)
	client := twitter.NewClient(httpClient)

	// Get the profile:
	user, _, err := client.Users.Show(&twitter.UserShowParams{
		ScreenName: "joescharf",
	})
	fmt.Println("User Profile Description:\n", user.Description)

	// Update the profile
	newDesc := user.Description + "\nHello World."
	app.UpdateProfileDesc(httpClient, newDesc)
}
