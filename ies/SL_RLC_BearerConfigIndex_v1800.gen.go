// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-RLC-BearerConfigIndex-v1800 (line 28152).
// SL-RLC-BearerConfigIndex-v1800 ::=                  INTEGER (maxSL-LCID-Plus1-r18..maxSL-LCID-r18)

var sLRLCBearerConfigIndexV1800Constraints = per.Constrained(common.MaxSL_LCID_Plus1_r18, common.MaxSL_LCID_r18)

type SL_RLC_BearerConfigIndex_v1800 struct {
	Value int64
}

func (v *SL_RLC_BearerConfigIndex_v1800) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLRLCBearerConfigIndexV1800Constraints)
}

func (v *SL_RLC_BearerConfigIndex_v1800) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLRLCBearerConfigIndexV1800Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
