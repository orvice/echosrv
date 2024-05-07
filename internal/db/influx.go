package db

import (
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var (
	influxCli influxdb2.Client
)

func InitInflux() {
	token := os.Getenv("INFLUXDB_TOKEN")
	url := os.Getenv("INFLUXDB_URL")
	influxCli = influxdb2.NewClient(url, token)
}
