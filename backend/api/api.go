package api

import (
	"encoding/json"
	"fmt"
	"log"
	"main/core"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/JammUtkarsh/depth"
)

type Options struct {
	Version bool   `json:"version"` // core
	Path    string `json:"path"`    // extra
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

func resolveRepo(repoURL string, opts Options) ([]byte, error) {
	dirName := repoURL[strings.LastIndex(repoURL, "/")+1:]
	os.RemoveAll(dirName)
	if err := exec.Command("git", "clone", repoURL).Run(); err != nil {
		return nil, err
	}
	var t depth.Tree = depth.Tree{
		Root:         &depth.Pkg{},
		OutputStdLib: opts.StdLib,
		Version:      opts.Version,
	}
	log.Printf("Resolving for %s with %+v\n", dirName, opts)
	os.Chdir(dirName)
	if v, err := core.HandlePkgs(&t, opts.Path); err != nil {
		os.RemoveAll(dirName)
		return nil, err
	} else {
		os.Chdir("..")
		os.RemoveAll(dirName)
		return v, nil
	}
}

func (p *Payload) checks() error {
	p.RepoURL = strings.TrimSpace(p.RepoURL)
	if p.RepoURL == "" {
		return fmt.Errorf("repoURL is required")
	}
	if err := isValidURL(p.RepoURL); err != nil {
		return fmt.Errorf("%v is an invalid url", p.RepoURL)
	}
	if p.Opts.Path != "" {
		p.Opts.Path = strings.TrimSpace(p.Opts.Path)
		p.Opts.Path = strings.Join([]string{".", p.Opts.Path}, "/")
	} else {
		p.Opts.Path = "."
	}
	return nil
}

func isValidURL(input string) error {
	if !strings.HasPrefix(input, "http") {
		return fmt.Errorf("invalid url")
	}
	if _, err := url.Parse(input); err != nil {
		return err
	}
	return nil
}
