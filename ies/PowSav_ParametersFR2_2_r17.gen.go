// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PowSav-ParametersFR2-2-r17 (line 23467).

var powSavParametersFR22R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "maxBW-Preference-r17", Optional: true},
		{Name: "maxMIMO-LayerPreference-r17", Optional: true},
	},
}

const (
	PowSav_ParametersFR2_2_r17_MaxBW_Preference_r17_Supported = 0
)

var powSavParametersFR22R17MaxBWPreferenceR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PowSav_ParametersFR2_2_r17_MaxBW_Preference_r17_Supported},
}

const (
	PowSav_ParametersFR2_2_r17_MaxMIMO_LayerPreference_r17_Supported = 0
)

var powSavParametersFR22R17MaxMIMOLayerPreferenceR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PowSav_ParametersFR2_2_r17_MaxMIMO_LayerPreference_r17_Supported},
}

type PowSav_ParametersFR2_2_r17 struct {
	MaxBW_Preference_r17        *int64
	MaxMIMO_LayerPreference_r17 *int64
}

func (ie *PowSav_ParametersFR2_2_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(powSavParametersFR22R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxBW_Preference_r17 != nil, ie.MaxMIMO_LayerPreference_r17 != nil}); err != nil {
		return err
	}

	// 3. maxBW-Preference-r17: enumerated
	{
		if ie.MaxBW_Preference_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxBW_Preference_r17, powSavParametersFR22R17MaxBWPreferenceR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. maxMIMO-LayerPreference-r17: enumerated
	{
		if ie.MaxMIMO_LayerPreference_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxMIMO_LayerPreference_r17, powSavParametersFR22R17MaxMIMOLayerPreferenceR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PowSav_ParametersFR2_2_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(powSavParametersFR22R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. maxBW-Preference-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(powSavParametersFR22R17MaxBWPreferenceR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxBW_Preference_r17 = &idx
		}
	}

	// 4. maxMIMO-LayerPreference-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(powSavParametersFR22R17MaxMIMOLayerPreferenceR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxMIMO_LayerPreference_r17 = &idx
		}
	}

	return nil
}
