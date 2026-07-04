// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AreaConfiguration-v1800 (line 26059).

var areaConfigurationV1800Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cag-ConfigList-r18"},
		{Name: "snpn-ConfigList-r18"},
	},
}

const (
	AreaConfiguration_v1800_Cag_ConfigList_r18  = 0
	AreaConfiguration_v1800_Snpn_ConfigList_r18 = 1
)

type AreaConfiguration_v1800 struct {
	Choice              int
	Cag_ConfigList_r18  *CAG_ConfigList_r18
	Snpn_ConfigList_r18 *SNPN_ConfigList_r18
}

func (ie *AreaConfiguration_v1800) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(areaConfigurationV1800Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case AreaConfiguration_v1800_Cag_ConfigList_r18:
		if err := ie.Cag_ConfigList_r18.Encode(e); err != nil {
			return err
		}
	case AreaConfiguration_v1800_Snpn_ConfigList_r18:
		if err := ie.Snpn_ConfigList_r18.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "AreaConfiguration-v1800",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *AreaConfiguration_v1800) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(areaConfigurationV1800Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "AreaConfiguration-v1800", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case AreaConfiguration_v1800_Cag_ConfigList_r18:
		ie.Cag_ConfigList_r18 = new(CAG_ConfigList_r18)
		if err := ie.Cag_ConfigList_r18.Decode(d); err != nil {
			return err
		}
	case AreaConfiguration_v1800_Snpn_ConfigList_r18:
		ie.Snpn_ConfigList_r18 = new(SNPN_ConfigList_r18)
		if err := ie.Snpn_ConfigList_r18.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "AreaConfiguration-v1800", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
