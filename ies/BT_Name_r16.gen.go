// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BT-Name-r16 (line 26124).
// BT-Name-r16 ::=                    OCTET STRING (SIZE (1..248))

var bTNameR16Constraints = per.SizeRange(1, 248)

type BT_Name_r16 struct {
	Value []byte
}

func (v *BT_Name_r16) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(v.Value, bTNameR16Constraints)
}

func (v *BT_Name_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeOctetString(bTNameR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
