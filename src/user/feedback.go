package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Feedback struct {
	Text  string `json:"text"`
	Grade int    `json:"grade"`
	Card  int    `json:"card"`
}

type Card int

const (
	NOCARD     int = 0
	YELLOWCARD     = 1
	REDCARD        = 2
)

//Constructor
func NewFeedback(text string, grade int, card int) Feedback {
	return Feedback{text, grade, card}
}

//Returns wether Feedback has red Card - for usage in Template Logic
func (fb Feedback) IsRed() bool {
	if fb.Card == 2 {
		return true
	}
	return false
}

//Returns wether Feedback has yellow Card - for usage in Template Logic
func (fb Feedback) IsYellow() bool {
	if fb.Card == 1 {
		return true
	}
	return false
}

//Returns Feedback struct from given Matrikel and Postnumber
func GetFeedback(matrikel, postNr int) (Feedback, error) {
	nrStr, _ := intToString(postNr)
	var fb Feedback
	jsondata, err := ioutil.ReadFile(fmt.Sprintf("./coursedata/users/%v/post_%s_feedback.json", matrikel, nrStr))
	if err != nil {
		return fb, err
	}
	err = json.Unmarshal(jsondata, &fb)
	if err != nil {
		return fb, err
	}
	return fb, err
}

//Stores Feedback struct to PostNumber in User Folder
func StoreFeedback(matrikel int, postNr int, fb Feedback) error {
	jsondata, err := json.Marshal(fb)
	if err != nil {
		return err
	}
	postNrStr, _ := intToString(postNr)
	path := fmt.Sprintf("./coursedata/users/%v/post_%s_feedback.json", matrikel, postNrStr)
	err = ioutil.WriteFile(path, jsondata, 0777)
	if err != nil {
		return err
	}
	return nil
}
