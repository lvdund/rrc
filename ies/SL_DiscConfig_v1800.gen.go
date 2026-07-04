// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-DiscConfig-v1800 (line 27023).

var sLDiscConfigV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RelayUE-ConfigU2U-r18", Optional: true},
		{Name: "sl-RemoteUE-ConfigU2U-r18", Optional: true},
	},
}

var sL_DiscConfig_v1800SlRelayUEConfigU2UR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_DiscConfig_v1800_Sl_RelayUE_ConfigU2U_r18_Release = 0
	SL_DiscConfig_v1800_Sl_RelayUE_ConfigU2U_r18_Setup   = 1
)

var sL_DiscConfig_v1800SlRemoteUEConfigU2UR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_DiscConfig_v1800_Sl_RemoteUE_ConfigU2U_r18_Release = 0
	SL_DiscConfig_v1800_Sl_RemoteUE_ConfigU2U_r18_Setup   = 1
)

type SL_DiscConfig_v1800 struct {
	Sl_RelayUE_ConfigU2U_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_RelayUE_ConfigU2U_r18
	}
	Sl_RemoteUE_ConfigU2U_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_RemoteUE_ConfigU2U_r18
	}
}

func (ie *SL_DiscConfig_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLDiscConfigV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_RelayUE_ConfigU2U_r18 != nil, ie.Sl_RemoteUE_ConfigU2U_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-RelayUE-ConfigU2U-r18: choice
	{
		if ie.Sl_RelayUE_ConfigU2U_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_DiscConfig_v1800SlRelayUEConfigU2UR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_RelayUE_ConfigU2U_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_RelayUE_ConfigU2U_r18).Choice {
			case SL_DiscConfig_v1800_Sl_RelayUE_ConfigU2U_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_DiscConfig_v1800_Sl_RelayUE_ConfigU2U_r18_Setup:
				if err := (*ie.Sl_RelayUE_ConfigU2U_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_RelayUE_ConfigU2U_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. sl-RemoteUE-ConfigU2U-r18: choice
	{
		if ie.Sl_RemoteUE_ConfigU2U_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_DiscConfig_v1800SlRemoteUEConfigU2UR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_RemoteUE_ConfigU2U_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_RemoteUE_ConfigU2U_r18).Choice {
			case SL_DiscConfig_v1800_Sl_RemoteUE_ConfigU2U_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_DiscConfig_v1800_Sl_RemoteUE_ConfigU2U_r18_Setup:
				if err := (*ie.Sl_RemoteUE_ConfigU2U_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_RemoteUE_ConfigU2U_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *SL_DiscConfig_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLDiscConfigV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-RelayUE-ConfigU2U-r18: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_RelayUE_ConfigU2U_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_RelayUE_ConfigU2U_r18
			}{}
			choiceDec := d.NewChoiceDecoder(sL_DiscConfig_v1800SlRelayUEConfigU2UR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_RelayUE_ConfigU2U_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_DiscConfig_v1800_Sl_RelayUE_ConfigU2U_r18_Release:
				(*ie.Sl_RelayUE_ConfigU2U_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_DiscConfig_v1800_Sl_RelayUE_ConfigU2U_r18_Setup:
				(*ie.Sl_RelayUE_ConfigU2U_r18).Setup = new(SL_RelayUE_ConfigU2U_r18)
				if err := (*ie.Sl_RelayUE_ConfigU2U_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-RemoteUE-ConfigU2U-r18: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_RemoteUE_ConfigU2U_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_RemoteUE_ConfigU2U_r18
			}{}
			choiceDec := d.NewChoiceDecoder(sL_DiscConfig_v1800SlRemoteUEConfigU2UR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_RemoteUE_ConfigU2U_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_DiscConfig_v1800_Sl_RemoteUE_ConfigU2U_r18_Release:
				(*ie.Sl_RemoteUE_ConfigU2U_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_DiscConfig_v1800_Sl_RemoteUE_ConfigU2U_r18_Setup:
				(*ie.Sl_RemoteUE_ConfigU2U_r18).Setup = new(SL_RemoteUE_ConfigU2U_r18)
				if err := (*ie.Sl_RemoteUE_ConfigU2U_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
