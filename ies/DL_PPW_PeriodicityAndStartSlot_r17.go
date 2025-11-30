package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DL_PPW_PeriodicityAndStartSlot_r17_Choice_nothing uint64 = iota
	DL_PPW_PeriodicityAndStartSlot_r17_Choice_Scs15
	DL_PPW_PeriodicityAndStartSlot_r17_Choice_Scs30
	DL_PPW_PeriodicityAndStartSlot_r17_Choice_Scs60
	DL_PPW_PeriodicityAndStartSlot_r17_Choice_Scs120
)

type DL_PPW_PeriodicityAndStartSlot_r17 struct {
	Choice uint64
	Scs15  *DL_PPW_PeriodicityAndStartSlot_r17_scs15
	Scs30  *DL_PPW_PeriodicityAndStartSlot_r17_scs30
	Scs60  *DL_PPW_PeriodicityAndStartSlot_r17_scs60
	Scs120 *DL_PPW_PeriodicityAndStartSlot_r17_scs120
}

func (ie *DL_PPW_PeriodicityAndStartSlot_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_Scs15:
		if err = ie.Scs15.Encode(w); err != nil {
			err = utils.WrapError("Encode Scs15", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_Scs30:
		if err = ie.Scs30.Encode(w); err != nil {
			err = utils.WrapError("Encode Scs30", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_Scs60:
		if err = ie.Scs60.Encode(w); err != nil {
			err = utils.WrapError("Encode Scs60", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_Scs120:
		if err = ie.Scs120.Encode(w); err != nil {
			err = utils.WrapError("Encode Scs120", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DL_PPW_PeriodicityAndStartSlot_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_Scs15:
		ie.Scs15 = new(DL_PPW_PeriodicityAndStartSlot_r17_scs15)
		if err = ie.Scs15.Decode(r); err != nil {
			return utils.WrapError("Decode Scs15", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_Scs30:
		ie.Scs30 = new(DL_PPW_PeriodicityAndStartSlot_r17_scs30)
		if err = ie.Scs30.Decode(r); err != nil {
			return utils.WrapError("Decode Scs30", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_Scs60:
		ie.Scs60 = new(DL_PPW_PeriodicityAndStartSlot_r17_scs60)
		if err = ie.Scs60.Decode(r); err != nil {
			return utils.WrapError("Decode Scs60", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_Scs120:
		ie.Scs120 = new(DL_PPW_PeriodicityAndStartSlot_r17_scs120)
		if err = ie.Scs120.Decode(r); err != nil {
			return utils.WrapError("Decode Scs120", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
