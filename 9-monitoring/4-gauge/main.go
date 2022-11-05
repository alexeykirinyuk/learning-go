package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var gauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace: "acme",
	Name:      "acme_gauge",
	Help:      "This is my gauge",
})

func init() {
	rand.Seed(time.Now().UnixNano())
	prometheus.MustRegister(gauge)
}

func main() {
	per := 2
	val := [per]int64{}

	_ = val

	go func() {
		for {
			gauge.Add(rand.Float64()*15 - 5)
			time.Sleep(time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())

	log.Print("Start listen port :2112")
	if err := http.ListenAndServe("127.0.0.1:2112", nil); err != nil {
		log.Fatalf("error %v", err.Error())
	}
}
