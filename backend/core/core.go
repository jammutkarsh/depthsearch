package core

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/JammUtkarsh/depth"
)

func HandlePkgs(t *depth.Tree, pkg string) ([]byte, error) {
	err := t.Resolve(pkg)
	if err != nil {
		fmt.Printf("'%v': FATAL: %v\n", pkg, err)
		return nil, err
	}
	res := ToJSON(t)
	return res, nil
}

func ToJSON(t *depth.Tree) []byte {
	var buf bytes.Buffer
	e := json.NewEncoder(&buf)
	e.SetIndent("", "  ")
	e.Encode(t.Root)
	return buf.Bytes()
}
