// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-BetaOffsets-r16 (line 28104).
// SL-BetaOffsets-r16 ::=                 INTEGER (0..31)

var sLBetaOffsetsR16Constraints = per.Constrained(0, 31)

type SL_BetaOffsets_r16 struct {
	Value int64
}

func (v *SL_BetaOffsets_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLBetaOffsetsR16Constraints)
}

func (v *SL_BetaOffsets_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLBetaOffsetsR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
