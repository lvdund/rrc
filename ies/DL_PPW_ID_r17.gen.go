// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DL-PPW-ID-r17 (line 7651).
// DL-PPW-ID-r17 ::= INTEGER  (0..maxNrofPPW-ID-1-r17)

var dLPPWIDR17Constraints = per.Constrained(0, common.MaxNrofPPW_ID_1_r17)

type DL_PPW_ID_r17 struct {
	Value int64
}

func (v *DL_PPW_ID_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, dLPPWIDR17Constraints)
}

func (v *DL_PPW_ID_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(dLPPWIDR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
