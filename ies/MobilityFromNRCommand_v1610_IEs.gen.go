// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MobilityFromNRCommand-v1610-IEs (line 828).

var mobilityFromNRCommandV1610IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "voiceFallbackIndication-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	MobilityFromNRCommand_v1610_IEs_VoiceFallbackIndication_r16_True = 0
)

var mobilityFromNRCommandV1610IEsVoiceFallbackIndicationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MobilityFromNRCommand_v1610_IEs_VoiceFallbackIndication_r16_True},
}

type MobilityFromNRCommand_v1610_IEs struct {
	VoiceFallbackIndication_r16 *int64
	NonCriticalExtension        *struct{}
}

func (ie *MobilityFromNRCommand_v1610_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mobilityFromNRCommandV1610IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.VoiceFallbackIndication_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. voiceFallbackIndication-r16: enumerated
	{
		if ie.VoiceFallbackIndication_r16 != nil {
			if err := e.EncodeEnumerated(*ie.VoiceFallbackIndication_r16, mobilityFromNRCommandV1610IEsVoiceFallbackIndicationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *MobilityFromNRCommand_v1610_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mobilityFromNRCommandV1610IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. voiceFallbackIndication-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mobilityFromNRCommandV1610IEsVoiceFallbackIndicationR16Constraints)
			if err != nil {
				return err
			}
			ie.VoiceFallbackIndication_r16 = &idx
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
