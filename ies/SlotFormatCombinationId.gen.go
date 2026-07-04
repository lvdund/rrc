// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SlotFormatCombinationId (line 15156).
// SlotFormatCombinationId ::=         INTEGER (0..maxNrofSlotFormatCombinationsPerSet-1)

var slotFormatCombinationIdConstraints = per.Constrained(0, common.MaxNrofSlotFormatCombinationsPerSet_1)

type SlotFormatCombinationId struct {
	Value int64
}

func (v *SlotFormatCombinationId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, slotFormatCombinationIdConstraints)
}

func (v *SlotFormatCombinationId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(slotFormatCombinationIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
