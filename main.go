package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/thilo-le/promeser/metric"
)

const (
	url = "localhost:8080"
)

func main() {
	log.Printf("Start metric server on http://%s/metrics", url)

	reg := prometheus.NewRegistry()
	metric.RegisterMetricAsync(reg)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	http.ListenAndServe(url, nil)
}
