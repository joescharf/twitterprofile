package templates

import (
	"embed"
	"html/template"
	"io"
)

//go:embed *
var files embed.FS

var (
	home    = parse("index.html")
	profile = parse("profile.html")
)

type HomeParams struct {
}

func Home(w io.Writer, p HomeParams) error {
	return home.Execute(w, p)
}

type ProfileParams struct {
	Description string
}

func Profile(w io.Writer, p ProfileParams) error {
	return profile.Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", "nav.html", file))
}
