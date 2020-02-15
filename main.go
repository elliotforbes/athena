package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	fmt.Println("Port Scanning")
	open := scanPort("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
	open = scanPort("tcp", "localhost", 1314)
	fmt.Printf("Port Open: %t\n", open)
}
