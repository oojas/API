package Controllers

import (
	"encoding/json"
	"net/http"
	"placementCracker_api/Database"
	"placementCracker_api/Database/Courses"
	"placementCracker_api/Database/youtubeChannels"
)

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
func GetFreeCourseChannels(w http.ResponseWriter, r *http.Request) {
	courses := youtubeChannels.GetFreeCourse()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func GetComputerSubjects(w http.ResponseWriter, r *http.Request) {
	subjects := youtubeChannels.GetSubjectChannels()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subjects)
}
func GetWebDev(w http.ResponseWriter, r *http.Request) {
	web := Courses.GetWebAppData()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(web)
}
func GetCyber(w http.ResponseWriter, r *http.Request) {
	cyber := Courses.GetCyberSecurityData()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cyber)
}
func GetMachineLearning(w http.ResponseWriter, r *http.Request) {
	machine := Courses.GetMachineLearningData()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(machine)
}
func GetCloudComputing(w http.ResponseWriter, r *http.Request) {
	cloud := Courses.GetCloudComputingData()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cloud)
}
