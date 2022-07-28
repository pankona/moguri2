package main

import "github.com/pankona/moguri2/moguri"

type globalState struct {
	currentInteraction moguri.Interacter
	characterInfo      *moguri.CharacterInfo
}

func (e *globalState) GetCurrentInteraction() (moguri.Interacter, error) {
	return e.currentInteraction, nil
}

func (e *globalState) GetCharacterInfo() (*moguri.CharacterInfo, error) {
	return e.characterInfo, nil
}
