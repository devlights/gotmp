package signals

type Ready2 struct {
	*ready2
}

type ready2 struct {
	ch chan struct{}
}

func NewReady2() *Ready2 {
	r := new(ready2)
	r.ch = make(chan struct{})

	return &Ready2{r}
}

func (me *Ready2) Signal() {
	close(me.ch)
}

func (me *Ready2) Wait() {
	<-me.ch
}
