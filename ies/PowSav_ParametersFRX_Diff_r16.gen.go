// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PowSav-ParametersFRX-Diff-r16 (line 23461).

var powSavParametersFRXDiffR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "maxBW-Preference-r16", Optional: true},
		{Name: "maxMIMO-LayerPreference-r16", Optional: true},
	},
}

const (
	PowSav_ParametersFRX_Diff_r16_MaxBW_Preference_r16_Supported = 0
)

var powSavParametersFRXDiffR16MaxBWPreferenceR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PowSav_ParametersFRX_Diff_r16_MaxBW_Preference_r16_Supported},
}

const (
	PowSav_ParametersFRX_Diff_r16_MaxMIMO_LayerPreference_r16_Supported = 0
)

var powSavParametersFRXDiffR16MaxMIMOLayerPreferenceR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PowSav_ParametersFRX_Diff_r16_MaxMIMO_LayerPreference_r16_Supported},
}

type PowSav_ParametersFRX_Diff_r16 struct {
	MaxBW_Preference_r16        *int64
	MaxMIMO_LayerPreference_r16 *int64
}

func (ie *PowSav_ParametersFRX_Diff_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(powSavParametersFRXDiffR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxBW_Preference_r16 != nil, ie.MaxMIMO_LayerPreference_r16 != nil}); err != nil {
		return err
	}

	// 3. maxBW-Preference-r16: enumerated
	{
		if ie.MaxBW_Preference_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MaxBW_Preference_r16, powSavParametersFRXDiffR16MaxBWPreferenceR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. maxMIMO-LayerPreference-r16: enumerated
	{
		if ie.MaxMIMO_LayerPreference_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MaxMIMO_LayerPreference_r16, powSavParametersFRXDiffR16MaxMIMOLayerPreferenceR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PowSav_ParametersFRX_Diff_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(powSavParametersFRXDiffR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. maxBW-Preference-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(powSavParametersFRXDiffR16MaxBWPreferenceR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxBW_Preference_r16 = &idx
		}
	}

	// 4. maxMIMO-LayerPreference-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(powSavParametersFRXDiffR16MaxMIMOLayerPreferenceR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxMIMO_LayerPreference_r16 = &idx
		}
	}

	return nil
}
