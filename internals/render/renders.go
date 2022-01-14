package render

import (
	"bytes"
	"fmt"
	"github.com/jerevick83/HOTEL-MGT/internals/config"
	"github.com/jerevick83/HOTEL-MGT/internals/models"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(data *models.TemplateData, req *http.Request) *models.TemplateData {
	data.CSRFToken = nosurf.Token(req)
	return data
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, req *http.Request, gohtml string, data *models.TemplateData) {
	var templateCache map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	templateC, ok := templateCache[gohtml]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	data = AddDefaultData(data, req)
	_ = templateC.Execute(buf, data)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error passing template", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
			myCache[name] = templateSet
		}
	}
	return myCache, nil
}
