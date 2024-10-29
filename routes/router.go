package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	controllers_ "tiktok-clone-api/app/controllers"
	library_ "tiktok-clone-api/library"
)

func Routergo() {
	router := mux.NewRouter().StrictSlash(false)
	routes := router.PathPrefix("/api/go/").Subrouter()

	// test connection
	routes.HandleFunc("/ping", library_.PingHandler).Methods("GET", "POST")
	// get preformance
	routes.HandleFunc("/get-video", controllers_.GetListVideo).Methods("GET")
	routes.HandleFunc("/search-video", controllers_.GetSearchVideo).Methods("GET")
	routes.HandleFunc("/stream-video/{filename}", controllers_.GetStreamVideo).Methods("GET")

	// Set up CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000","http://localhost:1234"}, // Change this to your allowed origins
		// AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	// Use the CORS handler
	handler := c.Handler(routes)

	http.Handle("/", handler)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", handler))
}
