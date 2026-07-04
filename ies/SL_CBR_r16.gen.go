// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-CBR-r16 (line 26920).
// SL-CBR-r16 ::=                        INTEGER (0..100)

var sLCBRR16Constraints = per.Constrained(0, 100)

type SL_CBR_r16 struct {
	Value int64
}

func (v *SL_CBR_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLCBRR16Constraints)
}

func (v *SL_CBR_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLCBRR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
