package main

import (
	"heatpump/internal/cli"
	"heatpump/internal/task"
	"io/ioutil"
	"log"
)

func main() {
	cliOptions := cli.Options()

	if !cliOptions.Verbose {
		log.SetOutput(ioutil.Discard)
	}

	task.Start()
}
