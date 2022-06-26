package main

import (
	"fmt"

	"github.com/pankona/moguri2/moguri"
)

type pondInteraction struct {
	message string
	choices []string
}

func NewPondInteraction() *pondInteraction {
	return &pondInteraction{
		message: "example message",
		choices: []string{"choice 1", "choice 2", "choice 3"},
	}
}
func (e *pondInteraction) GetCurrentMessage() string {
	return e.message
}

func (e *pondInteraction) GetCurrentChoices() []string {
	return e.choices
}

func (e *pondInteraction) ValidateInput(state moguri.State, action int) (bool, error) {
	if action < 0 || action >= len(e.choices) {
		return false, nil
	}
	return true, nil
}

func (e *pondInteraction) Interact(state moguri.State, action int) (moguri.Interacter, error) {
	return &pondInteraction{
		message: fmt.Sprintf("example action: %d", action),
		choices: []string{"ok"},
	}, nil
}
