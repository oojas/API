package Courses

import (
	"placementCracker_api/Modals/Courses"
	"placementCracker_api/imageProcessing"
)

func GetMachineLearningData() []Courses.MachineLearning {
	data := []Courses.MachineLearning{
		{imageProcessing.ImageToBase64("images/courses/machine/c1.png"), imageProcessing.ImageToBase64("images/courses/machine/t1.png"),
			"Machine Learning(Stanford University)", "This course starts by laying down the mathematical foundations of machine learning. It begins with a review of linear algebra and univariate linear regression before moving on to multivariate and logistic regression.\n\nIt then jumps from topic to topic each week to cover a wide variety of machine learning techniques and models. These include deep learning, support-vector machines, and principal component analysis.\n\nFinally, it touches on practical aspects such as how to design and leverage large-scale machine learning projects.",
			"https://www.classcentral.com/course/machine-learning-835"},
		{imageProcessing.ImageToBase64("images/courses/machine/c2.png"), imageProcessing.ImageToBase64("images/courses/machine/t2.png"),
			"Machine Learning Foundations:A Case Study Approach(University of Washington)", "The course starts by contextualizing machine learning: explaining what machine learning is, going over some of its applications, and making a case for its importance in the future.\n\nThe course introduction also takes the time to cover Python fundamentals as well as the rudiments of tools like Jupyter Notebooks.\n\nThe course then moves from case study to case study, using each one to illustrate a particular facet of machine learning: you use regressions to predict house prices, you use classification to evaluate sentiments in user reviews, you use clustering for grouping related articles, you use deep learning to identify objects in images, and so on.",
			"https://www.classcentral.com/course/ml-foundations-4352"},
		{imageProcessing.ImageToBase64("images/courses/machine/c3.png"), imageProcessing.ImageToBase64("images/courses/machine/t3.png"),
			"Machine Learning for All(University of London)", "This course starts by explaining what artificial intelligence and machine learning are and how these disciplines are connected.\n\nIt discusses various real-world applications of machine learning, including AlphaGo, a machine learning program capable of beating the best Go players in the world. It explains data representation, how to set up a machine learning project, and some of the opportunities and ethical considerations of machine learning.\n\nFinally, the course invites you to implement a machine learning project by collecting data, training a model, and putting it to the test.",
			"https://www.classcentral.com/course/uol-machine-learning-for-all-17124"},
		{imageProcessing.ImageToBase64("images/courses/machine/c4.png"), imageProcessing.ImageToBase64("images/courses/machine/t4.png"),
			"Machine Learning with Python(IBM)", "The course starts by covering machine learning fundamentals and applications in fields such as healthcare, banking, and telecommunications. And it explains the difference between supervised and unsupervised learning, and goes over which type of learning is suitable for which type of task.\n\nEach week is dedicated to one of the broad machine learning tasks — regression, clustering, and classification — and the various methods that can be used to implement them, such as decision trees, support vector machines, and k-means.",
			"https://www.classcentral.com/course/machine-learning-with-python-11714"},
		{imageProcessing.ImageToBase64("images/courses/machine/c5.png"), imageProcessing.ImageToBase64("images/courses/machine/t5.jpg"),
			"Machine Learning(Georgia Tech)", "This course is divided into three broad machine learning tasks.\n\nFirst, it covers supervised learning, discussing decision trees, regression and classification, and neural networks. Then, it covers unsupervised learning, discussing clustering, feature selection, and randomized optimization. Finally, it covers reinforcement learning, discussing markov decision processes, game theory, and decision making.\n\nBy the end of the course, you’ll have a comprehensive understanding of supervised, unsupervised, and reinforcement learning, and the differences between them.",
			"https://www.classcentral.com/course/udacity-machine-learning-1020"},
		{imageProcessing.ImageToBase64("images/courses/machine/c6.png"), imageProcessing.ImageToBase64("images/courses/machine/t6.png"),
			"Machine Learning Crash Course with TensorFlow APIs(Google)", "The crash course begins by asking you about your background in machine learning. Depending on your answer, it will orient you toward the appropriate resources, so you can make the best use of your time.\n\nAssuming you’re a complete beginner, you’ll start from square one. So your learning path will cover fundamental machine learning concepts, including regressions, loss functions, and gradient descent.\n\nThe course uses TensorFlow, Google’s popular machine learning library. So rapidly the low-level details will be abstracted away by leveraging the library functions.",
			"https://www.classcentral.com/course/independent-machine-learning-crash-course-with-tensorflow-apis-10503"},
		{imageProcessing.ImageToBase64("images/courses/machine/c7.png"), imageProcessing.ImageToBase64("images/courses/webapp/t1.png"),
			"Machine Learning A-Z:Hands-On Python & R in Data Science", "The course starts by covering various types of regression, classification, and clustering models. It discusses reinforcement learning as well as natural language processing, and it covers the fundamentals of artificial neural networks.\n\nThe course uses the Python and R programming languages, and the TensorFlow machine learning library.",
			"https://www.classcentral.com/course/udemy-machinelearning-23826"},
		{imageProcessing.ImageToBase64("images/courses/machine/c8.png"), imageProcessing.ImageToBase64("images/courses/machine/t8.png"),
			"Introduction to Machine Learning in Production(DeepLearning.AI)", "This course starts by discussing the lifecycle of a machine learning project and how to deploy production-ready machine learning systems.\n\nThen, the course explains strategies to pick adequate models and train them, as well as some of the pitfalls to avoid when dealing with skewed data sets.\n\nFinally, the course covers how to handle classification problems and how to establish a baseline to assess your models performance.",
			"https://www.classcentral.com/course/introduction-to-machine-learning-in-production-43546"},
		{imageProcessing.ImageToBase64("images/courses/machine/c9.png"), imageProcessing.ImageToBase64("images/courses/webapp/t1.png"),
			"Python for Data Science and Machine Learning Bootcamp", "In the first part of the course, after setting up your development environment, you’ll jump into a Python crash course. You'll learn the fundamentals of the programming language as well as a plethora of widely used libraries, such as NumPy, Pandas, and Matplotlib.\n\nOnce you've integrated these skills, you’ll be equipped to tackle the second part of the course, which is entirely dedicated to machine learning.",
			"https://www.udemy.com/course/python-for-data-science-and-machine-learning-bootcamp/?ranMID=39197&ranEAID=SAyYsTvLiGQ&ranSiteID=SAyYsTvLiGQ-nJGHccqxZCS.Lee08a51xw&LSNPUBID=SAyYsTvLiGQ&utm_source=aff-campaign&utm_medium=udemyads"},
	}
	return data
}
