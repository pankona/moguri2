package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	sdStore := &SaveDataStore{}
	if err := initializeSaveData(sdStore); err != nil {
		panic(err)
	}

	http.HandleFunc("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("pong")) }))
	http.HandleFunc("/current_interaction", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		sd, err := sdStore.Load(r.Context())
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to load save data: %v", err), http.StatusInternalServerError)
			return
		}

		ci, err := getCurrentInteraction(sd)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get current interaction: %v", err), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(ci); err != nil {
			http.Error(w, fmt.Sprintf("failed to encode json body: %v", err), http.StatusInternalServerError)
			return
		}
	}))
	http.HandleFunc("/interact", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		type Request struct {
			Choice string `json:"choice"`
		}

		req := Request{}
		json.NewDecoder(r.Body).Decode(&req)

		_, err := sdStore.Load(r.Context())
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to load save data: %v", err), http.StatusInternalServerError)
			return
		}

		//ci, err := getCurrentInteraction(sd)
		//if err != nil {
		//	http.Error(w, fmt.Sprintf("failed to get current interaction: %v", err), http.StatusInternalServerError)
		//	return
		//}

		//nextInteractionID, err := ci.interact(req.Choice)
		//if err != nil {
		//	http.Error(w, fmt.Sprintf("failed to interact: %v", err), http.StatusInternalServerError)
		//	return
		//}

		//fmt.Println(nextInteractionID)

		// update savedata
	}))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Serve closed: %v", err)
		}
	}
}
