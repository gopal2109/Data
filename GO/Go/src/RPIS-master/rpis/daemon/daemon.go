package daemon

import (
	//"time"
	"fmt"
)

type DaemonType interface {
	Run()
}

type Daemon struct {
	daemons []DaemonType
}

func (d Daemon) Start() {
	for _, daemon := range d.daemons {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Failed to start daemon")
				fmt.Println(err)
			}
		}()
		go daemon.Run()
	}
}

type ExampleDaemon struct{}

func (d ExampleDaemon) Run() {
	// for {
	// 	fmt.Println("Daemon: Checking DB")
	// 	time.Sleep(5 * time.Second)
	// }
}

type ExampleDaemon1 struct{}

func (d ExampleDaemon1) Run() {
	//fmt.Println("Daemon: Rotate Log")
}

func NewDaemon() Daemon {
	var d []DaemonType
	d = append(d, ExampleDaemon{})
	d = append(d, ExampleDaemon1{})
	return Daemon{daemons: d}
}
