package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	http.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)

		_, _ = fmt.Fprintln(w, "ok")
	})

	server := http.Server{Addr: "127.0.0.1:8080"}

	go func() {
		srvErr := server.ListenAndServe() // Стартуем сервер и блокируем горутину
		if !errors.Is(srvErr, http.ErrServerClosed) {
			log.Fatal(srvErr)
		}

		log.Print("server successfully stopped")
	}()

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	termSig := <-termChan // Блокируем основную горутину тут

	log.Println("graceful shutdown starter with signal", termSig)

	// Создаем контекст заврешения программы с таймаутом
	closeCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Просим сервер завершить работу
	err := server.Shutdown(closeCtx)

	if err != nil {
		log.Println("server.Shutdown() failed:", err)
	}
}
