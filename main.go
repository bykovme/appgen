package main

import (
	"log"
	"net/http"
)

// App - main application
type App struct {
	WebHandler  *WebHandler
	FileHandler *FileHandler
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)
	switch head {
	case "assets":
		app.FileHandler.ServeHTTP(w, r)
	default:
		app.WebHandler.ServeHTTP(w, r)
	}
}

func main() {
	app := &App{
		WebHandler:  new(WebHandler),
		FileHandler: new(FileHandler),
	}
	log.Println("App started on the port " + cPort)
	log.Fatal(http.ListenAndServe(":"+cPort, app))
}
