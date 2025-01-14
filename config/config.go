package config

import (
	"os"
)

type Config struct {
	Twitter TwitterConfig
	Gemini  GeminiConfig
}

type TwitterConfig struct {
	Client_id           string ``
	Client_secret       string ``
	Consumer_key        string ``
	Consumer_secret     string ``
	Access_token        string ``
	Access_token_secret string ``
}

type GeminiConfig struct {
	Api_key string ``
}

var config Config

func Init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("err loading: %v", err)
	// }
	config.Twitter.Client_id = os.Getenv("CLIENT_ID")
	config.Twitter.Client_secret = os.Getenv("CLIENT_SECRET")
	config.Twitter.Consumer_key = os.Getenv("CONSUMER_KEY")
	config.Twitter.Consumer_secret = os.Getenv("CONSUMER_SECRET")
	config.Twitter.Access_token = os.Getenv("ACCESS_TOKEN")
	config.Twitter.Access_token_secret = os.Getenv("ACCESS_TOKEN_SECRET")
	config.Gemini.Api_key = os.Getenv("API_KEY")

}

func GetConfig() Config {
	return config
}
