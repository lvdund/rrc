// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReselectionThreshold (line 13993).
// ReselectionThreshold ::=                INTEGER (0..31)

var reselectionThresholdConstraints = per.Constrained(0, 31)

type ReselectionThreshold struct {
	Value int64
}

func (v *ReselectionThreshold) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, reselectionThresholdConstraints)
}

func (v *ReselectionThreshold) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(reselectionThresholdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
