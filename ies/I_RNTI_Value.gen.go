// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: I-RNTI-Value (line 8535).
// I-RNTI-Value ::=                        BIT STRING (SIZE(40))

var iRNTIValueConstraints = per.FixedSize(40)

type I_RNTI_Value struct {
	Value per.BitString
}

func (v *I_RNTI_Value) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, iRNTIValueConstraints)
}

func (v *I_RNTI_Value) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(iRNTIValueConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
