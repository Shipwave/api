package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Route("/song/1", func(r chi.Router) {
		router.Get("/", getSong)
		router.Post("/", saveSong)
		router.Delete("/", deleteSong)
	})
	http.ListenAndServe(":3000", router)
}

type Song struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	AudioBytes  string `json:"audio_bytes,omitempty"`
	CoverArtURL string `json:"cover_art_url,omitempty"`
}

func getSong(w http.ResponseWriter, r *http.Request) {
	audioBytesPath := "../assets/rxk.mp3"
	audioBytes, err := os.ReadFile(audioBytesPath)
	audioBase64 := base64.StdEncoding.EncodeToString(audioBytes)
	if err != nil {
		fmt.Print("could not get song")
		http.Error(w, "Could not find song", http.StatusInternalServerError)
		return
	}
	fmt.Println("got song !")
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

	// Set the correct Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data
	w.Write(jsonData)
}

func saveSong(w http.ResponseWriter, r *http.Request) {
}

func deleteSong(w http.ResponseWriter, r *http.Request) {

}
