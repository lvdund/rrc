// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-DiscConfig-v1830 (line 27037).

var sLDiscConfigV1830Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RemoteUE-ConfigU2U-v1830", Optional: true},
	},
}

var sL_DiscConfig_v1830SlRemoteUEConfigU2UV1830Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_DiscConfig_v1830_Sl_RemoteUE_ConfigU2U_v1830_Release = 0
	SL_DiscConfig_v1830_Sl_RemoteUE_ConfigU2U_v1830_Setup   = 1
)

type SL_DiscConfig_v1830 struct {
	Sl_RemoteUE_ConfigU2U_v1830 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_RemoteUE_ConfigU2U_v1830
	}
}

func (ie *SL_DiscConfig_v1830) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLDiscConfigV1830Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_RemoteUE_ConfigU2U_v1830 != nil}); err != nil {
		return err
	}

	// 2. sl-RemoteUE-ConfigU2U-v1830: choice
	{
		if ie.Sl_RemoteUE_ConfigU2U_v1830 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_DiscConfig_v1830SlRemoteUEConfigU2UV1830Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_RemoteUE_ConfigU2U_v1830).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_RemoteUE_ConfigU2U_v1830).Choice {
			case SL_DiscConfig_v1830_Sl_RemoteUE_ConfigU2U_v1830_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_DiscConfig_v1830_Sl_RemoteUE_ConfigU2U_v1830_Setup:
				if err := (*ie.Sl_RemoteUE_ConfigU2U_v1830).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_RemoteUE_ConfigU2U_v1830).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *SL_DiscConfig_v1830) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLDiscConfigV1830Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-RemoteUE-ConfigU2U-v1830: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_RemoteUE_ConfigU2U_v1830 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_RemoteUE_ConfigU2U_v1830
			}{}
			choiceDec := d.NewChoiceDecoder(sL_DiscConfig_v1830SlRemoteUEConfigU2UV1830Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_RemoteUE_ConfigU2U_v1830).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_DiscConfig_v1830_Sl_RemoteUE_ConfigU2U_v1830_Release:
				(*ie.Sl_RemoteUE_ConfigU2U_v1830).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_DiscConfig_v1830_Sl_RemoteUE_ConfigU2U_v1830_Setup:
				(*ie.Sl_RemoteUE_ConfigU2U_v1830).Setup = new(SL_RemoteUE_ConfigU2U_v1830)
				if err := (*ie.Sl_RemoteUE_ConfigU2U_v1830).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
