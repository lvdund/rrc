// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IMS-ParametersFR2-2-r17 (line 20884).

var iMSParametersFR22R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "voiceOverNR-r17", Optional: true},
	},
}

const (
	IMS_ParametersFR2_2_r17_VoiceOverNR_r17_Supported = 0
)

var iMSParametersFR22R17VoiceOverNRR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IMS_ParametersFR2_2_r17_VoiceOverNR_r17_Supported},
}

type IMS_ParametersFR2_2_r17 struct {
	VoiceOverNR_r17 *int64
}

func (ie *IMS_ParametersFR2_2_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iMSParametersFR22R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.VoiceOverNR_r17 != nil}); err != nil {
		return err
	}

	// 3. voiceOverNR-r17: enumerated
	{
		if ie.VoiceOverNR_r17 != nil {
			if err := e.EncodeEnumerated(*ie.VoiceOverNR_r17, iMSParametersFR22R17VoiceOverNRR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IMS_ParametersFR2_2_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iMSParametersFR22R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. voiceOverNR-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(iMSParametersFR22R17VoiceOverNRR17Constraints)
			if err != nil {
				return err
			}
			ie.VoiceOverNR_r17 = &idx
		}
	}

	return nil
}
