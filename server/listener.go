package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

type Event struct {
	Port Port
}

// startListener listens on a UDP port and sends an event whenever data is received
func startListener(event chan Event, port Port) {
	addr := net.UDPAddr{
		Port: int(port),
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal("error on port :", port, ":", err)
	}
	defer conn.Close()

	log.Printf("[%d] start listening", int(port))
	buffer := make([]byte, 1024)
	for {
		_, _, err = conn.ReadFrom(buffer)
		if err != nil {
			continue
		}

		// this is how we identify the current port
		// by sending it through the channel
		event <- Event{Port: port}
	}
}

type Port int

type StateMachine struct {
	Sequence []Port
	Index    int
}

// Reset sets the sequence index back to the initial position (Index = 0)
func (sm *StateMachine) Reset() {
	sm.Index = 0
}

func (sm *StateMachine) Next(p Port) bool {
	if sm.Index < len(sm.Sequence) && p == sm.Sequence[sm.Index] {
		sm.Index++
		return true
	}

	sm.Reset()
	return false
}

// Complete returns true if the sequence is completed
func (sm *StateMachine) Complete() bool {
	return sm.Index == len(sm.Sequence)
}

func main() {
	event := make(chan Event, 100)

	// channel used to signal when the timer has expired
	// Do not communicate by sharing memory; instead, share memory by communicating
	// this is the most important thing i learned during this project
	timeoutSignal := make(chan struct{}, 1)

	// port sequence
	sequence := []Port{7000, 9000, 8000}

	// start listening to the ports
	for _, port := range sequence {
		go startListener(event, port)
	}

	sm := &StateMachine{Sequence: sequence}

	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	for {
		select {
		case e := <-event:
			if sm.Index == 0 {
				if cancel != nil {
					cancel()
				}

				// initialize the timeout context
				ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

				// wait for the context to expire and emit a timeout signal
				go func(c context.Context) {
					<-ctx.Done()
					if errors.Is(ctx.Err(), context.DeadlineExceeded) {
						timeoutSignal <- struct{}{}
					}
				}(ctx)
			}

			if !sm.Next(e.Port) {
				fmt.Println("RESET (wrong sequence)")
				if cancel != nil {
					cancel()
				}
				continue
			}

			fmt.Println("STEP:", sm.Index)

			if sm.Complete() {
				fmt.Println("ACCESS GRANTED")
				sm.Reset()
				if cancel != nil {
					cancel()
				}
			}
		case <-timeoutSignal:
			fmt.Println("TIMEOUT -> RESET")
			sm.Reset()
			if cancel != nil {
				cancel()
			}
		}
	}
}
