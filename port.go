package main

import (
	"fmt"
	"net"
)

// Ask the kernel for a free open port that is ready to use, sourced from the
// freeport utility: https://github.com/phayes/freeport/blob/e7681b37614941bf73b404e0caa37f19e33b5fed/freeport.go
func getPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()

	tcpaddr, ok := l.Addr().(*net.TCPAddr)
	if !ok {
		return 0, fmt.Errorf("Address resolved is not a tcp address")
	}

	return tcpaddr.Port, nil
}
