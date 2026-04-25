package main

import (
	"log"
	"net"
)

func main() {
	addr := net.UDPAddr{
		Port: 9000,
		IP: net.ParseIP("127.0.0.1"),
	}

	conn, err := net.DialUDP("udp", nil, &addr)
	if err != nil {
		panic("cannot connect with localhost")
	}
	defer conn.Close()

	_, err = conn.Write([]byte("ana mohamed amine karouach"))
	if err != nil {
		log.Println("cannot write to localhost")
	}
}