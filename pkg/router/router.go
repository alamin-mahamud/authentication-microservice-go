package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc(
		"/api/auth", 
		middleware.LogRequest(
			authHttpHandler.Auth()
		),
	)
	return r
}
