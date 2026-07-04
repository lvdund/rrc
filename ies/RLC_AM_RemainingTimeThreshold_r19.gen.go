// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RLC-AM-RemainingTimeThreshold-r19 (line 11295).
// RLC-AM-RemainingTimeThreshold-r19 ::= INTEGER (1..64)

var rLCAMRemainingTimeThresholdR19Constraints = per.Constrained(1, 64)

type RLC_AM_RemainingTimeThreshold_r19 struct {
	Value int64
}

func (v *RLC_AM_RemainingTimeThreshold_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, rLCAMRemainingTimeThresholdR19Constraints)
}

func (v *RLC_AM_RemainingTimeThreshold_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(rLCAMRemainingTimeThresholdR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
