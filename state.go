package main

import "github.com/pankona/moguri2/moguri"

type globalState struct {
	currentInteraction moguri.Interacter
}

func (e *globalState) GetCurrentInteraction() (moguri.Interacter, error) {
	return e.currentInteraction, nil
}

func (e *globalState) GetCharacterInfo() (*moguri.CharacterInfo, error) {
	return &moguri.CharacterInfo{}, nil
}
