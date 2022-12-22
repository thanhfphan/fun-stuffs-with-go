package main

import (
	"fmt"
	"net"
)

const (
	googleDNSServer = "8.8.8.8:88"
)

func main() {
	conn, err := net.Dial("udp", googleDNSServer)
	if err != nil {
		panic(err)
	}

	addr := conn.LocalAddr()
	if err := conn.Close(); err != nil {
		panic(err)
	}

	udpAddr, ok := addr.(*net.UDPAddr)
	if !ok {
		panic(err)
	}

	fmt.Printf("ip: %s\n", udpAddr.IP.String())
}
