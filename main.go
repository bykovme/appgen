package main

import (
	"log"
	"net/http"
)

// App - main application
type App struct {
	WebHandler *WebHandler
}

var app *App

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.WebHandler.ServeHTTP(w, r)
}

func main() {

	app = &App{
		WebHandler: new(WebHandler),
	}

	log.Println("App started on the port " + cPort)

	err := http.ListenAndServe(":"+cPort, app)
	log.Fatal(err)
}
