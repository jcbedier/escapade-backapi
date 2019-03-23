package main

import "net/http"

type MPictures struct {
	StoryID             int      `json:"story_id"`
	Author              string   `json:"author"`
	File                string   `json:"file"`
	ZoomAvailable       bool     `json:"zoom_available"`
	IsUploadingDuration int      `json:"is_uploading_duration"`
	Filters             []string `json:"filters"`
	DelayBeforeMessages int      `json:"delay_before_messages"`
	Messages            []struct {
		Author           string `json:"author"`
		Content          string `json:"content"`
		IsTypingDuration int    `json:"is_typing_duration"`
		PauseAfter       int    `json:"pause_after"`
	} `json:"messages"`
}

func (a *App) AddPicture(w http.ResponseWriter, r *http.Request) {
	var item MPictures
	item.Author = "bob"
	item.File = "chose.png"
	item.ZoomAvailable = true

	err := a.Mongo.Insert(item)
	if err != nil {
		a.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	a.RespondWithJSON(w, http.StatusOK, item)
}
