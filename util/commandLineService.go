package util

import (
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/michaelfraser99/randomgeneratordiscordbot/structs"
)

var (
	apiKey string
)

//TODO: sort to random shit
var message discordgo.MessageSend

func ProcessArgument(m *discordgo.MessageCreate, Config structs.Config) *discordgo.MessageSend {

	if len(m.Content) > (len(Config.CliKey) + 1) {

		parameters := strings.Fields(m.Content)

		if parameters[1] == "country" {

			apiKey = getApiKey(Config.ApiKeys, "randommer")

			var numberOfCountries int

			if len(parameters) > 2 {

				if _, err := strconv.ParseInt(parameters[2], 10, 64); err == nil {
					numberOfCountries, err = strconv.Atoi(parameters[2])

					if err != nil {
						log.Fatalf("Failed %v quantity as integer: %v", parameters[2], err)
					}

					log.Printf("Sending country set")

					message.TTS = false
					message.Content = GetCountries(apiKey, numberOfCountries)
					return &message
				}
			}
			log.Printf("Sending single country")

			message.TTS = false
			message.Content = GetCountries(apiKey, 1)
			return &message

		} else if parameters[1] == "insult" {
			log.Printf("Sending insult")

			message.TTS = true
			message.Content = GetSingleInsult()
			return &message
		} else if parameters[1] == "wizard" {
			log.Printf("Wizard api")
			if len(parameters) > 2 {
				switch parameters[2] {
				case "elixer":
					log.Printf("Elixer selected")
					elixer := GetRandomElixer()

					message.TTS = false
					message.Content += elixer.Name

					return &message
				case "house":
					log.Printf("House selected")
					house := GetRandomHouse()

					message.TTS = false
					message.Content += house.Name

					return &message
				default:
					log.Printf("None valid value")
				}
			} else {
				message.TTS = false
				message.Content = "Select option for this command"
				return &message
			}
		}

	}

	message.TTS = false
	message.Content = Config.CliKey +
		"\n\t country [quantity]" +
		"\n\t insult" +
		"\n\t wizard [elixer,house]"
	return &message
}

func getApiKey(apiKeys []structs.ApiKey, api string) string {

	randommerApi, err := FindApi(apiKeys, api)

	if err != nil {
		log.Fatalf("Unable to retrieve %v api key", api)
	}

	return randommerApi.Key
}
