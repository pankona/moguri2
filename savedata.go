package main

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
	Location             Location `json:"location"`
	CurrentInteractionID string   `json:"current_interaction_id"`
	CurrentMessage       string   `json:"current_message"`
}

type CharacterInfo struct {
	HP int `json:"hp"`
}
