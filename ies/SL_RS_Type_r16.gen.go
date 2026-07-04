// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RS-Type-r16 (line 27934).
// SL-RS-Type-r16 ::=                    ENUMERATED {dmrs, sl-prs, spare2, spare1}

const (
	SL_RS_Type_r16_Dmrs   = 0
	SL_RS_Type_r16_Sl_Prs = 1
	SL_RS_Type_r16_Spare2 = 2
	SL_RS_Type_r16_Spare1 = 3
)

var sLRSTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_RS_Type_r16_Dmrs, SL_RS_Type_r16_Sl_Prs, SL_RS_Type_r16_Spare2, SL_RS_Type_r16_Spare1},
}

type SL_RS_Type_r16 struct {
	Value int64
}

func (v *SL_RS_Type_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, sLRSTypeR16Constraints)
}

func (v *SL_RS_Type_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(sLRSTypeR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
