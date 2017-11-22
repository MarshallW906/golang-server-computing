package service

import (
	"net/http"
)

func inDevelopment(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Sorry, this is in development."))
}
