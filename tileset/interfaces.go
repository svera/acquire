package tileset

type Interface interface {
	Draw() (Position, error)
}
