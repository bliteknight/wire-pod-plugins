package main

import (
	"strings"
)

var Utterances = []string{"simon says", "repeat"}
var Name = "Simon Says Plugin"

func stripOutTriggerWords(s string) string {
	result := strings.Replace(s, "simon says", "", 1)
	result = strings.Replace(result, "repeat", "", 1)
	return result
}

func CountWords(s string) int {
	return len(strings.Fields(s))
}

func Action(transcribedText string, botSerial string, guid string, target string) (string, string) {

  VECTOR_PHRASE := stripOutTriggerWords(transcribedText)

 /* Other method of getting Vector to talk
	localURL := "http://127.0.0.1:8080/api-sdk"
	fmt.Println("(in Simon Says) Transcribed text: " + transcribedText + ", serial number: " + botSerial)

	//Take Control Of Vector
	log.Printf("Taking control of Vector " + botSerial)
	resp, err := http.Get(localURL + "/assume_behavior_control?priority=high&serial=" + botSerial)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	responseText := string(body)
	log.Printf(responseText)
	log.Printf("Waiting for vector to be ready...2 secs")

	//We give Vector 2 seconds to get ready
	time.Sleep(2 * time.Second)
	


	//Tell vector to repeat the phrase
	log.Printf("Sending vector the phrase for vector to be ready...2 secs")
	resp, error := http.Get(localURL + "/say_text?text=" + url.QueryEscape(VECTOR_PHRASE) + "&serial=" + botSerial)
	if error != nil {
		log.Println(error)
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	//Wait X amount of time depending on words
	wc := CountWords(VECTOR_PHRASE)
	fmt.Println(wc)
	log.Println("Giving Vector " + strconv.FormatInt(int64(wc), 10) + " seconds to say text")
	time.Sleep(time.Duration(wc) * time.Second)
	log.Printf("All done, releasing Vector")

	// RELEASE 
	resp, err = http.Get(localURL + "/release_behavior_control?serial=" + botSerial)
	if err != nil {
		log.Println(err)
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}*/
	return "intent_imperative_praise", VECTOR_PHRASE
}
