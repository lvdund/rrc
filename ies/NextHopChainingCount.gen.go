// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NextHopChainingCount (line 10540).
// NextHopChainingCount ::=                    INTEGER (0..7)

var nextHopChainingCountConstraints = per.Constrained(0, 7)

type NextHopChainingCount struct {
	Value int64
}

func (v *NextHopChainingCount) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, nextHopChainingCountConstraints)
}

func (v *NextHopChainingCount) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(nextHopChainingCountConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
