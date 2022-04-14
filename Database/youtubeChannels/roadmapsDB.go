package youtubeChannels

import (
	"placementCracker_api/Modals/channels"
	"placementCracker_api/imageProcessing"
)

func GetRoadMaps() []channels.RoadMaps {
	roadmaps := []channels.RoadMaps{
		{"Love Babbar", imageProcessing.ImageToBase64("images/babbar.jpg"), "Love Babbar has made many roadmaps on Whimsical.com, which made the life of students easy.", "https://tenowl.com/love-babbar-roadmaps/"},
		{"Anuj Kumar Sharma", imageProcessing.ImageToBase64("images/anuj.jpg"), "Passionate about Coding and online teaching beginners how to code efficiently.Ex-Amazon, Ex-Urban Company Software Engineer with nearly 3 years of Industry Experience.", "https://www.youtube.com/c/AnujBhaiya"},
	}
	return roadmaps
}
