package main

import (
	"fmt"

	"github.com/tsamantanis/port"
)

func main() {
	fmt.Println("Port Scanner")

	open := port.ScanPort("tcp", "localhost", 8080)
	fmt.Printf("Port Open: %t\n", open)

}
