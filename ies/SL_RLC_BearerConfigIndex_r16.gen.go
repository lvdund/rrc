// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-RLC-BearerConfigIndex-r16 (line 28150).
// SL-RLC-BearerConfigIndex-r16 ::=                    INTEGER (1..maxSL-LCID-r16)

var sLRLCBearerConfigIndexR16Constraints = per.Constrained(1, common.MaxSL_LCID_r16)

type SL_RLC_BearerConfigIndex_r16 struct {
	Value int64
}

func (v *SL_RLC_BearerConfigIndex_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLRLCBearerConfigIndexR16Constraints)
}

func (v *SL_RLC_BearerConfigIndex_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLRLCBearerConfigIndexR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
