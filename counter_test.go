package micrometer

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func BenchmarkCounter(b *testing.B) {
	c := NewCounter(nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Increment(0.1)
		}
	})
}

func TestCounter(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	c := NewCounter(nil)
	var sum float64
	for range [10000]struct{}{} {
		v := r.Float64()
		c.Increment(v)
		sum += v
	}
	require.Equal(t, sum, c.Count(), "bad count")
	for _, value := range c.Measure() {
		log.Println(value)
	}
}
