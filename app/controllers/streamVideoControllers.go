package controllers

import (
	"net/http"
	"os"
	"log"
	"path/filepath"
	"github.com/gorilla/mux"
	env "tiktok-clone-api/config"
)

// Handler for streaming video
func GetStreamVideo(w http.ResponseWriter, r *http.Request) {
	// Print the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Specify the directory you want to list
	relativePath := env.VIDEO_PATH

	// Combine the current working directory with the relative path
	uploadDir := filepath.Join(cwd, relativePath)

    vars := mux.Vars(r)
    filename := vars["filename"]
    filePath := filepath.Join(uploadDir, filename)

    // Check if file exists
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }

    // Serve the video file
    w.Header().Set("Content-Type", "video/mp4")
    http.ServeFile(w, r, filePath)
}