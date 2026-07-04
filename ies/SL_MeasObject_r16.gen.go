// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-MeasObject-r16 (line 27562).

var sLMeasObjectR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyInfoSL-r16"},
	},
}

type SL_MeasObject_r16 struct {
	FrequencyInfoSL_r16 ARFCN_ValueNR
}

func (ie *SL_MeasObject_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLMeasObjectR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. frequencyInfoSL-r16: ref
	{
		if err := ie.FrequencyInfoSL_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_MeasObject_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLMeasObjectR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. frequencyInfoSL-r16: ref
	{
		if err := ie.FrequencyInfoSL_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
