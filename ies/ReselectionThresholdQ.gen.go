// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReselectionThresholdQ (line 13998).
// ReselectionThresholdQ ::=           INTEGER (0..31)

var reselectionThresholdQConstraints = per.Constrained(0, 31)

type ReselectionThresholdQ struct {
	Value int64
}

func (v *ReselectionThresholdQ) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, reselectionThresholdQConstraints)
}

func (v *ReselectionThresholdQ) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(reselectionThresholdQConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
