package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  string
	State string
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
		return result
	}

	defer conn.Close()

	result.State = "Open"
	return result
}

func InitialScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 1; i <= 1024; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
		results = append(results, ScanPort("udp", hostname, i))
	}

	return results
}

func ScanAllTCP(hostname string, channelTCP chan []ScanResult) {
	var results []ScanResult

	for i := 1; i <= 49152; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}

	channelTCP <- results
}

func ScanAllUDP(hostname string, channelUDP chan []ScanResult) {
	var results []ScanResult

	for i := 1; i <= 49152; i++ {
		results = append(results, ScanPort("udp", hostname, i))
	}

	channelUDP <- results

}

func ScanAll(hostname string) []ScanResult {
	channelTCP := make(chan []ScanResult)
	channelUDP := make(chan []ScanResult)
	go ScanAllTCP(hostname, channelTCP)
	go ScanAllUDP(hostname, channelUDP)

	var result []ScanResult
	resultTCP := <-channelTCP
	resultUDP := <-channelUDP
	result = append(result, resultTCP...)
	result = append(result, resultUDP...)
	return result
}
