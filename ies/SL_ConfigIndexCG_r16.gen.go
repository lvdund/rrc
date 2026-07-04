// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-ConfigIndexCG-r16 (line 27077).
// SL-ConfigIndexCG-r16 ::=          INTEGER (0..maxNrofCG-SL-1-r16)

var sLConfigIndexCGR16Constraints = per.Constrained(0, common.MaxNrofCG_SL_1_r16)

type SL_ConfigIndexCG_r16 struct {
	Value int64
}

func (v *SL_ConfigIndexCG_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLConfigIndexCGR16Constraints)
}

func (v *SL_ConfigIndexCG_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLConfigIndexCGR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
