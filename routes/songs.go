package routes

import (
	"net/http"

	"github.com/campbell-frost/Shipwave/server/services"
)

type Song struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
}

func CreateSong(w http.ResponseWriter, r *http.Request) {
	err := services.CreateSong()
	if err != nil {
		http.Error(w, "could not create song", http.StatusInternalServerError)
	}
}

func GetSongFromDb(w http.ResponseWriter, r *http.Request) {
	err := services.GetSongs()
	if err != nil {
		http.Error(w, "Could not get song from db", http.StatusInternalServerError)
	}
}
