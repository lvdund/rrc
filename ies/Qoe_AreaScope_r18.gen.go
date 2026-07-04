// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Qoe-AreaScope-r18 (line 25993).

var qoeAreaScopeR18Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalIdList"},
		{Name: "trackingAreaCodeList"},
		{Name: "trackingAreaIdentityList"},
		{Name: "plmn-IdentityList"},
	},
}

const (
	Qoe_AreaScope_r18_CellGlobalIdList         = 0
	Qoe_AreaScope_r18_TrackingAreaCodeList     = 1
	Qoe_AreaScope_r18_TrackingAreaIdentityList = 2
	Qoe_AreaScope_r18_Plmn_IdentityList        = 3
)

type Qoe_AreaScope_r18 struct {
	Choice                   int
	CellGlobalIdList         *CellGlobalIdList_r16
	TrackingAreaCodeList     *TrackingAreaCodeList_r16
	TrackingAreaIdentityList *TrackingAreaIdentityList_r16
	Plmn_IdentityList        *PLMN_IdentityList2_r16
}

func (ie *Qoe_AreaScope_r18) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(qoeAreaScopeR18Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case Qoe_AreaScope_r18_CellGlobalIdList:
		if err := ie.CellGlobalIdList.Encode(e); err != nil {
			return err
		}
	case Qoe_AreaScope_r18_TrackingAreaCodeList:
		if err := ie.TrackingAreaCodeList.Encode(e); err != nil {
			return err
		}
	case Qoe_AreaScope_r18_TrackingAreaIdentityList:
		if err := ie.TrackingAreaIdentityList.Encode(e); err != nil {
			return err
		}
	case Qoe_AreaScope_r18_Plmn_IdentityList:
		if err := ie.Plmn_IdentityList.Encode(e); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "Qoe-AreaScope-r18",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *Qoe_AreaScope_r18) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(qoeAreaScopeR18Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "Qoe-AreaScope-r18", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case Qoe_AreaScope_r18_CellGlobalIdList:
		ie.CellGlobalIdList = new(CellGlobalIdList_r16)
		if err := ie.CellGlobalIdList.Decode(d); err != nil {
			return err
		}
	case Qoe_AreaScope_r18_TrackingAreaCodeList:
		ie.TrackingAreaCodeList = new(TrackingAreaCodeList_r16)
		if err := ie.TrackingAreaCodeList.Decode(d); err != nil {
			return err
		}
	case Qoe_AreaScope_r18_TrackingAreaIdentityList:
		ie.TrackingAreaIdentityList = new(TrackingAreaIdentityList_r16)
		if err := ie.TrackingAreaIdentityList.Decode(d); err != nil {
			return err
		}
	case Qoe_AreaScope_r18_Plmn_IdentityList:
		ie.Plmn_IdentityList = new(PLMN_IdentityList2_r16)
		if err := ie.Plmn_IdentityList.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "Qoe-AreaScope-r18", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
