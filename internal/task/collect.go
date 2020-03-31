package task

import (
	"heatpump/internal/heatpump"
	"heatpump/internal/metric"
	"heatpump/internal/sink"
	"log"
)

func collect() {
	data, err := heatpump.ReadData()
	if err == nil {
		metrics := metric.Parse(data)

		if err := sink.SendMetric(metrics.Reader()); err != nil {
			log.Printf("could not send metric - %v", err)
			metric.Backup(metrics)
		}
	} else {
		log.Printf("could not read heatpump metrics %v", err)
	}
}
