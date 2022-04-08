package Controllers

import (
	"encoding/json"
	"net/http"
	"placementCracker_api/Database"
)

func GetCourses(w http.ResponseWriter, r *http.Request) {
	courses := Database.GetDataCourses()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(courses)

}
func GetJobs(w http.ResponseWriter, r *http.Request) {
	jobs := Database.GetDataJobs()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(jobs)
}
