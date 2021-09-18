package util

import (
	"errors"

	"github.com/michaelfraser99/randomgeneratordiscordbot/structs"
)

func FindApi(apiKeys []structs.ApiKey, api string) (structs.ApiKey, error) {

	for _, apiKey := range apiKeys {
		if apiKey.Api == api {
			return apiKey, nil
		}
	}
	return structs.ApiKey{}, errors.New("no api key")
}
