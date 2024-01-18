package main

import (
	"strconv"
	"time"
)

var Utterances = []string{"what day is it", "date today", "date", "what's today's date"}
var Name = "Correct Date"

func Action(transcribedText string, botSerial string, guid string, target string) (string, string) {
	year, month, day := time.Now().Date()
	yearSring := strconv.FormatInt(int64(year), 10)

	VECTOR_PHRASE := "The date is " + month.String() + " " + strconv.FormatInt(int64(day), 10) + ", " + yearSring + " "

	return "intent_imperative_praise", VECTOR_PHRASE
}
