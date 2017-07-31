package duktape

import (
	"log"
)

type Dispatcher struct {
	context *Context
	ch      chan func(*Context)
	running bool
}

func CreateDispatcher(context *Context) *Dispatcher {
	d := &Dispatcher{
		context: context,
		ch:      make(chan func(*Context)),
		running: false,
	}
	d.Start()

	return d
}

func (d *Dispatcher) Start() {
	d.running = true
	go func() {
		for {
			select {
			case fn, ok := <-d.ch:
				if !ok || !d.running {
					log.Printf("dispatcher quit")
					break
				}
				fn(d.context)
			}
		}
	}()
}

func (d *Dispatcher) Stop() {
	d.running = false
	close(d.ch)
}
