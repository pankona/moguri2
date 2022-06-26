package main

import (
	"context"

	"github.com/pankona/moguri2/moguri"
)

type globalStateStore struct {
	state map[string]moguri.State
}

func (s *globalStateStore) LoadState(ctx context.Context, characterID string) (moguri.State, error) {
	return s.state[characterID], nil
}

func (s *globalStateStore) UpdateCurrentInteraction(ctx context.Context, characterID string, state moguri.State, interaction moguri.Interacter) error {
	es := state.(*globalState)
	es.currentInteraction = interaction
	s.state[characterID] = es
	return nil
}
