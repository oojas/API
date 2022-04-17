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
	// for jobs
	router.HandleFunc("/jobs", Controllers.GetJobs).Methods(http.MethodGet)
	// for articles
	router.HandleFunc("/articles", Controllers.GetArticles).Methods(http.MethodGet)
	// For Campus Ambassador programs
	router.HandleFunc("/programs", Controllers.GetProgram).Methods(http.MethodGet)
	// For DSA Channels
	router.HandleFunc("/dsaChannels", Controllers.GetDSAChannels).Methods(http.MethodGet)
	// For RoadMaps
	router.HandleFunc("/roadMaps", Controllers.GetRoadMapChannel).Methods(http.MethodGet)
	// for free youtube channels offering courses
	router.HandleFunc("/freeCourses", Controllers.GetFreeCourseChannels).Methods(http.MethodGet)
	// for computer subjects
	router.HandleFunc("/subjects", Controllers.GetComputerSubjects).Methods(http.MethodGet)
	// for web app development
	router.HandleFunc("/webdev", Controllers.GetWebDev).Methods(http.MethodGet)
	// For cyber security
	router.HandleFunc("/cybersecurity", Controllers.GetCyber).Methods(http.MethodGet)
	// For machine learning
	router.HandleFunc("/machinelearning", Controllers.GetMachineLearning).Methods(http.MethodGet)
	// For cloud computing
	router.HandleFunc("/cloud", Controllers.GetCloudComputing).Methods(http.MethodGet) //
	// For Big Data
	router.HandleFunc("/bigdata", Controllers.GetBigData).Methods(http.MethodGet)
	// For Devops
	router.HandleFunc("/devops", Controllers.GetDevops).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(p, router))
}
