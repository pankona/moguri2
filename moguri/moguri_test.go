package moguri

import (
	"context"
	"fmt"
)

type exampleState struct {
	currentInteraction Interacter
}

func (e *exampleState) GetCurrentInteraction() (Interacter, error) {
	return e.currentInteraction, nil
}

func (e *exampleState) GetCharacterInfo() (*CharacterInfo, error) {
	return &CharacterInfo{}, nil
}

type exampleInteraction struct {
	message string
	choices []string
}

func (e *exampleInteraction) GetCurrentMessage() string {
	return e.message
}

func (e *exampleInteraction) GetCurrentChoices() []string {
	return e.choices
}

func (e *exampleInteraction) ValidateInput(state State, action int) (bool, error) {
	if action < 0 || action >= len(e.choices) {
		return false, nil
	}
	return true, nil
}

func (e *exampleInteraction) Interact(state State, action int) (Interacter, error) {
	return &exampleInteraction{
		message: fmt.Sprintf("example action: %d", action),
		choices: []string{"ok"},
	}, nil
}

type exampleStateStore struct {
	state map[string]State
}

func (s *exampleStateStore) LoadState(ctx context.Context, characterID string) (State, error) {
	return s.state[characterID], nil
}

func (s *exampleStateStore) UpdateCurrentInteraction(ctx context.Context, characterID string, state State, interaction Interacter) error {
	es := state.(*exampleState)
	es.currentInteraction = interaction
	s.state[characterID] = es
	return nil
}

func ExampleMoguri() {
	m := &Moguri{
		StateStore: &exampleStateStore{
			state: map[string]State{
				"example_character_id": &exampleState{
					// 初期 interaction
					currentInteraction: &exampleInteraction{
						message: "example message",
						choices: []string{"choice 1", "choice 2", "choice 3"},
					},
				},
			},
		},
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
	fmt.Println(i.GetCurrentMessage())
	fmt.Println(i.GetCurrentChoices())

	// Output:
	// example message
	// [choice 1 choice 2 choice 3]
	// example action: 0
	// [ok]
}
