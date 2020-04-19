package sink

import (
	"encoding/base64"
	"fmt"
	"heatpump/internal/cli"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

const timeout = 1 * time.Second

var url string
var basicAuth string

func init() {
	cliOptions := cli.Options()

	url = fmt.Sprintf("%s/write?precision=s&db=%s", cliOptions.Target, cliOptions.Db)
	basicAuth = "Basic " + base64.StdEncoding.EncodeToString([]byte(cliOptions.Username+":"+cliOptions.Password))
}

func newHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: timeout,
			}).DialContext,
			ForceAttemptHTTP2:     false,
			MaxIdleConns:          1,
			IdleConnTimeout:       timeout,
			TLSHandshakeTimeout:   timeout,
			ExpectContinueTimeout: timeout,
			DisableKeepAlives:     true,
			MaxConnsPerHost:       1,
		},
		Timeout: timeout,
	}
}

func SendMetric(reader io.Reader) error {
	log.Printf("sending metric to %v", url)

	req, err := http.NewRequest("POST", url, reader)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", basicAuth)

	client := newHttpClient()
	defer client.CloseIdleConnections()

	resp, err := client.Do(req)

	if err == nil && resp.StatusCode != 204 {
		err = fmt.Errorf("%v", resp)
	}

	return err
}
