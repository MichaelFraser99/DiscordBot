package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const COUNTRY_GENERATION = "https://randommer.io/api/Phone/Countries"

var (
	API_KEY string
)

type Country struct {
	Name        string
	CallingCode string
	CountryCode string
}

func RandomCountry(quantity int) chan []Country {
	r := make(chan []Country)

	fmt.Println("Retrieving countries")

	go func() {
		var countriesString = GetCountryList()

		var countryList []Country

		json.Unmarshal(countriesString, &countryList)

		var countriesToReturn []Country

		fmt.Printf("Found %v countries... \n", quantity)
		i := 0
		for i < quantity {
			countriesToReturn = append(countriesToReturn, countryList[GetRandomNumber(len(countryList)-1, i)])
			i++
		}

		fmt.Println("Returning...")
		r <- countriesToReturn
	}()

	return r
}

func GetCountryList() []byte {

	//Create and configure client
	client := &http.Client{}

	//Set up request
	req, err := http.NewRequest("GET", COUNTRY_GENERATION, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("accept", "*/*")
	req.Header.Add("X-Api-Key", API_KEY)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func GetCountries(apiKey string, quantity int) string {
	var countries []string
	API_KEY = apiKey

	channel := RandomCountry(quantity)

	returnedCountries := <-channel

	for _, s := range returnedCountries {
		countries = append(countries, s.Name)
	}

	return strings.Join(countries, ", ")
}
