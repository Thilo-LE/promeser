package metric

import (
	"log"
	"math/rand"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	labels = []string{"name", "state"}
)

type syncMetrics struct {
	mu sync.RWMutex

	valOne *prometheus.CounterVec
	valTwo *prometheus.GaugeVec
}

func NewSyncMetrics() *syncMetrics {

	return &syncMetrics{
		valOne: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "sync_value_one",
				Help:      "Increments by a random number [0,100]",
			}, labels),

		valTwo: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "sync_level_two",
				Help:      "Random float of [0,1]",
			}, labels),
	}

}

func (m *syncMetrics) Describe(ch chan<- *prometheus.Desc) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	m.valOne.Describe(ch)
	m.valTwo.Describe(ch)

}

func (m *syncMetrics) Collect(ch chan<- prometheus.Metric) {
	log.Print("call synchron values (sequencial mode)\n")

	m.mu.RLock()
	defer m.mu.RUnlock()

	// collect values
	sv := generateSyncVal()

	m.valOne.With(prometheus.Labels{"name": "lorem", "state": "true"}).Add(float64(sv.randomOne))
	m.valTwo.With(prometheus.Labels{"name": "ipsum", "state": "true"}).Set(sv.randomTwo)

	m.valOne.Collect(ch)
	m.valTwo.Collect(ch)
}

type syncVal struct {
	randomOne int
	randomTwo float64
}

func generateSyncVal() syncVal {
	s := syncVal{}

	s.randomOne = rand.Intn(100)
	s.randomTwo = rand.Float64()

	return s
}
