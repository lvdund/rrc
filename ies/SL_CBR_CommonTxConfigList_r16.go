package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_CBR_CommonTxConfigList_r16 struct {
	Sl_CBR_RangeConfigList_r16    []SL_CBR_LevelsConfig_r16   `lb:1,ub:maxCBR_Config_r16,optional`
	Sl_CBR_PSSCH_TxConfigList_r16 []SL_CBR_PSSCH_TxConfig_r16 `lb:1,ub:maxTxConfig_r16,optional`
}

func (ie *SL_CBR_CommonTxConfigList_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Sl_CBR_RangeConfigList_r16) > 0, len(ie.Sl_CBR_PSSCH_TxConfigList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Sl_CBR_RangeConfigList_r16) > 0 {
		tmp_Sl_CBR_RangeConfigList_r16 := utils.NewSequence[*SL_CBR_LevelsConfig_r16]([]*SL_CBR_LevelsConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxCBR_Config_r16}, false)
		for _, i := range ie.Sl_CBR_RangeConfigList_r16 {
			tmp_Sl_CBR_RangeConfigList_r16.Value = append(tmp_Sl_CBR_RangeConfigList_r16.Value, &i)
		}
		if err = tmp_Sl_CBR_RangeConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_CBR_RangeConfigList_r16", err)
		}
	}
	if len(ie.Sl_CBR_PSSCH_TxConfigList_r16) > 0 {
		tmp_Sl_CBR_PSSCH_TxConfigList_r16 := utils.NewSequence[*SL_CBR_PSSCH_TxConfig_r16]([]*SL_CBR_PSSCH_TxConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxTxConfig_r16}, false)
		for _, i := range ie.Sl_CBR_PSSCH_TxConfigList_r16 {
			tmp_Sl_CBR_PSSCH_TxConfigList_r16.Value = append(tmp_Sl_CBR_PSSCH_TxConfigList_r16.Value, &i)
		}
		if err = tmp_Sl_CBR_PSSCH_TxConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_CBR_PSSCH_TxConfigList_r16", err)
		}
	}
	return nil
}

func (ie *SL_CBR_CommonTxConfigList_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_CBR_RangeConfigList_r16Present bool
	if Sl_CBR_RangeConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CBR_PSSCH_TxConfigList_r16Present bool
	if Sl_CBR_PSSCH_TxConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_CBR_RangeConfigList_r16Present {
		tmp_Sl_CBR_RangeConfigList_r16 := utils.NewSequence[*SL_CBR_LevelsConfig_r16]([]*SL_CBR_LevelsConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxCBR_Config_r16}, false)
		fn_Sl_CBR_RangeConfigList_r16 := func() *SL_CBR_LevelsConfig_r16 {
			return new(SL_CBR_LevelsConfig_r16)
		}
		if err = tmp_Sl_CBR_RangeConfigList_r16.Decode(r, fn_Sl_CBR_RangeConfigList_r16); err != nil {
			return utils.WrapError("Decode Sl_CBR_RangeConfigList_r16", err)
		}
		ie.Sl_CBR_RangeConfigList_r16 = []SL_CBR_LevelsConfig_r16{}
		for _, i := range tmp_Sl_CBR_RangeConfigList_r16.Value {
			ie.Sl_CBR_RangeConfigList_r16 = append(ie.Sl_CBR_RangeConfigList_r16, *i)
		}
	}
	if Sl_CBR_PSSCH_TxConfigList_r16Present {
		tmp_Sl_CBR_PSSCH_TxConfigList_r16 := utils.NewSequence[*SL_CBR_PSSCH_TxConfig_r16]([]*SL_CBR_PSSCH_TxConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxTxConfig_r16}, false)
		fn_Sl_CBR_PSSCH_TxConfigList_r16 := func() *SL_CBR_PSSCH_TxConfig_r16 {
			return new(SL_CBR_PSSCH_TxConfig_r16)
		}
		if err = tmp_Sl_CBR_PSSCH_TxConfigList_r16.Decode(r, fn_Sl_CBR_PSSCH_TxConfigList_r16); err != nil {
			return utils.WrapError("Decode Sl_CBR_PSSCH_TxConfigList_r16", err)
		}
		ie.Sl_CBR_PSSCH_TxConfigList_r16 = []SL_CBR_PSSCH_TxConfig_r16{}
		for _, i := range tmp_Sl_CBR_PSSCH_TxConfigList_r16.Value {
			ie.Sl_CBR_PSSCH_TxConfigList_r16 = append(ie.Sl_CBR_PSSCH_TxConfigList_r16, *i)
		}
	}
	return nil
}
