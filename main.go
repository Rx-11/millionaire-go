package main

import (
	"fmt"

	"github.com/Rx-11/millionaire-go/config"
	"github.com/Rx-11/millionaire-go/gemini"
	"github.com/Rx-11/millionaire-go/twitter"
)

func main() {
	config.Init()
	gemini.Init()
	body := gemini.Gem.Generate("You are @Millionaire1357 on twitter. You job is to grow this account. You can use the internet to check over the progress of the account. Now 'create a tweet giving a tip to become a millionaire' and grow the account")
	twitter.Init()
	twitter.X.CreateTweet(body)
	fmt.Println("Successfully Tweeted!")
}
