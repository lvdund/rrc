// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DiscardTimerExt2-r17 (line 11283).
// DiscardTimerExt2-r17 ::= ENUMERATED {ms2000, spare3, spare2, spare1}

const (
	DiscardTimerExt2_r17_Ms2000 = 0
	DiscardTimerExt2_r17_Spare3 = 1
	DiscardTimerExt2_r17_Spare2 = 2
	DiscardTimerExt2_r17_Spare1 = 3
)

var discardTimerExt2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DiscardTimerExt2_r17_Ms2000, DiscardTimerExt2_r17_Spare3, DiscardTimerExt2_r17_Spare2, DiscardTimerExt2_r17_Spare1},
}

type DiscardTimerExt2_r17 struct {
	Value int64
}

func (v *DiscardTimerExt2_r17) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, discardTimerExt2R17Constraints)
}

func (v *DiscardTimerExt2_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(discardTimerExt2R17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
