package Courses

import (
	"placementCracker_api/Modals/Courses"
	"placementCracker_api/imageProcessing"
)

func GetCloudComputingData() []Courses.CloudComputing {
	data := []Courses.CloudComputing{
		{imageProcessing.ImageToBase64("images/courses/cloud/c1.png"), imageProcessing.ImageToBase64("images/courses/webapp/t1.png"),
			"Introduction to Cloud Computing on Amazon AWS for Beginners", "This course explains key concepts of clouds like Iaas, PaaS, and Saas with AWS examples, making it easy to learn what Cloud is and what benefits it offers.this course is created by Infinite Skills and has, on average, 4.2 ratings from close to 2400 students, which is excellent. I recommend these books for absolute beginners who want to learn about Cloud computing in general and with AWS.",
			"https://www.udemy.com/course/introduction-to-cloud-computing-on-amazon-aws-for-beginners/?ranMID=39197&ranEAID=JVFxdTr9V80&ranSiteID=JVFxdTr9V80-iRozzpnWE7CUA9.EcYl6fw&LSNPUBID=JVFxdTr9V80&utm_source=aff-campaign&utm_medium=udemyads"},
		{imageProcessing.ImageToBase64("images/courses/cloud/c2.jpg"), imageProcessing.ImageToBase64("images/courses/cloud/t2.png"),
			"Cloud Computing Concepts by Coursera", "The course material is excellent, and instructor Indranil Gupta is phenomenal. He is exceptionally thorough, and his delivery is impressive, making learning a joyful experience.\n\nOn top of this, exams are marvelous and help you to retain the knowledge you have learned. This course is also part of Coursera's Cloud Computing Specialization, which is nothing but a collection of individual sessions to learn different areas of Cloud computing.",
			"https://www.coursera.org/learn/cloud-computing?ranMID=40328&ranEAID=JVFxdTr9V80&ranSiteID=JVFxdTr9V80-IZhyJgF9.XnqFwZgoGN22A&siteID=JVFxdTr9V80-IZhyJgF9.XnqFwZgoGN22A&utm_content=10&utm_medium=partners&utm_source=linkshare&utm_campaign=JVFxdTr9V80"},
		{imageProcessing.ImageToBase64("images/courses/cloud/c3.jpg"), imageProcessing.ImageToBase64("images/courses/webapp/et2.png"),
			"Cloud Computing 101,Master the Fundamentals", "This is another beginner-level, text-based interactive course on Cloud Computing, which I often recommend to my readers. This course will provide you with a fundamental understanding of what Cloud computing is and explain its essential characteristics.\n\nAs the title suggests, this best Educative cloud computing course is a 101 on cloud computing which means you will learn cloud computing from scratch, from basics to get enough knowledge to develop and deploy your own application son Cloud.\n\nIt also explains 3 main service models like IaaS (Infrastructure as Service), SaaS (Software as Service), and PaaS (Platform as Service), along with 4 cloud deployment models like private, public, hybrid, and community models.",
			"https://www.educative.io/courses/cloud-computing-101-master-the-fundamentals?affiliate_id=5073518643380224"},
		{imageProcessing.ImageToBase64("images/courses/cloud/c4.jpg"), imageProcessing.ImageToBase64("images/courses/cloud/t4.png"),
			"Cloud Computing,The Big Picture By David Chappell", "This course provides a good overview of cloud platforms, including Amazon Web Services and Microsoft Azure, and private clouds (bringing cloud technology on-premises).\n\nBy the time you're done, you will know what cloud computing is all about and be ready to start exploring specific implementations.",
			"https://www.pluralsight.com/courses/cloud-computing-big-picture?clickid=zZVUWSTKGxyIUsURIE30RT9pUkGXyFyUZ2jxXA0&irgwc=1&mpid=1193463&aid=7010a000001xAKZAA2&utm_medium=digital_affiliate&utm_campaign=1193463&utm_source=impactradius"},
		{imageProcessing.ImageToBase64("images/courses/cloud/c5.jpg"), imageProcessing.ImageToBase64("images/courses/webapp/t1.png"),
			"Ultimate AWS Certified Solutions Architect Associate by Stephane Maarek", "This is also the best course to pass the AWS Solutions Architect Associate Exam, but it also teaches you a lot of small details about cloud computing with AWS. You learn how the services you use on a daily basis on aws cloud works like S3, IAM, etc.",
			"https://www.udemy.com/course/aws-certified-solutions-architect-associate-saa-c02/?ranMID=39197&ranEAID=CuIbQrBnhiw&ranSiteID=CuIbQrBnhiw-1BPfGVtzepBKxBICrE1mRw&LSNPUBID=CuIbQrBnhiw&utm_source=aff-campaign&utm_medium=udemyads"},
	}
	return data
}
