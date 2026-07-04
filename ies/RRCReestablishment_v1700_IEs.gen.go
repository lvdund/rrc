// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReestablishment-v1700-IEs (line 905).

var rRCReestablishmentV1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-L2RemoteUE-Config-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCReestablishment_v1700_IEsSlL2RemoteUEConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReestablishment_v1700_IEs_Sl_L2RemoteUE_Config_r17_Release = 0
	RRCReestablishment_v1700_IEs_Sl_L2RemoteUE_Config_r17_Setup   = 1
)

type RRCReestablishment_v1700_IEs struct {
	Sl_L2RemoteUE_Config_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_L2RemoteUE_Config_r17
	}
	NonCriticalExtension *struct{}
}

func (ie *RRCReestablishment_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReestablishmentV1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_L2RemoteUE_Config_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. sl-L2RemoteUE-Config-r17: choice
	{
		if ie.Sl_L2RemoteUE_Config_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReestablishment_v1700_IEsSlL2RemoteUEConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_L2RemoteUE_Config_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_L2RemoteUE_Config_r17).Choice {
			case RRCReestablishment_v1700_IEs_Sl_L2RemoteUE_Config_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReestablishment_v1700_IEs_Sl_L2RemoteUE_Config_r17_Setup:
				if err := (*ie.Sl_L2RemoteUE_Config_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_L2RemoteUE_Config_r17).Choice), Constraint: "index out of range"}
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

func (ie *RRCReestablishment_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReestablishmentV1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-L2RemoteUE-Config-r17: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_L2RemoteUE_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_L2RemoteUE_Config_r17
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReestablishment_v1700_IEsSlL2RemoteUEConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_L2RemoteUE_Config_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReestablishment_v1700_IEs_Sl_L2RemoteUE_Config_r17_Release:
				(*ie.Sl_L2RemoteUE_Config_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReestablishment_v1700_IEs_Sl_L2RemoteUE_Config_r17_Setup:
				(*ie.Sl_L2RemoteUE_Config_r17).Setup = new(SL_L2RemoteUE_Config_r17)
				if err := (*ie.Sl_L2RemoteUE_Config_r17).Setup.Decode(d); err != nil {
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
