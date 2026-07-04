// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: P0-PUCCH-Id (line 12255).
// P0-PUCCH-Id ::=                         INTEGER (1..8)

var p0PUCCHIdConstraints = per.Constrained(1, 8)

type P0_PUCCH_Id struct {
	Value int64
}

func (v *P0_PUCCH_Id) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, p0PUCCHIdConstraints)
}

func (v *P0_PUCCH_Id) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(p0PUCCHIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
