package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func TimingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ww := newWrappedResponseWriter(w)
		defer func() {
			duration := time.Since(start)
			fmt.Printf("Processed request [%s] %s in %v. Status: %d, Size: %d bytes\n", r.Method, r.RequestURI, duration, ww.statusCode, ww.size)		
		} ()
	
	next.ServeHTTP(ww, r)
})
}

type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
	size int
}

func newWrappedResponseWriter(w http.ResponseWriter) *wrappedResponseWriter {
	return &wrappedResponseWriter{w, http.StatusOK, 0}
}

func (w *wrappedResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *wrappedResponseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.size += size
	return size, err
}