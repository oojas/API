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
