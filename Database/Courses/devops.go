package Courses

import (
	"placementCracker_api/Modals/Courses"
	"placementCracker_api/imageProcessing"
)

func GetDevopsData() []Courses.Devops {
	data := []Courses.Devops{
		{imageProcessing.ImageToBase64("images/courses/devops/c1.jpg"), imageProcessing.ImageToBase64("images/courses/webapp/t1.png"),
			"Learn DevOps,The Complete Kubernetes Course", "This course will help you to gain understanding how to deploy, use, and maintain your applications on Kubernetes. If you are into DevOps, this is a technology you need to master. Kubernetes has gained a lot of popularity lately and it is a well sought skill by companies.\n\nThis course is updated frequently to include the features of latest releases!",
			"https://www.udemy.com/course/learn-devops-the-complete-kubernetes-course/"},
		{
			imageProcessing.ImageToBase64("images/courses/devops/c2.png"), imageProcessing.ImageToBase64("images/courses/cloud/t4.png"),
			"DevOps,The Big Picture", "This course on Pluralsight shows you the methodologies behind the DevOps term like what is DevOps and how your organization transforms its workflow of the development and operations team to work together in the product build and deployment.\n\nThis course focused on the problems the DevOps can solve and transform the whole organization to work under these workflow terms and the different technologies and tools that used to transform the work inside the development and operations team.",
			"https://www.pluralsight.com/courses/devops-big-picture?clickid=zZVUWSTKGxyIUsURIE30RT9pUkGXRnTsZ2jxXA0&irgwc=1&mpid=1193463&aid=7010a000001xAKZAA2&utm_medium=digital_affiliate&utm_campaign=1193463&utm_source=impactradius",
		},
		{
			imageProcessing.ImageToBase64("images/courses/devops/c3.png"), imageProcessing.ImageToBase64("images/courses/cloud/t2.png"),
			"DevOps Culture and Mindset", "This course gives you the basic foundational principles of DevOps with a particular focus on culture and the DevOps mindset. We’ll learn about how DevOps is grounded in lean principles, and how it can help improve collaboration between developers and operations team members. We'll learn about ideas regarding systems thinking, feedback loops, continuous improvement, loosely coupled architecture and teams, managing risk, and dealing with unplanned work.",
			"https://www.coursera.org/learn/devops-culture-and-mindset?ranMID=40328&ranEAID=JVFxdTr9V80&ranSiteID=JVFxdTr9V80-waMezbjhHa_azBqQs_LFag&siteID=JVFxdTr9V80-waMezbjhHa_azBqQs_LFag&utm_content=10&utm_medium=partners&utm_source=linkshare&utm_campaign=JVFxdTr9V80",
		},
		{imageProcessing.ImageToBase64("images/courses/devops/c4.png"), imageProcessing.ImageToBase64("images/courses/cloud/t2.png"),
			"Continuous Delivery &DevOps", "Amazon famously delivers new code every 11.6 seconds. Just a few years ago, this was unthinkable: many ‘cutting edge’ firms would release software quarterly. When it comes to digital innovation, velocity is critical and many would say it’s the most reliable determinant of success. \n\nBringing an organization to the state of the art (or even functional capability) in this area requires strong work in a combination of disciplines and a combination of both technical and managerial skills.",
			"https://www.coursera.org/learn/uva-darden-continous-delivery-devops?ranMID=40328&ranEAID=JVFxdTr9V80&ranSiteID=JVFxdTr9V80-iubw7wHQWTitoXwkfOfPyQ&siteID=JVFxdTr9V80-iubw7wHQWTitoXwkfOfPyQ&utm_content=10&utm_medium=partners&utm_source=linkshare&utm_campaign=JVFxdTr9V80"},
		{imageProcessing.ImageToBase64("images/courses/devops/c5.png"), imageProcessing.ImageToBase64("images/courses/webapp/t1.png"),
			"Learn DevOps,Infrastructure Automation With Terraform", "Terraform has gained a lot in popularity lately and is the tool you need to master if you are or about to get into a Ops / DevOps role. You typically use a technology like Ansible, Chef, or Puppet to automate the provisioning of software. Terraform starts from the same principle, infrastructure as code, but focusses on the automation of the infrastructure itself. Your whole Cloud infrastructure (instances, volumes, networking, IPs) can be described in terraform.",
			"https://www.udemy.com/course/learn-devops-infrastructure-automation-with-terraform/?ranMID=39197&ranEAID=JVFxdTr9V80&ranSiteID=JVFxdTr9V80-ZryouHPZZwqLT5fmhgLNRA&LSNPUBID=JVFxdTr9V80&utm_source=aff-campaign&utm_medium=udemyads"},
		{
			imageProcessing.ImageToBase64("images/courses/devops/c6.jpg"), imageProcessing.ImageToBase64("images/courses/webapp/t1.png"),
			"Docker Mastery,with Kubernetes +Swarm from a Docker Captain", "Some of the many cool things you'll do in this course:\n\nEdit web code on your machine while it's served up in a container\n\nLock down your apps in private networks that only expose necessary ports\n\nCreate a 3-node Swarm cluster in the cloud\n\nInstall Kubernetes and learn the leading server cluster tools\n\nUse Virtual IP's for built-in load balancing in your cluster\n\nOptimize your Dockerfiles for faster building and tiny deploys\n\nBuild/Publish your own custom application images\n\nLearn the differences between Kubernetes and Swarm\n\nCreate your own image registry\n\nUse Swarm Secrets to encrypt your environment configs, even on disk\n\nDeploy container updates in a rolling always up design\n\nCreate the config utopia of a single set of YAML files for local dev, CI testing, and prod cluster deploys",
			"https://www.udemy.com/course/docker-mastery/?ranMID=39197&ranEAID=JVFxdTr9V80&ranSiteID=JVFxdTr9V80-xIXTQRA64yjKSOag7K94Ug&LSNPUBID=JVFxdTr9V80&utm_source=aff-campaign&utm_medium=udemyads",
		},
	}
	return data
}
