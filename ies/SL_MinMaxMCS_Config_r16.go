package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_MinMaxMCS_Config_r16 struct {
	Sl_MCS_Table_r16    SL_MinMaxMCS_Config_r16_sl_MCS_Table_r16 `madatory`
	Sl_MinMCS_PSSCH_r16 int64                                    `lb:0,ub:27,madatory`
	Sl_MaxMCS_PSSCH_r16 int64                                    `lb:0,ub:31,madatory`
}

func (ie *SL_MinMaxMCS_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Sl_MCS_Table_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_MCS_Table_r16", err)
	}
	if err = w.WriteInteger(ie.Sl_MinMCS_PSSCH_r16, &uper.Constraint{Lb: 0, Ub: 27}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_MinMCS_PSSCH_r16", err)
	}
	if err = w.WriteInteger(ie.Sl_MaxMCS_PSSCH_r16, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_MaxMCS_PSSCH_r16", err)
	}
	return nil
}

func (ie *SL_MinMaxMCS_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Sl_MCS_Table_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_MCS_Table_r16", err)
	}
	var tmp_int_Sl_MinMCS_PSSCH_r16 int64
	if tmp_int_Sl_MinMCS_PSSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 27}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_MinMCS_PSSCH_r16", err)
	}
	ie.Sl_MinMCS_PSSCH_r16 = tmp_int_Sl_MinMCS_PSSCH_r16
	var tmp_int_Sl_MaxMCS_PSSCH_r16 int64
	if tmp_int_Sl_MaxMCS_PSSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_MaxMCS_PSSCH_r16", err)
	}
	ie.Sl_MaxMCS_PSSCH_r16 = tmp_int_Sl_MaxMCS_PSSCH_r16
	return nil
}
