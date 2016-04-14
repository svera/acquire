package interfaces

// Bot defines which methods must be implemented by a bot implementation.
type Bot interface {
	Update(st interface{})
	Play() interface{}
}
