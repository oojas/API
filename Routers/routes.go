package Routers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"placementCracker_api/Controllers"
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
	router.HandleFunc("/courses", Controllers.GetCourses).Methods(http.MethodGet)
	// for jobs
	router.HandleFunc("/jobs", Controllers.GetJobs).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(p, router))
}
