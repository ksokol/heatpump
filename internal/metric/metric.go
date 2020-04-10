package metric

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

const (
	t   = "temperature"
	om  = "operation_mode"
	o   = "output"
	oh  = "operating_hour"
	f   = "fault"
	s   = "shutdown"
	th  = "thermal" //kelvin
	st  = "state"
	c   = "count"
	p   = "pressure"
	e   = "energy"
	v   = "volt"
	ops = "operation_state"
)

type DataPoint struct {
	timestamp int64
	key       string
	value     interface{}
}

type CategoryData map[string][]DataPoint

type Metric struct {
	timestamp int64
	data      CategoryData
}

func (m *Metric) Reader() *strings.Reader {
	var buffer bytes.Buffer

	for category, dataPoints := range m.data {
		for _, dataPoint := range dataPoints {
			buffer.WriteString(
				fmt.Sprintf(
					"%s %s=%v %d\n",
					category,
					dataPoint.key,
					dataPoint.value,
					dataPoint.timestamp,
				),
			)
		}
	}

	return strings.NewReader(buffer.String())
}

func (m *Metric) addTemperature(field string, value interface{}) {
	m.add(t, field, value)
}

func (m *Metric) addOperationMode(field string, value interface{}) {
	m.add(om, field, value)
}

func (m *Metric) addOutput(field string, value interface{}) {
	m.add(o, field, value)
}

func (m *Metric) addOperatingHour(field string, value interface{}) {
	m.add(oh, field, value)
}

func (m *Metric) addFault(field string, value interface{}, timestamp int64) {
	m.addAtTime(f, field, value, timestamp)
}

func (m *Metric) addShutdown(field string, value interface{}, timestamp int64) {
	m.addAtTime(s, field, value, timestamp)
}

func (m *Metric) addEnergy(field string, value interface{}) {
	m.add(e, field, value)
}

func (m *Metric) addState(field string, value interface{}) {
	m.add(st, field, value)
}

func (m *Metric) addCount(field string, value interface{}) {
	m.add(c, field, value)
}

func (m *Metric) addPressure(field string, value interface{}) {
	m.add(p, field, value)
}

func (m *Metric) addThermal(field string, value interface{}) {
	m.add(th, field, value)
}

func (m *Metric) addVolt(field string, value interface{}) {
	m.add(v, field, value)
}

func (m *Metric) addOperationState(field string, value interface{}) {
	m.add(ops, field, value)
}

func (m *Metric) add(category string, field string, value interface{}) {
	m.addAtTime(category, field, value, m.timestamp)
}

func (m *Metric) addAtTime(category string, field string, value interface{}, timestamp int64) {
	m.data[category] = append(
		m.data[category],
		DataPoint{
			timestamp,
			field,
			value,
		},
	)
}

func New() *Metric {
	return &Metric{
		time.Now().UTC().Unix(),
		CategoryData{
			t:  make([]DataPoint, 0),
			om: make([]DataPoint, 0),
			o:  make([]DataPoint, 0),
			oh: make([]DataPoint, 0),
			f:  make([]DataPoint, 0),
			s:  make([]DataPoint, 0),
			th: make([]DataPoint, 0),
			st: make([]DataPoint, 0),
		},
	}
}
