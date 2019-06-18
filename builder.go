package micrometer

import "fmt"

var builder = MeterBuilder{}

type MeterBuilder struct {
}

func (b MeterBuilder) Counter(name string) *CounterBuilder {
	return &CounterBuilder{
		name: name,
	}
}

type CounterBuilder struct {
	name string
	tags []string
	desc string
	unit string
}

func (p *CounterBuilder) Description(desc string) *CounterBuilder {
	p.desc = desc
	return p
}

func (p *CounterBuilder) Tag(k, v string) *CounterBuilder {
	p.tags = append(p.tags, k, v)
	return p
}

func (p *CounterBuilder) Tags(tags ...string) *CounterBuilder {
	if len(tags)&1 != 0 {
		panic(fmt.Errorf("micrometer: invalid tags amount, must be even"))
	}
	p.tags = append(p.tags, tags...)
	return p
}

func (p *CounterBuilder) Unit(unit string) *CounterBuilder {
	p.unit = unit
	return p
}

func (p *CounterBuilder) Register(registry MeterRegistry) (Counter, error) {
	id := &ID{
		Name: p.name,
		Type: TypeCounter,
		Desc: p.desc,
		Unit: p.unit,
	}
	for i, l := 0, len(p.tags); i < l; i += 2 {
		id.Tags = append(id.Tags, &Tag{
			K: p.tags[i],
			V: p.tags[i+1],
		})
	}
	return registry.Counter(id)
}

func Builder() MeterBuilder {
	return builder
}
