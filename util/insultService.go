package util

import (
	"encoding/json"

	"github.com/michaelfraser99/randomgeneratordiscordbot/structs"
)

var (
	Insults structs.Insults
)

func GetInsults() structs.Insults {
	insultsJson := ReadJson("./data/insults.json")

	json.Unmarshal(insultsJson, &Insults)

	return Insults
}

func GetSingleInsult() string {
	return GetInsults().Insults[GetRandomNumber(len(Insults.Insults), 0)].Insult
}
