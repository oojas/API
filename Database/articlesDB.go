package Database

import (
	"placementCracker_api/Modals"
	"placementCracker_api/imageProcessing"
)

func GetDataArticles() []Modals.Articles {
	articles := []Modals.Articles{
		{"Perspectives in Leadership: The data engineering leaders need to prioritize", imageProcessing.ImageToBase64("images/art9.png"), "https://www.pluralsight.com/blog/software-development/pil-data-engineering-leaders-must-prioritize"},
		{"Flutter packages which I use in every project.", imageProcessing.ImageToBase64("images/art1.jpeg"), "https://medium.com/@dawletataoglu/flutter-packages-which-i-use-in-every-project-d8bf14c43622"},
		{"Http Server Performance: NodeJS vs. Go", imageProcessing.ImageToBase64("images/art2.jpg"), "https://medium.com/better-programming/http-server-performance-nodejs-vs-go-397751e8d275"},
		{"Clean Architecture + MVVM", imageProcessing.ImageToBase64("images/art3.png"), "https://medium.com/@hemanthappu006/clean-architecture-mvvm-9f3c495c154c"},
		{"Making $300 Per Month With A Telegram Bot?", imageProcessing.ImageToBase64("images/art4.jpg"), "https://medium.com/@amir-tech/making-300-per-month-with-a-telegram-bot-1556c41903ee"},
		{"I Will Reject Your Resume If I See These 4 Things", imageProcessing.ImageToBase64("images/art5.jpeg"), "https://medium.com/startup-stash/i-will-reject-your-resume-if-i-see-these-4-things-4da5b59453c8"},
		{"WEB 3.0", imageProcessing.ImageToBase64("images/art6.jpg"), "https://medium.com/@vijay.viru146/web-3-0-a4def4f4328f"},
		{"How I built my tech startup as a solo developer", imageProcessing.ImageToBase64("images/art7.png"), "https://medium.com/dreamwod-tech/how-i-built-my-tech-startup-as-a-solo-developer-45390f460002"},
		{"How I got a FAANG offer without grinding Leetcode", imageProcessing.ImageToBase64("images/art8.png"), "https://medium.com/@contrapasso/how-i-got-faang-offers-without-grinding-leetcode-7e556243e9ce"},
	}
	return articles
}
