package Courses

import (
	"placementCracker_api/Modals/Courses"
	"placementCracker_api/imageProcessing"
)

func GetCyberSecurityData() []Courses.CyberSecurity {
	data := []Courses.CyberSecurity{
		{
			imageProcessing.ImageToBase64("images/courses/cyberSecurity/c1.jpg"), imageProcessing.ImageToBase64("images/courses/cyberSecurity/t1.png"),
			"Master Certificate in Cyber Security (Red Team) – Jigsaw Academy with HackerU", "Jigsaw Academy’s Master Certificate in Cyber Security (Red Team), is the only program on offensive technology in India. The program is intensive in delivery and extensive in technology coverage and is delivered in collaboration with/by HackerU, Israel’s premier cybersecurity training institute. HackerU has more than 20 years of experience in providing cybersecurity solutions and training in the US, Singapore, Russia, Australia, and other geographies in the US and European market. The course covers more than 14 modules in 3 different phases focusing on network fundamentals, Windows, Linux Administration, applicative hacking and penetration testing on emerging technologies like IoT.",
			"https://www.jigsawacademy.com/master-certificate-in-cyber-security-red-team/",
		},
		{imageProcessing.ImageToBase64("images/courses/cyberSecurity/c2.jpeg"), imageProcessing.ImageToBase64("images/courses/cyberSecurity/t2.png"), "Stanford Advanced Computer Security Program –Great Learning",
			"Some of the salient features of the program are:\n\nA certificate of achievement from Stanford Engineering\nRegular mentorship from industry experts in cybersecurity\nHands-on practice through a series of labs and projects that allows participants to put what they learned to practice.", "https://www.mygreatlearning.com/online-cyber-security-course"},
		{imageProcessing.ImageToBase64("images/courses/cyberSecurity/c3.jpg"), imageProcessing.ImageToBase64("images/courses/cyberSecurity/t3.png"),
			"PGP in Cybersecurity – Praxis Business School", "The program is designed to create industry-ready cyber warriors by addressing the three aspects of the cybersecurity ecosystem — people, processes and technology, with special emphasis on governance and compliance. Praxis has forged extensive industry partnerships with CISCO, Fortinet, ISACA (Kolkata Chapter), British Standards Institute and Infosec Foundation to make the program relevant and effective. The curriculum with 525 hours of lectures, lab work, case studies and projects is distributed over three trimesters with specialisations in SOC analysis, digital forensics, security incident handling and information cyber audit.",
			"https://praxis.ac.in/px_program_details/pgp-in-cyber-security/"},
		{imageProcessing.ImageToBase64("images/courses/cyberSecurity/c4.jpg"), imageProcessing.ImageToBase64("images/courses/cyberSecurity/t4.png"),
			"Certified Ethical Hacker and Certified Information System Security Professional – Simplilearn", "This course will immerse the learner into a “hacker mindset” to teach how to think like a hacker, and better defend against future attacks. It also offers a hands-on training environment employing a systematic ethical hacking process. The course covers five phases of ethical hacking, diving into reconnaissance, gaining access, enumeration, maintaining access, and covering your tracks. Simplilearn’s CISSP certification training is aligned to the (ISC)² CBK latest requirements. The course trains you in the industry’s most recent best practices which will help you pass the exam in the first attempt.",
			"https://www.simplilearn.com/cyber-security/ceh-certification"},
		{imageProcessing.ImageToBase64("images/courses/cyberSecurity/c5.jpg"), imageProcessing.ImageToBase64("images/courses/cyberSecurity/t5.png"), "Cybersecurity Certification Course – Edureka",
			" Edurekas Cybersecurity Certification Course will help learners master the basic concepts of cybersecurity along with the methodologies that must be practised to ensure information security of an organization. Starting from the Ground level security essentials, this course will lead one through cryptography, computer networks and security, application security, data and endpoint security, idAM (Identity and Access Management), cloud security, cyber-attacks and various security practices for businesses. This course is designed to cover a holistic and a wide variety of foundation topics in cybersecurity which will prepare freshers as well as IT professionals for the next level of choice such as ethical hacking/ audit and compliance/GRC/Security Architecture and so on.",
			"https://www.edureka.co/cybersecurity-certification-training"},
	}
	return data
}
