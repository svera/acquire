package tile

type Interface interface {
	Number() int
	Letter() string
	SetContent(content TileContent)
	Content() TileContent
}
