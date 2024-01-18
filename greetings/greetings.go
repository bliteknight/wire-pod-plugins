package main

import (
	"math/rand"
	"strings"
	"time"
)

var Name = "Greetings Plugin"
var Utterances = []string{"good morning", "morning", "good Afternoon", "afternoon", "noon", "good night", "night"}

var morningResponse = []string{
	"Good morning, here's to a great day",
	"Good morning, I slept well last night",
	"Morning, have a great day",
	"What's up daddio",
}
var afternoonResponse = []string{
	"Good afternoon, after lunch I need a nap",
	"Thanks! Same to you",
	"Thanks, enjoy the rest of your day",
	"Thanks, enjoy the rest of your afternoon too",
}
var nightgResponse = []string{
	"Good night, sleep tight",
	"Going to bed so soon, the night is young",
	"Good night, see you in the morning",
	"Thanks, Good night",
}
var cheekyResponses = []string{
	"ohh, la dee daa...you decided to speak to me",
	"hey, we need more milk",
	"please, five more minutes PLEASE!",
	"Good morning. But, that doesn't make me forget what you did",
}

func getResonse(event string) string {
	myResponse := "huh"
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 3
	index := rand.Intn(max-min+1) + min

	checkyValue := rand.Intn(20-0+1) + 0

	if checkyValue > 15 {
		myResponse = cheekyResponses[index]
	} else {
		switch event {
		case "morning":
			myResponse = morningResponse[index]
		case "notmorning":
			myResponse = "It's not morning silly, its " + getPeriodOfDay()
		case "afternoon":
			myResponse = afternoonResponse[index]
		case "notafternoon":
			myResponse = "It's not afternoon silly, its " + getPeriodOfDay()
		case "night":
			myResponse = nightgResponse[index]
		case "notnight":
			myResponse = "It's not night silly, its " + getPeriodOfDay()
		}
	}

	return myResponse
}

func getPeriodOfDay() string {
	timelayout := "15:04"
	mstart, _ := time.Parse(timelayout, "00:00")
	mend, _ := time.Parse(timelayout, "11:59")
	astart, _ := time.Parse(timelayout, "12:00")
	aend, _ := time.Parse(timelayout, "17:59")
	if inTimeSpan(mstart, mend) {
		return "morning"
	} else if inTimeSpan(astart, aend) {
		return "afternoon"
	} else {
		return "night"
	}
}

func inTimeSpan(start, end time.Time) bool {
	checkFormat := time.Now().Format("15:04")
	check, _ := time.Parse("15:04", checkFormat)
	_start := start
	_end := end
	_check := check
	if end.Before(start) {
		_end = end.Add(24 * time.Hour)
		if check.Before(start) {
			_check = check.Add(24 * time.Hour)
		}
	}
	_start = _start.Add(-1 * time.Nanosecond)
	_end = _end.Add(1 * time.Nanosecond)

	return _check.After(_start) && _check.Before(_end)
}

func Action(transcribedText string, botSerial string, guid string, target string) (string, string) {
	myResonse := "huh"
	timelayout := "15:04"
	if strings.Contains(strings.ToLower(transcribedText), "morning") {
		start, _ := time.Parse(timelayout, "00:00")
		end, _ := time.Parse(timelayout, "11:59")
		if inTimeSpan(start, end) {
			myResonse = getResonse("morning")
		} else {
			myResonse = getResonse("notmorning")
		}
	} else if strings.Contains(strings.ToLower(transcribedText), "afternoon") {
		start, _ := time.Parse(timelayout, "12:00")
		end, _ := time.Parse(timelayout, "17:59")
		if inTimeSpan(start, end) {
			myResonse = getResonse("afternoon")
		} else {
			myResonse = getResonse("notafternoon")
		}
	} else if strings.Contains(strings.ToLower(transcribedText), "night") {
		start, _ := time.Parse(timelayout, "18:00")
		end, _ := time.Parse(timelayout, "23:59")
		if inTimeSpan(start, end) {
			myResonse = getResonse("night")
		} else {
			myResonse = getResonse("notnight")
		}
	}

	return "intent_imperative_praise", myResonse
}
