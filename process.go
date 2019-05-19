package main

import (
	"fmt"
	"time"
)

// Intraprocess locks
var locks = make(map[int]chan bool)

// Process represents a process in the distributed system.
type Process struct {
	pid       int
	timestamp int
	interval  time.Duration
	inChan    chan int
	outChan   chan int
}

// Start the process
func (p *Process) Start() {
	// Init lock
	locks[p.pid] = make(chan bool)

	go p.runOp(p.recv)
	go p.runOp(p.send)
	go p.scheduleOp(p.event, p.interval)

	locks[p.pid] <- true
}

func (p *Process) runOp(op func()) {
	for {
		// Acquire lock, run operation, put back lock
		<-locks[p.pid]
		op()
		locks[p.pid] <- true
	}
}

func (p *Process) scheduleOp(op func(), interval time.Duration) {
	for {
		// Acquire lock, run operation, put back lock
		<-locks[p.pid]
		op()
		locks[p.pid] <- true

		// Sleep before reacquiring the lock
		time.Sleep(interval)
	}
}

// Operations

// Fire a non-descript event incrementing the timestamp
func (p *Process) event() {
	p.timestamp++
	fmt.Printf("%d: EVENT (timestamp: %d)\n", p.pid, p.timestamp)
}

// Receive on the process' incoming channel
func (p *Process) recv() {
	var recvValue int
	select {
	case recvValue = <-p.inChan:
	default:
		return
	}

	// Update timestamp if the received one was bigger
	if p.timestamp < recvValue {
		p.timestamp = recvValue
		fmt.Printf("%d: updated timestamp to %d\n", p.pid, recvValue)
	}

	p.timestamp++
}

// Send a message on the process' outgoing channels
func (p *Process) send() {
	// Only send on every fourth iteration
	shouldSend := p.timestamp%4 == 0
	if !shouldSend {
		return
	}

	p.timestamp++

	p.outChan <- p.timestamp
	fmt.Printf("%d: sent value %d\n", p.pid, p.timestamp)
}
