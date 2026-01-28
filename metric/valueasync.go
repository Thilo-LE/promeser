package metric

import (
	"log"
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type asyncMetrics struct {
	valOne prometheus.Counter
	valTwo prometheus.GaugeVec
}

func newMetric(reg prometheus.Registerer) *asyncMetrics {
	m := &asyncMetrics{
		valOne: promauto.With(reg).NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "async_value_one",
			Help:      "This values will be shown the count of calculation",
		}),
		valTwo: *promauto.With(reg).NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "async_value_random",
				Help:      "Generate a random integer between 1000 and 2000.",
			},
			[]string{"name"},
		),
	}
	return m
}

func recordMetric(m *asyncMetrics) {
	go func() {
		for {
			m.valOne.Inc()
			x := CollectMetric()
			m.valTwo.With(prometheus.Labels{"name": "lorem"}).Set(float64(x.RandVal))
			time.Sleep(sleepCycle * time.Second)
		}
	}()
}

func RegisterMetricAsync(reg *prometheus.Registry) {
	m := newMetric(reg)
	recordMetric(m)
}

type AsyncVal struct {
	RandVal   int
	Timestamp string
}

func CollectMetric() AsyncVal {
	m := AsyncVal{}

	m.RandVal = rand.Intn(1000) + 1000

	log.Print("call asynchron values (concurrent mode)")

	return m
}
