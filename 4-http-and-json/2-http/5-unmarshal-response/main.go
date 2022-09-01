package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type getQuoteContract struct {
	Quote string `json:"quote"`
}

func main() {
	fmt.Println()

	ctx := context.Background()

	for i := 0; i < 10; i++ {
		quote := getQuote(ctx)

		fmt.Println(quote)
		fmt.Println()
	}
}

func getQuote(ctx context.Context) string {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.kanye.rest/", nil)
	if err != nil {
		log.Fatal(err)
	}

	httpClient := http.Client{}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	contract := getQuoteContract{}
	err = json.Unmarshal(respBody, &contract)
	if err != nil {
		log.Fatal(err)
	}

	return contract.Quote
}
