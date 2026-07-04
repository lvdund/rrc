// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DiscardTimerExt-r16 (line 11281).
// DiscardTimerExt-r16 ::= ENUMERATED {ms0dot5, ms1, ms2, ms4, ms6, ms8, spare2, spare1}

const (
	DiscardTimerExt_r16_Ms0dot5 = 0
	DiscardTimerExt_r16_Ms1     = 1
	DiscardTimerExt_r16_Ms2     = 2
	DiscardTimerExt_r16_Ms4     = 3
	DiscardTimerExt_r16_Ms6     = 4
	DiscardTimerExt_r16_Ms8     = 5
	DiscardTimerExt_r16_Spare2  = 6
	DiscardTimerExt_r16_Spare1  = 7
)

var discardTimerExtR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DiscardTimerExt_r16_Ms0dot5, DiscardTimerExt_r16_Ms1, DiscardTimerExt_r16_Ms2, DiscardTimerExt_r16_Ms4, DiscardTimerExt_r16_Ms6, DiscardTimerExt_r16_Ms8, DiscardTimerExt_r16_Spare2, DiscardTimerExt_r16_Spare1},
}

type DiscardTimerExt_r16 struct {
	Value int64
}

func (v *DiscardTimerExt_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, discardTimerExtR16Constraints)
}

func (v *DiscardTimerExt_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(discardTimerExtR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
