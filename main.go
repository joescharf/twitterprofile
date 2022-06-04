// https://blog.radwell.codes/2022/01/go-program-for-a-unique-twitter-profile-banner/

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"

	"github.com/davecgh/go-spew/spew"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"

	"github.com/michimani/gotwi"
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
	OA1RequestToken  string
	OA1RequestSecret string
	CodeVerifier     string
	Token            *oauth2.Token
	Token1           *oauth1.Token
}

var app *App

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// **** INITIALIZE
	ctx := context.Background()
	app = &App{
		Tp:           &TP{},
		Server:       &http.Server{},
		Oauth2Config: &oauth2.Config{},
		Oauth1Config: &oauth1.Config{},
		Token:        &oauth2.Token{},
	}

	app.Tp = &TP{
		APIKey:       os.Getenv("TP_API_KEY"),
		APISecret:    os.Getenv("TP_API_KEY_SECRET"),
		AccessToken:  os.Getenv("TP_ACCESS_TOKEN"),
		AccessSecret: os.Getenv("TP_ACCESS_TOKEN_SECRET"),
		ClientID:     os.Getenv("TP_CLIENT_ID"),
	}

	// AUTHENTICATE TWITTER
	app.AuthTwitter2(ctx)

	spew.Dump("TOKEN:", app.Token)

	// DO SHIT

	// Init the gotwi client (APIv2)
	in := &gotwi.NewClientWithAccessTokenInput{
		AccessToken: app.Token.AccessToken,
	}

	c, err := gotwi.NewClientWithAccessToken(in)
	if err != nil {
		os.Exit(1)
	}

	// Get profile with goTWI:

	u, err := GetProfile(ctx, c, "joescharf")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("ID:          ", gotwi.StringValue(u.Data.ID))
	fmt.Println("Name:        ", gotwi.StringValue(u.Data.Name))
	fmt.Println("Username:    ", gotwi.StringValue(u.Data.Username))
	fmt.Println("CreatedAt:   ", u.Data.CreatedAt)
	fmt.Println("Description: ", gotwi.StringValue(u.Data.Description))
	if u.Includes.Tweets != nil {
		for _, t := range u.Includes.Tweets {
			fmt.Println("PinnedTweet: ", gotwi.StringValue(t.Text))
		}
	}

	// Init go-twitter (APIv1.1)
	httpClient := getTwitterOauth1Client(app.Tp.APIKey, app.Tp.APISecret, app.Tp.AccessToken, app.Tp.AccessSecret)
	// Twitter client
	goTwitterClient := twitter.NewClient(httpClient)

	// User Show
	user, _, err := goTwitterClient.Users.Show(&twitter.UserShowParams{
		ScreenName: "joescharf",
	})

	fmt.Println("User Oauth 1.1 Desc:", user.Description)

	// **** 3 legged oauth1:
	err = app.AuthTwitter1()
	if err != nil {
		fmt.Println("AuthTwitter1 Error: ", err)
	}

}

// cleanup closes the HTTP server
func cleanup(server *http.Server) {
	// we run this as a goroutine so that this function falls through and
	// the socket to the browser gets flushed/closed before the server goes away
	go server.Close()
}
