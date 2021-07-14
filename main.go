package main

import (
	"fmt"

	"github.com/tsamantanis/port-scanner/port"
)

func main() {
	fmt.Println("Port Scanner")

	open := port.ScanPort("tcp", "localhost", 8080)
	fmt.Printf("%q: %q\n", open.Port, open.State)

	result := port.InitialScan("localhost")
	for _, value := range result {
		fmt.Printf("%q: %q\n", value.Port, value.State)
	}
}
