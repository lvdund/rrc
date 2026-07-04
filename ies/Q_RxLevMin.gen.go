// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Q-RxLevMin (line 12769).
// Q-RxLevMin ::=                      INTEGER (-70..-22)

var qRxLevMinConstraints = per.Constrained(-70, -22)

type Q_RxLevMin struct {
	Value int64
}

func (v *Q_RxLevMin) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, qRxLevMinConstraints)
}

func (v *Q_RxLevMin) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(qRxLevMinConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
