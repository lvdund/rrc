// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RAN-NotificationAreaInfo (line 1344).

var rANNotificationAreaInfoConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellList"},
		{Name: "ran-AreaConfigList"},
	},
}

const (
	RAN_NotificationAreaInfo_CellList           = 0
	RAN_NotificationAreaInfo_Ran_AreaConfigList = 1
)

type RAN_NotificationAreaInfo struct {
	Choice             int
	CellList           *PLMN_RAN_AreaCellList
	Ran_AreaConfigList *PLMN_RAN_AreaConfigList
}

func (ie *RAN_NotificationAreaInfo) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(rANNotificationAreaInfoConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case RAN_NotificationAreaInfo_CellList:
		if err := ie.CellList.Encode(e); err != nil {
			return err
		}
	case RAN_NotificationAreaInfo_Ran_AreaConfigList:
		if err := ie.Ran_AreaConfigList.Encode(e); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "RAN-NotificationAreaInfo",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *RAN_NotificationAreaInfo) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(rANNotificationAreaInfoConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "RAN-NotificationAreaInfo", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case RAN_NotificationAreaInfo_CellList:
		ie.CellList = new(PLMN_RAN_AreaCellList)
		if err := ie.CellList.Decode(d); err != nil {
			return err
		}
	case RAN_NotificationAreaInfo_Ran_AreaConfigList:
		ie.Ran_AreaConfigList = new(PLMN_RAN_AreaConfigList)
		if err := ie.Ran_AreaConfigList.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "RAN-NotificationAreaInfo", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
