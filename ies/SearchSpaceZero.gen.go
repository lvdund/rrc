// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SearchSpaceZero (line 14549).
// SearchSpaceZero ::=                 INTEGER (0..15)

var searchSpaceZeroConstraints = per.Constrained(0, 15)

type SearchSpaceZero struct {
	Value int64
}

func (v *SearchSpaceZero) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, searchSpaceZeroConstraints)
}

func (v *SearchSpaceZero) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(searchSpaceZeroConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
