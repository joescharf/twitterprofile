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
	ScreenName  string
	Description string
}

func Profile(w io.Writer, p ProfileParams) error {
	p.LayoutParams = layoutParams
	return profile.Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", "nav.html", file))
}
