package main

import (
	"fmt"
	"log"

	"github.com/pankona/moguri2/moguri"
)

var interactions = map[string]func(state moguri.State, action int) (moguri.State, error){
	"choiceRoom":  choiceRoom,
	"choiceRoom1": choiceRoom1,
	"Pond":        Pond,
	"Pond0":       Pond0,
	"Pond1":       Pond1,
	"Vegit":       Vegit,
	"Vegit0":      Vegit0,
	"Vegit1":      Vegit1,
}

var roomlist = map[string]func() moguri.Interacter{
	"泉の部屋":  NewPondInteraction,
	"野菜の部屋": NewVegitInteraction,
}

// とりあえずランダムに 3 つほど部屋の候補を返す
func GetRoomInteractionRandom() (moguri.Interacter, error) {
	choices := make([]string, 0, 3)
	// ランダムに取り出す
	for i := 0; i < cap(choices); i++ {
		for k := range roomlist {
			choices = append(choices, k)
			break
		}
	}

	return &choiceRoomInteraction{
		Message:      "どの部屋を選ぶ？",
		Choices:      choices,
		NextInteract: "choiceRoom",
	}, nil
}

func choiceRoom(state moguri.State, action int) (moguri.State, error) {
	c, err := state.GetCharacterInfo()
	if err != nil {
		return nil, err
	}

	i, err := state.GetCurrentInteraction()
	if err != nil {
		return nil, err
	}

	return &globalState{
		CharacterInfo: c,
		CurrentInteraction: &choiceRoomInteraction{
			Message:      fmt.Sprintf("%sへ向かった。", i.GetCurrentChoices()[action]),
			Choices:      []string{"ok"},
			NextInteract: "choiceRoom1",
			LastChoice:   i.GetCurrentChoices()[action],
		},
	}, nil
}

func choiceRoom1(state moguri.State, action int) (moguri.State, error) {
	c, err := state.GetCharacterInfo()
	if err != nil {
		return nil, err
	}

	ci, err := state.GetCurrentInteraction()
	if err != nil {
		return nil, err
	}

	nextInteraction := roomlist[ci.(*choiceRoomInteraction).LastChoice]()

	return &globalState{
		CharacterInfo:      c,
		CurrentInteraction: nextInteraction,
	}, nil
}

type choiceRoomInteraction struct {
	Message      string   `json:"message"`
	Choices      []string `json:"choices"`
	NextInteract string   `json:"next_interaction"`
	LastChoice   string   `json:"last_choice"`
}

func (c *choiceRoomInteraction) GetCurrentMessage() string {
	return c.Message
}

func (c *choiceRoomInteraction) GetCurrentChoices() []string {
	return c.Choices
}

func (c *choiceRoomInteraction) ValidateInput(state moguri.State, action int) (bool, error) {
	if action < 0 || action >= len(c.GetCurrentChoices()) {
		return false, nil
	}
	return true, nil
}

func (c *choiceRoomInteraction) Interact(state moguri.State, action int) (moguri.State, error) {
	return interactions[c.NextInteract](state, action)
}

type savedata struct {
	Message      string   `json:"message"`
	Choices      []string `json:"choices"`
	NextInteract string   `json:"next_interaction"`
	LastChoice   string   `json:"last_choice"`
}

type pondInteraction struct {
	Message      string   `json:"message"`
	Choices      []string `json:"choices"`
	NextInteract string   `json:"next_interaction"`
	LastChoice   string   `json:"last_choice"`
}

func Pond(state moguri.State, action int) (moguri.State, error) {
	c, err := state.GetCharacterInfo()
	if err != nil {
		return nil, err
	}

	switch action {
	case 0: // 飲んでみる
		c.HP -= 5

		return &globalState{
			CharacterInfo: c,
			CurrentInteraction: &pondInteraction{
				Message:      fmt.Sprintf("元気になった。"),
				Choices:      []string{"ok"},
				NextInteract: "Pond0",
			},
		}, nil
	case 1: // 立ち去る
		return &globalState{
			CharacterInfo: c,
			CurrentInteraction: &pondInteraction{
				Message:      fmt.Sprintf("立ち去った。"),
				Choices:      []string{"ok"},
				NextInteract: "Pond1",
			},
		}, nil
	default:
		panic(fmt.Sprintf("invalid choice: %d", action))
	}
}

