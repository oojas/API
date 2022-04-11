package Controllers

import (
	"encoding/json"
	"net/http"
	"placementCracker_api/Database"
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
