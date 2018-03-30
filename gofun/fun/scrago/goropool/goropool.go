package goropool

type GoroPool interface {
	Submit(f func())
}
