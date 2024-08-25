package routes

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
)

type Song struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	AudioBytes  string `json:"audio_bytes,omitempty"`
	CoverArtURL string `json:"cover_art_url,omitempty"`
}

func GetSong(w http.ResponseWriter, r *http.Request) {
	audioBytesPath := "./assets/rxk.mp3"
	audioBytes, err := os.ReadFile(audioBytesPath)
	audioBase64 := base64.StdEncoding.EncodeToString(audioBytes)
	if err != nil {
		http.Error(w, "Could not find song", http.StatusInternalServerError)
		return
	}
	response := Song{
		Id:         1,
		Title:      "Ginger Claps",
		Artist:     "Machine Girl",
		Album:      "Wlfgrl",
		AudioBytes: audioBase64,
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
