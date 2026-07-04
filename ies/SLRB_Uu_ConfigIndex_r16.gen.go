// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SLRB-Uu-ConfigIndex-r16 (line 28382).
// SLRB-Uu-ConfigIndex-r16 ::=                    INTEGER (1..maxNrofSLRB-r16)

var sLRBUuConfigIndexR16Constraints = per.Constrained(1, common.MaxNrofSLRB_r16)

type SLRB_Uu_ConfigIndex_r16 struct {
	Value int64
}

func (v *SLRB_Uu_ConfigIndex_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLRBUuConfigIndexR16Constraints)
}

func (v *SLRB_Uu_ConfigIndex_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLRBUuConfigIndexR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
