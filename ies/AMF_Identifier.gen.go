// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AMF-Identifier (line 4988).
// AMF-Identifier ::=                      BIT STRING (SIZE (24))

var aMFIdentifierConstraints = per.FixedSize(24)

type AMF_Identifier struct {
	Value per.BitString
}

func (v *AMF_Identifier) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, aMFIdentifierConstraints)
}

func (v *AMF_Identifier) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(aMFIdentifierConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
