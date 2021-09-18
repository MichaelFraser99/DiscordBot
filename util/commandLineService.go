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
		}

	}

	message.TTS = false
	message.Content = Config.CliKey +
		"\n\t country [quantity]" +
		"\n\t insult"
	return &message
}

func getApiKey(apiKeys []structs.ApiKey, api string) string {

	randommerApi, err := FindApi(apiKeys, api)

	if err != nil {
		log.Fatalf("Unable to retrieve %v api key", api)
	}

	return randommerApi.Key
}
