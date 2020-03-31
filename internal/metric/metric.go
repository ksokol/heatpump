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
	hw  = "hot_water"
	c   = "cooling"
	h   = "heating"
	com = "cooling_operation_mode"
	vom = "ventilation_operation_mode"
	so  = "solar"
	o   = "output"
	oh  = "operating_hours"
	f   = "fault"
	sd  = "shutdowns"
	th  = "thermal"
	st  = "state"
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

func (m *Metric) addCoolingOperationMode(field string, value interface{}) {
	m.add(com, field, value)
}

func (m *Metric) addHotWater(field string, value interface{}) {
	m.add(hw, field, value)
}

func (m *Metric) addCooling(field string, value interface{}) {
	m.add(c, field, value)
}

func (m *Metric) addHeating(field string, value interface{}) {
	m.add(h, field, value)
}

func (m *Metric) addVentilationOperationMode(field string, value interface{}) {
	m.add(vom, field, value)
}

func (m *Metric) addSolar(field string, value interface{}) {
	m.add(so, field, value)
}

func (m *Metric) addOutput(field string, value interface{}) {
	m.add(o, field, value)
}

func (m *Metric) addOperatingHours(field string, value interface{}) {
	m.add(oh, field, value)
}

func (m *Metric) addFault(field string, value interface{}, timestamp int64) {
	m.addAtTime(f, field, value, timestamp)
}

func (m *Metric) addShutdowns(field string, value interface{}, timestamp int64) {
	m.addAtTime(sd, field, value, timestamp)
}

func (m *Metric) addThermal(field string, value interface{}) {
	m.add(th, field, value)
}

func (m *Metric) addState(field string, value interface{}) {
	m.add(st, field, value)
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
			t:   make([]DataPoint, 0),
			om:  make([]DataPoint, 0),
			hw:  make([]DataPoint, 0),
			c:   make([]DataPoint, 0),
			h:   make([]DataPoint, 0),
			com: make([]DataPoint, 0),
			vom: make([]DataPoint, 0),
			so:  make([]DataPoint, 0),
			o:   make([]DataPoint, 0),
			oh:  make([]DataPoint, 0),
			f:   make([]DataPoint, 0),
			sd:  make([]DataPoint, 0),
			th:  make([]DataPoint, 0),
			st:  make([]DataPoint, 0),
		},
	}
}
