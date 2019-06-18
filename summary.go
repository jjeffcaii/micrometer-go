package micrometer

type Summary interface {
	Meter
	Record(v float64)
	Count() int64
	Amount() float64
	Mean() float64
	Max() float64
}

type distributionSummary struct {
}

func (p *distributionSummary) Measure() []Measurement {
	panic("implement me")
}

func (p *distributionSummary) Record(v float64) {
	panic("implement me")
}

func (p *distributionSummary) Count() int64 {
	panic("implement me")
}

func (p *distributionSummary) Amount() float64 {
	panic("implement me")
}

func (p *distributionSummary) Mean() float64 {
	panic("implement me")
}

func (p *distributionSummary) Max() float64 {
	panic("implement me")
}

func (p *distributionSummary) ID() *ID {
	panic("implement me")
}
