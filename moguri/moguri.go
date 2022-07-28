package moguri

import (
	"context"
	"errors"
)

type Moguri struct {
	StateStore StateStore
}

var ErrValidateInput = errors.New("invalid input")

func (m *Moguri) GetCurrentInteraction(ctx context.Context, characterID string) (Interacter, error) {
	// current state
	cs, err := m.StateStore.LoadState(ctx, characterID)
	if err != nil {
		return nil, err
	}

	i, err := cs.GetCurrentInteraction()
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (m *Moguri) Interact(ctx context.Context, characterID string, action int) error {
	// current state
	cs, err := m.StateStore.LoadState(ctx, characterID)
	if err != nil {
		return err
	}

	i, err := cs.GetCurrentInteraction()
	if err != nil {
		return err
	}

	ok, err := i.ValidateInput(cs, action)
	if err != nil {
		return err
	}
	if !ok {
		return ErrValidateInput
	}

	// next interaction
	ns, err := i.Interact(cs, action)
	if err != nil {
		return err
	}

	err = m.StateStore.UpdateCurrentInteraction(ctx, characterID, ns)
	if err != nil {
		return err
	}

	return nil
}
