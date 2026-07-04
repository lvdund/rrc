// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TDD-UL-DL-SlotIndex (line 16135).
// TDD-UL-DL-SlotIndex ::=             INTEGER (0..maxNrofSlots-1)

var tDDULDLSlotIndexConstraints = per.Constrained(0, common.MaxNrofSlots_1)

type TDD_UL_DL_SlotIndex struct {
	Value int64
}

func (v *TDD_UL_DL_SlotIndex) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, tDDULDLSlotIndexConstraints)
}

func (v *TDD_UL_DL_SlotIndex) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(tDDULDLSlotIndexConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
