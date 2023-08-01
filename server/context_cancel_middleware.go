package main

import (
	"context"
	"log"
	"net/http"
)

func ContextCancellationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// ハンドラー実行前にキャンセルが検知された場合の処理
		defer func() {
			if err := ctx.Err(); err == context.Canceled {
				log.Println("Request canceled in middleware:", err)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
