package main

import (
	"github.com/rs/zerolog/log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":2112", nil); err != nil {
		log.Fatal().Err(err).Msg("http.ListenAndServe()")
	}
}
