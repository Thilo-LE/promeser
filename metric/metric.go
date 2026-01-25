package metric

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const namespace = "ths"

type metrics struct {
	keyValueOne    prometheus.Counter
	keyValueRandom prometheus.GaugeVec
}

func newMetricOne(reg prometheus.Registerer) *metrics {
	m := &metrics{
		keyValueOne: promauto.With(reg).NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "async_value_one",
			Help:      "This values will be increment every 2 seconds",
		}),
	}
	return m
}

func newMetricRandom(reg prometheus.Registerer) *metrics {
	m := &metrics{
		keyValueRandom: *promauto.With(reg).NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "async_value_random",
				Help:      "Every 5 seconds will be generate a random integer between 0 and 100.",
			},
			[]string{"timestamp"},
		),
	}
	return m
}

func recordMetricOne(m *metrics) {
	go func() {
		for {
			m.keyValueOne.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func recordMetricRandom(m *metrics) {
	go func() {
		for {
			x := CollectMetric()
			m.keyValueRandom.With(prometheus.Labels{"timestamp": x.Timestamp}).Set(float64(x.RandVal))
			time.Sleep(5 * time.Second)
		}
	}()
}

func RegisterMetricAsync(reg *prometheus.Registry) {
	m := newMetricOne(reg)
	recordMetricOne(m)

	n := newMetricRandom(reg)
	recordMetricRandom(n)
}
