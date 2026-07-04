// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: T-Reselection (line 15945).
// T-Reselection ::=                   INTEGER (0..7)

var tReselectionConstraints = per.Constrained(0, 7)

type T_Reselection struct {
	Value int64
}

func (v *T_Reselection) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, tReselectionConstraints)
}

func (v *T_Reselection) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(tReselectionConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
