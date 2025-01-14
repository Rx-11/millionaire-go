package twitter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Rx-11/millionaire-go/config"
	"github.com/dghubble/oauth1"
)

type Twitter struct {
	client *http.Client
}

var X *Twitter

func Init() {
	X = GetTwitter()
}

func GetTwitter() *Twitter {
	return &Twitter{client: oauth1.NewClient(oauth1.NoContext, oauth1.NewConfig(config.GetConfig().Twitter.Consumer_key, config.GetConfig().Twitter.Consumer_secret),
		oauth1.NewToken(config.GetConfig().Twitter.Access_token, config.GetConfig().Twitter.Access_token_secret))}
}

func (t *Twitter) CreateTweet(body string) {
	tweetURL := "https://api.x.com/2/tweets"

	tweetBody := map[string]string{"text": body}
	jsonBody, err := json.Marshal(tweetBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	resp, err := t.client.Post(tweetURL, "application/json", strings.NewReader(string(jsonBody)))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Twitter API returned non-201 status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		var responseBody map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&responseBody)
		fmt.Println(responseBody)
	}
}
