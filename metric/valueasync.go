package metric

import (
	"log"
	"math/rand"
	"time"
)

type AsyncVal struct {
	RandVal   int
	Timestamp string
}

func CollectMetric() AsyncVal {
	m := AsyncVal{}

	t := time.Now()
	m.Timestamp = "Zeit ist " + t.Format(time.UnixDate)

	m.RandVal = rand.Intn(100)

	log.Print("call CollectMetric()")

	return m
}
