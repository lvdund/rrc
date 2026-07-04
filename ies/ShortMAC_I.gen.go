// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ShortMAC-I (line 15005).
// ShortMAC-I ::=                      BIT STRING (SIZE (16))

var shortMACIConstraints = per.FixedSize(16)

type ShortMAC_I struct {
	Value per.BitString
}

func (v *ShortMAC_I) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, shortMACIConstraints)
}

func (v *ShortMAC_I) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(shortMACIConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
