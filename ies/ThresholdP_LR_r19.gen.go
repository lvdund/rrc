// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ThresholdP-LR-r19 (line 16140).
// ThresholdP-LR-r19 ::=               INTEGER (-80..0)

var thresholdPLRR19Constraints = per.Constrained(-80, 0)

type ThresholdP_LR_r19 struct {
	Value int64
}

func (v *ThresholdP_LR_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, thresholdPLRR19Constraints)
}

func (v *ThresholdP_LR_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(thresholdPLRR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
