package main

import (
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/gorilla/sessions"
	"github.com/joescharf/twitterprofile/v2/ent"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type TP struct {
	APIKey       string
	APISecret    string
	AccessToken  string
	AccessSecret string
	ClientID     string
}
type TwitterInfo struct {
	UserID int64
}

type App struct {
	Tp            *TP
	Server        *http.Server
	Store         *sessions.CookieStore
	Oauth2Config  *oauth2.Config
	Oauth1Config  *oauth1.Config
	CodeVerifier  string
	Token         *oauth2.Token
	HttpClient1   *http.Client
	TwitterClient *twitter.Client
	SessionKey    string
	DB            *ent.Client
	Logger        *zap.SugaredLogger
	Env           string
}
