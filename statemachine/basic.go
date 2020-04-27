package statemachine // {{{1

import (
	"fmt"
	"log"
)

// StateMachine provides a simple state machine for testing the Executor. {{{1
/*
type StateMachine struct {
	err       bool
	callTrace []string
}
*/
// StateMachineBasic has only one state function, Consume, that loops for as long
// as it lives. All it does is consuming requests to invoke functions used
// by another state machine, whose State is being kept in this StateMachineBasic. A
// Request is being consumed by the Consuming channel.
//
// Create and start new StateMachineBasic with:
//
//   c := make(chan Request)
//   &StateMachineBasic{ Consuming: c }.Consume()
//
// To stop and destroy a StateMachineBasic, associated with channel c, run:
//
//   close(c)
//
type StateMachineBasic struct {
	//StateMachine
	State     int
	Consuming chan Request
}

// Request validates and invokes a function used by another state machine. Then
// it updates that machine's state in StateMachineBasic.
type Request interface {
	Validate(*StateMachineBasic)
	Invoke(*StateMachineBasic)
	Update(*StateMachineBasic)
}

// State function Consume reads a Request from the consuming channel.
func (sm *StateMachineBasic) Consume() {
	for r := range sm.Consuming {
		r.Validate(sm)
		r.Invoke(sm)
		r.Update(sm)
	}
}

type Consumer struct { // {{{1
	Listener chan Event
}

type Event interface {
	Act()
}

func (c *Consumer) ConsumeEvents() { // {{{2
	go func() {
		for e := range c.Listener {
			log.Printf("- e %+v\n", e)
			e.Act()
		}
		log.Printf("- c %+v\n", c)
	}()
}

type EventImpl struct { // {{{2
	Next    *EventImpl
	Pipe    chan Event
	visited bool
	Msg     string
}

func (e *EventImpl) Act() {
	if e.visited {
		return
	}
	e.visited = true

	fmt.Printf("%s ", e.Msg)
	e.Pipe <- e.Next
	close(e.Pipe)
} // }}}2

type Exchange struct { // {{{1
	LocalListener  chan OfferOut
	RemoteListener chan OfferIn
}

type OfferOut interface {
	Event
}

type OfferIn interface {
	Event
}

func (e *Exchange) ConsumeOffers() { // {{{2
	go func() {
		for {
			select {
			case oo := <-e.LocalListener:
				oo.Act()
			case oi := <-e.RemoteListener:
				oi.Act()
			}
		}
	}()
}

type Add struct {
	RemoteListener chan OfferIn
	Offer          OfferIn
}

func (add *Add) Act() {
	add.RemoteListener <- add.Offer
}
