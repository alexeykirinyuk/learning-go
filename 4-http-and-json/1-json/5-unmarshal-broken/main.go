package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
)

//go:embed data.json
var jsonBytes []byte

func main() {
	fmt.Println()

	data := new(contract)
	err := json.Unmarshal(jsonBytes, data)
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range data.Posts {
		fmt.Printf("%d: %s\n", post.ID, post.Title)
	}
}
