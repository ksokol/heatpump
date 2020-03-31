package sink

import (
	"fmt"
	"heatpump/internal/cli"
	"io"
	"log"
	"net/http"
)

const contentType = "application/x-www-form-urlencoded"

var url string

func init() {
	cliOptions := cli.Options()

	url = fmt.Sprintf("%s/write?precision=s&db=%s", cliOptions.Target, cliOptions.Db)
}

func SendMetric(reader io.Reader) error {
	log.Printf("sending metric to %v", url)
	resp, err := http.Post(url, contentType, reader)

	if err == nil && resp.StatusCode != 204 {
		err = fmt.Errorf("%v", resp)
	}

	return err
}
