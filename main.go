package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/thilo-le/promeser/metric"
)

const (
	url                 = "localhost:8080"
	withGoStdCollectors = false
)

func main() {
	log.Printf("Start metric server on http://%s/metrics", url)

	reg := prometheus.NewRegistry()

	metric.RegisterMetricAsync(reg)
	reg.MustRegister(metric.NewSyncMetrics())

	if withGoStdCollectors {
		reg.MustRegister(collectors.NewBuildInfoCollector())
		reg.MustRegister(collectors.NewGoCollector())
	}

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

	http.ListenAndServe(url, nil)
}
