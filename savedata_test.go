package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestSaveData(t *testing.T) {
	sd := SaveData{
		Structure: Structure{
			Rooms: []Room{
				{
					Name: "entrance",
					Location: Location{
						Depth:     0,
						RoomIndex: 0,
					},
					Next: []int{0, 1},
				},
				{
					Name: "pond",
					Location: Location{
						Depth:     1,
						RoomIndex: 0,
					},
					Next: []int{0},
				},
				{
					Name: "vegit",
					Location: Location{
						Depth:     1,
						RoomIndex: 1,
					},
					Next: []int{1},
				},
				{
					Name: "pond",
					Location: Location{
						Depth:     2,
						RoomIndex: 0,
					},
					Next: []int{0},
				},
				{
					Name: "vegit",
					Location: Location{
						Depth:     2,
						RoomIndex: 1,
					},
					Next: []int{0},
				},
				{
					Name: "boss",
					Location: Location{
						Depth:     3,
						RoomIndex: 0,
					},
					Next: []int{0},
				},
				{
					Name: "dead_end",
					Location: Location{
						Depth:     4,
						RoomIndex: 0,
					},
					Next: nil,
				},
			},
		},
		Progress: Progress{
			Location: Location{
				Depth:     0,
				RoomIndex: 0,
			},
			CurrentInteractionID: "",
			CurrentMessage:       "",
		},
		CharacterInfo: CharacterInfo{
			CharacterID: "character_id",
			Name:        "character_name",
			HP:          0,
		},
	}

	rawJSON := &bytes.Buffer{}
	if err := json.NewEncoder(rawJSON).Encode(sd); err != nil {
		t.Fatalf("failed to encode JSON: %v", err)
	}

	indentedJSON := &bytes.Buffer{}
	if err := json.Indent(indentedJSON, rawJSON.Bytes(), "", "  "); err != nil {
		t.Fatalf("failed to indent JSON: %v", err)
	}
	t.Log(indentedJSON.String())
}
