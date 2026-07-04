// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DiscardTimerForLowImportance-r18 (line 11293).
// DiscardTimerForLowImportance-r18 ::= ENUMERATED {ms0, ms2, ms4, ms6, ms8, ms10, ms12, ms14, ms18, ms22, ms26, ms30, ms40, ms50, ms75, ms100}

const (
	DiscardTimerForLowImportance_r18_Ms0   = 0
	DiscardTimerForLowImportance_r18_Ms2   = 1
	DiscardTimerForLowImportance_r18_Ms4   = 2
	DiscardTimerForLowImportance_r18_Ms6   = 3
	DiscardTimerForLowImportance_r18_Ms8   = 4
	DiscardTimerForLowImportance_r18_Ms10  = 5
	DiscardTimerForLowImportance_r18_Ms12  = 6
	DiscardTimerForLowImportance_r18_Ms14  = 7
	DiscardTimerForLowImportance_r18_Ms18  = 8
	DiscardTimerForLowImportance_r18_Ms22  = 9
	DiscardTimerForLowImportance_r18_Ms26  = 10
	DiscardTimerForLowImportance_r18_Ms30  = 11
	DiscardTimerForLowImportance_r18_Ms40  = 12
	DiscardTimerForLowImportance_r18_Ms50  = 13
	DiscardTimerForLowImportance_r18_Ms75  = 14
	DiscardTimerForLowImportance_r18_Ms100 = 15
)

var discardTimerForLowImportanceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DiscardTimerForLowImportance_r18_Ms0, DiscardTimerForLowImportance_r18_Ms2, DiscardTimerForLowImportance_r18_Ms4, DiscardTimerForLowImportance_r18_Ms6, DiscardTimerForLowImportance_r18_Ms8, DiscardTimerForLowImportance_r18_Ms10, DiscardTimerForLowImportance_r18_Ms12, DiscardTimerForLowImportance_r18_Ms14, DiscardTimerForLowImportance_r18_Ms18, DiscardTimerForLowImportance_r18_Ms22, DiscardTimerForLowImportance_r18_Ms26, DiscardTimerForLowImportance_r18_Ms30, DiscardTimerForLowImportance_r18_Ms40, DiscardTimerForLowImportance_r18_Ms50, DiscardTimerForLowImportance_r18_Ms75, DiscardTimerForLowImportance_r18_Ms100},
}

type DiscardTimerForLowImportance_r18 struct {
	Value int64
}

func (v *DiscardTimerForLowImportance_r18) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, discardTimerForLowImportanceR18Constraints)
}

func (v *DiscardTimerForLowImportance_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(discardTimerForLowImportanceR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
