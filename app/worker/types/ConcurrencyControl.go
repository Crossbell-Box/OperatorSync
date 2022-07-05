package types

type ConcurrencyControl struct {
	ch chan struct{}
}

func NewCtrl(limit int) *ConcurrencyControl {
	return &ConcurrencyControl{ch: make(chan struct{}, limit)}
}
func (cc *ConcurrencyControl) Request() {
	cc.ch <- struct{}{}
}
func (cc *ConcurrencyControl) Done() {
	<-cc.ch
}
