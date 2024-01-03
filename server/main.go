package main

import (
	"Demo/initialize"
	"Demo/router"
	"fmt"
	"net/http"
	"time"
)

func main() {
	initialize.InitConfig()
	Router := router.InitRouter()

	server := &http.Server{
		Addr:           "0.0.0.0:9090",
		Handler:        Router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
