package structs

type Greetings struct {
	Enabled         bool     `json:"enabled"`
	Tts             bool     `json:"tts"`
	AccountsIgnored []string `json:"accountsIgnored"`
}
