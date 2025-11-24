package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RAN_NotificationAreaInfo_Choice_nothing uint64 = iota
	RAN_NotificationAreaInfo_Choice_CellList
	RAN_NotificationAreaInfo_Choice_Ran_AreaConfigList
)

type RAN_NotificationAreaInfo struct {
	Choice             uint64
	CellList           *PLMN_RAN_AreaCellList
	Ran_AreaConfigList *PLMN_RAN_AreaConfigList
}

func (ie *RAN_NotificationAreaInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RAN_NotificationAreaInfo_Choice_CellList:
		if err = ie.CellList.Encode(w); err != nil {
			err = utils.WrapError("Encode CellList", err)
		}
	case RAN_NotificationAreaInfo_Choice_Ran_AreaConfigList:
		if err = ie.Ran_AreaConfigList.Encode(w); err != nil {
			err = utils.WrapError("Encode Ran_AreaConfigList", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RAN_NotificationAreaInfo) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RAN_NotificationAreaInfo_Choice_CellList:
		ie.CellList = new(PLMN_RAN_AreaCellList)
		if err = ie.CellList.Decode(r); err != nil {
			return utils.WrapError("Decode CellList", err)
		}
	case RAN_NotificationAreaInfo_Choice_Ran_AreaConfigList:
		ie.Ran_AreaConfigList = new(PLMN_RAN_AreaConfigList)
		if err = ie.Ran_AreaConfigList.Decode(r); err != nil {
			return utils.WrapError("Decode Ran_AreaConfigList", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
