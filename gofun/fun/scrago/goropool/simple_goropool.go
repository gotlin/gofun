package goropool

type SimpleGoroPool struct {
	PoolSize int
	PoolChan chan struct{}
}

func New(poolSize int) *SimpleGoroPool {
	return &SimpleGoroPool{poolSize, make(chan struct{}, poolSize)}
}

func (p *SimpleGoroPool) Submit(f func()) {

	p.PoolChan <- struct{}{}

	go f()

	<-p.PoolChan

}
