package main

import "flag"
import "fmt"
import "github.com/amissine/golib/statemachine"
import "github.com/golang/glog"
import "os"

type StateObj struct{}

func (p *StateObj) PrintHello() (statemachine.StateFn, error) {
	fmt.Printf("%s ", "Hello")
	return p.PrintWorld, nil
}
func (p *StateObj) PrintWorld() (statemachine.StateFn, error) {
	fmt.Println("World")
	return nil, nil
}
func main() {
	so := &StateObj{}
	log := func(s string, i ...interface{}) {
		glog.Infof(s, i...)
	}
	exec := statemachine.New("helloWorld", so.PrintHello, statemachine.LogFacility(log))
	exec.Log(true)
	if len(os.Args) > 1 {
		flag.Parse()
		log("logger", os.Args)
	}

	if e := exec.Execute(); e != nil {
		fmt.Printf("\nERROR: %s\n\n", e)
	}
}
