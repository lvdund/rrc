// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-CBR-DedicatedSL-PRS-RP-r18 (line 26940).
// SL-CBR-DedicatedSL-PRS-RP-r18 ::= INTEGER (0..100)

var sLCBRDedicatedSLPRSRPR18Constraints = per.Constrained(0, 100)

type SL_CBR_DedicatedSL_PRS_RP_r18 struct {
	Value int64
}

func (v *SL_CBR_DedicatedSL_PRS_RP_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLCBRDedicatedSLPRSRPR18Constraints)
}

func (v *SL_CBR_DedicatedSL_PRS_RP_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLCBRDedicatedSLPRSRPR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
