// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NID-r16 (line 10576).
// NID-r16 ::=                      BIT STRING (SIZE (44))

var nIDR16Constraints = per.FixedSize(44)

type NID_r16 struct {
	Value per.BitString
}

func (v *NID_r16) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, nIDR16Constraints)
}

func (v *NID_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(nIDR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
