// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DormantBWP-Config-r16 (line 14784).

var dormantBWPConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dormantBWP-Id-r16", Optional: true},
		{Name: "withinActiveTimeConfig-r16", Optional: true},
		{Name: "outsideActiveTimeConfig-r16", Optional: true},
	},
}

var dormantBWP_Config_r16WithinActiveTimeConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	DormantBWP_Config_r16_WithinActiveTimeConfig_r16_Release = 0
	DormantBWP_Config_r16_WithinActiveTimeConfig_r16_Setup   = 1
)

var dormantBWP_Config_r16OutsideActiveTimeConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	DormantBWP_Config_r16_OutsideActiveTimeConfig_r16_Release = 0
	DormantBWP_Config_r16_OutsideActiveTimeConfig_r16_Setup   = 1
)

type DormantBWP_Config_r16 struct {
	DormantBWP_Id_r16          *BWP_Id
	WithinActiveTimeConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *WithinActiveTimeConfig_r16
	}
	OutsideActiveTimeConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *OutsideActiveTimeConfig_r16
	}
}

func (ie *DormantBWP_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dormantBWPConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DormantBWP_Id_r16 != nil, ie.WithinActiveTimeConfig_r16 != nil, ie.OutsideActiveTimeConfig_r16 != nil}); err != nil {
		return err
	}

	// 2. dormantBWP-Id-r16: ref
	{
		if ie.DormantBWP_Id_r16 != nil {
			if err := ie.DormantBWP_Id_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. withinActiveTimeConfig-r16: choice
	{
		if ie.WithinActiveTimeConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(dormantBWP_Config_r16WithinActiveTimeConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.WithinActiveTimeConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.WithinActiveTimeConfig_r16).Choice {
			case DormantBWP_Config_r16_WithinActiveTimeConfig_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case DormantBWP_Config_r16_WithinActiveTimeConfig_r16_Setup:
				if err := (*ie.WithinActiveTimeConfig_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.WithinActiveTimeConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. outsideActiveTimeConfig-r16: choice
	{
		if ie.OutsideActiveTimeConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(dormantBWP_Config_r16OutsideActiveTimeConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.OutsideActiveTimeConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.OutsideActiveTimeConfig_r16).Choice {
			case DormantBWP_Config_r16_OutsideActiveTimeConfig_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case DormantBWP_Config_r16_OutsideActiveTimeConfig_r16_Setup:
				if err := (*ie.OutsideActiveTimeConfig_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.OutsideActiveTimeConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *DormantBWP_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dormantBWPConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dormantBWP-Id-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.DormantBWP_Id_r16 = new(BWP_Id)
			if err := ie.DormantBWP_Id_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. withinActiveTimeConfig-r16: choice
	{
		if seq.IsComponentPresent(1) {
			ie.WithinActiveTimeConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *WithinActiveTimeConfig_r16
			}{}
			choiceDec := d.NewChoiceDecoder(dormantBWP_Config_r16WithinActiveTimeConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.WithinActiveTimeConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case DormantBWP_Config_r16_WithinActiveTimeConfig_r16_Release:
				(*ie.WithinActiveTimeConfig_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case DormantBWP_Config_r16_WithinActiveTimeConfig_r16_Setup:
				(*ie.WithinActiveTimeConfig_r16).Setup = new(WithinActiveTimeConfig_r16)
				if err := (*ie.WithinActiveTimeConfig_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. outsideActiveTimeConfig-r16: choice
	{
		if seq.IsComponentPresent(2) {
			ie.OutsideActiveTimeConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *OutsideActiveTimeConfig_r16
			}{}
			choiceDec := d.NewChoiceDecoder(dormantBWP_Config_r16OutsideActiveTimeConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.OutsideActiveTimeConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case DormantBWP_Config_r16_OutsideActiveTimeConfig_r16_Release:
				(*ie.OutsideActiveTimeConfig_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case DormantBWP_Config_r16_OutsideActiveTimeConfig_r16_Setup:
				(*ie.OutsideActiveTimeConfig_r16).Setup = new(OutsideActiveTimeConfig_r16)
				if err := (*ie.OutsideActiveTimeConfig_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
