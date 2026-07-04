// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-DestinationIdentity-r16 (line 27111).
// SL-DestinationIdentity-r16 ::=           BIT STRING (SIZE (24))

var sLDestinationIdentityR16Constraints = per.FixedSize(24)

type SL_DestinationIdentity_r16 struct {
	Value per.BitString
}

func (v *SL_DestinationIdentity_r16) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, sLDestinationIdentityR16Constraints)
}

func (v *SL_DestinationIdentity_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(sLDestinationIdentityR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
