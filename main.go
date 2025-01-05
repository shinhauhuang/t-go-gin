package main

import (
	"fmt"
	"log"
	"net/http"
	"t-go-gin/internal/settings"
	"t-go-gin/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	settings.Setup()
}

func main() {
	gin.SetMode(settings.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := settings.ServerSetting.ReadTimeout
	writeTimeout := settings.ServerSetting.WriteTimeout
	httpPort := fmt.Sprintf(":%d", settings.ServerSetting.HTTPPort)

	server := &http.Server{
		Addr:         httpPort,
		Handler:      routersInit,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Printf("[info] start http server listening %s", httpPort)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
