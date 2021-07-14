package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  string
	State bool
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = false
		return result
	}

	defer conn.Close()

	result.State = true
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

func ScanPorts(hostname string, protocol string, channel chan []ScanResult) {
	var results []ScanResult

	for i := 1; i <= 49152; i++ {
		port := ScanPort(protocol, hostname, i)
		if port.State == true {
			results = append(results, port)
		}
	}

	channel <- results
}

func ScanAll(hostname string) []ScanResult {
	channelTCP := make(chan []ScanResult)
	channelUDP := make(chan []ScanResult)
	go ScanPorts(hostname, "tcp", channelTCP)
	go ScanPorts(hostname, "udp", channelUDP)

	var result []ScanResult
	resultTCP := <-channelTCP
	resultUDP := <-channelUDP
	result = append(result, resultTCP...)
	result = append(result, resultUDP...)
	return result
}
