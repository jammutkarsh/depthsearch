package api

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		return err
	}
	return nil
}

func resolve(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(v)
		return
	}
}

