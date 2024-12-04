package main

import (
	"time"
)

var Utterances = []string{"what time is it", "what's the time", "what time is", "time"}
var Name = "Correct Time" //Name variable is important

func Action(transcribedText string, botSerial string, guid string, target string) (string, string) {

	now := time.Now()

	currentTime := now.Format("3:04 PM")

	VECTOR_PHRASE := "The time is " + currentTime + " "

	return "intent_imperative_praise", VECTOR_PHRASE
}
