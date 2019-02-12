package main

import (
	"net/http"
	"os"

	"github.com/alamin-mahamud/authentication-microservice-go/config"
	"github.com/alamin-mahamud/authentication-microservice-go/pkg/router"
)

func main() {
	initAuthHandler()
	panic(initServer())

}

func initAuthHandler() {
	privateKey, err := config.InitPrivateKey()
	if err != nil {
		os.Exit(1)
	}
}

func initServer() {
	router := router.InitRouter()
	http.ListenAndServe(":80123", nil)
}
