package Database

import (
	"placementCracker_api/Modals"
	"placementCracker_api/imageProcessing"
)

func GetDataPrograms() []Modals.Programs {
	programs := []Modals.Programs{
		{"Microsoft Student Partner", imageProcessing.ImageToBase64("images/microsoftAmbassador.png"), "https://studentambassadors.microsoft.com/en-us"},
		{"Mozilla Campus Club", imageProcessing.ImageToBase64("images/mozilla.png"), "https://campus.mozilla.community/"},
		{"Student Ambassador By IBM", imageProcessing.ImageToBase64("images/ibm.png"), "https://developer.ibm.com/champions/"},
		{"Github Campus Experts", imageProcessing.ImageToBase64("images/github.jpg"), "https://education.github.com/students/experts"},
		{"Campus Ambassador At Coding Ninjas", imageProcessing.ImageToBase64("images/codingNinjas.png"), "https://www.codingninjas.com/blog/tag/campus-ambassador/"},
		{"HackerEarth Student Ambassador", imageProcessing.ImageToBase64("images/hackerearth.jpg"), "https://www.hackerearth.com/docs/wiki/campus/introduction/"},
		{"OnePlus Student Ambassador", imageProcessing.ImageToBase64("images/oneplus.png"), "https://www.oneplus.in/campus"},
		{"JetBrains Campus Ambassador", imageProcessing.ImageToBase64("images/jetbrains.png"), "https://blog.jetbrains.com/blog/2015/06/30/jetbrains-for-universities-public-events-in-passau-and-munich-july-9th-10th/#:~:text=JetBrains%20Campus%20Ambassador%20Program&text=The%20Campus%20Ambassador%20Program%20aims,funding%2C%20giveaways%2C%20and%20more."},
	}
	return programs
}
