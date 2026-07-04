// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-Thres-RSRP-r16 (line 28340).
// SL-Thres-RSRP-r16 ::=         INTEGER (0..66)

var sLThresRSRPR16Constraints = per.Constrained(0, 66)

type SL_Thres_RSRP_r16 struct {
	Value int64
}

func (v *SL_Thres_RSRP_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLThresRSRPR16Constraints)
}

func (v *SL_Thres_RSRP_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLThresRSRPR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
