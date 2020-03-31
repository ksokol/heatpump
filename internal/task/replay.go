package task

import (
	"bytes"
	"heatpump/internal/metric"
	"heatpump/internal/sink"
	"io/ioutil"
	"log"
	"os"
)

func replay() {
	files := metric.ReadBackupFiles()

	for _, file := range files {
		log.Printf("replay metric %s", file)

		if fileBytes, err := ioutil.ReadFile(file); err == nil {
			if err := sink.SendMetric(bytes.NewReader(fileBytes)); err != nil {
				log.Printf("could not send replay file %s due to %v", file, err)
			} else {
				log.Printf("deleting backup file %v", file)
				if err := os.Remove(file); err != nil {
					log.Printf("could not delete backup file %v due to %v", file, err)
				}
			}
		} else {
			log.Printf("could not read file %s due to %v", file, err)
		}
	}
}
