package main // {{{1

//import "flag"
import "fmt"
import "github.com/amissine/golib/statemachine"
import "time"

//import "github.com/golang/glog"
//import "os"

type StateObj struct{}

func (p *StateObj) PrintHello() (statemachine.StateFn, error) {
	fmt.Printf("%s ", "Hello")
	return p.PrintWorld, nil
}
func (p *StateObj) PrintWorld() (statemachine.StateFn, error) {
	fmt.Println("World")
	return nil, nil
}
func main() { // if using glog, run with 'go run main.go -logtostderr=true' {{{1
	/* {{{2
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
	*/
	c1 := &statemachine.Consumer{Listener: make(chan statemachine.Event)} // {{{2
	c1.ConsumeEvents()
	c2 := &statemachine.Consumer{Listener: make(chan statemachine.Event)}
	c2.ConsumeEvents()

	e1 := &statemachine.EventImpl{Pipe: c2.Listener, Msg: "Hello"}
	e2 := &statemachine.EventImpl{Pipe: c1.Listener, Next: e1, Msg: "World"}
	e1.Next = e2

	c1.Listener <- e1
	time.Sleep(1 * time.Second)
	/*
		alik@mba ~/go/src/github.com/amissine/golib (master) $ go run main.go 2>/dev/null
		Hello World alik@mba ~/go/src/github.com/amissine/golib (master) $
		alik@mba ~/go/src/github.com/amissine/golib (master) $ go run main.go >/dev/null
		2020/04/26 17:44:00 - e &{Next:0xc0000841b0 Pipe:0xc0000800c0 visited:false Msg:Hello}
		2020/04/26 17:44:00 - e &{Next:0xc000084180 Pipe:0xc000080060 visited:false Msg:World}
		2020/04/26 17:44:00 - c &{Listener:0xc0000800c0}
		2020/04/26 17:44:00 - e &{Next:0xc0000841b0 Pipe:0xc0000800c0 visited:true Msg:Hello}
		2020/04/26 17:44:00 - c &{Listener:0xc000080060}
		alik@mba ~/go/src/github.com/amissine/golib (master) $
	}}}2 */
}
