// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SNPN-ConfigList-r18 (line 26092).

var sNPNConfigListR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "snpn-ConfigCellIdList-r18"},
		{Name: "snpn-ConfigTAI-List-r18"},
		{Name: "snpn-ConfigID-List-r18"},
	},
}

const (
	SNPN_ConfigList_r18_Snpn_ConfigCellIdList_r18 = 0
	SNPN_ConfigList_r18_Snpn_ConfigTAI_List_r18   = 1
	SNPN_ConfigList_r18_Snpn_ConfigID_List_r18    = 2
)

type SNPN_ConfigList_r18 struct {
	Choice                    int
	Snpn_ConfigCellIdList_r18 *SNPN_ConfigCellIdList_r18
	Snpn_ConfigTAI_List_r18   *SNPN_ConfigTAI_List_r18
	Snpn_ConfigID_List_r18    *SNPN_ConfigID_List_r18
}

func (ie *SNPN_ConfigList_r18) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sNPNConfigListR18Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SNPN_ConfigList_r18_Snpn_ConfigCellIdList_r18:
		if err := ie.Snpn_ConfigCellIdList_r18.Encode(e); err != nil {
			return err
		}
	case SNPN_ConfigList_r18_Snpn_ConfigTAI_List_r18:
		if err := ie.Snpn_ConfigTAI_List_r18.Encode(e); err != nil {
			return err
		}
	case SNPN_ConfigList_r18_Snpn_ConfigID_List_r18:
		if err := ie.Snpn_ConfigID_List_r18.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SNPN-ConfigList-r18",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SNPN_ConfigList_r18) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sNPNConfigListR18Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SNPN-ConfigList-r18", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SNPN_ConfigList_r18_Snpn_ConfigCellIdList_r18:
		ie.Snpn_ConfigCellIdList_r18 = new(SNPN_ConfigCellIdList_r18)
		if err := ie.Snpn_ConfigCellIdList_r18.Decode(d); err != nil {
			return err
		}
	case SNPN_ConfigList_r18_Snpn_ConfigTAI_List_r18:
		ie.Snpn_ConfigTAI_List_r18 = new(SNPN_ConfigTAI_List_r18)
		if err := ie.Snpn_ConfigTAI_List_r18.Decode(d); err != nil {
			return err
		}
	case SNPN_ConfigList_r18_Snpn_ConfigID_List_r18:
		ie.Snpn_ConfigID_List_r18 = new(SNPN_ConfigID_List_r18)
		if err := ie.Snpn_ConfigID_List_r18.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "SNPN-ConfigList-r18", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
