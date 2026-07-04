// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LCG-Id-r18 (line 9083).
// LCG-Id-r18 ::= INTEGER (0..maxLCG-ID)

var lCGIdR18Constraints = per.Constrained(0, common.MaxLCG_ID)

type LCG_Id_r18 struct {
	Value int64
}

func (v *LCG_Id_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, lCGIdR18Constraints)
}

func (v *LCG_Id_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(lCGIdR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
