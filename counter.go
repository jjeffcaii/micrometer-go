package micrometer

import (
	ga "github.com/linxGnu/go-adder"
)

type (
	Counter interface {
		Meter
		Increment(amount float64)
		Count() float64
	}

	counter struct {
		id *ID
		v  ga.Float64Adder
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
	return p.v.Sum()
}

func NewCounter(id *ID) Counter {
	return &counter{
		id: id,
		v:  ga.NewJDKF64Adder(),
	}
}
