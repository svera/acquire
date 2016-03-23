package interfaces

type Bot interface {
	Update(st interface{})
	Play() interface{}
}
