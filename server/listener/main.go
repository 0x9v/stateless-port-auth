package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
)

type Event struct {
	Port int
}

func startListener(event chan Event, port int) {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal("error on port :", port, ":", err)
	}
	defer conn.Close()

	log.Printf("[%d] start listening", port)
	buffer := make([]byte, 1024)
	for {
		_, _, err = conn.ReadFrom(buffer)
		if err != nil {
			continue
		}
		event <- Event{Port: port}
	}
}

var state = 0

func resetState() {
	state = 0
}

func watchTimeout(ctx context.Context, timeoutChan chan struct{}) {
	<-ctx.Done()
	if ctx.Err() == context.DeadlineExceeded {
		select {
		case timeoutChan <- struct{}{}:
		default:
		}
	}
}

func main() {
	event := make(chan Event, 100)
	timeoutChan := make(chan struct{}, 1)

	go startListener(event, 7000)
	go startListener(event, 9000)
	go startListener(event, 8000)

	ctx, cancel := context.WithCancel(context.Background())

	for {
		select {
		case e := <-event:
			switch state {
			case 0:
				if e.Port == 7000 {
					state = 1
					fmt.Println("STATE 1")
					cancel()
					ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
					go watchTimeout(ctx, timeoutChan)
				}
			case 1:
				if e.Port == 9000 {
					state = 2
					fmt.Println("STATE 2")
				} else {
					resetState()
					cancel()
				}
			case 2:
				if e.Port == 8000 {
					fmt.Println("ACCESS GRANTED")
					resetState()
					cancel()
				}
			}
		case <-timeoutChan:
			fmt.Println("TIMEOUT -> RESET")
			resetState()
		}
	}

}
