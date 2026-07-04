// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AvailableSlotOffset-r17 (line 15386).
// AvailableSlotOffset-r17 ::=   INTEGER (0..7)

var availableSlotOffsetR17Constraints = per.Constrained(0, 7)

type AvailableSlotOffset_r17 struct {
	Value int64
}

func (v *AvailableSlotOffset_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, availableSlotOffsetR17Constraints)
}

func (v *AvailableSlotOffset_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(availableSlotOffsetR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
