package main

import "github.com/pankona/moguri2/moguri"

type globalState struct {
	CurrentInteraction moguri.Interacter     `json:"current_interaction"`
	CharacterInfo      *moguri.CharacterInfo `json:"character_info"`
}

func (e *globalState) GetCurrentInteraction() (moguri.Interacter, error) {
	return e.CurrentInteraction, nil
}

func (e *globalState) GetCharacterInfo() (*moguri.CharacterInfo, error) {
	return e.CharacterInfo, nil
}
