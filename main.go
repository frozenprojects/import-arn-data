package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
	"github.com/jinzhu/copier"
	"github.com/mssola/user_agent"
)

var wg sync.WaitGroup

func convertUser(o *OldUser) *arn.User {
	user := new(arn.User)
	copier.Copy(&user, &o)

	user.Gender = arn.FixGender(o.Gender)

	if user.Location.CountryName == "-" && user.Location.CountryCode == "-" {
		// Reset location data to an empty struct
		user.Location = arn.UserLocation{}
	} else {
		user.Location.Latitude, _ = strconv.ParseFloat(o.Location.Latitude, 64)
		user.Location.Longitude, _ = strconv.ParseFloat(o.Location.Longitude, 64)
	}

	user.Accounts.Facebook.ID = o.Accounts.Facebook
	user.Accounts.Google.ID = o.Accounts.Google
	user.Accounts.Twitter.Nick = strings.TrimPrefix(o.Twitter, "@")
	user.Accounts.Osu.Nick = o.OsuDetails.Nick

	copier.Copy(&user.Accounts.Osu, &o.OsuDetails)

	if math.IsNaN(user.Accounts.Osu.PP) {
		user.Accounts.Osu.PP = 0
	}

	if math.IsNaN(user.Accounts.Osu.Level) {
		user.Accounts.Osu.Level = 0
	}

	if math.IsNaN(user.Accounts.Osu.Accuracy) {
		user.Accounts.Osu.Accuracy = 0
	}

	user.Accounts.AniList.Nick = o.ListProviders.AniList.UserName
	user.Accounts.MyAnimeList.Nick = o.ListProviders.MyAnimeList.UserName
	user.Accounts.AnimePlanet.Nick = o.ListProviders.AnimePlanet.UserName

	if o.Accounts.Twitter != 0 {
		user.Accounts.Twitter.ID = fmt.Sprint(o.Accounts.Twitter)
	}

	ua := user_agent.New(o.Agent.Source)

	name, version := ua.Browser()
	user.UserAgent = o.Agent.Source
	user.Browser.Name = name
	user.Browser.Version = version
	copier.Copy(&user.OS, ua.OSInfo())

	return user
}

func exportUsers() {
	defer wg.Done()

	userStream := make(chan *OldUser)
	err := arn.Scan("Users", userStream)

	if err != nil {
		panic(err)
	}

	allUsers := make([]*arn.User, 0)
	allSettings := make([]*arn.Settings, 0)

	count := 0
	for old := range userStream {
		count++

		// User
		user := convertUser(old)
		allUsers = append(allUsers, user)

		_, err := json.MarshalIndent(user, "", "    ")

		if err != nil {
			fmt.Printf("%+v\n", user)
			color.Red(err.Error())
		}

		// fmt.Println(string(data))

		// Settings
		settings := new(arn.Settings)
		settings.UserID = user.ID
		settings.TitleLanguage = old.TitleLanguage
		settings.SortBy = old.SortBy
		copier.Copy(&settings.Providers, &old.Providers)
		allSettings = append(allSettings, settings)
	}

	writeJSONFile("User.json", allUsers)
	writeJSONFile("Settings.json", allSettings)
}

func exportThreads() {
	defer wg.Done()

	stream := make(chan *arn.Thread)
	err := arn.Scan("Threads", stream)

	if err != nil {
		panic(err)
	}

	list := make([]*arn.Thread, 0)
	for obj := range stream {
		list = append(list, obj)
	}

	writeJSONFile("Thread.json", list)
}

func writeJSONFile(name string, data interface{}) {
	json, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(name, json, 0644)
}

func main() {
	wg.Add(2)

	go exportUsers()
	go exportThreads()

	wg.Wait()
	fmt.Println("Finished.")
}
