package main

import (
	"context"
	"fmt"

	"github.com/pankona/moguri2/moguri"
)

const characterID = "character_id"

func main() {
	m := &moguri.Moguri{
		StateStore: &globalStateStore{
			state: map[string]moguri.State{
				characterID: &globalState{
					currentInteraction: &choiceRoomInteraction{},
				},
			},
		},
	}

	ctx := context.Background()

	i, err := m.GetCurrentInteraction(ctx, characterID)
	if err != nil {
		panic(err)
	}
	fmt.Println(i.GetCurrentMessage())
	fmt.Println(i.GetCurrentChoices())

	actionNum := 0
	err = m.Interact(ctx, characterID, actionNum)
	if err != nil {
		panic(err)
	}

	i, err = m.GetCurrentInteraction(ctx, characterID)
	if err != nil {
		panic(err)
	}
	fmt.Println(i.GetCurrentMessage())
	fmt.Println(i.GetCurrentChoices())
}
