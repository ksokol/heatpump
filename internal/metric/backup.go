package metric

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const metricsDir = "metrics"
const sep = string(filepath.Separator)

func Backup(metrics *Metric) {
	if err := ensureBackupDir(); err != nil {
		log.Printf("could not create dir metrics due to %v", err)
	} else {
		if bytes, err := ioutil.ReadAll(metrics.Reader()); err == nil {
			fileName := fmt.Sprintf("%s%s%d", metricsDir, sep, metrics.timestamp)
			if err := ioutil.WriteFile(fileName, bytes, 0770); err == nil {
				log.Printf("metrics written to file %v", fileName)
			} else {
				log.Printf("could not write metrics to file %v due to %v", fileName, err)
			}
		} else {
			log.Printf("could not read metrics %v", err)
		}
	}
}

func ReadBackupFiles() []string {
	var files []string

	err := filepath.Walk(metricsDir, func(path string, info os.FileInfo, err error) error {
		if path != metricsDir {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Printf("could not read metric files from %v due to %v", metricsDir, err)
	}

	return files
}

func ensureBackupDir() error {
	if _, err := os.Stat(metricsDir); !os.IsNotExist(err) {
		return err
	}
	return os.Mkdir(metricsDir, 0770)
}
