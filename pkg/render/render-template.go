package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/rakibullahsazib/hotel-booking-api/pkg/config"
	"github.com/rakibullahsazib/hotel-booking-api/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	var err error

	// get template cache from app config
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal("Error Creating Template Cache")
		}
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Error Finding Template")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
		return
	}
}

var functions = template.FuncMap{}

// CreateTemplateCache creates a template map for all pages in templates folder which maps page name(string) to template
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/pages/*.html")
	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob("./templates/layouts/*.html")
		if err != nil {
			return templateCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/layouts/*.html")
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = ts
	}

	return templateCache, nil
}
