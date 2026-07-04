// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ScalingFactorSidelink-r16 (line 17123).
// ScalingFactorSidelink-r16 ::=       ENUMERATED {f0p4, f0p75, f0p8, f1}

const (
	ScalingFactorSidelink_r16_F0p4  = 0
	ScalingFactorSidelink_r16_F0p75 = 1
	ScalingFactorSidelink_r16_F0p8  = 2
	ScalingFactorSidelink_r16_F1    = 3
)

var scalingFactorSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ScalingFactorSidelink_r16_F0p4, ScalingFactorSidelink_r16_F0p75, ScalingFactorSidelink_r16_F0p8, ScalingFactorSidelink_r16_F1},
}

type ScalingFactorSidelink_r16 struct {
	Value int64
}

func (v *ScalingFactorSidelink_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, scalingFactorSidelinkR16Constraints)
}

func (v *ScalingFactorSidelink_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(scalingFactorSidelinkR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
