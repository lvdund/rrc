// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-C-RNTI (line 8300).
// EUTRA-C-RNTI ::=                      BIT STRING (SIZE (16))

var eUTRACRNTIConstraints = per.FixedSize(16)

type EUTRA_C_RNTI struct {
	Value per.BitString
}

func (v *EUTRA_C_RNTI) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, eUTRACRNTIConstraints)
}

func (v *EUTRA_C_RNTI) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(eUTRACRNTIConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
