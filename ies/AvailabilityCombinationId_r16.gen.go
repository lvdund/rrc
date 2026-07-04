// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AvailabilityCombinationId-r16 (line 5071).
// AvailabilityCombinationId-r16 ::=       INTEGER (0..maxNrofAvailabilityCombinationsPerSet-1-r16)

var availabilityCombinationIdR16Constraints = per.Constrained(0, common.MaxNrofAvailabilityCombinationsPerSet_1_r16)

type AvailabilityCombinationId_r16 struct {
	Value int64
}

func (v *AvailabilityCombinationId_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, availabilityCombinationIdR16Constraints)
}

func (v *AvailabilityCombinationId_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(availabilityCombinationIdR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
