// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-MeasConfigInfo-r16 (line 27512).

var sLMeasConfigInfoR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DestinationIndex-r16"},
		{Name: "sl-MeasConfig-r16"},
	},
}

type SL_MeasConfigInfo_r16 struct {
	Sl_DestinationIndex_r16 SL_DestinationIndex_r16
	Sl_MeasConfig_r16       SL_MeasConfig_r16
}

func (ie *SL_MeasConfigInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLMeasConfigInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sl-DestinationIndex-r16: ref
	{
		if err := ie.Sl_DestinationIndex_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-MeasConfig-r16: ref
	{
		if err := ie.Sl_MeasConfig_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_MeasConfigInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLMeasConfigInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sl-DestinationIndex-r16: ref
	{
		if err := ie.Sl_DestinationIndex_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-MeasConfig-r16: ref
	{
		if err := ie.Sl_MeasConfig_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
