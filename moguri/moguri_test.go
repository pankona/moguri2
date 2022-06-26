package moguri

import (
	"context"
	"fmt"
)

type exampleStateStore struct{}

type exampleState struct {
	currentInteraction Interacter
}

func (e *exampleState) GetCurrentInteraction() (Interacter, error) {
	return &exampleInteraction{}, nil
}

func (e *exampleState) GetCharacterInfo() (*CharacterInfo, error) {
	return &CharacterInfo{}, nil
}

var globalState = map[string]State{
	"example_character_id": &exampleState{
		currentInteraction: &exampleInteraction{},
	},
}

type exampleInteraction struct {
}

func (e *exampleInteraction) GetCurrentMessage() string {
	return "example message"
}

func (e *exampleInteraction) GetCurrentChoices() []string {
	return []string{"choice 1", "choice 2", "choice 3"}
}

func (e *exampleInteraction) ValidateInput(state State, action int) (bool, error) {
	if action < 0 || action > 3 {
		return false, nil
	}
	return true, nil
}
func (e *exampleInteraction) Interact(state State, action int) (State, error) {
	// ここで last action result のための文言を作る
	return state, nil
}

func (e *exampleInteraction) GetLastActionResult() string {
	return "example action result"
}

func (ss *exampleStateStore) LoadState(ctx context.Context, characterID string) (State, error) {
	return globalState[characterID], nil
}

func (ss *exampleStateStore) SaveState(ctx context.Context, characterID string, state State) error {
	globalState[characterID] = state
	return nil
}

func ExampleMoguri() {
	m := &Moguri{
		StateStore: &exampleStateStore{},
	}

	var (
		ctx         = context.Background()
		characterID = "example_character_id"
	)

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

	fmt.Println(i.GetLastActionResult())

	// Output:
	// example message
	// [choice 1 choice 2 choice 3]
	// example action result
}
