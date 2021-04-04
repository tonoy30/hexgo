package middlewares

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() { log.Printf("[%s]\t%s\t%s\n", r.Method, r.URL.Path, time.Since(start)) }()
		next.ServeHTTP(w, r)
	})
}
