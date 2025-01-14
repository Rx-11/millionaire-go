package gemini

import (
	"context"
	"fmt"

	"github.com/Rx-11/millionaire-go/config"
	"google.golang.org/genai"
)

type Gemini struct {
	client *genai.Client
}

var Gem *Gemini

func Init() {
	Gem = NewGemini()
}

func NewGemini() *Gemini {
	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey:  config.GetConfig().Gemini.Api_key,
		Backend: genai.BackendGoogleAI,
	})
	if err != nil {
		fmt.Println(err)
	}
	return &Gemini{client: client}
}

func (g *Gemini) Generate(prompt string) string {
	resp, err := g.client.Models.GenerateContent(context.Background(), "gemini-1.5-flash", genai.Text(prompt), &genai.GenerateContentConfig{
		SystemInstruction: genai.Text("Create a tweet in a way that it is destined to go viral and structure it in a way that it will increase user interaction. Give only the tweet and nothing else ONLY TWEET! Keep length of tweet inside 280 characters strictly")[0],
		Temperature:       func(f float64) *float64 { return &f }(2),
	})
	if err != nil {
		fmt.Println(err)
	}
	return string(resp.Candidates[0].Content.Parts[0].Text)
}
