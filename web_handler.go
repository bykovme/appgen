package main

import (
	"net/http"
)

// APIHandler - all api logic is here
type WebHandler struct {
}

func (web *WebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)
	switch head {
	case "features":
		web.handleFeature(w, r)
	case "finish":
		web.handleFinalStep(w, r)
	default:
		switch r.Method {
		case "GET":
			web.handleGet(w)
		default:
			http.Error(w, "Only GET requets are allowed", http.StatusMethodNotAllowed)
		}

	}

}

func (web *WebHandler) handleFeature(w http.ResponseWriter, r *http.Request) {
	//ShiningTreeJSON(w)
}

func (web *WebHandler) handleGet(w http.ResponseWriter) {
	//ShiningTreeJSON(w)
}

func (web *WebHandler) handleFinalStep(w http.ResponseWriter, r *http.Request) {

}
