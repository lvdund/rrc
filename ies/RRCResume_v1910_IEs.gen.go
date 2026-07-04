// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCResume-v1910-IEs (line 1581).

var rRCResumeV1910IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "assisted-SSB-MTC-MG-Config-r19", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCResume_v1910_IEsAssistedSSBMTCMGConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCResume_v1910_IEs_Assisted_SSB_MTC_MG_Config_r19_Release = 0
	RRCResume_v1910_IEs_Assisted_SSB_MTC_MG_Config_r19_Setup   = 1
)

type RRCResume_v1910_IEs struct {
	Assisted_SSB_MTC_MG_Config_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *Assisted_SSB_MTC_MG_Config_r19
	}
	NonCriticalExtension *struct{}
}

func (ie *RRCResume_v1910_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeV1910IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Assisted_SSB_MTC_MG_Config_r19 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. assisted-SSB-MTC-MG-Config-r19: choice
	{
		if ie.Assisted_SSB_MTC_MG_Config_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCResume_v1910_IEsAssistedSSBMTCMGConfigR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Assisted_SSB_MTC_MG_Config_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Assisted_SSB_MTC_MG_Config_r19).Choice {
			case RRCResume_v1910_IEs_Assisted_SSB_MTC_MG_Config_r19_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCResume_v1910_IEs_Assisted_SSB_MTC_MG_Config_r19_Setup:
				if err := (*ie.Assisted_SSB_MTC_MG_Config_r19).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Assisted_SSB_MTC_MG_Config_r19).Choice), Constraint: "index out of range"}
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

func (ie *RRCResume_v1910_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeV1910IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. assisted-SSB-MTC-MG-Config-r19: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Assisted_SSB_MTC_MG_Config_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *Assisted_SSB_MTC_MG_Config_r19
			}{}
			choiceDec := d.NewChoiceDecoder(rRCResume_v1910_IEsAssistedSSBMTCMGConfigR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Assisted_SSB_MTC_MG_Config_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCResume_v1910_IEs_Assisted_SSB_MTC_MG_Config_r19_Release:
				(*ie.Assisted_SSB_MTC_MG_Config_r19).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCResume_v1910_IEs_Assisted_SSB_MTC_MG_Config_r19_Setup:
				(*ie.Assisted_SSB_MTC_MG_Config_r19).Setup = new(Assisted_SSB_MTC_MG_Config_r19)
				if err := (*ie.Assisted_SSB_MTC_MG_Config_r19).Setup.Decode(d); err != nil {
					return err
				}
			}
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
