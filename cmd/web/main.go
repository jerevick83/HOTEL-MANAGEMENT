package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/jerevick83/HOTEL-MGT/internals/config"
	"github.com/jerevick83/HOTEL-MGT/internals/handlers"
	"github.com/jerevick83/HOTEL-MGT/internals/helpers"
	"github.com/jerevick83/HOTEL-MGT/internals/models"
	"github.com/jerevick83/HOTEL-MGT/internals/render"
	"github.com/jerevick83/HOTEL-MGT/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager

var L = log.New(os.Stdout, "HOTEL MANAGEMENT API", log.LstdFlags)
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main func that is run when the application starts
func main() {
	// storing values in a Session
	err := run()
	if err != nil {
		log.Fatal()
	}
	fmt.Println("Server running on Port", utils.PortName)

	server := &http.Server{
		Addr:         utils.PortName,
		Handler:      routes(),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err = server.ListenAndServe()
		log.Fatal(err)
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	L.Println("Received terminate, graceful shutdown", sig)
	shutdownCtx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	err = server.Shutdown(shutdownCtx)
	if err != nil {
		L.Fatal(err)
	}
}

func run() error {
	//what's to put in the session
	gob.Register(models.Reservation{})
	app.InProduction = false
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// Handling site session
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog
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
	helpers.NewHelpers(&app)
	return nil
}
