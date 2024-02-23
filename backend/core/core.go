package core

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/JammUtkarsh/depth"
)

type Summary struct {
	NumInternal int `json:"internal"`
	NumExternal int `json:"external"`
	NumTesting  int `json:"testing"`
}

func HandlePkgs(t *depth.Tree, pkg string) ([]byte, error) {
	err := t.Resolve(pkg)
	if err != nil {
		fmt.Printf("'%v': FATAL: %v\n", pkg, err)
		return nil, err
	}
	res := struct {
		Tree    *depth.Pkg `json:"tree"`
		Summary Summary    `json:"summary"`
	}{
		Tree:    t.Root,
		Summary: WritePkgSummary(*t.Root),
	}
	return ToJSON(res)
}

// writePkgSummary writes a summary of all packages in a tree
func WritePkgSummary(pkg depth.Pkg) Summary {
	var sum Summary
	set := make(map[string]struct{})
	for _, p := range pkg.Deps {
		collectSummary(&sum, p, set)
	}
	return sum
}

func collectSummary(sum *Summary, pkg depth.Pkg, nameSet map[string]struct{}) {
	if _, ok := nameSet[pkg.Name]; !ok {
		nameSet[pkg.Name] = struct{}{}
		if pkg.Internal {
			sum.NumInternal++
		} else {
			sum.NumExternal++
		}
		if pkg.Test {
			sum.NumTesting++
		}
		for _, p := range pkg.Deps {
			collectSummary(sum, p, nameSet)
		}
	}
}

func ToJSON(v any) ([]byte, error) {
	var buf bytes.Buffer
	e := json.NewEncoder(&buf)
	e.SetIndent("", "  ")
	if err := e.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
