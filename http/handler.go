package http

import (
	"net/http"
	"strings"
)

type Handler struct {
	DialHandler *DialHandler
}

// ServeHTTP delegates request to appropriate subhandler
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api/dials") {
		h.DialHandler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}
