// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PowSav-ParametersCommon-r16 (line 23452).

var powSavParametersCommonR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drx-Preference-r16", Optional: true},
		{Name: "maxCC-Preference-r16", Optional: true},
		{Name: "releasePreference-r16", Optional: true},
		{Name: "minSchedulingOffsetPreference-r16", Optional: true},
	},
}

const (
	PowSav_ParametersCommon_r16_Drx_Preference_r16_Supported = 0
)

var powSavParametersCommonR16DrxPreferenceR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PowSav_ParametersCommon_r16_Drx_Preference_r16_Supported},
}

const (
	PowSav_ParametersCommon_r16_MaxCC_Preference_r16_Supported = 0
)

var powSavParametersCommonR16MaxCCPreferenceR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PowSav_ParametersCommon_r16_MaxCC_Preference_r16_Supported},
}

const (
	PowSav_ParametersCommon_r16_ReleasePreference_r16_Supported = 0
)

var powSavParametersCommonR16ReleasePreferenceR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PowSav_ParametersCommon_r16_ReleasePreference_r16_Supported},
}

const (
	PowSav_ParametersCommon_r16_MinSchedulingOffsetPreference_r16_Supported = 0
)

var powSavParametersCommonR16MinSchedulingOffsetPreferenceR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PowSav_ParametersCommon_r16_MinSchedulingOffsetPreference_r16_Supported},
}

type PowSav_ParametersCommon_r16 struct {
	Drx_Preference_r16                *int64
	MaxCC_Preference_r16              *int64
	ReleasePreference_r16             *int64
	MinSchedulingOffsetPreference_r16 *int64
}

func (ie *PowSav_ParametersCommon_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(powSavParametersCommonR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Drx_Preference_r16 != nil, ie.MaxCC_Preference_r16 != nil, ie.ReleasePreference_r16 != nil, ie.MinSchedulingOffsetPreference_r16 != nil}); err != nil {
		return err
	}

	// 3. drx-Preference-r16: enumerated
	{
		if ie.Drx_Preference_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Drx_Preference_r16, powSavParametersCommonR16DrxPreferenceR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. maxCC-Preference-r16: enumerated
	{
		if ie.MaxCC_Preference_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MaxCC_Preference_r16, powSavParametersCommonR16MaxCCPreferenceR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. releasePreference-r16: enumerated
	{
		if ie.ReleasePreference_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ReleasePreference_r16, powSavParametersCommonR16ReleasePreferenceR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. minSchedulingOffsetPreference-r16: enumerated
	{
		if ie.MinSchedulingOffsetPreference_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MinSchedulingOffsetPreference_r16, powSavParametersCommonR16MinSchedulingOffsetPreferenceR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PowSav_ParametersCommon_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(powSavParametersCommonR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. drx-Preference-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(powSavParametersCommonR16DrxPreferenceR16Constraints)
			if err != nil {
				return err
			}
			ie.Drx_Preference_r16 = &idx
		}
	}

	// 4. maxCC-Preference-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(powSavParametersCommonR16MaxCCPreferenceR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxCC_Preference_r16 = &idx
		}
	}

	// 5. releasePreference-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(powSavParametersCommonR16ReleasePreferenceR16Constraints)
			if err != nil {
				return err
			}
			ie.ReleasePreference_r16 = &idx
		}
	}

	// 6. minSchedulingOffsetPreference-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(powSavParametersCommonR16MinSchedulingOffsetPreferenceR16Constraints)
			if err != nil {
				return err
			}
			ie.MinSchedulingOffsetPreference_r16 = &idx
		}
	}

	return nil
}
