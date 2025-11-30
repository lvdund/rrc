package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasTriggerQuantityUTRA_FDD_r16_Choice_nothing uint64 = iota
	MeasTriggerQuantityUTRA_FDD_r16_Choice_Utra_FDD_RSCP_r16
	MeasTriggerQuantityUTRA_FDD_r16_Choice_Utra_FDD_EcN0_r16
)

type MeasTriggerQuantityUTRA_FDD_r16 struct {
	Choice            uint64
	Utra_FDD_RSCP_r16 int64 `lb:-5,ub:91,madatory`
	Utra_FDD_EcN0_r16 int64 `lb:0,ub:49,madatory`
}

func (ie *MeasTriggerQuantityUTRA_FDD_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantityUTRA_FDD_r16_Choice_Utra_FDD_RSCP_r16:
		if err = w.WriteInteger(int64(ie.Utra_FDD_RSCP_r16), &aper.Constraint{Lb: -5, Ub: 91}, false); err != nil {
			err = utils.WrapError("Encode Utra_FDD_RSCP_r16", err)
		}
	case MeasTriggerQuantityUTRA_FDD_r16_Choice_Utra_FDD_EcN0_r16:
		if err = w.WriteInteger(int64(ie.Utra_FDD_EcN0_r16), &aper.Constraint{Lb: 0, Ub: 49}, false); err != nil {
			err = utils.WrapError("Encode Utra_FDD_EcN0_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasTriggerQuantityUTRA_FDD_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantityUTRA_FDD_r16_Choice_Utra_FDD_RSCP_r16:
		var tmp_int_Utra_FDD_RSCP_r16 int64
		if tmp_int_Utra_FDD_RSCP_r16, err = r.ReadInteger(&aper.Constraint{Lb: -5, Ub: 91}, false); err != nil {
			return utils.WrapError("Decode Utra_FDD_RSCP_r16", err)
		}
		ie.Utra_FDD_RSCP_r16 = tmp_int_Utra_FDD_RSCP_r16
	case MeasTriggerQuantityUTRA_FDD_r16_Choice_Utra_FDD_EcN0_r16:
		var tmp_int_Utra_FDD_EcN0_r16 int64
		if tmp_int_Utra_FDD_EcN0_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 49}, false); err != nil {
			return utils.WrapError("Decode Utra_FDD_EcN0_r16", err)
		}
		ie.Utra_FDD_EcN0_r16 = tmp_int_Utra_FDD_EcN0_r16
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
