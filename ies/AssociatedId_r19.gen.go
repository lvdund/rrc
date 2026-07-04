// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AssociatedId-r19 (line 5034).
// AssociatedId-r19 ::=        BIT STRING (SIZE (24))

var associatedIdR19Constraints = per.FixedSize(24)

type AssociatedId_r19 struct {
	Value per.BitString
}

func (v *AssociatedId_r19) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, associatedIdR19Constraints)
}

func (v *AssociatedId_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(associatedIdR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
