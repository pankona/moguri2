package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pankona/moguri2/moguri"
)

const characterID = "character_id"

func main() {
	m := &moguri.Moguri{
		StateStore: &globalStateStore{
			state: map[string]moguri.State{
				characterID: &globalState{
					currentInteraction: NewChoiceRoomInteraction(),
				},
			},
		},
	}

	r := mux.NewRouter()

	r.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			Pong string `json:"pong"`
		}{
			Pong: "Hello!",
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("error: %v", err)
		}
	})).Methods(http.MethodGet)

	r.Handle("/current_interaction", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		i, err := m.GetCurrentInteraction(r.Context(), characterID)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %v", err), http.StatusInternalServerError)
			return
		}

		resp := struct {
			Message string   `json:"message"`
			Choices []string `json:"choices"`
		}{
			Message: i.GetCurrentMessage(),
			Choices: i.GetCurrentChoices(),
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("error: %v", err)
		}
	})).Methods(http.MethodGet)

	r.Handle("/interact", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := struct {
			ActionNum int `json:"action_num"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("error: %v", err), http.StatusInternalServerError)
			return
		}

		if err := m.Interact(r.Context(), characterID, req.ActionNum); err != nil {
			http.Error(w, fmt.Sprintf("error: %v", err), http.StatusInternalServerError)
			return
		}
	})).Methods(http.MethodPost)

	if err := http.ListenAndServe(":3000", r); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Printf("error occurred: %v", err)
		}
	}
}
