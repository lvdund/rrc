// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-DiscConfig-r17 (line 27018).

var sLDiscConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RelayUE-Config-r17", Optional: true},
		{Name: "sl-RemoteUE-Config-r17", Optional: true},
	},
}

var sL_DiscConfig_r17SlRelayUEConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_DiscConfig_r17_Sl_RelayUE_Config_r17_Release = 0
	SL_DiscConfig_r17_Sl_RelayUE_Config_r17_Setup   = 1
)

var sL_DiscConfig_r17SlRemoteUEConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_DiscConfig_r17_Sl_RemoteUE_Config_r17_Release = 0
	SL_DiscConfig_r17_Sl_RemoteUE_Config_r17_Setup   = 1
)

type SL_DiscConfig_r17 struct {
	Sl_RelayUE_Config_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_RelayUE_Config_r17
	}
	Sl_RemoteUE_Config_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_RemoteUE_Config_r17
	}
}

func (ie *SL_DiscConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLDiscConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_RelayUE_Config_r17 != nil, ie.Sl_RemoteUE_Config_r17 != nil}); err != nil {
		return err
	}

	// 2. sl-RelayUE-Config-r17: choice
	{
		if ie.Sl_RelayUE_Config_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_DiscConfig_r17SlRelayUEConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_RelayUE_Config_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_RelayUE_Config_r17).Choice {
			case SL_DiscConfig_r17_Sl_RelayUE_Config_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_DiscConfig_r17_Sl_RelayUE_Config_r17_Setup:
				if err := (*ie.Sl_RelayUE_Config_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_RelayUE_Config_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. sl-RemoteUE-Config-r17: choice
	{
		if ie.Sl_RemoteUE_Config_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_DiscConfig_r17SlRemoteUEConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_RemoteUE_Config_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_RemoteUE_Config_r17).Choice {
			case SL_DiscConfig_r17_Sl_RemoteUE_Config_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_DiscConfig_r17_Sl_RemoteUE_Config_r17_Setup:
				if err := (*ie.Sl_RemoteUE_Config_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_RemoteUE_Config_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *SL_DiscConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLDiscConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-RelayUE-Config-r17: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_RelayUE_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_RelayUE_Config_r17
			}{}
			choiceDec := d.NewChoiceDecoder(sL_DiscConfig_r17SlRelayUEConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_RelayUE_Config_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_DiscConfig_r17_Sl_RelayUE_Config_r17_Release:
				(*ie.Sl_RelayUE_Config_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_DiscConfig_r17_Sl_RelayUE_Config_r17_Setup:
				(*ie.Sl_RelayUE_Config_r17).Setup = new(SL_RelayUE_Config_r17)
				if err := (*ie.Sl_RelayUE_Config_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-RemoteUE-Config-r17: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_RemoteUE_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_RemoteUE_Config_r17
			}{}
			choiceDec := d.NewChoiceDecoder(sL_DiscConfig_r17SlRemoteUEConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_RemoteUE_Config_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_DiscConfig_r17_Sl_RemoteUE_Config_r17_Release:
				(*ie.Sl_RemoteUE_Config_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_DiscConfig_r17_Sl_RemoteUE_Config_r17_Setup:
				(*ie.Sl_RemoteUE_Config_r17).Setup = new(SL_RemoteUE_Config_r17)
				if err := (*ie.Sl_RemoteUE_Config_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
