// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PRS-ThresRSRP-r18 (line 27699).
// SL-PRS-ThresRSRP-r18 ::=       INTEGER (0..66)

var sLPRSThresRSRPR18Constraints = per.Constrained(0, 66)

type SL_PRS_ThresRSRP_r18 struct {
	Value int64
}

func (v *SL_PRS_ThresRSRP_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLPRSThresRSRPR18Constraints)
}

func (v *SL_PRS_ThresRSRP_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLPRSThresRSRPR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
