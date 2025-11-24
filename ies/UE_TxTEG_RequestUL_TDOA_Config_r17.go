package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UE_TxTEG_RequestUL_TDOA_Config_r17_Choice_nothing uint64 = iota
	UE_TxTEG_RequestUL_TDOA_Config_r17_Choice_OneShot_r17
	UE_TxTEG_RequestUL_TDOA_Config_r17_Choice_PeriodicReporting_r17
)

type UE_TxTEG_RequestUL_TDOA_Config_r17 struct {
	Choice                uint64
	OneShot_r17           uper.NULL `madatory`
	PeriodicReporting_r17 *UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17
}

func (ie *UE_TxTEG_RequestUL_TDOA_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UE_TxTEG_RequestUL_TDOA_Config_r17_Choice_OneShot_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode OneShot_r17", err)
		}
	case UE_TxTEG_RequestUL_TDOA_Config_r17_Choice_PeriodicReporting_r17:
		if err = ie.PeriodicReporting_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode PeriodicReporting_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UE_TxTEG_RequestUL_TDOA_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UE_TxTEG_RequestUL_TDOA_Config_r17_Choice_OneShot_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode OneShot_r17", err)
		}
	case UE_TxTEG_RequestUL_TDOA_Config_r17_Choice_PeriodicReporting_r17:
		ie.PeriodicReporting_r17 = new(UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17)
		if err = ie.PeriodicReporting_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PeriodicReporting_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
