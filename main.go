package main

import "flag"
import "fmt"
import "github.com/amissine/golib/statemachine"

//import "github.com/golang/glog"
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
func main() { // if using glog, run with 'go run main.go -logtostderr=true'
	// TODO remove the unrelated stuff:
	mode := flag.String("mode", "sim", "default is 'sim'; set 'trade' if you do not want to simulate")

	so := &StateObj{}
	log := func(s string, i ...interface{}) {
		//glog.Infof(s, i...)
		fmt.Printf(s+"\n", i...)
	}
	exec := statemachine.New("helloworld", so.PrintHello, statemachine.LogFacility(log))
	exec.Log(true)
	//if len(os.Args) > 1 {
	flag.Parse()
	log("logger %v %s\n", os.Args, *mode) // TODO remove the unrelated stuff
	//}

	if e := exec.Execute(); e != nil {
		fmt.Printf("\nERROR: %s\n\n", e)
	}
}
