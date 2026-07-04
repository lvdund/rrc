// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NeedForGapsNR-r16 (line 10467).

var needForGapsNRR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandNR-r16"},
		{Name: "gapIndication-r16"},
	},
}

const (
	NeedForGapsNR_r16_GapIndication_r16_Gap    = 0
	NeedForGapsNR_r16_GapIndication_r16_No_Gap = 1
)

var needForGapsNRR16GapIndicationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NeedForGapsNR_r16_GapIndication_r16_Gap, NeedForGapsNR_r16_GapIndication_r16_No_Gap},
}

type NeedForGapsNR_r16 struct {
	BandNR_r16        FreqBandIndicatorNR
	GapIndication_r16 int64
}

func (ie *NeedForGapsNR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(needForGapsNRR16Constraints)
	_ = seq

	// 1. bandNR-r16: ref
	{
		if err := ie.BandNR_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. gapIndication-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.GapIndication_r16, needForGapsNRR16GapIndicationR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NeedForGapsNR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(needForGapsNRR16Constraints)
	_ = seq

	// 1. bandNR-r16: ref
	{
		if err := ie.BandNR_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. gapIndication-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(needForGapsNRR16GapIndicationR16Constraints)
		if err != nil {
			return err
		}
		ie.GapIndication_r16 = v1
	}

	return nil
}
