package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/animenotifier/arn"
)

func importUsers(set string) {
	file := "../" + set + ".json"
	fmt.Println("Importing", file, "into", set)

	var allObjects []*arn.User
	data, _ := ioutil.ReadFile(file)
	json.Unmarshal(data, &allObjects)
	println(len(allObjects), set)

	for _, obj := range allObjects {
		arn.SetObject(set, obj.ID, obj)
	}
}

func importSettings(set string) {
	file := "../" + set + ".json"
	fmt.Println("Importing", file, "into", set)

	var allObjects []*arn.Settings
	data, _ := ioutil.ReadFile(file)
	json.Unmarshal(data, &allObjects)
	println(len(allObjects), set)

	for _, obj := range allObjects {
		arn.SetObject(set, obj.UserID, obj)
	}
}

func importPosts(set string) {
	file := "../" + set + ".json"
	fmt.Println("Importing", file, "into", set)

	var allObjects []*arn.Post
	data, _ := ioutil.ReadFile(file)
	json.Unmarshal(data, &allObjects)
	println(len(allObjects), set)

	for _, obj := range allObjects {
		arn.SetObject(set, obj.ID, obj)
	}
}

func importThreads(set string) {
	file := "../" + set + ".json"
	fmt.Println("Importing", file, "into", set)

	var allObjects []*arn.Thread
	data, _ := ioutil.ReadFile(file)
	json.Unmarshal(data, &allObjects)
	println(len(allObjects), set)

	for _, obj := range allObjects {
		arn.SetObject(set, obj.ID, obj)
	}
}

func main() {
	importUsers("User")
	importSettings("Settings")
	importPosts("Post")
	importThreads("Thread")
}
