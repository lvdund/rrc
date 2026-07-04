// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: GapPriority-r17 (line 8472).
// GapPriority-r17 ::=                       INTEGER (1..maxNrOfGapPri-r17)

var gapPriorityR17Constraints = per.Constrained(1, common.MaxNrOfGapPri_r17)

type GapPriority_r17 struct {
	Value int64
}

func (v *GapPriority_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, gapPriorityR17Constraints)
}

func (v *GapPriority_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(gapPriorityR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
