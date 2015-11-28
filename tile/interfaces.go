package tile

type Owner interface {
	Type() string
}

type Interface interface {
	Number() int
	Letter() string
	SetOwner(owner Owner) Interface
	Owner() Owner
}
