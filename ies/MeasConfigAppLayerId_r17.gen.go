// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasConfigAppLayerId-r17 (line 26295).
// MeasConfigAppLayerId-r17 ::= INTEGER (0..maxNrofAppLayerMeas-1-r17)

var measConfigAppLayerIdR17Constraints = per.Constrained(0, common.MaxNrofAppLayerMeas_1_r17)

type MeasConfigAppLayerId_r17 struct {
	Value int64
}

func (v *MeasConfigAppLayerId_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, measConfigAppLayerIdR17Constraints)
}

func (v *MeasConfigAppLayerId_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(measConfigAppLayerIdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
