package main

import (
	"fmt"
	"github.com/amissine/golib/statemachine"
	"time"
)

func Example() {
	fmt.Println("Hello")
	// Output: Hello
}

func ExampleHelloWorld() {

	c1 := &statemachine.Consumer{Listener: make(chan statemachine.Event)} // {{{2
	c1.ConsumeEvents()
	c2 := &statemachine.Consumer{Listener: make(chan statemachine.Event)}
	c2.ConsumeEvents()

	e1 := &statemachine.EventImpl{Pipe: c2.Listener, Msg: "Hello"}
	e2 := &statemachine.EventImpl{Pipe: c1.Listener, Next: e1, Msg: "World"}
	e1.Next = e2

	c1.Listener <- e1
	time.Sleep(1 * time.Second)

	// Output: Hello World
}
