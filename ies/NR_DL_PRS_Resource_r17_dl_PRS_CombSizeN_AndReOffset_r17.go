package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_nothing uint64 = iota
	NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_N2_r17
	NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_N4_r17
	NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_N6_r17
	NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_N12_r17
)

type NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17 struct {
	Choice  uint64
	N2_r17  int64 `lb:0,ub:1,madatory`
	N4_r17  int64 `lb:0,ub:3,madatory`
	N6_r17  int64 `lb:0,ub:5,madatory`
	N12_r17 int64 `lb:0,ub:11,madatory`
}

func (ie *NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_N2_r17:
		if err = w.WriteInteger(int64(ie.N2_r17), &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			err = utils.WrapError("Encode N2_r17", err)
		}
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_N4_r17:
		if err = w.WriteInteger(int64(ie.N4_r17), &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode N4_r17", err)
		}
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_N6_r17:
		if err = w.WriteInteger(int64(ie.N6_r17), &uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
			err = utils.WrapError("Encode N6_r17", err)
		}
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_N12_r17:
		if err = w.WriteInteger(int64(ie.N12_r17), &uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
			err = utils.WrapError("Encode N12_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_N2_r17:
		var tmp_int_N2_r17 int64
		if tmp_int_N2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode N2_r17", err)
		}
		ie.N2_r17 = tmp_int_N2_r17
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_N4_r17:
		var tmp_int_N4_r17 int64
		if tmp_int_N4_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode N4_r17", err)
		}
		ie.N4_r17 = tmp_int_N4_r17
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_N6_r17:
		var tmp_int_N6_r17 int64
		if tmp_int_N6_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
			return utils.WrapError("Decode N6_r17", err)
		}
		ie.N6_r17 = tmp_int_N6_r17
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_N12_r17:
		var tmp_int_N12_r17 int64
		if tmp_int_N12_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
			return utils.WrapError("Decode N12_r17", err)
		}
		ie.N12_r17 = tmp_int_N12_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
