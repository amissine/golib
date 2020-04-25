package statemachine

// StateMachine provides a simple state machine for testing the Executor.
type StateMachine struct {
	err       bool
	callTrace []string
}

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
	StateMachine
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
