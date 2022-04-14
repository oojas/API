package youtubeChannels

import (
	"placementCracker_api/Modals/channels"
	"placementCracker_api/imageProcessing"
)

func GetFreeCourse() []channels.FreeCourse {
	course := []channels.FreeCourse{

		{"freeCodeCamp.org", imageProcessing.ImageToBase64("images/freeCodecamp.png"), "The non-profit community on YoTube offers coding and tech learning. The channel provides free videos, articles, coding interview questions, and interacting learning lessons. You can get certification after completion of coding challenges. The Python course for beginners, SQL tutorial, and JavaScript full course are the most popular courses on this channel. freeCodeCamp.org has 1.92 million subscribers", "https://www.youtube.com/c/Freecodecamp"},
		{"Edureka", imageProcessing.ImageToBase64("images/edureka.jpg"), "The e-learning platform offers instructor-led courses, webinars, and lectures. It also has a Youtube Channel that hosts most of these courses. The most popular technologies covered by the channel include big data, DevOps, Data Science, Hadoop, Apache Spark, Python, Selenium, Blockchain, Tableau, Artificial Intelligence (AI), AWS, and digital marketing. The YouTube channel has 1.68 million subscribers.", "https://www.youtube.com/c/edurekaIN"},
		{" ProgrammingKnowledge", imageProcessing.ImageToBase64("images/programmingKnowledge.jpg"), "This channel offers fundamental knowledge one arias programming topics. It talks about online programming tutorials, coding strategies, installation of open-source software etc. The topics covered in the videos include Elastic Stack, Python, Android, Flutter, Socket Programming, MongoDB etc.", "https://www.youtube.com/c/ProgrammingKnowledge"},
		{"Telusko", imageProcessing.ImageToBase64("images/telusko.jpg"), "This channel offers free tutorials from beginner to advanced level. The technical topics covered in their videos include Python, Blockchain, Android, JavaScript, Rest API, Kotlin, Scala, Spring Framework, Networking etc. Telusko also offers motivation videos and online sessions with industry experts. The channel has 993,000 subscribers.", "https://www.youtube.com/playlist?list=PLsyeobzWxl7oRKwDi7wjrANsbhTX0IK0J"},
		{" Intellipaat", imageProcessing.ImageToBase64("images/Intellipaat.png"), "This YouTube channel offers free courses in big data, data science, and artificial intelligence. The video content also helps professionals in making a career decision. Some videos talk about how to assist corporate clients in upskilling their workforce. The channel has 496,000 subscribers", "https://www.youtube.com/c/Intellipaat"},
	}
	return course
}
