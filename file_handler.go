package main

import (
	"net/http"
)

// APIHandler - all api logic is here
type FileHandler struct {
}

func (file *FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/"+r.URL.Path[1:])
}
