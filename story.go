package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/globalsign/mgo/bson"
	uuid "github.com/satori/go.uuid"
)

type MStory struct {
	Name string
	Type string
	UUID string
}

type Story struct {
	Name string
	UUID string
}

func (a *App) AddStory(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		a.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var msg Story
	err = json.Unmarshal(b, &msg)
	if err != nil {
		a.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var item MStory
	uuid, err := uuid.NewV4()
	if err != nil {
		a.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	item.Name = msg.Name
	item.UUID = uuid.String()
	item.Type = "story"

	err = a.Mongo.Insert(item)
	if err != nil {
		a.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.RespondWithJSON(w, http.StatusOK, item)
}

func (a *App) GetStories(w http.ResponseWriter, r *http.Request) {
	var item []Story
	err := a.Mongo.Find(bson.M{"type": "story"}).All(&item)
	if err != nil {
		panic(err)
	}

	a.RespondWithJSON(w, http.StatusOK, item)
}

// 	// Query One
// 	result := Person{}
// 	err = c.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0}).One(&result)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Phone", result)
