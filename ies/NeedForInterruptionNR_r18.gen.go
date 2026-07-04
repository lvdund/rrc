// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NeedForInterruptionNR-r18 (line 10533).

var needForInterruptionNRR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "interruptionIndication-r18", Optional: true},
	},
}

const (
	NeedForInterruptionNR_r18_InterruptionIndication_r18_No_Gap_With_Interruption = 0
	NeedForInterruptionNR_r18_InterruptionIndication_r18_No_Gap_No_Interruption   = 1
)

var needForInterruptionNRR18InterruptionIndicationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NeedForInterruptionNR_r18_InterruptionIndication_r18_No_Gap_With_Interruption, NeedForInterruptionNR_r18_InterruptionIndication_r18_No_Gap_No_Interruption},
}

type NeedForInterruptionNR_r18 struct {
	InterruptionIndication_r18 *int64
}

func (ie *NeedForInterruptionNR_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(needForInterruptionNRR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InterruptionIndication_r18 != nil}); err != nil {
		return err
	}

	// 2. interruptionIndication-r18: enumerated
	{
		if ie.InterruptionIndication_r18 != nil {
			if err := e.EncodeEnumerated(*ie.InterruptionIndication_r18, needForInterruptionNRR18InterruptionIndicationR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NeedForInterruptionNR_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(needForInterruptionNRR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. interruptionIndication-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(needForInterruptionNRR18InterruptionIndicationR18Constraints)
			if err != nil {
				return err
			}
			ie.InterruptionIndication_r18 = &idx
		}
	}

	return nil
}
