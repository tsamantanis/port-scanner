package main

import (
	"fmt"

	"github.com/tsamantanis/port-scanner/port"
)

func main() {
	fmt.Println("Port Scanner")

	open := port.ScanPort("tcp", "localhost", 8080)
	fmt.Printf("%q: %t\n", open.Port, open.State)

	result := port.ScanAll("localhost")
	PrintResults(result)
}

func PrintResults(res []port.ScanResult) {
	fmt.Println("List of open ports")
	for _, value := range res {
		fmt.Printf("%q: %t\n", value.Port, value.State)
	}
}
