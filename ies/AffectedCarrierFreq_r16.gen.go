// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AffectedCarrierFreq-r16 (line 2506).

var affectedCarrierFreqR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq-r16"},
		{Name: "interferenceDirection-r16"},
	},
}

const (
	AffectedCarrierFreq_r16_InterferenceDirection_r16_Nr    = 0
	AffectedCarrierFreq_r16_InterferenceDirection_r16_Other = 1
	AffectedCarrierFreq_r16_InterferenceDirection_r16_Both  = 2
	AffectedCarrierFreq_r16_InterferenceDirection_r16_Spare = 3
)

var affectedCarrierFreqR16InterferenceDirectionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AffectedCarrierFreq_r16_InterferenceDirection_r16_Nr, AffectedCarrierFreq_r16_InterferenceDirection_r16_Other, AffectedCarrierFreq_r16_InterferenceDirection_r16_Both, AffectedCarrierFreq_r16_InterferenceDirection_r16_Spare},
}

type AffectedCarrierFreq_r16 struct {
	CarrierFreq_r16           ARFCN_ValueNR
	InterferenceDirection_r16 int64
}

func (ie *AffectedCarrierFreq_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(affectedCarrierFreqR16Constraints)
	_ = seq

	// 1. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. interferenceDirection-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.InterferenceDirection_r16, affectedCarrierFreqR16InterferenceDirectionR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *AffectedCarrierFreq_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(affectedCarrierFreqR16Constraints)
	_ = seq

	// 1. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. interferenceDirection-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(affectedCarrierFreqR16InterferenceDirectionR16Constraints)
		if err != nil {
			return err
		}
		ie.InterferenceDirection_r16 = v1
	}

	return nil
}
