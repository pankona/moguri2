package main

import "context"

func initializeSaveData(sdStore *SaveDataStore) error {
	s, err := GenerateStructure()
	if err != nil {
		return err
	}

	p, err := initializeProgress(s)
	if err != nil {
		return err
	}

	c, err := initializeCharacterInfo()
	if err != nil {
		return err
	}

	return sdStore.Save(context.Background(),
		&SaveData{
			Structure:     *s,
			Progress:      *p,
			CharacterInfo: *c,
		})
}

func GenerateStructure() (*Structure, error) {
	return &Structure{
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
	}, nil
}

func initializeProgress(s *Structure) (*Progress, error) {
	return &Progress{
		Location: Location{
			Depth:     0,
			RoomIndex: 0,
		},
		CurrentInteractionID: "entrance-0",
		CurrentMessage:       "",
	}, nil
}

func initializeCharacterInfo() (*CharacterInfo, error) {
	return &CharacterInfo{
		CharacterID: "character_id",
		Name:        "character_name",
		HP:          10,
	}, nil
}
