package main

import (
	"context"
	"errors"
	"fmt"
)

var _ Interactor = &SampleRoom{}

type SampleRoom struct{}

func (r *SampleRoom) InitialState() State {
	return State{
		"id":      0,
		"message": "サンプルの部屋だ",
		"choices": []string{"選択肢1", "選択肢2", "選択肢3"},
	}
}

func (r *SampleRoom) Show(currentState State) {
	fmt.Println(currentState["message"].(string))
}

func (r *SampleRoom) Interact(currentState State, action int) (State, error) {
	return nil, nil
}

func (r *SampleRoom) HasNext(state State) bool {
	return true
}

func (r *SampleRoom) ValidateInput(state State, action int) (bool, error) {
	return true, nil
}

// interaction のドメインロジック

type State map[string]any

type Interactor interface {
	InitialState() State
	Show(state State)
	Interact(state State, action int) (State, error)
	HasNext(State) bool
	ValidateInput(state State, action int) (bool, error)
	GetPlayerChoice() (int, error)
}

func Do(interactor Interactor) {
	var err error
	state := interactor.InitialState()

	for interactor.HasNext(state) {
		interactor.Show(state)

		// 適切な入力をしてもらえるまでループする
		var action int
		for {
			// プレイヤーにアクションを選んでもらう
			action, err = interactor.GetPlayerChoice()
			if err != nil {
				panic(err)
			}
			ok, err := interactor.ValidateInput(state, action)
			if err != nil {
				panic(err)
			}
			if ok {
				break
			}
		}

		state, err = interactor.Interact(state, action)
		if err != nil {
			panic(err)
		}
	}
}

type InteractService struct {
	loadState     func(ctx context.Context, characterID string) (any, error)
	saveState     func(ctx context.Context, characterID string, state any) error
	validateInput func(ctx context.Context, state any, action int) (bool, error)
	interact      func(ctx context.Context, state any, action int) (any, error)
}

var ErrInvalidInput = errors.New("invalid input")

func (i *InteractService) CurrentState(ctx context.Context, characterID string) (any, error) {
	return i.loadState(ctx, characterID)
}

func (i *InteractService) Interact(ctx context.Context, characterID string, action int) error {
	state, err := i.loadState(ctx, characterID)
	if err != nil {
		return err
	}

	ok, err := i.validateInput(ctx, state, action)
	if err != nil {
		return err
	}
	if !ok {
		return ErrInvalidInput
	}

	nextState, err := i.interact(ctx, state, action)
	if err != nil {
		return err
	}

	if err := i.saveState(ctx, characterID, nextState); err != nil {
		return err
	}

	return nil
}

type interact struct {
	InteractService *InteractService
}
