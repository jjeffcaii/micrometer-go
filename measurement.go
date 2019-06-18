package micrometer

import "fmt"

const (
	StatisticUnknown Statistic = iota
	StatisticTotal
	StatisticTotalTime
	StatisticCount
	StatisticMax
	StatisticValue
	StatisticActiveTasks
	StatisticDuration
)

type (
	Statistic int8

	Measurement interface {
		Statistic() Statistic
		Value() float64
	}

	measurement struct {
		g func() float64
		s Statistic
	}
)

func (m measurement) String() string {
	return fmt.Sprintf("Measurement{statistic='%s', value=%f}", m.s, m.g())
}

func (s Statistic) String() string {
	v, ok := statisticMap[s]
	if ok {
		return v
	}
	return statisticMap[StatisticUnknown]
}

var statisticMap = map[Statistic]string{
	StatisticUnknown:     "UNKNOWN",
	StatisticTotal:       "TOTAL",
	StatisticTotalTime:   "TOTAL_TIME",
	StatisticCount:       "COUNT",
	StatisticMax:         "MAX",
	StatisticValue:       "VALUE",
	StatisticActiveTasks: "ACTIVE_TASKS",
	StatisticDuration:    "DURATION",
}

func (m measurement) Statistic() Statistic {
	return m.s
}

func (m measurement) Value() float64 {
	return m.g()
}

func newMeasurement(g func() float64, s Statistic) Measurement {
	return measurement{
		g: g,
		s: s,
	}
}
