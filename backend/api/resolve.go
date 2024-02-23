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

var pwd string

func init() {
	pwd, _ = os.Getwd()
}

func resolveRepo(repoURL string, opts Options) ([]byte, error) {
	repoDir := repoURL[strings.LastIndex(repoURL, "/")+1:]
	os.RemoveAll(repoDir)
	if err := exec.Command("git", "clone", repoURL).Run(); err != nil {
		return nil, err
	}
	var t depth.Tree = depth.Tree{
		Root:         &depth.Pkg{},
		OutputStdLib: opts.StdLib,
		Version:      opts.Version,
	}
	log.Printf("Resolving for %s with %+v\n", repoDir, opts)
	// Instead of mentioneding the pacakge that needs to be resolved, we cd into the a specefic directory and then resolve the packages
	if err := os.Chdir(repoDir + string(os.PathSeparator) + opts.Path); err != nil {
		return nil, err
	}
	v, err := core.HandlePkgs(&t, ".")
	fmt.Println(pwd)
	os.Chdir(pwd)
	os.RemoveAll(repoDir)
	return v, err
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
		p.Opts.Path = strings.Join([]string{p.Opts.Path}, string(os.PathSeparator))
	} else {
		p.Opts.Path = "."
	}
	return nil
}

// The function `isValidURL` in Go checks if a given input is a valid URL by verifying the prefix and
// parsing the URL.
func isValidURL(input string) error {
	if !strings.HasPrefix(input, "http") {
		return fmt.Errorf("invalid url")
	}
	if _, err := url.Parse(input); err != nil {
		return err
	}
	return nil
}
