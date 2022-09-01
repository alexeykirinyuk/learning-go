package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("https://example.com/")
	if err != nil {
		panic(err) // nак делать нельзя, лучше Fatal
	}

	fmt.Println("StatusCode", res.StatusCode)
}
