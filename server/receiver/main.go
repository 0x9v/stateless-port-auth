package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Event struct {
	Port int
}

func startListener(events chan Event, port int) {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal("error on port :", port, ":", err)
	}
	defer conn.Close()

	buffer := make([]byte, 2000)
	for {
		_, _, err = conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("cannot reading :", err)
			continue
		}

		events <- Event{Port: port}
	}
}

func main() {
	state := 0
	events := make(chan Event, 100)

	go startListener(events, 7000)
	go startListener(events, 9000)
	go startListener(events, 8000)

	timer := time.NewTimer(0)
	timer.Stop()
	resetTimer := func() {
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
		timer.Reset(5 * time.Second)
	}

	resetAll := func() {
		state = 0
		timer.Stop()
		<-timer.C
	}

	for {
		select {
		case e := <-events:
			switch state {
			case 0:
				if e.Port == 7000 {
					state = 1
					fmt.Println("STATE 1")
					resetTimer()
				}
			case 1:
				if e.Port == 9000 {
					state = 2
					fmt.Println("STATE 2")
				} else {
					resetAll()
				}
			case 2:
				if e.Port == 8000 {
					fmt.Println("ACTION SUCCESS")
				}
				resetAll()
			}
		case <-timer.C:
			resetAll()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
