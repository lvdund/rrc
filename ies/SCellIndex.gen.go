// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SCellIndex (line 14245).
// SCellIndex ::=                      INTEGER (1..31)

var sCellIndexConstraints = per.Constrained(1, 31)

type SCellIndex struct {
	Value int64
}

func (v *SCellIndex) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sCellIndexConstraints)
}

func (v *SCellIndex) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sCellIndexConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
