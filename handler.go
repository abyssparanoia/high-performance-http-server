package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func routing(r chi.Router) {
	r.Get("/", handler)
}

func handler(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	time.Sleep(time.Millisecond * 500)
	io.WriteString(w, "Limited")
	log.Printf("request completed")
}
