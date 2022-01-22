package render

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Masterminds/sprig"
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

// AddDefaultData adds data for all templates
func AddDefaultData(data *models.TemplateData, req *http.Request) *models.TemplateData {
	data.FlashMessage = app.Session.PopString(req.Context(), "flash")
	data.ErrorMessage = app.Session.PopString(req.Context(), "error")
	data.WarningMessage = app.Session.PopString(req.Context(), "warning")
	data.CSRFToken = nosurf.Token(req)
	return data
}

// RenderingTemplate renders a template
func RenderingTemplate(w http.ResponseWriter, req *http.Request, gohtml string, data *models.TemplateData) error {
	var templateCache map[string]*template.Template
	if app.UseCache {
		tc, err := CreateTemplateCache()
		if err != nil {
			log.Println(err)
		}
		templateCache = tc
	} else {
		// get the template cache from the app config
		templateCache = app.TemplateCache
	}

	templateC, ok := templateCache[gohtml]

	if !ok {
		//log.Fatal("Could not get template from template cache")
		return errors.New("can't get template from cache")
	}

	buf := new(bytes.Buffer)
	data = AddDefaultData(data, req)
	_ = templateC.Execute(buf, data)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error passing template", err)
		return err
	}
	return nil
}

var pathToTemplate = "./templates"

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.gohtml", pathToTemplate))
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(sprig.FuncMap()).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.gohtml", pathToTemplate))
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob(fmt.Sprintf("%s/*.layout.gohtml", pathToTemplate))
			if err != nil {
				return myCache, err
			}
			myCache[name] = templateSet
		}
	}
	return myCache, nil
}
