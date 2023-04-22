package main

import (
	"flag"
	"fmt"
	"os"
)

func main()  {
	// videos get command
	getCmd := flag.NewFlagSet("get",flag.ExitOnError)

	// inputs for videos get command
	getAll := getCmd.Bool("all", false, "Get all videos")
	getID := getCmd.String("id", "", "Get video by ID")

	//videos add command
	addCmd := flag.NewFlagSet("add",flag.ExitOnError)

	addID := addCmd.String("id", "", "Video ID")
	addTitle := addCmd.String("title", "", "Video title")
	addUrl := addCmd.String("url", "", "Video URL")
	addDescription := addCmd.String("description", "", "Video description")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAll, getID)
	case "add":
		HandleAdd(addCmd, addID, addTitle, addUrl, addDescription)
	default:
		fmt.Println("expected 'get' or 'add' subcommands")
	}
}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string)  {
	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Print("id is required or specify --all to get all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		videos := getVideos()

		fmt.Printf("ID \t Title \t Description \t ImageUrl \t Url\n")

		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v\n", video.Id, video.Title, video.Description, video.ImageUrl, video.Url)
		}

		return
	}

	if *id != "" {
		videos := getVideos()

		id := *id
		for _, video := range videos {
			if video.Id == id {
				fmt.Printf("ID \t Title \t Description \t ImageUrl \t Url\n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v\n", video.Id, video.Title, video.Description, video.ImageUrl, video.Url)
			}
		}
	}
}

func ValidateVideo(addCmd *flag.FlagSet, id *string, title *string, url *string, description *string)  {
	if *id == "" || *title == "" || *url == "" || *description == "" {
		fmt.Println("All fields are required")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}

func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, description *string)  {
	ValidateVideo(addCmd, id, title, url, description)
	video := Video {
		Id: *id,
		Title: *title,
		Url: *url,
		Description: *description,
	}

	videos := getVideos()

	videos = append(videos, video)
	
}
