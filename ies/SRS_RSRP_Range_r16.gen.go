// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-RSRP-Range-r16 (line 15826).
// SRS-RSRP-Range-r16 ::=                      INTEGER(0..98)

var sRSRSRPRangeR16Constraints = per.Constrained(0, 98)

type SRS_RSRP_Range_r16 struct {
	Value int64
}

func (v *SRS_RSRP_Range_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sRSRSRPRangeR16Constraints)
}

func (v *SRS_RSRP_Range_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sRSRSRPRangeR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
