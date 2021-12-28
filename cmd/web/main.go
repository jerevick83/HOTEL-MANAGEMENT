package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/jerevick83/HOTEL-MGT/pkg/config"
	"github.com/jerevick83/HOTEL-MGT/pkg/handlers"
	"github.com/jerevick83/HOTEL-MGT/pkg/render"
	"github.com/jerevick83/HOTEL-MGT/utils"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager

// main is the main func that is run when the application starts
func main() {
	app.InProduction = false

	// Handling site session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = session
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = templateCache
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println("Server started on Port", utils.PortName)

	server := &http.Server{
		Addr:    utils.PortName,
		Handler: routes(&app),
	}
	err = server.ListenAndServe()
	log.Fatal(err)
}