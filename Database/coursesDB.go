package Database

import (
	"placementCracker_api/Modals"
	"placementCracker_api/imageProcessing"
)

func GetDataCourses() []Modals.Courses {
	courses := []Modals.Courses{
		{"OJas", imageProcessing.ImageToBase64("images/business.jpg"), "acniousnsaioucn"},
		{"OJas", imageProcessing.ImageToBase64("images/business.jpg"), "acniousnsaioucn"},
		{"OJas", imageProcessing.ImageToBase64("images/business.jpg"), "acniousnsaioucn"},
	}
	return courses
}

func GetDataJobs() []Modals.Jobs {
	jobs := []Modals.Jobs{
		{"HCL", "HR Senior Executive", imageProcessing.ImageToBase64("images/hcl.jpg"), "https://bit.ly/37qz7Sg"},
		{"Genpact", "Process Associate", imageProcessing.ImageToBase64("images/Genpact.png"), "https://bit.ly/3DR8cuS"},
		{"TCS", "TCS Digital Hiring", imageProcessing.ImageToBase64("images/tcs.png"), "https://jobforfreshar.in/index.php/tcs-recruitment-2022-for-digital-hiring/"},
		{"Wipro", "Associate", imageProcessing.ImageToBase64("images/wipro.jpg"), "https://bit.ly/3NoOz1Q "},
		{"Tata Steel", "Engineer Trainee", imageProcessing.ImageToBase64("images/tata.png"), "https://alljobsinindia.in/tata-steel-recruitment-2022/"},
	}
	return jobs
}
