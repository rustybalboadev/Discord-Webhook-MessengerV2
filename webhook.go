package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/akamensky/argparse"
)

//JSONConfig struct for config.json file
type JSONConfig struct {
	WebhookURL      string `json:"webhook_url"`
	WebhookUsername string `json:"webhook_username"`
	AvatarURL       string `json:"avatar_url"`
}

//ErrorResponse struct for Discord Ratelimit response
type ErrorResponse struct {
	Message    string `json:"message"`
	RetryAfter int    `json:"retry_after"`
}

var (
	webhook string
	avatar  string
	message *string
	timeout *int
	amount  *int
	client  *http.Client = &http.Client{}
)

func readConfig() (string, string, string) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}

	data := JSONConfig{}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data.WebhookURL, data.WebhookUsername, data.AvatarURL
}

func spamWebhook(webhook string, avatar string, message *string, username string, timeout *int, amount *int) {
	if *amount == 0 {
		for {
			var e ErrorResponse
			jsonStr := []byte(fmt.Sprintf(`{"content": "%s", "username": "%s", "avatar_url": "%s"}`, *message, username, avatar))

			req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")

			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
			}
			err = json.Unmarshal(body, &e)
			if err == nil {
				fmt.Println("You Are Being Ratelimited Waiting " + strconv.Itoa(e.RetryAfter) + " Milliseconds")
				time.Sleep(time.Duration(e.RetryAfter+5) * time.Millisecond)
			}
			time.Sleep(time.Duration(*timeout) * time.Second)
		}
	} else {
		for i := 0; i < *amount; i++ {
			var e ErrorResponse
			jsonStr := []byte(fmt.Sprintf(`{"content": "%s", "username": "%s", "avatar_url": "%s"}`, *message, username, avatar))

			req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")

			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
			}
			err = json.Unmarshal(body, &e)
			if err == nil {
				fmt.Println("You Are Being Ratelimited Waiting " + strconv.Itoa(e.RetryAfter) + " Milliseconds")
				time.Sleep(time.Duration(e.RetryAfter+5) * time.Millisecond)
			}
			time.Sleep(time.Duration(*timeout) * time.Second)
		}
	}
}

func main() {

	parser := argparse.NewParser("Webhook Sender", "Spam a webhook with a message")
	message = parser.String("m", "message", &argparse.Options{Required: true, Help: "Message to Send Through Webhook"})
	timeout = parser.Int("t", "timeout", &argparse.Options{Required: false, Help: "Time Between Messages In Seconds"})
	amount = parser.Int("a", "amount", &argparse.Options{Required: false, Help: "Amount Of Times To Send Message"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		os.Exit(0)
	}
	fmt.Println("Webhook message sender | Made by: cookie#0003")
	webhook, username, avatar := readConfig()

	spamWebhook(webhook, avatar, message, username, timeout, amount)
}
