package main

import (
	"encoding/json"
	"os"
)

type Video struct {
	Id string
	Title string
	Description string
	ImageUrl string
	Url string
}

func getVideos()(videos []Video) {
	fileBytes, err := os.ReadFile("videos.json")

	if err != nil {
		panic(err)	
	}

	err = json.Unmarshal(fileBytes, &videos)

	if err != nil {
		panic(err)	
	}

	return videos
}

func saveVideos(videos []Video)  {
	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./videos-updated.json", videoBytes, 0644)

	if err != nil {
		panic(err)
	}

	
}