package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"snippetbox/internal/models"
	"snippetbox/ui"
	"time"
)

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {

		name := filepath.Base(page)

		patters := []string{
			"html/base.tmpl.html", "html/partials/*.tmpl.html", page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patters...)
		if err != nil {
			return nil, err
		}
		cache[name] = ts

	}
	return cache, nil
}
