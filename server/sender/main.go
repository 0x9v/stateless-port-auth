package main

import (
	"log"
	"net"
	"time"
)

func sendUDP(port int, msg string) {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("127.0.0.1"),
	}

	conn, err := net.DialUDP("udp", nil, &addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(msg))
	if err != nil {
		log.Println("error sending:", err)
	}
}

func main() {
	log.Println("Start test sequence...")

	sendUDP(7000, "step 1")
	log.Println("sent 7000")

	time.Sleep(1 * time.Second)

	sendUDP(9000, "step 2")
	log.Println("sent 9000")

	time.Sleep(1 * time.Second)

	sendUDP(8000, "step 3")
	log.Println("sent 8000")

	log.Println("Done")
}
