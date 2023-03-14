package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Raw_All(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/raw" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(all_data)
}

func Raw_Artist(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	if !valid_url(id) {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(all_data[id-1])
}
