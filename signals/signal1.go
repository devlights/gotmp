package signals

import "sync"

type Ready1 struct {
	*ready1
}

type ready1 struct {
	cond *sync.Cond
	done bool
}

func NewReady1() *Ready1 {
	r := new(ready1)
	r.cond = sync.NewCond(new(sync.Mutex))
	r.done = false

	return &Ready1{r}
}

func (me *Ready1) Signal() {
	me.cond.L.Lock()
	if !me.done {
		me.cond.Broadcast()
	}
	me.cond.L.Unlock()
}

func (me *Ready1) Wait() {
	me.cond.L.Lock()
	for !me.done {
		me.cond.Wait()
		me.done = true
	}
	me.cond.L.Unlock()
}
