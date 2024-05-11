package middlewares

import (
	"io"
	"net/http"

	"github.com/andybalholm/brotli"
)

func CompressResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wc := brotli.HTTPCompressor(w, r)
		defer wc.Close()
		
		w = &compressResponseWriter{ResponseWriter: w, WriteCloser: wc}

		next.ServeHTTP(w, r)
	
	})
}

type compressResponseWriter struct {
	http.ResponseWriter
	io.WriteCloser
}

func (w *compressResponseWriter) Write(b []byte) (int, error) {
	return w.WriteCloser.Write(b)
}