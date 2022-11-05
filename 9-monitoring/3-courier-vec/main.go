package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

var counterVec = promauto.NewCounterVec(prometheus.CounterOpts{
	Namespace: "acme",
	Name:      "counter_vec",
	Help:      "Counter with labels",
}, []string{"label_name"})

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	for _, labelName := range []string{"a", "b", "c", "d", "e"} {
		go job(labelName)
	}

	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe("127.0.0.1:2112", nil)
}

func job(labelName string) {
	for {
		counterVec.WithLabelValues(labelName).Inc()

		delay := time.Duration(1000+rand.Intn(1000)) * time.Millisecond
		time.Sleep(delay)
	}
}
