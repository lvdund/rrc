// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-DiscConfig-v1900 (line 27045).

var sLDiscConfigV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RelayUE-ConfigMH-r19", Optional: true},
	},
}

var sL_DiscConfig_v1900SlRelayUEConfigMHR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_DiscConfig_v1900_Sl_RelayUE_ConfigMH_r19_Release = 0
	SL_DiscConfig_v1900_Sl_RelayUE_ConfigMH_r19_Setup   = 1
)

type SL_DiscConfig_v1900 struct {
	Sl_RelayUE_ConfigMH_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_RelayUE_ConfigMH_r19
	}
}

func (ie *SL_DiscConfig_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLDiscConfigV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_RelayUE_ConfigMH_r19 != nil}); err != nil {
		return err
	}

	// 2. sl-RelayUE-ConfigMH-r19: choice
	{
		if ie.Sl_RelayUE_ConfigMH_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_DiscConfig_v1900SlRelayUEConfigMHR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_RelayUE_ConfigMH_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_RelayUE_ConfigMH_r19).Choice {
			case SL_DiscConfig_v1900_Sl_RelayUE_ConfigMH_r19_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_DiscConfig_v1900_Sl_RelayUE_ConfigMH_r19_Setup:
				if err := (*ie.Sl_RelayUE_ConfigMH_r19).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_RelayUE_ConfigMH_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *SL_DiscConfig_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLDiscConfigV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-RelayUE-ConfigMH-r19: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_RelayUE_ConfigMH_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_RelayUE_ConfigMH_r19
			}{}
			choiceDec := d.NewChoiceDecoder(sL_DiscConfig_v1900SlRelayUEConfigMHR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_RelayUE_ConfigMH_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_DiscConfig_v1900_Sl_RelayUE_ConfigMH_r19_Release:
				(*ie.Sl_RelayUE_ConfigMH_r19).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_DiscConfig_v1900_Sl_RelayUE_ConfigMH_r19_Setup:
				(*ie.Sl_RelayUE_ConfigMH_r19).Setup = new(SL_RelayUE_ConfigMH_r19)
				if err := (*ie.Sl_RelayUE_ConfigMH_r19).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
