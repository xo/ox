// Package hcl registers the HCL config file format for ox. Import it for its
// side effect:
//
//	import _ "github.com/xo/ox/hcl"
//
// Then a program reads an HCL config file with [ox.ConfigFiles], and a flag
// default reads a key with ${hcl::key}.
package hcl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"maps"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	ctyjson "github.com/zclconf/go-cty/cty/json"

	"github.com/xo/ox"
)

func init() {
	ox.RegisterConfigFormat(ox.HCLT, Decode, "hcl")
}

// Decode decodes HCL data as an [ox.Config]. An attribute is a key, and a
// block nests its body under its type, then under each of its labels. So the
// key of `port` in `server "db" { port = 5432 }` is "server.db.port". An
// expression can use no variable and no function.
func Decode(data []byte) (ox.Config, error) {
	f, diags := hclsyntax.ParseConfig(data, "config.hcl", hcl.InitialPos)
	if diags.HasErrors() {
		return nil, diags
	}
	body, ok := f.Body.(*hclsyntax.Body)
	if !ok {
		return nil, fmt.Errorf("%w: %T is not an HCL body", ox.ErrInvalidType, f.Body)
	}
	m, err := decodeBody(body)
	if err != nil {
		return nil, err
	}
	return ox.MapConfig(m), nil
}

// decodeBody decodes the attributes and the blocks of a body.
func decodeBody(body *hclsyntax.Body) (map[string]any, error) {
	m := make(map[string]any)
	for name, attr := range body.Attributes {
		v, diags := attr.Expr.Value(nil)
		if diags.HasErrors() {
			return nil, diags
		}
		// cty writes JSON that encoding/json reads as the JSON format of ox
		// does, with each number as a json.Number
		buf, err := ctyjson.Marshal(v, v.Type())
		if err != nil {
			return nil, fmt.Errorf("%s: %w", name, err)
		}
		dec := json.NewDecoder(bytes.NewReader(buf))
		dec.UseNumber()
		var val any
		if err := dec.Decode(&val); err != nil {
			return nil, fmt.Errorf("%s: %w", name, err)
		}
		m[name] = val
	}
	for _, block := range body.Blocks {
		sub, err := decodeBody(block.Body)
		if err != nil {
			return nil, err
		}
		parent := m
		keys := append([]string{block.Type}, block.Labels...)
		for _, key := range keys[:len(keys)-1] {
			next, ok := parent[key].(map[string]any)
			if !ok {
				next = make(map[string]any)
				parent[key] = next
			}
			parent = next
		}
		// a second block of the same type and labels adds to the first
		key := keys[len(keys)-1]
		if prev, ok := parent[key].(map[string]any); ok {
			maps.Copy(prev, sub)
			continue
		}
		parent[key] = sub
	}
	return m, nil
}
