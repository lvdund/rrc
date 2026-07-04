// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ThresholdQ-LR-r19 (line 16145).
// ThresholdQ-LR-r19 ::=               INTEGER (-34..0)

var thresholdQLRR19Constraints = per.Constrained(-34, 0)

type ThresholdQ_LR_r19 struct {
	Value int64
}

func (v *ThresholdQ_LR_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, thresholdQLRR19Constraints)
}

func (v *ThresholdQ_LR_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(thresholdQLRR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