func Pond0(state moguri.State, action int) (moguri.State, error) {
	c, err := state.GetCharacterInfo()
	if err != nil {
		return nil, err
	}

	nextInteraction, err := GetRoomInteractionRandom()
	if err != nil {
		return nil, err
	}

	return &globalState{
		CharacterInfo:      c,
		CurrentInteraction: nextInteraction,
	}, nil
}

func Pond1(state moguri.State, action int) (moguri.State, error) {
	c, err := state.GetCharacterInfo()
	if err != nil {
		return nil, err
	}

	nextInteraction, err := GetRoomInteractionRandom()
	if err != nil {
		return nil, err
	}

	return &globalState{
		CharacterInfo:      c,
		CurrentInteraction: nextInteraction,
	}, nil
}

func NewPondInteraction() moguri.Interacter {
	return &pondInteraction{
		Message:      "泉がある",
		Choices:      []string{"飲んでみる", "立ち去る"},
		NextInteract: "Pond",
	}
}

func (p *pondInteraction) GetCurrentMessage() string {
	return p.Message
}

func (p *pondInteraction) GetCurrentChoices() []string {
	return p.Choices
}

func (p *pondInteraction) ValidateInput(state moguri.State, action int) (bool, error) {
	if action < 0 || action >= len(p.Choices) {
		return false, nil
	}
	return true, nil
}

func interact(fn string, state moguri.State, action int) (moguri.State, error) {
	log.Printf("next: %s", fn)
	if _, ok := interactions[fn]; !ok {
		return nil, fmt.Errorf("interaction is missing: %s", fn)
	}
	return interactions[fn](state, action)
}

func (p *pondInteraction) Interact(state moguri.State, action int) (moguri.State, error) {
	return interact(p.NextInteract, state, action)
}

type vegitInteraction struct {
	Message      string   `json:"message"`
	Choices      []string `json:"choices"`
	NextInteract string   `json:"next_interaction"`
	LastChoice   string   `json:"last_choice"`
}

func Vegit(state moguri.State, action int) (moguri.State, error) {
	c, err := state.GetCharacterInfo()
	if err != nil {
		return nil, err
	}

	switch action {
	case 0: // 食べる
		c.HP -= 5

		return &globalState{
			CharacterInfo: c,
			CurrentInteraction: &vegitInteraction{
				Message:      fmt.Sprintf("元気になった。"),
				Choices:      []string{"ok"},
				NextInteract: "Vegit0",
			},
		}, nil
	case 1: // 立ち去る
		return &globalState{
			CharacterInfo: c,
			CurrentInteraction: &vegitInteraction{
				Message:      fmt.Sprintf("立ち去った。"),
				Choices:      []string{"ok"},
				NextInteract: "Vegit1",
			},
		}, nil
	default:
		panic(fmt.Sprintf("invalid choice: %d", action))
	}
}

func Vegit0(state moguri.State, action int) (moguri.State, error) {
	c, err := state.GetCharacterInfo()
	if err != nil {
		return nil, err
	}

	nextInteraction, err := GetRoomInteractionRandom()
	if err != nil {
		return nil, err
	}

	return &globalState{
		CharacterInfo:      c,
		CurrentInteraction: nextInteraction,
	}, nil
}

func Vegit1(state moguri.State, action int) (moguri.State, error) {
	c, err := state.GetCharacterInfo()
	if err != nil {
		return nil, err
	}

	nextInteraction, err := GetRoomInteractionRandom()
	if err != nil {
		return nil, err
	}

	return &globalState{
		CharacterInfo:      c,
		CurrentInteraction: nextInteraction,
	}, nil
}

func NewVegitInteraction() moguri.Interacter {
	return &vegitInteraction{
		Message:      "草が生えている",
		Choices:      []string{"食べる", "立ち去る"},
		NextInteract: "Vegit",
	}
}

func (v *vegitInteraction) GetCurrentMessage() string {
	return v.Message
}

func (v *vegitInteraction) GetCurrentChoices() []string {
	return v.Choices
}

func (v *vegitInteraction) ValidateInput(state moguri.State, action int) (bool, error) {
	if action < 0 || action >= len(v.Choices) {
		return false, nil
	}
	return true, nil
}

func (v *vegitInteraction) Interact(state moguri.State, action int) (moguri.State, error) {
	return interact(v.NextInteract, state, action)
}
