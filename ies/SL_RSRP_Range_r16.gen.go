// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RSRP-Range-r16 (line 28327).
// SL-RSRP-Range-r16 ::=                  INTEGER (0..13)

var sLRSRPRangeR16Constraints = per.Constrained(0, 13)

type SL_RSRP_Range_r16 struct {
	Value int64
}

func (v *SL_RSRP_Range_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLRSRPRangeR16Constraints)
}

func (v *SL_RSRP_Range_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLRSRPRangeR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
