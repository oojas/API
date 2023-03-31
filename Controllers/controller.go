package Controllers

import (
	"encoding/json"
	"net/http"
	"placementCracker_api/Database"
	"placementCracker_api/Database/Courses"
	"placementCracker_api/Database/youtubeChannels"
)

func GetJobs(w http.ResponseWriter, _ *http.Request) {
	jobs := Database.GetDataJobs()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}
func GetArticles(w http.ResponseWriter, _ *http.Request) {
	articles := Database.GetDataArticles()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}
func GetProgram(w http.ResponseWriter, _ *http.Request) {
	programs := Database.GetDataPrograms()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(programs)
}
func GetDSAChannels(w http.ResponseWriter, _ *http.Request) {
	channels := youtubeChannels.GetDSAChannels()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(channels)
}
func GetRoadMapChannel(w http.ResponseWriter, _ *http.Request) {
	maps := youtubeChannels.GetRoadMaps()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(maps)
}
func GetFreeCourseChannels(w http.ResponseWriter, _ *http.Request) {
	courses := youtubeChannels.GetFreeCourse()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func GetComputerSubjects(w http.ResponseWriter, _ *http.Request) {
	subjects := youtubeChannels.GetSubjectChannels()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subjects)
}
func GetWebDev(w http.ResponseWriter, _ *http.Request) {
	web := Courses.GetWebAppData()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(web)
}
func GetCyber(w http.ResponseWriter, _ *http.Request) {
	cyber := Courses.GetCyberSecurityData()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cyber)
}
func GetMachineLearning(w http.ResponseWriter, _ *http.Request) {
	machine := Courses.GetMachineLearningData()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(machine)
}
func GetCloudComputing(w http.ResponseWriter, _ *http.Request) {
	cloud := Courses.GetCloudComputingData()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cloud)
}
func GetBigData(w http.ResponseWriter, _ *http.Request) {
	data := Courses.GetBigDataData()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func GetDevops(w http.ResponseWriter, _*http.Request) {
	dev := Courses.GetDevopsData()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dev)
}
