package main

import (
	"encoding/json"
	"os/exec"
)

type speedTestOutput struct {
	Type      string
	ISP       string
	Timestamp string
	Ping      speedTestPing
	Download  speedTestDownload
	Upload    speedTestUpload
	Interface speedTestInterface
	Server    speedTestServer
	Result    speedTestResult
}

type speedTestPing struct {
	Jitter  float32
	Latency float32
}
type speedTestDownload struct {
	Bandwidth int
	Bytes     int
	Elapsed   int
}
type speedTestUpload struct {
	Bandwidth int
	Bytes     int
	Elapsed   int
}
type speedTestInterface struct {
	InternalIP string
	Name       string
	MacAddr    string
	IsVPN      bool
	ExternalIP string
}
type speedTestServer struct {
	ID       int
	Name     string
	Location string
	Country  string
	Host     string
	Port     int
	IP       string
}
type speedTestResult struct {
	ID  string
	URL string
}

func runOoklaSpeedTest() (speedTestOutput, error) {
	res := speedTestOutput{}
	out, err := exec.Command("./speedtest", "--accept-gdpr", "--accept-license", "--format=json").Output()
	if err != nil {
		return res, err
	}
	json.Unmarshal(out, &res)
	return res, nil
}
