package main

import (
	"fmt"

	"github.com/elliotforbes/athena/port"
)

func main() {
	fmt.Println("Port Scanning")
	open := port.ScanPort("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
	open = port.ScanPort("tcp", "localhost", 1314)
	fmt.Printf("Port Open: %t\n", open)
}
