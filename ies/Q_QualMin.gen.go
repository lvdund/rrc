// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Q-QualMin (line 12764).
// Q-QualMin ::=                       INTEGER (-43..-12)

var qQualMinConstraints = per.Constrained(-43, -12)

type Q_QualMin struct {
	Value int64
}

func (v *Q_QualMin) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, qQualMinConstraints)
}

func (v *Q_QualMin) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(qQualMinConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
