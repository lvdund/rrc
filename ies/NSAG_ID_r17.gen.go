// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NSAG-ID-r17 (line 10767).
// NSAG-ID-r17 ::=                      BIT STRING (SIZE (8))

var nSAGIDR17Constraints = per.FixedSize(8)

type NSAG_ID_r17 struct {
	Value per.BitString
}

func (v *NSAG_ID_r17) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, nSAGIDR17Constraints)
}

func (v *NSAG_ID_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(nSAGIDR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
