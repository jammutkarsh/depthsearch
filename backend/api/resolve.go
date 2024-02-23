package api

import (
	"fmt"
	"log"
	"main/core"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/JammUtkarsh/depth"
)

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
	// non-empty repoURL is required
	if p.RepoURL == "" {
		return fmt.Errorf("repoURL is required")
	}
	// Given URL is a valid URL or not
	if err := isValidURL(p.RepoURL); err != nil {
		return fmt.Errorf("%v is an invalid url", p.RepoURL)
	}
	// Given path is a relative path to the repo
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
