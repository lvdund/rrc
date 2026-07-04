// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AreaConfig-r16 (line 26064).

var areaConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalIdList-r16"},
		{Name: "trackingAreaCodeList-r16"},
		{Name: "trackingAreaIdentityList-r16"},
	},
}

const (
	AreaConfig_r16_CellGlobalIdList_r16         = 0
	AreaConfig_r16_TrackingAreaCodeList_r16     = 1
	AreaConfig_r16_TrackingAreaIdentityList_r16 = 2
)

type AreaConfig_r16 struct {
	Choice                       int
	CellGlobalIdList_r16         *CellGlobalIdList_r16
	TrackingAreaCodeList_r16     *TrackingAreaCodeList_r16
	TrackingAreaIdentityList_r16 *TrackingAreaIdentityList_r16
}

func (ie *AreaConfig_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(areaConfigR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case AreaConfig_r16_CellGlobalIdList_r16:
		if err := ie.CellGlobalIdList_r16.Encode(e); err != nil {
			return err
		}
	case AreaConfig_r16_TrackingAreaCodeList_r16:
		if err := ie.TrackingAreaCodeList_r16.Encode(e); err != nil {
			return err
		}
	case AreaConfig_r16_TrackingAreaIdentityList_r16:
		if err := ie.TrackingAreaIdentityList_r16.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "AreaConfig-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *AreaConfig_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(areaConfigR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "AreaConfig-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case AreaConfig_r16_CellGlobalIdList_r16:
		ie.CellGlobalIdList_r16 = new(CellGlobalIdList_r16)
		if err := ie.CellGlobalIdList_r16.Decode(d); err != nil {
			return err
		}
	case AreaConfig_r16_TrackingAreaCodeList_r16:
		ie.TrackingAreaCodeList_r16 = new(TrackingAreaCodeList_r16)
		if err := ie.TrackingAreaCodeList_r16.Decode(d); err != nil {
			return err
		}
	case AreaConfig_r16_TrackingAreaIdentityList_r16:
		ie.TrackingAreaIdentityList_r16 = new(TrackingAreaIdentityList_r16)
		if err := ie.TrackingAreaIdentityList_r16.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "AreaConfig-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
