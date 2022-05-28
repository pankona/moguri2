package main

import "fmt"

type Interacter interface {
	InitialState() any
	Interact(currentState any, action int) (any, error)
}

type SampleRoom struct{}

func (r *SampleRoom) InitialState() map[string]any {
	return map[string]any{
		"id":      0,
		"message": "サンプルの部屋だ",
		"choices": []string{"選択肢1", "選択肢2", "選択肢3"},
	}
}

func (r *SampleRoom) Show(currentState map[string]any) {
	fmt.Println(currentState["message"].(string))
}

func (r *SampleRoom) Interact(currentState map[string]any, action int) (any, error) {
	return nil, nil
}
