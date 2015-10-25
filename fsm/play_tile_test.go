package fsm

import "testing"

func TestTransitionToPlayTile(t *testing.T) {
	state := &PlayTile{}
	if _, err := state.ToPlayTile(); err == nil {
		t.Errorf("Transition from PlayTile must return an error, returned nil")
	}
}

func TestTransitionToFoundCorp(t *testing.T) {
	state := &PlayTile{}
	_, ok := state.(*FoundCorp)
	if state, err := state.ToFoundCorp(); ok {
		t.Errorf("Transition from PlayTile must be of type FoundCorp")
	}
}
