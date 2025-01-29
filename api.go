package main

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	GITHUBURL = "https://github.com/Abdulrasheed1729/hng12-stage0"
	EMAIL     = "fawomath@gmail.com"
)

type profilehandler struct{}

type Response struct {
	Email           string `json:"email"`
	CurrentDatetime string `json:"current_datetime"`
	GithubUrl       string `json:"github_url"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (p *profilehandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Format the time to ISO 8601 (UTC) format
	currentTime := time.Now().UTC().Format(time.RFC3339)

	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	response := Response{
		Email:           EMAIL,
		CurrentDatetime: currentTime,
		GithubUrl:       GITHUBURL}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Request method not allowed"))
		return
	}

	err := WriteJSON(w, http.StatusOK, response)

	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
