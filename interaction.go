package main

import (
	"errors"
	"fmt"
	"strings"
)

type Attraction interface {
	Interactions() []*Interaction
}

type Interaction struct {
	ID       InteractionID `json:"id"`
	Message  string        `json:"message"`
	Choices  []string      `json:"choices"`
	interact func(sd *SaveData, choice string) (Progress, error)
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
			interact: func(sd *SaveData, _ string) (Progress, error) {
				return Progress{
					Location:             sd.Progress.Location,
					CurrentInteractionID: "choice-next",
					CurrentMessage:       "Choice next room.",
				}, nil
			},
		},
	},
}

type Pond struct {
	interactions []*Interaction
}

func (p *Pond) Interactions() []*Interaction {
	return p.interactions
}

var pond = &Pond{
	interactions: []*Interaction{
		{
			ID:      "pond-0",
			Message: "Pond is there",
			Choices: nil,
			interact: func(sd *SaveData, _ string) (Progress, error) {
				return Progress{
					Location:             sd.Progress.Location,
					CurrentInteractionID: "choice-next",
					CurrentMessage:       "Choice next room.",
				}, nil
			},
		},
	},
}

type Veget struct {
	interactions []*Interaction
}

func (v *Veget) Interactions() []*Interaction {
	return v.interactions
}

var veget = &Veget{
	interactions: []*Interaction{
		{
			ID:      "veget-0",
			Message: "Vegetable is there",
			Choices: nil,
			interact: func(sd *SaveData, _ string) (Progress, error) {
				return Progress{
					Location:             sd.Progress.Location,
					CurrentInteractionID: "choice-next",
					CurrentMessage:       "Choice next room.",
				}, nil
			},
		},
	},
}

var attractionByName = map[string]Attraction{
	"entrance": entrance,
	"pond":     pond,
	"veget":    veget,
}

var (
	ErrCurrentInteractionMissing = errors.New("current interaction missing")
	ErrInvalidChoice             = errors.New("invalid choice")
)

func choiceRoomInteractions(sd *SaveData) ([]*Interaction, error) {
	return []*Interaction{
		{
			ID:      "choice-next",
			Message: sd.Progress.CurrentMessage,
			Choices: func() []string {
				var candidates []string
				for _, r := range sd.Structure.Rooms {
					if r.Location.Depth == sd.Progress.Location.Depth+1 {
						candidates = append(candidates, r.Name)
					}
				}
				return candidates
			}(),
			interact: func(sd *SaveData, choice string) (Progress, error) {
				var candidates []Room
				for _, r := range sd.Structure.Rooms {
					if r.Location.Depth == sd.Progress.Location.Depth+1 {
						candidates = append(candidates, r)
					}
				}
				for _, c := range candidates {
					if c.Name == choice {
						// validation OK
						return Progress{
							Location:             c.Location,
							CurrentInteractionID: "choice-next-1",
							CurrentMessage:       fmt.Sprintf("Go to %s", c.Name),
						}, nil
					}
				}
				// validation NG
				return Progress{}, fmt.Errorf("failed to interact: %w", ErrInvalidChoice)
			},
		},
		{
			ID:      "choice-next-1",
			Message: sd.Progress.CurrentMessage,
			Choices: nil,
			interact: func(sd *SaveData, _ string) (Progress, error) {
				var currentAttraction Attraction
				for _, r := range sd.Structure.Rooms {
					if r.Location.Depth == sd.Progress.Location.Depth &&
						r.Location.RoomIndex == sd.Progress.Location.RoomIndex {
						currentAttraction = attractionByName[r.Name]
					}
				}
				if currentAttraction == nil {
					return Progress{}, fmt.Errorf("current attraction is missing: %w", ErrCurrentInteractionMissing)
				}
				i := currentAttraction.Interactions()[0]
				return Progress{
					Location:             sd.Progress.Location,
					CurrentInteractionID: i.ID,
					CurrentMessage:       i.Message,
				}, nil
			},
		},
	}, nil
}

func getCurrentInteraction(sd *SaveData) (*Interaction, error) {
	if strings.HasPrefix(string(sd.Progress.CurrentInteractionID), "choice-next") {
		interactions, err := choiceRoomInteractions(sd)
		if err != nil {
			return nil, err
		}
		var currentInteraction *Interaction
		for _, i := range interactions {
			if sd.Progress.CurrentInteractionID == i.ID {
				currentInteraction = i
			}
		}
		if currentInteraction == nil {
			return nil, fmt.Errorf("current interaction is missing: %w", ErrCurrentInteractionMissing)
		}
		return currentInteraction, nil
	}

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
