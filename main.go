package main

import (
	"log"
)

func main() {
	log.Println("Configuring the application")
	conf, err := configure()
	ok(err)

	log.Println("Authorising with Google API")
	client, err := getClient(conf.GoogleAuth)
	ok(err)

	log.Println("Ookla Speedtest running now (this takes about 20 seconds)")
	result, err := runOoklaSpeedTest()
	ok(err)

	log.Println("Uploading results to Google Spreadsheet")
	values := [][]interface{}{{result.Timestamp, result.ISP, result.Ping.Latency, result.Download.Bandwidth, result.Upload.Bandwidth, result.Server.Country, result.Server.Location, result.Server.Host}}
	appendSpreadsheet(conf.GoogleSpreadsheet, client, values)
}

func ok(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
