package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BWP_UplinkDedicated_ul_TCI_StateList_r17_Choice_nothing uint64 = iota
	BWP_UplinkDedicated_ul_TCI_StateList_r17_Choice_Explicitlist
	BWP_UplinkDedicated_ul_TCI_StateList_r17_Choice_UnifiedTCI_StateRef_r17
)

type BWP_UplinkDedicated_ul_TCI_StateList_r17 struct {
	Choice                  uint64
	Explicitlist            *BWP_UplinkDedicated_ul_TCI_StateList_r17_explicitlist
	UnifiedTCI_StateRef_r17 *ServingCellAndBWP_Id_r17
}

func (ie *BWP_UplinkDedicated_ul_TCI_StateList_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BWP_UplinkDedicated_ul_TCI_StateList_r17_Choice_Explicitlist:
		if err = ie.Explicitlist.Encode(w); err != nil {
			err = utils.WrapError("Encode Explicitlist", err)
		}
	case BWP_UplinkDedicated_ul_TCI_StateList_r17_Choice_UnifiedTCI_StateRef_r17:
		if err = ie.UnifiedTCI_StateRef_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode UnifiedTCI_StateRef_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BWP_UplinkDedicated_ul_TCI_StateList_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BWP_UplinkDedicated_ul_TCI_StateList_r17_Choice_Explicitlist:
		ie.Explicitlist = new(BWP_UplinkDedicated_ul_TCI_StateList_r17_explicitlist)
		if err = ie.Explicitlist.Decode(r); err != nil {
			return utils.WrapError("Decode Explicitlist", err)
		}
	case BWP_UplinkDedicated_ul_TCI_StateList_r17_Choice_UnifiedTCI_StateRef_r17:
		ie.UnifiedTCI_StateRef_r17 = new(ServingCellAndBWP_Id_r17)
		if err = ie.UnifiedTCI_StateRef_r17.Decode(r); err != nil {
			return utils.WrapError("Decode UnifiedTCI_StateRef_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
