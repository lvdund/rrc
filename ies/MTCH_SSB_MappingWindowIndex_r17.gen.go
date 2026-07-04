// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MTCH-SSB-MappingWindowIndex-r17 (line 28520).
// MTCH-SSB-MappingWindowIndex-r17  ::= INTEGER (0..maxNrofMTCH-SSB-MappingWindow-1-r17)

var mTCHSSBMappingWindowIndexR17Constraints = per.Constrained(0, common.MaxNrofMTCH_SSB_MappingWindow_1_r17)

type MTCH_SSB_MappingWindowIndex_r17 struct {
	Value int64
}

func (v *MTCH_SSB_MappingWindowIndex_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, mTCHSSBMappingWindowIndexR17Constraints)
}

func (v *MTCH_SSB_MappingWindowIndex_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(mTCHSSBMappingWindowIndexR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
