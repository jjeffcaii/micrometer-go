package micrometer

type Type int8

const (
	_ Type = iota
	TypeCounter
	TypeGauge
	TypeLongTaskTimer
	TypeTimer
	TypeDistributionSummary
	TypeOther
)

type Tag struct {
	K string
	V string
}

type ID struct {
	Name string
	Tags []*Tag
	Type Type
	Unit string
	Desc string
}
