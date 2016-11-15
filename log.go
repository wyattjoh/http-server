package main

import (
	"log"
	"net/http"
	"time"
)

type loggedResponse struct {
	code int

	http.ResponseWriter
}

func (lr *loggedResponse) WriteHeader(code int) {
	lr.code = code
	lr.ResponseWriter.WriteHeader(code)
}

func logWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		rw := loggedResponse{ResponseWriter: res}
		start := time.Now()
		defer func() {
			if rw.code == 0 {
				rw.code = http.StatusOK
			}

			log.Printf("%s %s %d - %s", req.Method, req.URL.Path, rw.code, time.Since(start).String())
		}()

		next.ServeHTTP(&rw, req)
	})
}
