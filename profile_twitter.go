package main

import (
	"context"

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
