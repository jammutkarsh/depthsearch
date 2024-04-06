package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

type Options struct {
	Version bool   `json:"version"` // core
	Path    string `json:"path"`    // core
	StdLib  bool   `json:"stdlib"`  // semi-core
	Branch  string `json:"branch"`  // extra
}

type Payload struct {
	RepoURL string  `json:"repoURL" binding:"required"`
	Opts    Options `json:"options"`
}

func Server() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/resolve", resolve)
	handler := cors.Default().Handler(mux)

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		return err
	}
	return nil
}

func resolve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}
	var payload Payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := payload.checks(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if v, err := resolveRepo(payload.RepoURL, payload.Opts); err != nil {
		http.Error(w, "unable to resolve the given path", http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(v)
		return
	}
}

