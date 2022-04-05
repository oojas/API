package Controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"placementCracker_api/Database"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":3000"
}

func Start() {
	p := getPort()
	router := mux.NewRouter()
	// for courses
	router.HandleFunc("/courses", getCourses).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(p, router))
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	courses := Database.GetDataCourses()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(courses)

}
