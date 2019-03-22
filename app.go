package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(c *Config) {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	//var pattern string
	// INIT
	//a.Router.HandleFunc("/v1/init", a.Director.RootInit).Methods("GET")

	a.Router.Handle("/metrics", promhttp.Handler())
}

func (a *App) Run(addr string) {
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type", "authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(a.Router)))
}
