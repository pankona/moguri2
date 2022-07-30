package main

import (
	"context"
	"errors"
	"fmt"
)

type Attraction interface {
	Interactions() []*Interaction
}

type Interaction struct {
	id       string
	message  string
	choices  []string
	interact func() (string, error)
}

type Entrance struct {
	interactions []*Interaction
}

func (e *Entrance) Interactions() []*Interaction {
	return e.interactions
}

var ErrNoSuchInteraction = errors.New("no such interaction")

func (e *Entrance) IntractionByID(id string) (*Interaction, error) {
	for _, v := range e.interactions {
		if v.id == id {
			return v, nil
		}
	}
	return nil, fmt.Errorf("%s is missing in interaction list: %w", id, ErrNoSuchInteraction)
}

var entrance = &Entrance{
	interactions: []*Interaction{
		{
			id:      "entrance-0",
			message: "Welcome to entrance.",
			interact: func() (string, error) {
				return "entrance-1", nil
			},
		},
		{
			id:      "entrance-1",
			message: "Let's choice next door.",
			interact: func() (string, error) {
				return "choice_next_room", nil
			},
		},
	},
}

var roomByName = map[string]Attraction{
	"entrance": entrance,
}

func loadSaveData(sdStore *SaveDataStore) (*SaveData, error) {
	return sdStore.Load(context.Background())
}

var ErrCurrentInteractionMissing = errors.New("current interaction missing")

func loadCurrentInteraction(sdStore *SaveDataStore) (*Interaction, error) {
	sd, err := loadSaveData(sdStore)
	if err != nil {
		panic(err)
	}

	var currentAttraction Attraction
	for _, r := range sd.Structure.Rooms {
		if r.Location.Depth == sd.Progress.Location.Depth &&
			r.Location.RoomIndex == sd.Progress.Location.RoomIndex {
			currentAttraction = roomByName[r.Name]
		}
	}
	if currentAttraction == nil {
		return nil, fmt.Errorf("current attraction is missing: %w", ErrCurrentInteractionMissing)
	}

	var currentInteraction *Interaction
	for _, i := range currentAttraction.Interactions() {
		if sd.Progress.CurrentInteractionID == i.id {
			currentInteraction = i
		}
	}
	if currentInteraction == nil {
		return nil, fmt.Errorf("current interaction is missing: %w", ErrCurrentInteractionMissing)
	}

	return currentInteraction, nil
}

func main() {
	sdStore := &SaveDataStore{}

	if err := initializeSaveData(sdStore); err != nil {
		panic(err)
	}

	interaction, err := loadCurrentInteraction(sdStore)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", interaction)
}
