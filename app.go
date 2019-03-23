package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/globalsign/mgo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type App struct {
	Router *mux.Router
	Mongo  *mgo.Collection
}

func (a *App) Initialize(c *MyConfig) {
	hosts := "127.0.0.1"
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: "escapade",
		Username: "root",
		Password: "example",
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	//defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Collection People
	a.Mongo = session.DB("escapade").C("actions")

	// Index
	// index := mgo.Index{
	// 	Key:        []string{"name", "phone"},
	// 	Unique:     false,
	// 	DropDups:   true,
	// 	Background: true,
	// 	Sparse:     true,
	// }

	// err = a.Mongo.EnsureIndex(index)
	// if err != nil {
	// 	panic(err)
	// }

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/v1/story/add", a.AddStory).Methods("POST")
	a.Router.HandleFunc("/v1/picture/add", a.AddPicture).Methods("POST")
	a.Router.HandleFunc("/v1/story/all", a.GetStories).Methods("GET")

	a.Router.Handle("/metrics", promhttp.Handler())
}

func (a *App) Run(addr string) {
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type", "authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(a.Router)))
}

func (a *App) RespondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	a.RespondWithJSON(w, code, map[string]string{"error": message})
}

func (a *App) RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(code)
	w.Write(response)
}
