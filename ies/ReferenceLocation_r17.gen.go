// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReferenceLocation-r17 (line 13316).
// ReferenceLocation-r17 ::= OCTET STRING

var referenceLocationR17Constraints = per.SizeConstraints{}

type ReferenceLocation_r17 struct {
	Value []byte
}

func (v *ReferenceLocation_r17) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(v.Value, referenceLocationR17Constraints)
}

func (v *ReferenceLocation_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeOctetString(referenceLocationR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
