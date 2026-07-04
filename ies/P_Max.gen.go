// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: P-Max (line 10918).
// P-Max ::=                           INTEGER (-30..33)

var pMaxConstraints = per.Constrained(-30, 33)

type P_Max struct {
	Value int64
}

func (v *P_Max) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pMaxConstraints)
}

func (v *P_Max) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pMaxConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
