package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CrossCarrierSchedulingConfig_schedulingCellInfo_Choice_nothing uint64 = iota
	CrossCarrierSchedulingConfig_schedulingCellInfo_Choice_Own
	CrossCarrierSchedulingConfig_schedulingCellInfo_Choice_Other
)

type CrossCarrierSchedulingConfig_schedulingCellInfo struct {
	Choice uint64
	Own    *CrossCarrierSchedulingConfig_schedulingCellInfo_own
	Other  *CrossCarrierSchedulingConfig_schedulingCellInfo_other
}

func (ie *CrossCarrierSchedulingConfig_schedulingCellInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CrossCarrierSchedulingConfig_schedulingCellInfo_Choice_Own:
		if err = ie.Own.Encode(w); err != nil {
			err = utils.WrapError("Encode Own", err)
		}
	case CrossCarrierSchedulingConfig_schedulingCellInfo_Choice_Other:
		if err = ie.Other.Encode(w); err != nil {
			err = utils.WrapError("Encode Other", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CrossCarrierSchedulingConfig_schedulingCellInfo) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CrossCarrierSchedulingConfig_schedulingCellInfo_Choice_Own:
		ie.Own = new(CrossCarrierSchedulingConfig_schedulingCellInfo_own)
		if err = ie.Own.Decode(r); err != nil {
			return utils.WrapError("Decode Own", err)
		}
	case CrossCarrierSchedulingConfig_schedulingCellInfo_Choice_Other:
		ie.Other = new(CrossCarrierSchedulingConfig_schedulingCellInfo_other)
		if err = ie.Other.Decode(r); err != nil {
			return utils.WrapError("Decode Other", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
