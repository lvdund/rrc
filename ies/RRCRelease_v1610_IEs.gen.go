// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCRelease-v1610-IEs (line 1246).

var rRCReleaseV1610IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "voiceFallbackIndication-r16", Optional: true},
		{Name: "measIdleConfig-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCRelease_v1610_IEs_VoiceFallbackIndication_r16_True = 0
)

var rRCReleaseV1610IEsVoiceFallbackIndicationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCRelease_v1610_IEs_VoiceFallbackIndication_r16_True},
}

var rRCRelease_v1610_IEsMeasIdleConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCRelease_v1610_IEs_MeasIdleConfig_r16_Release = 0
	RRCRelease_v1610_IEs_MeasIdleConfig_r16_Setup   = 1
)

type RRCRelease_v1610_IEs struct {
	VoiceFallbackIndication_r16 *int64
	MeasIdleConfig_r16          *struct {
		Choice  int
		Release *struct{}
		Setup   *MeasIdleConfigDedicated_r16
	}
	NonCriticalExtension *RRCRelease_v1650_IEs
}

func (ie *RRCRelease_v1610_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReleaseV1610IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.VoiceFallbackIndication_r16 != nil, ie.MeasIdleConfig_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. voiceFallbackIndication-r16: enumerated
	{
		if ie.VoiceFallbackIndication_r16 != nil {
			if err := e.EncodeEnumerated(*ie.VoiceFallbackIndication_r16, rRCReleaseV1610IEsVoiceFallbackIndicationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. measIdleConfig-r16: choice
	{
		if ie.MeasIdleConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCRelease_v1610_IEsMeasIdleConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.MeasIdleConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.MeasIdleConfig_r16).Choice {
			case RRCRelease_v1610_IEs_MeasIdleConfig_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCRelease_v1610_IEs_MeasIdleConfig_r16_Setup:
				if err := (*ie.MeasIdleConfig_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.MeasIdleConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCRelease_v1610_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReleaseV1610IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. voiceFallbackIndication-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rRCReleaseV1610IEsVoiceFallbackIndicationR16Constraints)
			if err != nil {
				return err
			}
			ie.VoiceFallbackIndication_r16 = &idx
		}
	}

	// 3. measIdleConfig-r16: choice
	{
		if seq.IsComponentPresent(1) {
			ie.MeasIdleConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MeasIdleConfigDedicated_r16
			}{}
			choiceDec := d.NewChoiceDecoder(rRCRelease_v1610_IEsMeasIdleConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MeasIdleConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCRelease_v1610_IEs_MeasIdleConfig_r16_Release:
				(*ie.MeasIdleConfig_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCRelease_v1610_IEs_MeasIdleConfig_r16_Setup:
				(*ie.MeasIdleConfig_r16).Setup = new(MeasIdleConfigDedicated_r16)
				if err := (*ie.MeasIdleConfig_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(RRCRelease_v1650_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
