package Database

import (
	"placementCracker_api/Modals"
	"placementCracker_api/imageProcessing"
)

func GetDataPrograms() []Modals.Programs {
	programs := []Modals.Programs{
		{"Microsoft Student Partner", imageProcessing.ImageToBase64("images/microsoftAmbassador.png"), "https://studentambassadors.microsoft.com/en-us", "Microsoft Student Partners is a program to sponsor students who are majoring in technology-related disciplines. The student partners are a group of on-campus ambassadors who helps fellow students, lead in their local tech community as well as develop technical and career skills for the future. "},
		{"Mozilla Campus Club", imageProcessing.ImageToBase64("images/mozilla.png"), "https://campus.mozilla.community/", "Mozilla Campus Club is an initiative where a group of students from universities and colleges clubs together in order to build and innovate open-source projects and webs for the community. "},
		{"Student Ambassador By IBM", imageProcessing.ImageToBase64("images/ibm.png"), "https://developer.ibm.com/champions/", "The Student Ambassador at IBM will have the opportunity to work closely with IBM technical experts, recruitment leaders, and marketing teams to raise IBM’s profile on campus and introduce IBM to students and student groups across the respective campus."},
		{"Github Campus Experts", imageProcessing.ImageToBase64("images/github.jpg"), "https://education.github.com/students/experts", "The GitHub Campus Expert program includes learning various types of technical and professional skills such as technical writing, community leadership, public speaking and most of all, software development. "},
		{"Cisco DevNet Student Ambassador Program", imageProcessing.ImageToBase64("images/cisco.jpg"), "https://developer.cisco.com/site/student/", "Cisco DevNet is a community to help developers build on Cisco APIs and platforms. It provides the tools to learn, code, get inspired and connect with the Cisco technologies."},
		{"Adobe Creative ambassador", imageProcessing.ImageToBase64("images/adobe.png"), "https://airtable.com/shrV9AxjwHxv1tkWX", "You will help spread the word about Adobe’s wide product range, help out at Adobe events and hand out great freebies! "},
		{"Campus Ambassador At Coding Ninjas", imageProcessing.ImageToBase64("images/codingNinjas.png"), "https://www.codingninjas.com/blog/tag/campus-ambassador/", "Coding Ninjas is a diverse mix of educators-turned-programmers as well as programmers-turned-educators. The responsibilities of Campus Ambassador at Coding Ninjas include organise hackathons and talks(pre-placement) which are to be conducted at colleges, creating programming contests on CodeZen, maintain relevant databases, and other such"},
		{"HackerEarth Student Ambassador", imageProcessing.ImageToBase64("images/hackerearth.jpg"), "https://www.hackerearth.com/docs/wiki/campus/introduction/", "The HackerEarth University Ambassador Program is a platform for students who are passionate about programming to start and run an active programming club in their university."},
		{"OnePlus Student Ambassador", imageProcessing.ImageToBase64("images/oneplus.png"), "https://www.oneplus.in/campus", "OnePlus Student Ambassador Program is a six-month-long campus connect program which aims to form a network of like-minded students across India. The goal of the student ambassadors will be serving as the OnePlus community specialists, brand evangelists and technology experts in their respective institutions."},
		{"JetBrains Campus Ambassador", imageProcessing.ImageToBase64("images/jetbrains.png"), "https://blog.jetbrains.com/blog/2015/06/30/jetbrains-for-universities-public-events-in-passau-and-munich-july-9th-10th/#:~:text=JetBrains%20Campus%20Ambassador%20Program&text=The%20Campus%20Ambassador%20Program%20aims,funding%2C%20giveaways%2C%20and%20more.", "The Campus Ambassador Program by JetBrains aims to support students interested in JetBrains tools at the university. The program provides technical assistance and guidance, sponsorship of various activities and funding, giveaways, and more. "},
	}
	return programs
}
