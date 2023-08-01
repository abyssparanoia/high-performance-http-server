package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func routing(r chi.Router) {
	r.Use(ContextCancellationMiddleware)
	r.Get("/", handler)
}

func handler(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	time.Sleep(time.Millisecond * 5000)
	io.WriteString(w, "Limited")
	log.Printf("request completed")

	// ctx := r.Context()

	// // 3秒かかる処理をシミュレート
	// select {
	// case <-time.After(3 * time.Second):
	// 	fmt.Fprintf(w, "Request completed successfully.")
	// case <-ctx.Done():
	// 	// context canceledエラーをクライアントに返さないようにする
	// 	return
	// }
}
