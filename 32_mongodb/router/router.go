package router

import (
	"encoding/json"
	"mongodbinstallation/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api", healthCheck).Methods("GET")
	router.HandleFunc("/api/movies", controllers.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateMovies).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.WatchedMovies).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteOneMovies).Methods("DELETE")
	router.HandleFunc("/api/movies", controllers.DeleteAllMovies).Methods("DELETE")

	return router

}
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	json.NewEncoder(w).Encode(map[string]bool{"status":true})
	return
}
