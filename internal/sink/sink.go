package sink

import (
	"encoding/base64"
	"fmt"
	"heatpump/internal/cli"
	"io"
	"log"
	"net/http"
	"time"
)

var url string
var basicAuth string

func init() {
	cliOptions := cli.Options()

	url = fmt.Sprintf("%s/write?precision=s&db=%s", cliOptions.Target, cliOptions.Db)
	basicAuth = "Basic " + base64.StdEncoding.EncodeToString([]byte(cliOptions.Username+":"+cliOptions.Password))
}

func SendMetric(reader io.Reader) error {
	log.Printf("sending metric to %v", url)

	req, err := http.NewRequest("POST", url, reader)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", basicAuth)

	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		client.CloseIdleConnections()
	}

	if err == nil && resp.StatusCode != 204 {
		err = fmt.Errorf("%v", resp)
	}

	return err
}
