package fsm

type EndGame struct {
	BaseState
}

func (s *EndGame) Type() string {
	return "EndGame"
}
