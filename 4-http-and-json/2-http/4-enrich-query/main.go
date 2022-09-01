package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println()

	req, err := http.NewRequest(http.MethodGet, "https://reqres.com/api/users", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Set("CuStOm", "CuStOm") // key is  non case-sensetive

	q := url.Values{} // map[string]string
	q.Add("page", "2")
	q.Add("pAgE", "2") // key is case-sensetive
	req.URL.RawQuery = q.Encode()

	fmt.Println()
	fmt.Println("Headers", req.Header)
	fmt.Println("Url", req.URL)
	fmt.Println()
	fmt.Println()
}
