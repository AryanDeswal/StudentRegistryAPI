package controller

import "net/http"

// defaultHandler handles unmatched routes
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 not found", http.StatusNotFound)
}
