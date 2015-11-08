package Board

type Empty struct{}

func (e *Empty) ContentType() string {
	return "empty"
}
