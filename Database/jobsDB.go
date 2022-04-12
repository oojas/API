package Database

import (
	"placementCracker_api/Modals"
	"placementCracker_api/imageProcessing"
)

func GetDataJobs() []Modals.Jobs {
	jobs := []Modals.Jobs{
		{"HCL", "HR Senior Executive", imageProcessing.ImageToBase64("images/hcl.jpg"), "https://bit.ly/37qz7Sg"},
		{"Genpact", "Process Associate", imageProcessing.ImageToBase64("images/Genpact.png"), "https://bit.ly/3DR8cuS"},
		{"TCS", "TCS Digital Hiring", imageProcessing.ImageToBase64("images/tcs.png"), "https://jobforfreshar.in/index.php/tcs-recruitment-2022-for-digital-hiring/"},
		{"Wipro", "Associate", imageProcessing.ImageToBase64("images/wipro.jpg"), "https://bit.ly/3NoOz1Q "},
		{"Tata Steel", "Engineer Trainee", imageProcessing.ImageToBase64("images/tata.png"), "https://alljobsinindia.in/tata-steel-recruitment-2022/"},
		{"Eloelo", "Android developer", imageProcessing.ImageToBase64("images/eloelo.jpg"), "https://tinyurl.com/5xe32nke"},
		{"Tata Communications", "Software Engineer", imageProcessing.ImageToBase64("images/tata-comm.jpg"), "https://alljobsinindia.in/tata-communications-recruitment/"},
		{"OYO", "Mobility Manager", imageProcessing.ImageToBase64("images/oyo.jpg"), "https://bit.ly/3600XV0"},
	}
	return jobs
}
