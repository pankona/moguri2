package main

import (
	"fmt"

	"github.com/pankona/moguri2/moguri"
)

// とりあえずランダムに 3 つほど部屋の候補を返す
func GetRoomInteractionRandom() (moguri.Interacter, error) {
	var roomlist = map[string]func() moguri.Interacter{
		"泉の部屋": NewPondInteraction,
	}

	choices := make([]string, 0, 3)
	// ランダムに取り出す
	for k := range roomlist {
		choices = append(choices, k)
	}
	return &choiceRoomInteraction{
		message: "どの部屋を選ぶ？",
		choices: choices,
		interact: func(state moguri.State, action int) (moguri.Interacter, error) {
			return &choiceRoomInteraction{
				message: fmt.Sprintf("%sの部屋へ向かった。", choices[action]),
				choices: []string{"ok"},
				interact: func(state moguri.State, action int) (moguri.Interacter, error) {
					return roomlist[choices[action]](), nil
				},
			}, nil
		},
	}, nil
}

type choiceRoomInteraction struct {
	message  string
	choices  []string
	interact func(state moguri.State, action int) (moguri.Interacter, error)
}

func NewChoiceRoomInteraction() *choiceRoomInteraction {
	choices := []string{"左の部屋", "中央の部屋", "右の部屋"}
	return &choiceRoomInteraction{
		message: "どの部屋を選ぶ？",
		choices: choices,
		interact: func(state moguri.State, action int) (moguri.Interacter, error) {
			return &choiceRoomInteraction{
				message: fmt.Sprintf("%sの部屋へ向かった。", choices[action]),
				choices: []string{"ok"},
				interact: func(state moguri.State, action int) (moguri.Interacter, error) {
					return NewPondInteraction(), nil
				},
			}, nil
		},
	}
}

func (p *choiceRoomInteraction) GetCurrentMessage() string {
	return p.message
}

func (p *choiceRoomInteraction) GetCurrentChoices() []string {
	return p.choices
}

func (p *choiceRoomInteraction) ValidateInput(state moguri.State, action int) (bool, error) {
	if action < 0 || action >= len(p.GetCurrentChoices()) {
		return false, nil
	}
	return true, nil
}

func (p *choiceRoomInteraction) Interact(state moguri.State, action int) (moguri.Interacter, error) {
	return p.interact(state, action)
}

type pondInteraction struct {
	message  string
	choices  []string
	interact func(state moguri.State, action int) (moguri.Interacter, error)
}

func NewPondInteraction() moguri.Interacter {
	return &pondInteraction{
		message: "泉がある",
		choices: []string{"飲んでみる", "立ち去る"},
		interact: func(state moguri.State, action int) (moguri.Interacter, error) {
			switch action {
			case 0: // 飲んでみる
				return &pondInteraction{
					message: fmt.Sprintf("元気になった。"),
					choices: []string{"ok"},
					interact: func(state moguri.State, action int) (moguri.Interacter, error) {
						return GetRoomInteractionRandom()
					},
				}, nil
			case 1: // 立ち去る
				return &pondInteraction{
					message: fmt.Sprintf("立ち去った。"),
					choices: []string{"ok"},
					interact: func(state moguri.State, action int) (moguri.Interacter, error) {
						return GetRoomInteractionRandom()
					},
				}, nil
			default:
				panic(fmt.Sprintf("invalid choice: %d", action))
			}
		},
	}
}

func (p *pondInteraction) GetCurrentMessage() string {
	return p.message
}

func (p *pondInteraction) GetCurrentChoices() []string {
	return p.choices
}

func (p *pondInteraction) ValidateInput(state moguri.State, action int) (bool, error) {
	if action < 0 || action >= len(p.choices) {
		return false, nil
	}
	return true, nil
}

func (p *pondInteraction) Interact(state moguri.State, action int) (moguri.Interacter, error) {
	return p.interact(state, action)
}
