// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CarrierFreqEUTRA-v1610 (line 4153).

var carrierFreqEUTRAV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "highSpeedEUTRACarrier-r16", Optional: true},
	},
}

const (
	CarrierFreqEUTRA_v1610_HighSpeedEUTRACarrier_r16_True = 0
)

var carrierFreqEUTRAV1610HighSpeedEUTRACarrierR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CarrierFreqEUTRA_v1610_HighSpeedEUTRACarrier_r16_True},
}

type CarrierFreqEUTRA_v1610 struct {
	HighSpeedEUTRACarrier_r16 *int64
}

func (ie *CarrierFreqEUTRA_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(carrierFreqEUTRAV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.HighSpeedEUTRACarrier_r16 != nil}); err != nil {
		return err
	}

	// 2. highSpeedEUTRACarrier-r16: enumerated
	{
		if ie.HighSpeedEUTRACarrier_r16 != nil {
			if err := e.EncodeEnumerated(*ie.HighSpeedEUTRACarrier_r16, carrierFreqEUTRAV1610HighSpeedEUTRACarrierR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CarrierFreqEUTRA_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(carrierFreqEUTRAV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. highSpeedEUTRACarrier-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(carrierFreqEUTRAV1610HighSpeedEUTRACarrierR16Constraints)
			if err != nil {
				return err
			}
			ie.HighSpeedEUTRACarrier_r16 = &idx
		}
	}

	return nil
}
