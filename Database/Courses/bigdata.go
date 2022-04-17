package Courses

import (
	"placementCracker_api/Modals/Courses"
	"placementCracker_api/imageProcessing"
)

func GetBigDataData() []Courses.BigData {
	data := []Courses.BigData{
		{
			imageProcessing.ImageToBase64("images/courses/bigdata/c1.jpg"), imageProcessing.ImageToBase64("images/courses/webapp/t1.png"),
			"The Ultimate Hands-On Hadoop Course,Tame your Big Data!", "This is seriously the ultimate course on learning Hadoop and other Big Data technologies as it covers Hadoop, MapReduce, HDFS, Spark, Hive, Pig, HBase, MongoDB, Cassandra, Flume, etc.\nIn this course, you will learn to design distributed systems that manage a huge amount of data using Hadoop and related technology.\nYou will not only learn how to use Pig and Spark to create scripts to process data on the Hadoop cluster but also how to analyze non-relational data using HBase, Cassandra, and MongoDB.",
			"https://www.udemy.com/course/the-ultimate-hands-on-hadoop-tame-your-big-data/?LSNPUBID=JVFxdTr9V80&ranEAID=JVFxdTr9V80&ranMID=39197&ranSiteID=JVFxdTr9V80-zyzc91X05gaQUwnodwysgg&utm_medium=udemyads&utm_source=aff-campaign"},
		{imageProcessing.ImageToBase64("images/courses/bigdata/c2.jpeg"), imageProcessing.ImageToBase64("images/courses/cloud/t2.png"),
			"Big Data Specialization on Coursera", "There are 6 courses in this specialization:\nIntroduction to Big Data\nBig Data Modelling and Management Systems\nBig Data Integration and Processing\nMachine Learning with Big Data\nGraph Analytics for Big Data\nBig Data Capstone Project",
			"https://www.coursera.org/specializations/big-data?irclickid=zZVUWSTKGxyIUsURIE30RT9pUkGXUL2gZ2jxXA0&irgwc=1&utm_medium=partners&utm_source=impact&utm_campaign=3294490&utm_content=b2c"},
		{imageProcessing.ImageToBase64("images/courses/bigdata/c3.jpeg"), imageProcessing.ImageToBase64("images/courses/cloud/t4.png"),
			"The Building Blocks of Hadoop Course,HDFS,MapReduce,and YARN", "In this course first, you will learn about Hadoop architecture and then do some hands-on work by setting up a pseudo-distributed Hadoop environment.\nYou will submit and monitor tasks in that environment and slowly learn how to make configuration choices for stability, optimization, and scheduling of your distributed system.\nAt the end of this course, you should have complete knowledge of how Hadoop works and its individual building blocks e.g.HDFS,MapReduce,and YARN",
			"https://www.pluralsight.com/courses/building-blocks-hadoop-hdfs-mapreduce-yarn?clickid=zZVUWSTKGxyIUsURIE30RT9pUkGXULWAZ2jxXA0&irgwc=1&mpid=1193463&aid=7010a000001xAKZAA2&utm_medium=digital_affiliate&utm_campaign=1193463&utm_source=impactradius"},
		{imageProcessing.ImageToBase64("images/courses/bigdata/c4.jpg"), imageProcessing.ImageToBase64("images/courses/cloud/t4.png"),
			"SQL on Hadoop,Analyzing Big Data with Hive", "The course starts with explaining key Apache Hadoop concepts like distributed computing, MapReduce and then goes into great detail about Apache Hive.The course presents some real-world challenges to demonstrate how Hive makes that task easier to accomplish.\nIn short, a good course to learn how to use Hive query language to find the solution to common Big Data Problems.",
			"https://www.pluralsight.com/courses/sql-hadoop-analyzing-big-data-hive?clickid=zZVUWSTKGxyIUsURIE30RT9pUkGXUI2AZ2jxXA0&irgwc=1&mpid=1193463&aid=7010a000001xAKZAA2&utm_medium=digital_affiliate&utm_campaign=1193463&utm_source=impactradius"},
		{
			imageProcessing.ImageToBase64("images/courses/bigdata/c5.jpg"), imageProcessing.ImageToBase64("images/courses/webapp/t1.png"),
			"Learn Big Data,The Hadoop Ecosystem Masterclass", "The course is very hands on but comes with the right amount of theory. It contains more than 6 hours of lectures to teach you everything you need to know about Hadoop.\nYou will also learn how to install and configure the Hortonworks Data Platform or HDP. It provides demons that you can try out on your machine by setting up a Hadoop cluster on the virtual machine. Though, you need 8GB or more RAM for that.",
			"https://www.udemy.com/course/learn-big-data-the-hadoop-ecosystem-masterclass/?LSNPUBID=JVFxdTr9V80&ranEAID=JVFxdTr9V80&ranMID=39197&ranSiteID=JVFxdTr9V80-8eLi0Wb2UedxVdyUkqVtaQ&utm_medium=udemyads&utm_source=aff-campaign",
		},
	}
	return data
}
