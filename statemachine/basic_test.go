package statemachine

import (
	"fmt"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	c1 := &Consumer{Listener: make(chan Event)}
	c1.ConsumeEvents()
	c2 := &Consumer{Listener: make(chan Event)}
	c2.ConsumeEvents()

	e1 := &EventImpl{pipe: c2.Listener}
	e2 := &EventImpl{pipe: c1.Listener, next: e1}
	e1.next = e2

	c1.Listener <- e1
	//<-c1.Listener
	time.Sleep(1 * time.Second)
}
