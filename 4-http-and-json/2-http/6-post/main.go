package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	url         = "https://httpbin.org/post"
	contentType = "application/json"
	reqBody     = `{"id":999, "value":"content"}`
)

func main() {
	httpClient := &http.Client{Timeout: time.Second}

	{
		req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(reqBody))
		req.Header.Add("Content-Type", contentType)

		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))
	}

	{
		resp, err := httpClient.Post(url, contentType,
			strings.NewReader(reqBody)) // урезанный способ, смысла мало
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))
	}
}
