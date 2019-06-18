package micrometer

import (
	"fmt"
	"sync"
)

type Meter interface {
	ID() *ID
	Measure() []Measurement
}

type MeterRegistry interface {
	Counter(id *ID) (Counter, error)
	//Summary(id *ID) error
	//Gauge(id *ID) error
	//Timer(id *ID) error
}

type simpleMeterRegistry struct {
	locker sync.RWMutex
	m      map[string]Meter // id -> meter
}

func (p *simpleMeterRegistry) Counter(id *ID) (Counter, error) {
	p.locker.Lock()
	_, ok := p.m[id.Name]
	if !ok {
		p.locker.Unlock()
		return nil, fmt.Errorf("duplicated meter name %s", id.Name)
	}
	c := NewCounter(id)
	p.m[id.Name] = c
	p.locker.Unlock()
	return c, nil
}

func NewSimpleMeterRegistry() MeterRegistry {
	return &simpleMeterRegistry{
		locker: sync.RWMutex{},
		m:      make(map[string]Meter),
	}
}
