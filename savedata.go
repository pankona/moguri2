package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type SaveData struct {
	Structure     Structure     `json:"structure"`
	Progress      Progress      `json:"progress"`
	CharacterInfo CharacterInfo `json:"character_info"`
}

type Structure struct {
	Rooms []Room `json:"room"`
}

type Room struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
	Next     []int    `json:"next"`
}

type Location struct {
	Depth     int `json:"depth"`
	RoomIndex int `json:"room_index"`
}

type Progress struct {
	Location             Location      `json:"location"`
	CurrentInteractionID InteractionID `json:"current_interaction_id"`
	CurrentMessage       string        `json:"current_message"`
}

type CharacterInfo struct {
	CharacterID string `json:"character_id"`
	Name        string `json:"name"`
	HP          int    `json:"hp"`
}

type SaveDataStore struct {
	buf []byte
}

var ErrNilSaveData = errors.New("save data is nil")

func (s *SaveDataStore) Load(ctx context.Context) (*SaveData, error) {
	if s.buf == nil {
		return nil, fmt.Errorf("failed to load save data: %w", ErrNilSaveData)
	}
	ret := &SaveData{}
	if err := json.NewDecoder(bytes.NewBuffer(s.buf)).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *SaveDataStore) Save(ctx context.Context, sd *SaveData) error {
	if sd == nil {
		return fmt.Errorf("failed to save save data: %w", ErrNilSaveData)
	}
	buf, err := json.Marshal(sd)
	if err != nil {
		return err
	}
	s.buf = buf
	return nil
}
