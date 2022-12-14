package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	httpClient := http.Client{
		Transport:     nil,             // Механика выполнения запросов, RoundTripper интерфейс
		CheckRedirect: nil,             // Функция, определяющая политику выполнения редиректа
		Jar:           nil,             // Общее хранилище cookie
		Timeout:       2 * time.Second, // Лимит времени выполнения запроса
	}

	res, err := httpClient.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("StatusCode", res.StatusCode)
}
