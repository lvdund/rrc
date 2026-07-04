// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RepetitionSchemeConfig-r16 (line 13343).

var repetitionSchemeConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fdm-TDM-r16"},
		{Name: "slotBased-r16"},
	},
}

const (
	RepetitionSchemeConfig_r16_Fdm_TDM_r16   = 0
	RepetitionSchemeConfig_r16_SlotBased_r16 = 1
)

var repetitionSchemeConfigR16FdmTDMR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RepetitionSchemeConfig_r16_Fdm_TDM_r16_Release = 0
	RepetitionSchemeConfig_r16_Fdm_TDM_r16_Setup   = 1
)

var repetitionSchemeConfigR16SlotBasedR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RepetitionSchemeConfig_r16_SlotBased_r16_Release = 0
	RepetitionSchemeConfig_r16_SlotBased_r16_Setup   = 1
)

type RepetitionSchemeConfig_r16 struct {
	Choice      int
	Fdm_TDM_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *FDM_TDM_r16
	}
	SlotBased_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *SlotBased_r16
	}
}

func (ie *RepetitionSchemeConfig_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(repetitionSchemeConfigR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case RepetitionSchemeConfig_r16_Fdm_TDM_r16:
		choiceEnc := e.NewChoiceEncoder(repetitionSchemeConfigR16FdmTDMR16Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Fdm_TDM_r16).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Fdm_TDM_r16).Choice {
		case RepetitionSchemeConfig_r16_Fdm_TDM_r16_Release:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case RepetitionSchemeConfig_r16_Fdm_TDM_r16_Setup:
			if err := (*ie.Fdm_TDM_r16).Setup.Encode(e); err != nil {
				return err
			}
		}
	case RepetitionSchemeConfig_r16_SlotBased_r16:
		choiceEnc := e.NewChoiceEncoder(repetitionSchemeConfigR16SlotBasedR16Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.SlotBased_r16).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.SlotBased_r16).Choice {
		case RepetitionSchemeConfig_r16_SlotBased_r16_Release:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case RepetitionSchemeConfig_r16_SlotBased_r16_Setup:
			if err := (*ie.SlotBased_r16).Setup.Encode(e); err != nil {
				return err
			}
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "RepetitionSchemeConfig-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *RepetitionSchemeConfig_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(repetitionSchemeConfigR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "RepetitionSchemeConfig-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case RepetitionSchemeConfig_r16_Fdm_TDM_r16:
		ie.Fdm_TDM_r16 = &struct {
			Choice  int
			Release *struct{}
			Setup   *FDM_TDM_r16
		}{}
		choiceDec := d.NewChoiceDecoder(repetitionSchemeConfigR16FdmTDMR16Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Fdm_TDM_r16).Choice = int(index)
		switch index {
		case RepetitionSchemeConfig_r16_Fdm_TDM_r16_Release:
			(*ie.Fdm_TDM_r16).Release = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case RepetitionSchemeConfig_r16_Fdm_TDM_r16_Setup:
			(*ie.Fdm_TDM_r16).Setup = new(FDM_TDM_r16)
			if err := (*ie.Fdm_TDM_r16).Setup.Decode(d); err != nil {
				return err
			}
		}
	case RepetitionSchemeConfig_r16_SlotBased_r16:
		ie.SlotBased_r16 = &struct {
			Choice  int
			Release *struct{}
			Setup   *SlotBased_r16
		}{}
		choiceDec := d.NewChoiceDecoder(repetitionSchemeConfigR16SlotBasedR16Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.SlotBased_r16).Choice = int(index)
		switch index {
		case RepetitionSchemeConfig_r16_SlotBased_r16_Release:
			(*ie.SlotBased_r16).Release = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case RepetitionSchemeConfig_r16_SlotBased_r16_Setup:
			(*ie.SlotBased_r16).Setup = new(SlotBased_r16)
			if err := (*ie.SlotBased_r16).Setup.Decode(d); err != nil {
				return err
			}
		}
	default:
		return &per.DecodeError{TypeName: "RepetitionSchemeConfig-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
