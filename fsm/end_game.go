package fsm

type EndGame struct {
	ErrorState
}

func (s *EndGame) Name() string {
	return "EndGame"
}
