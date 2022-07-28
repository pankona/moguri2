package main

import (
	"fmt"

	"github.com/pankona/moguri2/moguri"
)

// とりあえずランダムに 3 つほど部屋の候補を返す
func GetRoomInteractionRandom() (moguri.Interacter, error) {
	var roomlist = map[string]func() moguri.Interacter{
		"泉の部屋":  NewPondInteraction,
		"野菜の部屋": NewVegitInteraction,
	}

	choices := make([]string, 0, 3)
	// ランダムに取り出す
	for i := 0; i < cap(choices); i++ {
		for k := range roomlist {
			choices = append(choices, k)
			break
		}
	}

	return &choiceRoomInteraction{
		message: "どの部屋を選ぶ？",
		choices: choices,
		interact: func(state moguri.State, action int) (moguri.State, error) {
			c, err := state.GetCharacterInfo()
			if err != nil {
				return nil, err
			}

			return &globalState{
				characterInfo: c,
				currentInteraction: &choiceRoomInteraction{
					message: fmt.Sprintf("%sへ向かった。", choices[action]),
					choices: []string{"ok"},
					interact: func(state moguri.State, action int) (moguri.State, error) {
						c, err := state.GetCharacterInfo()
						if err != nil {
							return nil, err
						}

						nextInteraction := roomlist[choices[action]]()

						return &globalState{
							characterInfo:      c,
							currentInteraction: nextInteraction,
						}, nil
					},
				},
			}, nil
		},
	}, nil
}

type choiceRoomInteraction struct {
	message  string
	choices  []string
	interact func(state moguri.State, action int) (moguri.State, error)
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

func (p *choiceRoomInteraction) Interact(state moguri.State, action int) (moguri.State, error) {
	return p.interact(state, action)
}

type pondInteraction struct {
	message  string
	choices  []string
	interact func(state moguri.State, action int) (moguri.State, error)
}

func NewPondInteraction() moguri.Interacter {
	return &pondInteraction{
		message: "泉がある",
		choices: []string{"飲んでみる", "立ち去る"},
		interact: func(state moguri.State, action int) (moguri.State, error) {
			c, err := state.GetCharacterInfo()
			if err != nil {
				return nil, err
			}

			switch action {
			case 0: // 飲んでみる
				c.HP -= 5

				return &globalState{
					characterInfo: c,
					currentInteraction: &pondInteraction{
						message: fmt.Sprintf("元気になった。"),
						choices: []string{"ok"},
						interact: func(state moguri.State, action int) (moguri.State, error) {
							c, err := state.GetCharacterInfo()
							if err != nil {
								return nil, err
							}

							nextInteraction, err := GetRoomInteractionRandom()
							if err != nil {
								return nil, err
							}

							return &globalState{
								characterInfo:      c,
								currentInteraction: nextInteraction,
							}, nil
						},
					},
				}, nil
			case 1: // 立ち去る
				return &globalState{
					characterInfo: c,
					currentInteraction: &pondInteraction{
						message: fmt.Sprintf("立ち去った。"),
						choices: []string{"ok"},
						interact: func(state moguri.State, action int) (moguri.State, error) {
							c, err := state.GetCharacterInfo()
							if err != nil {
								return nil, err
							}

							nextInteraction, err := GetRoomInteractionRandom()
							if err != nil {
								return nil, err
							}

							return &globalState{
								characterInfo:      c,
								currentInteraction: nextInteraction,
							}, nil
						},
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

func (p *pondInteraction) Interact(state moguri.State, action int) (moguri.State, error) {
	return p.interact(state, action)
}

type vegitInteraction struct {
	message  string
	choices  []string
	interact func(state moguri.State, action int) (moguri.State, error)
}

func NewVegitInteraction() moguri.Interacter {
	return &vegitInteraction{
		message: "草が生えている",
		choices: []string{"食べる", "立ち去る"},
		interact: func(state moguri.State, action int) (moguri.State, error) {
			c, err := state.GetCharacterInfo()
			if err != nil {
				return nil, err
			}

			switch action {
			case 0: // 食べる
				c.HP -= 5

				return &globalState{
					characterInfo: c,
					currentInteraction: &vegitInteraction{
						message: fmt.Sprintf("元気になった。"),
						choices: []string{"ok"},
						interact: func(state moguri.State, action int) (moguri.State, error) {
							c, err := state.GetCharacterInfo()
							if err != nil {
								return nil, err
							}

							nextInteraction, err := GetRoomInteractionRandom()
							if err != nil {
								return nil, err
							}

							return &globalState{
								characterInfo:      c,
								currentInteraction: nextInteraction,
							}, nil
						},
					},
				}, nil
			case 1: // 立ち去る
				return &globalState{
					characterInfo: c,
					currentInteraction: &vegitInteraction{
						message: fmt.Sprintf("立ち去った。"),
						choices: []string{"ok"},
						interact: func(state moguri.State, action int) (moguri.State, error) {
							c, err := state.GetCharacterInfo()
							if err != nil {
								return nil, err
							}

							nextInteraction, err := GetRoomInteractionRandom()
							if err != nil {
								return nil, err
							}

							return &globalState{
								characterInfo:      c,
								currentInteraction: nextInteraction,
							}, nil
						},
					},
				}, nil
			default:
				panic(fmt.Sprintf("invalid choice: %d", action))
			}
		},
	}
}

func (v *vegitInteraction) GetCurrentMessage() string {
	return v.message
}

func (v *vegitInteraction) GetCurrentChoices() []string {
	return v.choices
}

func (v *vegitInteraction) ValidateInput(state moguri.State, action int) (bool, error) {
	if action < 0 || action >= len(v.choices) {
		return false, nil
	}
	return true, nil
}

func (v *vegitInteraction) Interact(state moguri.State, action int) (moguri.State, error) {
	return v.interact(state, action)
}
