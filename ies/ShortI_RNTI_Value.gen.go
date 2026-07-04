// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ShortI-RNTI-Value (line 15000).
// ShortI-RNTI-Value ::=   BIT STRING (SIZE(24))

var shortIRNTIValueConstraints = per.FixedSize(24)

type ShortI_RNTI_Value struct {
	Value per.BitString
}

func (v *ShortI_RNTI_Value) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, shortIRNTIValueConstraints)
}

func (v *ShortI_RNTI_Value) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(shortIRNTIValueConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
