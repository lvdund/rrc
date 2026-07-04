// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IMS-ParametersCommon (line 20868).

var iMSParametersCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "voiceOverEUTRA-5GC", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	IMS_ParametersCommon_VoiceOverEUTRA_5GC_Supported = 0
)

var iMSParametersCommonVoiceOverEUTRA5GCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IMS_ParametersCommon_VoiceOverEUTRA_5GC_Supported},
}

const (
	IMS_ParametersCommon_Ext_VoiceOverSCG_BearerEUTRA_5GC_Supported = 0
)

var iMSParametersCommonExtVoiceOverSCGBearerEUTRA5GCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IMS_ParametersCommon_Ext_VoiceOverSCG_BearerEUTRA_5GC_Supported},
}

const (
	IMS_ParametersCommon_Ext_VoiceFallbackIndicationEPS_r16_Supported = 0
)

var iMSParametersCommonExtVoiceFallbackIndicationEPSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IMS_ParametersCommon_Ext_VoiceFallbackIndicationEPS_r16_Supported},
}

type IMS_ParametersCommon struct {
	VoiceOverEUTRA_5GC             *int64
	VoiceOverSCG_BearerEUTRA_5GC   *int64
	VoiceFallbackIndicationEPS_r16 *int64
}

func (ie *IMS_ParametersCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iMSParametersCommonConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.VoiceOverSCG_BearerEUTRA_5GC != nil
	hasExtGroup1 := ie.VoiceFallbackIndicationEPS_r16 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.VoiceOverEUTRA_5GC != nil}); err != nil {
		return err
	}

	// 3. voiceOverEUTRA-5GC: enumerated
	{
		if ie.VoiceOverEUTRA_5GC != nil {
			if err := e.EncodeEnumerated(*ie.VoiceOverEUTRA_5GC, iMSParametersCommonVoiceOverEUTRA5GCConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "voiceOverSCG-BearerEUTRA-5GC", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.VoiceOverSCG_BearerEUTRA_5GC != nil}); err != nil {
				return err
			}

			if ie.VoiceOverSCG_BearerEUTRA_5GC != nil {
				if err := ex.EncodeEnumerated(*ie.VoiceOverSCG_BearerEUTRA_5GC, iMSParametersCommonExtVoiceOverSCGBearerEUTRA5GCConstraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "voiceFallbackIndicationEPS-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.VoiceFallbackIndicationEPS_r16 != nil}); err != nil {
				return err
			}

			if ie.VoiceFallbackIndicationEPS_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.VoiceFallbackIndicationEPS_r16, iMSParametersCommonExtVoiceFallbackIndicationEPSR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *IMS_ParametersCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iMSParametersCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. voiceOverEUTRA-5GC: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(iMSParametersCommonVoiceOverEUTRA5GCConstraints)
			if err != nil {
				return err
			}
			ie.VoiceOverEUTRA_5GC = &idx
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "voiceOverSCG-BearerEUTRA-5GC", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(iMSParametersCommonExtVoiceOverSCGBearerEUTRA5GCConstraints)
			if err != nil {
				return err
			}
			ie.VoiceOverSCG_BearerEUTRA_5GC = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "voiceFallbackIndicationEPS-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(iMSParametersCommonExtVoiceFallbackIndicationEPSR16Constraints)
			if err != nil {
				return err
			}
			ie.VoiceFallbackIndicationEPS_r16 = &v
		}
	}

	return nil
}
