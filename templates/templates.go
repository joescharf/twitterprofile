package templates

import (
	"embed"
	"html/template"
	"io"
)

//go:embed *
var files embed.FS

type Flash struct {
	Title   string
	Message string
}
type StripeAccount struct {
	Name   string
	MRR    float32
	Status string
}
type UserInfo struct {
	ID          int
	TokenSecret string
}

var (
	layout       = parse("layout.html")
	home         = parse("index.html")
	profile      = parse("profile.html")
	layoutParams LayoutParams
)

type LayoutParams struct {
	ProfileImageURL string
}

func SetLayoutParams(p LayoutParams) {
	layoutParams = p
}

type HomeParams struct {
	LayoutParams
	Flashes []*Flash
}

func Home(w io.Writer, p HomeParams) error {
	p.LayoutParams = layoutParams
	return home.Execute(w, p)
}

type ProfileParams struct {
	LayoutParams
	UserInfo       UserInfo
	ScreenName     string
	Description    string
	Min            int32
	Max            int32
	StripeAccounts []*StripeAccount
}

func Profile(w io.Writer, p ProfileParams) error {
	p.LayoutParams = layoutParams

	// TODO: Add fabricated stripe accounts for now
	stripeAccounts := []*StripeAccount{}
	stripeAccount1 := &StripeAccount{
		Name:   "Widget Co",
		MRR:    100.0,
		Status: "Active",
	}
	stripeAccounts = append(stripeAccounts, stripeAccount1)
	stripeAccount2 := &StripeAccount{
		Name:   "NewCo",
		MRR:    10000.0,
		Status: "Active",
	}
	stripeAccounts = append(stripeAccounts, stripeAccount2)

	p.StripeAccounts = stripeAccounts

	return profile.Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", "nav.html", file))
}
