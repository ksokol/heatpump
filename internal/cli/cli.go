package cli

import (
	"flag"
	"time"
)

type OptionsData struct {
	HeatpumpIp      string
	Target          string
	Username        string
	Password        string
	Db              string
	CollectInterval time.Duration
	Verbose         bool
}

var heatpumpIp *string
var target *string
var username *string
var password *string
var db *string
var collectInterval *time.Duration
var verbose *bool

func init() {
	heatpumpIp = flag.String("source", "127.0.0.1", "heatpump IPv4")
	target = flag.String("target", "http://localhost", "influxdb base url")
	username = flag.String("username", "heatpump", "influxdb user name")
	password = flag.String("password", "password", "influxdb user password")
	db = flag.String("db", "heatpump_db", "influxdb database name used to persist metrics")
	collectInterval = flag.Duration("interval", 5*time.Second, "data collection interval")
	verbose = flag.Bool("verbose", false, "enable verbose logging")
	flag.Parse()
}

func Options() OptionsData {
	return OptionsData{
		*heatpumpIp,
		*target,
		*username,
		*password,
		*db,
		*collectInterval,
		*verbose,
	}
}
