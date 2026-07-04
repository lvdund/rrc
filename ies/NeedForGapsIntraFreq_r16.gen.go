// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NeedForGapsIntraFreq-r16 (line 10462).

var needForGapsIntraFreqR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "servCellId-r16"},
		{Name: "gapIndicationIntra-r16"},
	},
}

const (
	NeedForGapsIntraFreq_r16_GapIndicationIntra_r16_Gap    = 0
	NeedForGapsIntraFreq_r16_GapIndicationIntra_r16_No_Gap = 1
)

var needForGapsIntraFreqR16GapIndicationIntraR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NeedForGapsIntraFreq_r16_GapIndicationIntra_r16_Gap, NeedForGapsIntraFreq_r16_GapIndicationIntra_r16_No_Gap},
}

type NeedForGapsIntraFreq_r16 struct {
	ServCellId_r16         ServCellIndex
	GapIndicationIntra_r16 int64
}

func (ie *NeedForGapsIntraFreq_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(needForGapsIntraFreqR16Constraints)
	_ = seq

	// 1. servCellId-r16: ref
	{
		if err := ie.ServCellId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. gapIndicationIntra-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.GapIndicationIntra_r16, needForGapsIntraFreqR16GapIndicationIntraR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NeedForGapsIntraFreq_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(needForGapsIntraFreqR16Constraints)
	_ = seq

	// 1. servCellId-r16: ref
	{
		if err := ie.ServCellId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. gapIndicationIntra-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(needForGapsIntraFreqR16GapIndicationIntraR16Constraints)
		if err != nil {
			return err
		}
		ie.GapIndicationIntra_r16 = v1
	}

	return nil
}
