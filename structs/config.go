package structs

type Config struct {
	ApiKeys   []ApiKey  `json:"apiKeys"`
	CliKey    string    `json:"cliKey"`
	Greetings Greetings `json:"greetings"`
}
