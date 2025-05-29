package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ArcticOctopus/bookings/pkg/handlers"
	"github.com/ArcticOctopus/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
	"github.com/ArcticOctopus/bookings/pkg/config"
)

var app config.AppConfig

const portNumber = ":8080"

var session *scs.SessionManager

func main() {

	// change to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = false
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	//Use False for Development mode, True for production
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s \n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
	// _ = http.ListenAndServe(portNumber, nil)
}
