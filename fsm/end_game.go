package fsm

type EndGame struct {
	BaseState
}

func (s *EndGame) Name() string {
	return "EndGame"
}
