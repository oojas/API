package youtubeChannels

import (
	"placementCracker_api/Modals/channels"
	"placementCracker_api/imageProcessing"
)

func GetDSAChannels() []channels.DSATheory {
	channels := []channels.DSATheory{
		{"Tulesko Data structures using Java", imageProcessing.ImageToBase64("images/telusko.jpg"), "If you’re a Java developer, you may know or at the very least have heard of Tulesko. They provide some of the best easy to understand lessons on Java, Python, and JavaScript.It’s divided into 17 separate videos to make them easier to consume.From LinkedList data structure implementation to dynamic array and tree data structures, this tutorial provides a clear explanation with examples of the data structures.", "https://www.youtube.com/playlist?list=PLsyeobzWxl7oRKwDi7wjrANsbhTX0IK0J"},
		{"Caleb Curry Data Structures and Algorithms", imageProcessing.ImageToBase64("images/caleb.jpg"), "His series also includes 17 unique videos that guide learners from the basics of data structures and algorithms to advanced stages.The series will take you through data structures, algorithms, big O, and time complexity, searching, sorting, classes, objects, graphs, trees, and many more.", "https://www.youtube.com/playlist?list=PL_c9BZzLwBRLpDEpYRFXKBN-2ZCsAx0ps"},
		{"Freecodecamp Data Structures Easy to Advanced Course - Full Tutorial from a Google Engineer", imageProcessing.ImageToBase64("images/freeCodecamp.png"), "This is an 8-hour course meant to carry learners from the basic to advanced.It’s scary to think you’d need to sit down for eight hours to understand the concepts being taught but the video comes with time codes that will enable you to jump to specific sections.", "https://www.youtube.com/watch?v=RBSGKlAvoiM"},
		{"Dinesh Varyani Data structures and algorithms full course", imageProcessing.ImageToBase64("images/dinesh.png"), "This course focuses on Data structures and algorithms in Java. This doesn’t mean anyone learning other languages like C#, Javascript, C++, Python, and others can’t learn data structures from this course.The course provides training on implementing various data structures and algorithms and also covers interview room questions.", "https://www.youtube.com/playlist?list=PL6Zs6LgrJj3tDXv8a_elC6eT_4R5gfX4d"},
		{"CS Dojo Data structures and algorithms", imageProcessing.ImageToBase64("images/csDojo.jpg"), "This 13-part video course on data structures and algorithms provides a breakdown of the concept of data structures.The course is taught by YK Sugishita, a former software developer at Google. His lessons are well-explained to enable new programmers to follow along.", "https://www.youtube.com/playlist?list=PLBZBJbE_rGRV8D7XZ08LK6z-4zPoWzu5H"},
		{"Jenny’s lectures CS/IT NET&JRF –Data structures and algorithms", imageProcessing.ImageToBase64("images/jenny.jpg"), "This channel provides 112 video lessons on various concepts from understanding the concept of arrays and implementing LinkedList to trees and graphs.It breaks down these concepts and explains them from the basic level. The first lesson breaks down arrays by explaining what arrays are, how to declare and implement arrays, and how arrays are stored in memory.", "https://www.youtube.com/playlist?list=PLdo5W4Nhv31bbKJzrsKfMpo_grxuLl8LU"},
	}
	return channels
}
