package duktape

type Dispatcher struct {
	context *Context
	ch      chan func(*Context)
	running bool
}

func CreateDispatcher() *Dispatcher {
	d := &Dispatcher{
		ch: make(chan func(*Context)),
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
