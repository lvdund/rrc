// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UplinkHARQ-mode-r17 (line 12696).
// UplinkHARQ-mode-r17 ::=                 BIT STRING (SIZE (32))

var uplinkHARQModeR17Constraints = per.FixedSize(32)

type UplinkHARQ_Mode_r17 struct {
	Value per.BitString
}

func (v *UplinkHARQ_Mode_r17) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, uplinkHARQModeR17Constraints)
}

func (v *UplinkHARQ_Mode_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(uplinkHARQModeR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
