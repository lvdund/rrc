// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UPInterruptionTimeAtHO-r17 (line 3560).
// UPInterruptionTimeAtHO-r17 ::= INTEGER (0..1023)

var uPInterruptionTimeAtHOR17Constraints = per.Constrained(0, 1023)

type UPInterruptionTimeAtHO_r17 struct {
	Value int64
}

func (v *UPInterruptionTimeAtHO_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, uPInterruptionTimeAtHOR17Constraints)
}

func (v *UPInterruptionTimeAtHO_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(uPInterruptionTimeAtHOR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
