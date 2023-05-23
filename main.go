package zorp

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Zorp struct {
	Client        *resty.Client
	Configuration *Configuration
}

func New() *Zorp {
	client := resty.New().SetBaseURL("https://odguqqcioeujknbmifzp.functions.supabase.co")

	zorp := Zorp{
		Client: client,
	}

	return &zorp
}

type Metadata struct {
	Tags []string               `json:"tags"`
	Data map[string]interface{} `json:"data"`
}

type ActionStyle struct {
	Color           string `json:"color"`
	BackgroundColor string `json:"backgroundColor"`
}

type Action struct {
	Title string      `json:"title"`
	Type  string      `json:"type"`
	Url   string      `json:"url"`
	Color string      `json:"color"`
	Style ActionStyle `json:"style"`
}

type From struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	From    string      `json:"from"`
	Body    string      `json:"body"`
	To      string      `json:"to"`
	Tags    []string    `json:"tags"`
	Data    interface{} `json:"data"`
	Actions []Action    `json:"actions"`
}

type MessageGroup struct {
	From    string      `json:"from"`
	Body    string      `json:"body"`
	To      []string    `json:"to"`
	Tags    []string    `json:"tags"`
	Data    interface{} `json:"data"`
	Actions []Action    `json:"actions"`
}

type Configuration struct {
	Active   bool   `json:"active"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Type     string `json:"type"`
	Url      string `json:"url"`
}

func (z *Zorp) Configure(id string, configuration *Configuration) {
	configuration.Username = id

	body := map[string]interface{}{
		"id":            id,
		"configuration": configuration,
	}

	z.Client.R().SetBody(body).Post("setup")
	z.Configuration = configuration
}

type MessageFn func(message *Message)

func (z *Zorp) Message(message *Message) {
	if message.To == "" {
		message.To = "jack"
	}

	body := map[string]interface{}{
		"body":    message.Body,
		"to":      message.To,
		"from":    message.From,
		"tags":    message.Tags,
		"data":    message.Data,
		"actions": message.Actions,
	}

	if true {
		_, err := z.Client.R().SetBody(body).Post("push")

		if err != nil {
			fmt.Println("Failed to push", err.Error())
		}
	}
}
