package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PSSCH_TxParameters_r16 struct {
	Sl_MinMCS_PSSCH_r16          int64           `lb:0,ub:27,madatory`
	Sl_MaxMCS_PSSCH_r16          int64           `lb:0,ub:31,madatory`
	Sl_MinSubChannelNumPSSCH_r16 int64           `lb:1,ub:27,madatory`
	Sl_MaxSubchannelNumPSSCH_r16 int64           `lb:1,ub:27,madatory`
	Sl_MaxTxTransNumPSSCH_r16    int64           `lb:1,ub:32,madatory`
	Sl_MaxTxPower_r16            *SL_TxPower_r16 `optional`
}

func (ie *SL_PSSCH_TxParameters_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_MaxTxPower_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.Sl_MinMCS_PSSCH_r16, &uper.Constraint{Lb: 0, Ub: 27}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_MinMCS_PSSCH_r16", err)
	}
	if err = w.WriteInteger(ie.Sl_MaxMCS_PSSCH_r16, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_MaxMCS_PSSCH_r16", err)
	}
	if err = w.WriteInteger(ie.Sl_MinSubChannelNumPSSCH_r16, &uper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_MinSubChannelNumPSSCH_r16", err)
	}
	if err = w.WriteInteger(ie.Sl_MaxSubchannelNumPSSCH_r16, &uper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_MaxSubchannelNumPSSCH_r16", err)
	}
	if err = w.WriteInteger(ie.Sl_MaxTxTransNumPSSCH_r16, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_MaxTxTransNumPSSCH_r16", err)
	}
	if ie.Sl_MaxTxPower_r16 != nil {
		if err = ie.Sl_MaxTxPower_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MaxTxPower_r16", err)
		}
	}
	return nil
}

func (ie *SL_PSSCH_TxParameters_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_MaxTxPower_r16Present bool
	if Sl_MaxTxPower_r16Present, err = r.ReadBool(); err != nil {
		return err
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
	var tmp_int_Sl_MinSubChannelNumPSSCH_r16 int64
	if tmp_int_Sl_MinSubChannelNumPSSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_MinSubChannelNumPSSCH_r16", err)
	}
	ie.Sl_MinSubChannelNumPSSCH_r16 = tmp_int_Sl_MinSubChannelNumPSSCH_r16
	var tmp_int_Sl_MaxSubchannelNumPSSCH_r16 int64
	if tmp_int_Sl_MaxSubchannelNumPSSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_MaxSubchannelNumPSSCH_r16", err)
	}
	ie.Sl_MaxSubchannelNumPSSCH_r16 = tmp_int_Sl_MaxSubchannelNumPSSCH_r16
	var tmp_int_Sl_MaxTxTransNumPSSCH_r16 int64
	if tmp_int_Sl_MaxTxTransNumPSSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_MaxTxTransNumPSSCH_r16", err)
	}
	ie.Sl_MaxTxTransNumPSSCH_r16 = tmp_int_Sl_MaxTxTransNumPSSCH_r16
	if Sl_MaxTxPower_r16Present {
		ie.Sl_MaxTxPower_r16 = new(SL_TxPower_r16)
		if err = ie.Sl_MaxTxPower_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MaxTxPower_r16", err)
		}
	}
	return nil
}
