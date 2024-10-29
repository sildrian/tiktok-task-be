package controllers

import (
	"net/http"
	// "fmt"
	"os"
	"io/ioutil"
	"log"
	"path/filepath"
	library_ "tiktok-clone-api/library"
	env "tiktok-clone-api/config"
)

type Video struct {
    ID     int `json:"id"`
    Title  string `json:"title"`
    URL    string `json:"video"`
}

func GetListVideo(w http.ResponseWriter, r *http.Request){
	 // Example data
    titles := []string{"Pokemon", "Chocolate", "Ice Cream"}

	// Print the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Specify the directory you want to list
	relativePath := env.VIDEO_PATH

	// Combine the current working directory with the relative path
	dir := filepath.Join(cwd, relativePath)

	// Open the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the files and print their names
	// for _, file := range files {
	// 	fmt.Println(env.LOCAL_URL+env.STREAM_URL+file.Name())
	// }
	// Slice to hold Video structs
    var videos []Video

    // Loop over the data and populate the Video structs
    for i := 0; i < len(files); i++ {
        video := Video{
            ID:    i + 1,
            Title: titles[i],
            URL:   env.LOCAL_URL+env.STREAM_URL+files[i].Name(),
        }
        videos = append(videos, video)
    }
	
	library_.Res_200(w, "Success", videos)
}
