package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
)

func main() {
	
	webhook, err := ioutil.ReadFile("config.txt")
	if err == nil {
			fmt.Println("Webhook message sender | Made by: cookie#0003")
		}
	
	for {
		fmt.Printf("\nMessage: ")
		var MESSAGE string
		fmt.Scanln(&MESSAGE)
		
		username := "cookie v2"
		data := url.Values{
			"content": {MESSAGE},
			"username": {username},
			"avatar_url": {""},
		}
		http.PostForm(string(webhook), data)
	}

}