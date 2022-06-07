package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/user/userlookup"
	"github.com/michimani/gotwi/user/userlookup/types"
)

func GetProfile(ctx context.Context, client *gotwi.Client, username string) (*types.GetByUsernameOutput, error) {
	// Get tweets:
	p := &types.GetByUsernameInput{
		Username: username,
		Expansions: fields.ExpansionList{
			fields.ExpansionPinnedTweetID,
		},
		UserFields: fields.UserFieldList{
			fields.UserFieldCreatedAt,
			fields.UserFieldDescription,
		},
		TweetFields: fields.TweetFieldList{
			fields.TweetFieldCreatedAt,
		},
	}

	u, err := userlookup.GetByUsername(ctx, client, p)
	return u, err
}

func (a *App) UpdateProfileDesc(httpClient *http.Client, newDesc string) {
	apiURL := "https://api.twitter.com/1.1/account/update_profile.json"

	// Encode query params and create url
	u, err := url.Parse(apiURL)
	v := url.Values{}
	v.Add("description", newDesc)
	u.RawQuery = v.Encode()

	// setup the POST request the post
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Response Status:", res.StatusCode)
}
