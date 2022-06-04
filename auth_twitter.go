package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"

	"github.com/davecgh/go-spew/spew"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"

	pkce "github.com/grokify/go-pkce"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

func (a *App) AuthTwitter2(ctx context.Context) error {

	// **** OAUTH2 PKCE
	app.Oauth2Config = &oauth2.Config{
		ClientID:    app.Tp.ClientID,
		Scopes:      []string{"tweet.read", "users.read"},
		RedirectURL: "http://localhost:3000/auth/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://twitter.com/i/oauth2/authorize",
			TokenURL: "https://api.twitter.com/2/oauth2/token",
		},
	}
	app.CodeVerifier = pkce.NewCodeVerifier()
	codeChallenge := pkce.CodeChallengeS256(app.CodeVerifier)

	// Create authorization_code URL using `oauth2.Config`
	authURL := app.Oauth2Config.AuthCodeURL(
		"myState",
		oauth2.AccessTypeOffline,
		oauth2.SetAuthURLParam(pkce.ParamCodeChallenge, codeChallenge),
		oauth2.SetAuthURLParam(pkce.ParamCodeChallengeMethod, pkce.MethodS256))

	spew.Dump(authURL)

	// start a web server to listen on a callback URL
	app.Server = &http.Server{Addr: app.Oauth2Config.RedirectURL}
	http.HandleFunc("/", HandleAuthCallback)

	// parse the redirect URL for the port number
	u, err := url.Parse(app.Oauth2Config.RedirectURL)
	if err != nil {
		return err
	}

	// set up a listener on the redirect port
	port := fmt.Sprintf(":%s", u.Port())
	l, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	// open a browser window to the authorizationURL
	err = open.Start(authURL)
	if err != nil {
		return err
	}

	// start the blocking web server loop
	// this will exit when the handler gets fired and calls server.Close()
	app.Server.Serve(l)

	return nil
}

func HandleAuthCallback(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	// Get the authorization code from query parameters
	code := r.URL.Query().Get("code")
	fmt.Println("code:", code)
	if code == "" {
		io.WriteString(w, "Error: could not find 'code' URL parameter\n")

		// close the HTTP server and return
		cleanup(app.Server)
		return
	}

	// Exchange the authorization_code for a token with PKCE.
	token, _ := app.Oauth2Config.Exchange(
		ctx,
		code,
		oauth2.SetAuthURLParam(pkce.ParamCodeVerifier, app.CodeVerifier),
	)

	app.Token = token
	// close the HTTP server
	cleanup(app.Server)

}

func getTwitterOauth1Client(consumerKey string, consumerSecret string, accessToken string, accessSecret string) *http.Client {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	return config.Client(oauth1.NoContext, token)
}

func (a *App) AuthTwitter1() error {
	a.Oauth1Config = &oauth1.Config{
		ConsumerKey:    a.Tp.APIKey,
		ConsumerSecret: a.Tp.APISecret,
		CallbackURL:    "http://localhost:3000/auth/callback/1",
		Endpoint:       twitter.AuthorizeEndpoint,
	}
	requestToken, requestSecret, err := a.Oauth1Config.RequestToken()
	if err != nil {
		return err
	}
	a.OA1RequestToken = requestToken
	a.OA1RequestSecret = requestSecret
	authURL, err := a.Oauth1Config.AuthorizationURL(requestToken)

	// start a web server to listen on a callback URL
	app.Server = &http.Server{Addr: "http://localhost:3000/auth/callback"}
	http.HandleFunc("/1", HandleOAuth1Callback)

	// parse the redirect URL for the port number
	u, err := url.Parse(app.Oauth1Config.CallbackURL)
	if err != nil {
		return err
	}

	// set up a listener on the redirect port
	port := fmt.Sprintf(":%s", u.Port())
	l, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	// open a browser window to the authorizationURL
	err = open.Start(authURL.String())
	if err != nil {
		return err
	}

	// start the blocking web server loop
	// this will exit when the handler gets fired and calls server.Close()
	app.Server.Serve(l)

	return nil

}
func HandleOAuth1Callback(w http.ResponseWriter, r *http.Request) {
	// Get the authorization code from query parameters
	requestToken, verifier, err := oauth1.ParseAuthorizationCallback(r)
	if err != nil {
		return
	}
	accessToken, accessSecret, err := app.Oauth1Config.AccessToken(requestToken, app.OA1RequestSecret, verifier)
	// handle error
	token := oauth1.NewToken(accessToken, accessSecret)
	app.Token1 = token
	spew.Dump("OA1TOKEN:", token)

	// close the HTTP server
	cleanup(app.Server)

}
