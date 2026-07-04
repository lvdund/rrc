// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SearchSpaceId (line 14544).
// SearchSpaceId ::=                   INTEGER (0..maxNrofSearchSpaces-1)

var searchSpaceIdConstraints = per.Constrained(0, common.MaxNrofSearchSpaces_1)

type SearchSpaceId struct {
	Value int64
}

func (v *SearchSpaceId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, searchSpaceIdConstraints)
}

func (v *SearchSpaceId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(searchSpaceIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
