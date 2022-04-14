package youtubeChannels

import (
	"placementCracker_api/Modals/channels"
	"placementCracker_api/imageProcessing"
)

func GetSubjectChannels() []channels.ComputerSubjects {
	sub := []channels.ComputerSubjects{
		{"Khan Academy", imageProcessing.ImageToBase64("images/khan.png"), "Khan Academy is a not-for-profit educational organization started by Salman Khan in 2008.Their mission is to provide a free, world-class education to anyone, anywhere.", "https://www.khanacademy.org/computing/computer-science"},
		{"Neso Academy", imageProcessing.ImageToBase64("images/neso.jpg"), "Neso Academy offers world-class learning resources on engineering courses, school syllabus, competitive exams, and many more.", "https://www.youtube.com/c/nesoacademy"},
		{"Freecodecamp.org", imageProcessing.ImageToBase64("images/freeCodecamp.png"), "The non-profit community on YoTube offers coding and tech learning. The channel provides free videos, articles, coding interview questions, and interacting learning lessons. You can get certification after completion of coding challenges. The Python course for beginners, SQL tutorial, and JavaScript full course are the most popular courses on this channel. freeCodeCamp.org has 1.92 million subscribers", "https://www.youtube.com/c/Freecodecamp"},
		{"Jennys lectures CS/IT NET&JRF Data structures and algorithms", imageProcessing.ImageToBase64("images/jenny.jpg"), "This channel provides 112 video lessons on various concepts from understanding the concept of arrays and implementing LinkedList to trees and graphs.It breaks down these concepts and explains them from the basic level. The first lesson breaks down arrays by explaining what arrays are, how to declare and implement arrays, and how arrays are stored in memory.", "https://www.youtube.com/playlist?list=PLdo5W4Nhv31bbKJzrsKfMpo_grxuLl8LU"},
		{"Abdul Bari", imageProcessing.ImageToBase64("images/adbul.jpg"), "Abdul Bari is a Udemy instructor with educational courses available for enrollment. Check out the latest courses taught by Abdul Bari.", "https://www.youtube.com/channel/UCZCFT11CWBi3MHNlGf019nw/playlists"},
		{"Takeyouforward", imageProcessing.ImageToBase64("images/takeyou.jpg"), "Takeuforward is the best place to learn data structures, algorithms, most asked coding interview questions, real interview experiences free of cost.", "https://www.youtube.com/c/takeUforward"},
		{"Code with Harry", imageProcessing.ImageToBase64("images/codeharry.png"), "Welcome to Code With Harry. Code With Harry is my attempt to teach basics and those coding techniques to people in short time which took me ages to learn.", "https://www.youtube.com/c/CodeWithHarry"},
		{"KnowledgeGate", imageProcessing.ImageToBase64("images/knowledgeGate.png"), "Knowledge Gate is an online educational platform that helps you in making your career at a very affordable price.", "https://www.youtube.com/c/KNOWLEDGEGATE_kg/playlists"},
		{"CS50", imageProcessing.ImageToBase64("images/cs50.png"), "Harvard offers its popular CS50 computer science classes online through the e-learning platform edX. You can enroll for free, or pay a fee to receive a certificate to add to your resume or LinkedIn.", "https://www.youtube.com/c/cs50/playlists"},
	}
	return sub
}
