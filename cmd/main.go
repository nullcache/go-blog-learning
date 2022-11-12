package main

import (
	"github.com/nullcache/go-blog-learning/internal/router"
	"net/http"
	"time"
)

func main() {
	r := router.NewRouter()
	s := &http.Server{
		Addr:           ":8000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
