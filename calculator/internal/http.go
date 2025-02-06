package internal

import (
	"log/slog"
	"net/http"
	"os"
)

func RequestLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		slog.Info("request", "path", r.URL.Path, "hostname", hostname)
		h.ServeHTTP(w, r)
	})
}
