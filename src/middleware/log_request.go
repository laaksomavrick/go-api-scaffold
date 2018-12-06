package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/laaksomavrick/goals-api/src/config"
)

type logWriter struct {
	http.ResponseWriter
	status int
	length int
	body   []byte
}

func (w *logWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *logWriter) Write(body []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(body)
	w.length += n
	w.body = body
	return n, err

}

// LogRequest writes request and response metadata to std output
func LogRequest(next http.Handler, name string, config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lw := logWriter{ResponseWriter: w}
		next.ServeHTTP(&lw, r)
		duration := time.Since(start)
		userId := r.Context().Value("userId")

		if config.Env != "testing" {
			// todo log to /tmp/logs ?
			log.Printf("LOG\nHost: %s\nRemoteAddr: %s\nMethod: %s\nRequestURI: %s\nProto: %s\nStatus: %d\nContentLength: %d\nUserAgent: %s\nDuration: %s\nUserId: %d\nResBody: %s\n",
				r.Host,
				r.RemoteAddr,
				r.Method,
				r.RequestURI,
				r.Proto,
				lw.status,
				lw.length,
				r.Header.Get("User-Agent"),
				duration,
				userId,
				lw.body,
			)
		}
	}
}
