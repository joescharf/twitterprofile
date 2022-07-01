# TwitterProfile

- Also see (lol): [Add MRR Progress Bars to Twitter](https://chartmogul.com/blog/twitter-mrr-bars/)
- And [MRR Meter Generator](https://www.brandbird.app/tools/twitter-mrr-meter)
- [twitter-stripe-MRR Script](https://github.com/jsjoeio/twitter-stripe-mrr/blob/main/index.js) - really good script that does everything

## Text based mrr icons from twitter-stripe-MRR:
- 🟩🟩🟩🟩🟨⬜️⬜️⬜️⬜️⬜️ $1000/m
- 0 🟩🟩🟩🟩🟩🟨⬜⬜⬜⬜  5K
- 0 🟢🟢🟢🟢🟢🟡⚪⚪⚪⚪  5K

## Twitter Oauth

We have to use Twitter Oauth1 for now because it provides the read/write endpoints for the User Profile. Twitter hasn't ported the write capability to their APIv2 Yet.

I did go down the road of implementing Twitter Oauth2 in this library. Currently it's unused but the auth code is in auth_twitter.go. Requires TP_CLIENT_ID / TP_CLIENT_SECRET env vars in the .env file.

- [twitter-oauth1](https://github.com/dghubble/gologin#twitter-oauth1) Is a nice library that handles all the twitter Oauth1 including login and callback endpoints. 

## Setup for TailwindCSS
```shell
npm i
yarn dev
```

## ent
- `ent generate --feature sql/upsert ./ent/schema`