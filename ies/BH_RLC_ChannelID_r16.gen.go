// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BH-RLC-ChannelID-r16 (line 5213).
// BH-RLC-ChannelID-r16 ::=    BIT STRING (SIZE (16))

var bHRLCChannelIDR16Constraints = per.FixedSize(16)

type BH_RLC_ChannelID_r16 struct {
	Value per.BitString
}

func (v *BH_RLC_ChannelID_r16) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, bHRLCChannelIDR16Constraints)
}

func (v *BH_RLC_ChannelID_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(bHRLCChannelIDR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
