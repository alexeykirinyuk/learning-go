package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

var counter = promauto.NewCounter(prometheus.CounterOpts{
	Namespace: "acme",
	Name:      "counter",
	Help:      "This is my counter",
})

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	go job()

	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe("127.0.0.1:2112", nil)
}

func job() {
	for {
		counter.Inc()

		delay := time.Duration(1000+rand.Intn(4000)) * time.Millisecond
		time.Sleep(delay)
	}
}
