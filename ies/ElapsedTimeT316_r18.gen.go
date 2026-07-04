// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ElapsedTimeT316-r18 (line 3562).
// ElapsedTimeT316-r18 ::= INTEGER (0..2000)

var elapsedTimeT316R18Constraints = per.Constrained(0, 2000)

type ElapsedTimeT316_r18 struct {
	Value int64
}

func (v *ElapsedTimeT316_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, elapsedTimeT316R18Constraints)
}

func (v *ElapsedTimeT316_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(elapsedTimeT316R18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
