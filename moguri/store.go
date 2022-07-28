package moguri

import "context"

type Interacter interface {
	ValidateInput(state State, action int) (bool, error)
	Interact(state State, action int) (State, error)
	GetCurrentMessage() string
	GetCurrentChoices() []string
}

type State interface {
	GetCurrentInteraction() (Interacter, error)
	GetCharacterInfo() (*CharacterInfo, error)
}

type StateStore interface {
	LoadState(ctx context.Context, characterID string) (State, error)
	UpdateCurrentInteraction(ctx context.Context, characterID string, state State) error
}
