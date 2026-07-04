// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AbsoluteTimeInfo-r16 (line 25976).
// AbsoluteTimeInfo-r16 ::= BIT STRING (SIZE (48))

var absoluteTimeInfoR16Constraints = per.FixedSize(48)

type AbsoluteTimeInfo_r16 struct {
	Value per.BitString
}

func (v *AbsoluteTimeInfo_r16) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, absoluteTimeInfoR16Constraints)
}

func (v *AbsoluteTimeInfo_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(absoluteTimeInfoR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
