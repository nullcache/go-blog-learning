package main

import (
	"github.com/nullcache/go-blog-learning/global"
	"github.com/nullcache/go-blog-learning/internal/initDB"
	"github.com/nullcache/go-blog-learning/internal/router"
	"log"
	"net/http"
	"time"
)

func init() {

	err := global.SetupSetting()
	if err != nil {
		log.Panicf("init.setupSetting err: %v", err)
	}
	err = global.SetupLogger()
	if err != nil {
		log.Panicf("init.setupSetting err: %v", err)
	}
	err = initDB.SetupDBEngine()
	if err != nil {
		log.Panicf("init.setupDB err: %v", err)
	}
}

func main() {
	r := router.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        r,
		ReadTimeout:    global.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: global.ServerSetting.MaxHeaderBytes,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
