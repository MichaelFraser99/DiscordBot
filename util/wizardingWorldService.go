package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const API_URL = "https://wizard-world-api.herokuapp.com"

var ()

type House struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	HouseColours string  `json:"houseColours"`
	Founder      string  `json:"founder"`
	Animal       string  `json:"animal"`
	Element      string  `json:"element"`
	Ghost        string  `json:"ghost"`
	CommonRoom   string  `json:"commonRoom"`
	Heads        []Head  `json:"heads"`
	Traits       []Trait `json:"traits"`
}

type Head struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Trait struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Elixer struct {
	Id              string       `json:"id"`
	Name            string       `json:"name"`
	Effect          string       `json:"effect"`
	SideEffects     string       `json:"sideEffects"`
	Characteristics string       `json:"characteristics"`
	Time            string       `json:"time"`
	Difficulty      string       `json:"difficulty"`
	Ingredients     []Ingredient `json:"ingredients"`
	Inventors       string       `json:"inventors"`
	Manufacturer    string       `json:"manufacturer"`
}

type Ingredient struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetRandomElixer() Elixer {

	//Create and configure client
	client := &http.Client{}

	//Set up request
	req, err := http.NewRequest("GET", API_URL+"/Elixirs", nil)

	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("accept", "*/*")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var elixerList []Elixer
	json.Unmarshal(body, &elixerList)
	return elixerList[GetRandomNumber(len(elixerList), 0)]
}

func GetRandomHouse() House {

	//Create and configure client
	client := &http.Client{}

	//Set up request
	req, err := http.NewRequest("GET", API_URL+"/Houses", nil)

	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("accept", "*/*")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var houseList []House
	json.Unmarshal(body, &houseList)
	return houseList[GetRandomNumber(len(houseList), 0)]
}
