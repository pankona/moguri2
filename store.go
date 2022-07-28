package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/pankona/moguri2/moguri"
)

type globalStateStore struct{}

func (s *globalStateStore) LoadState(ctx context.Context, characterID string) (moguri.State, error) {
	f, err := os.Open("moguri_save.json")
	if err != nil {
		return nil, err
	}

	sd := savedata{}
	//ret := &globalState{
	//	CurrentInteraction: &savedata{},
	//	CharacterInfo:      &moguri.CharacterInfo{},
	//}
	if err := json.NewDecoder(f).Decode(sd); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *globalStateStore) UpdateCurrentInteraction(ctx context.Context, characterID string, state moguri.State) error {
	buf, err := json.Marshal(state)
	if err != nil {
		return err
	}
	if err := os.WriteFile("moguri_save.json", buf, 0644); err != nil {
		return err
	}
	return nil
}
