// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IMS-ParametersFRX-Diff (line 20879).

var iMSParametersFRXDiffConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "voiceOverNR", Optional: true},
	},
}

const (
	IMS_ParametersFRX_Diff_VoiceOverNR_Supported = 0
)

var iMSParametersFRXDiffVoiceOverNRConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IMS_ParametersFRX_Diff_VoiceOverNR_Supported},
}

type IMS_ParametersFRX_Diff struct {
	VoiceOverNR *int64
}

func (ie *IMS_ParametersFRX_Diff) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iMSParametersFRXDiffConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.VoiceOverNR != nil}); err != nil {
		return err
	}

	// 3. voiceOverNR: enumerated
	{
		if ie.VoiceOverNR != nil {
			if err := e.EncodeEnumerated(*ie.VoiceOverNR, iMSParametersFRXDiffVoiceOverNRConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IMS_ParametersFRX_Diff) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iMSParametersFRXDiffConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. voiceOverNR: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(iMSParametersFRXDiffVoiceOverNRConstraints)
			if err != nil {
				return err
			}
			ie.VoiceOverNR = &idx
		}
	}

	return nil
}
