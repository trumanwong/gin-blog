package main

import (
	"fmt"
	"gin-blog/pkg/settings"
	"gin-blog/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      settings.ReadTimeout,
		IdleTimeout:       settings.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	server.ListenAndServe()
}
