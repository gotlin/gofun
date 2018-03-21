package gorotinepool

type Gpool struct {
	num  int
	goro chan struct{}
}

func NewPool(num int) Gpool {
	return Gpool{num, make(chan struct{}, num)}
}

func (g *Gpool) Get() {
	g.goro <- struct{}{}
}

func (g *Gpool) Free() {
	<-g.goro
}

func (g *Gpool) Empty() bool {
	if 0 == len(g.goro) {
		return true
	}
	return false
}
