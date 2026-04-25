package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := net.UDPAddr{
		Port: 9000,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic("server cannot listen")
	}
	defer conn.Close()

	buffer := make([]byte, 2000)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("cannot read from client")
		}
		fmt.Println("Message :", string(buffer[:n]))
		fmt.Println("Client Address:", clientAddr)
	}

}
