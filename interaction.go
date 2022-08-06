package main

import (
	"errors"
	"fmt"
)

type Attraction interface {
	Interactions() []*Interaction
}

type Interaction struct {
	ID       InteractionID `json:"id"`
	Message  string        `json:"message"`
	Choices  []string      `json:"choices"`
	interact func(choice string) (InteractionID, error)
}

type Entrance struct {
	interactions []*Interaction
}

func (e *Entrance) Interactions() []*Interaction {
	return e.interactions
}

var ErrNoSuchInteraction = errors.New("no such interaction")

func (e *Entrance) IntractionByID(id InteractionID) (*Interaction, error) {
	for _, v := range e.interactions {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, fmt.Errorf("%s is missing in interaction list: %w", id, ErrNoSuchInteraction)
}

type InteractionID string

var entrance = &Entrance{
	interactions: []*Interaction{
		{
			ID:      "entrance-0",
			Message: "Welcome to entrance.",
			Choices: nil,
			interact: func(_ string) (InteractionID, error) {
				return "choice_next", nil
			},
		},
	},
}

type ChoiceNext struct{}

func (c *ChoiceNext) Interactions() []*Interaction {
	return []*Interaction{
		{
			ID:      "choice_next",
			Message: "choice next room.",
			Choices: nil,
			interact: func(choice string) (InteractionID, error) {
				return "", nil
			},
		},
	}
}

var choiceNext = &ChoiceNext{}

var attractionByName = map[string]Attraction{
	"choice_next": choiceNext,
	"entrance":    entrance,
}

var ErrCurrentInteractionMissing = errors.New("current interaction missing")

func getCurrentInteraction(sd *SaveData) (*Interaction, error) {
	var currentAttraction Attraction
	for _, r := range sd.Structure.Rooms {
		if r.Location.Depth == sd.Progress.Location.Depth &&
			r.Location.RoomIndex == sd.Progress.Location.RoomIndex {
			currentAttraction = attractionByName[r.Name]
		}
	}
	if currentAttraction == nil {
		return nil, fmt.Errorf("current attraction is missing: %w", ErrCurrentInteractionMissing)
	}

	var currentInteraction *Interaction
	for _, i := range currentAttraction.Interactions() {
		if InteractionID(sd.Progress.CurrentInteractionID) == i.ID {
			currentInteraction = i
		}
	}
	if currentInteraction == nil {
		return nil, fmt.Errorf("current interaction is missing: %w", ErrCurrentInteractionMissing)
	}

	return currentInteraction, nil
}

func interact(currentInteraction *Interaction, choice string) {
	// send choice
	// save updated current interaction
}
