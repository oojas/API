package Controllers

import (
	"encoding/json"
	"net/http"
	"placementCracker_api/Database"
	"placementCracker_api/Database/youtubeChannels"
)

func GetCourses(w http.ResponseWriter, r *http.Request) {
	courses := Database.GetDataCourses()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}
func GetJobs(w http.ResponseWriter, r *http.Request) {
	jobs := Database.GetDataJobs()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}
func GetArticles(w http.ResponseWriter, r *http.Request) {
	articles := Database.GetDataArticles()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}
func GetProgram(w http.ResponseWriter, r *http.Request) {
	programs := Database.GetDataPrograms()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(programs)
}
func GetDSAChannels(w http.ResponseWriter, r *http.Request) {
	channels := youtubeChannels.GetDSAChannels()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(channels)
}
func GetRoadMapChannel(w http.ResponseWriter, r *http.Request) {
	maps := youtubeChannels.GetRoadMaps()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(maps)
}
