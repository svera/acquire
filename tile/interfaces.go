package tile

type TileContent interface {
	Type() string
}

type Interface interface {
	Number() int
	Letter() string
	SetContent(content TileContent)
	Content() TileContent
}
