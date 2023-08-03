package micrometer

import (
	"go.uber.org/atomic"
)

type (
	Counter interface {
		Meter
		Increment(amount float64)
		Count() float64
	}

	counter struct {
		id *ID
		v  *atomic.Float64
	}
)

func (p *counter) ID() *ID {
	return p.id
}

func (p *counter) Measure() []Measurement {
	return []Measurement{
		newMeasurement(
			func() float64 {
				return p.Count()
			},
			StatisticCount,
		),
	}
}

func (p *counter) Increment(amount float64) {
	p.v.Add(amount)
}

func (p *counter) Count() float64 {
	return p.v.Load()
}

func NewCounter(id *ID) Counter {
	return &counter{
		id: id,
		v:  atomic.NewFloat64(0),
	}
}
