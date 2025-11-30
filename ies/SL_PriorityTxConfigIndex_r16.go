package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_PriorityTxConfigIndex_r16 struct {
	Sl_PriorityThreshold_r16    *int64                 `lb:1,ub:8,optional`
	Sl_DefaultTxConfigIndex_r16 *int64                 `lb:0,ub:maxCBR_Level_1_r16,optional`
	Sl_CBR_ConfigIndex_r16      *int64                 `lb:0,ub:maxCBR_Config_1_r16,optional`
	Sl_Tx_ConfigIndexList_r16   []SL_TxConfigIndex_r16 `lb:1,ub:maxCBR_Level_r16,optional`
}

func (ie *SL_PriorityTxConfigIndex_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_PriorityThreshold_r16 != nil, ie.Sl_DefaultTxConfigIndex_r16 != nil, ie.Sl_CBR_ConfigIndex_r16 != nil, len(ie.Sl_Tx_ConfigIndexList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_PriorityThreshold_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_PriorityThreshold_r16, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Sl_PriorityThreshold_r16", err)
		}
	}
	if ie.Sl_DefaultTxConfigIndex_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_DefaultTxConfigIndex_r16, &aper.Constraint{Lb: 0, Ub: maxCBR_Level_1_r16}, false); err != nil {
			return utils.WrapError("Encode Sl_DefaultTxConfigIndex_r16", err)
		}
	}
	if ie.Sl_CBR_ConfigIndex_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_CBR_ConfigIndex_r16, &aper.Constraint{Lb: 0, Ub: maxCBR_Config_1_r16}, false); err != nil {
			return utils.WrapError("Encode Sl_CBR_ConfigIndex_r16", err)
		}
	}
	if len(ie.Sl_Tx_ConfigIndexList_r16) > 0 {
		tmp_Sl_Tx_ConfigIndexList_r16 := utils.NewSequence[*SL_TxConfigIndex_r16]([]*SL_TxConfigIndex_r16{}, aper.Constraint{Lb: 1, Ub: maxCBR_Level_r16}, false)
		for _, i := range ie.Sl_Tx_ConfigIndexList_r16 {
			tmp_Sl_Tx_ConfigIndexList_r16.Value = append(tmp_Sl_Tx_ConfigIndexList_r16.Value, &i)
		}
		if err = tmp_Sl_Tx_ConfigIndexList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_Tx_ConfigIndexList_r16", err)
		}
	}
	return nil
}

func (ie *SL_PriorityTxConfigIndex_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_PriorityThreshold_r16Present bool
	if Sl_PriorityThreshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_DefaultTxConfigIndex_r16Present bool
	if Sl_DefaultTxConfigIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CBR_ConfigIndex_r16Present bool
	if Sl_CBR_ConfigIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_Tx_ConfigIndexList_r16Present bool
	if Sl_Tx_ConfigIndexList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_PriorityThreshold_r16Present {
		var tmp_int_Sl_PriorityThreshold_r16 int64
		if tmp_int_Sl_PriorityThreshold_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Sl_PriorityThreshold_r16", err)
		}
		ie.Sl_PriorityThreshold_r16 = &tmp_int_Sl_PriorityThreshold_r16
	}
	if Sl_DefaultTxConfigIndex_r16Present {
		var tmp_int_Sl_DefaultTxConfigIndex_r16 int64
		if tmp_int_Sl_DefaultTxConfigIndex_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxCBR_Level_1_r16}, false); err != nil {
			return utils.WrapError("Decode Sl_DefaultTxConfigIndex_r16", err)
		}
		ie.Sl_DefaultTxConfigIndex_r16 = &tmp_int_Sl_DefaultTxConfigIndex_r16
	}
	if Sl_CBR_ConfigIndex_r16Present {
		var tmp_int_Sl_CBR_ConfigIndex_r16 int64
		if tmp_int_Sl_CBR_ConfigIndex_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxCBR_Config_1_r16}, false); err != nil {
			return utils.WrapError("Decode Sl_CBR_ConfigIndex_r16", err)
		}
		ie.Sl_CBR_ConfigIndex_r16 = &tmp_int_Sl_CBR_ConfigIndex_r16
	}
	if Sl_Tx_ConfigIndexList_r16Present {
		tmp_Sl_Tx_ConfigIndexList_r16 := utils.NewSequence[*SL_TxConfigIndex_r16]([]*SL_TxConfigIndex_r16{}, aper.Constraint{Lb: 1, Ub: maxCBR_Level_r16}, false)
		fn_Sl_Tx_ConfigIndexList_r16 := func() *SL_TxConfigIndex_r16 {
			return new(SL_TxConfigIndex_r16)
		}
		if err = tmp_Sl_Tx_ConfigIndexList_r16.Decode(r, fn_Sl_Tx_ConfigIndexList_r16); err != nil {
			return utils.WrapError("Decode Sl_Tx_ConfigIndexList_r16", err)
		}
		ie.Sl_Tx_ConfigIndexList_r16 = []SL_TxConfigIndex_r16{}
		for _, i := range tmp_Sl_Tx_ConfigIndexList_r16.Value {
			ie.Sl_Tx_ConfigIndexList_r16 = append(ie.Sl_Tx_ConfigIndexList_r16, *i)
		}
	}
	return nil
}
