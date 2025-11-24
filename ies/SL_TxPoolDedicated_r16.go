package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_TxPoolDedicated_r16 struct {
	Sl_PoolToReleaseList_r16 []SL_ResourcePoolID_r16     `lb:1,ub:maxNrofTXPool_r16,optional`
	Sl_PoolToAddModList_r16  []SL_ResourcePoolConfig_r16 `lb:1,ub:maxNrofTXPool_r16,optional`
}

func (ie *SL_TxPoolDedicated_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Sl_PoolToReleaseList_r16) > 0, len(ie.Sl_PoolToAddModList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Sl_PoolToReleaseList_r16) > 0 {
		tmp_Sl_PoolToReleaseList_r16 := utils.NewSequence[*SL_ResourcePoolID_r16]([]*SL_ResourcePoolID_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofTXPool_r16}, false)
		for _, i := range ie.Sl_PoolToReleaseList_r16 {
			tmp_Sl_PoolToReleaseList_r16.Value = append(tmp_Sl_PoolToReleaseList_r16.Value, &i)
		}
		if err = tmp_Sl_PoolToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PoolToReleaseList_r16", err)
		}
	}
	if len(ie.Sl_PoolToAddModList_r16) > 0 {
		tmp_Sl_PoolToAddModList_r16 := utils.NewSequence[*SL_ResourcePoolConfig_r16]([]*SL_ResourcePoolConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofTXPool_r16}, false)
		for _, i := range ie.Sl_PoolToAddModList_r16 {
			tmp_Sl_PoolToAddModList_r16.Value = append(tmp_Sl_PoolToAddModList_r16.Value, &i)
		}
		if err = tmp_Sl_PoolToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PoolToAddModList_r16", err)
		}
	}
	return nil
}

func (ie *SL_TxPoolDedicated_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_PoolToReleaseList_r16Present bool
	if Sl_PoolToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PoolToAddModList_r16Present bool
	if Sl_PoolToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_PoolToReleaseList_r16Present {
		tmp_Sl_PoolToReleaseList_r16 := utils.NewSequence[*SL_ResourcePoolID_r16]([]*SL_ResourcePoolID_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofTXPool_r16}, false)
		fn_Sl_PoolToReleaseList_r16 := func() *SL_ResourcePoolID_r16 {
			return new(SL_ResourcePoolID_r16)
		}
		if err = tmp_Sl_PoolToReleaseList_r16.Decode(r, fn_Sl_PoolToReleaseList_r16); err != nil {
			return utils.WrapError("Decode Sl_PoolToReleaseList_r16", err)
		}
		ie.Sl_PoolToReleaseList_r16 = []SL_ResourcePoolID_r16{}
		for _, i := range tmp_Sl_PoolToReleaseList_r16.Value {
			ie.Sl_PoolToReleaseList_r16 = append(ie.Sl_PoolToReleaseList_r16, *i)
		}
	}
	if Sl_PoolToAddModList_r16Present {
		tmp_Sl_PoolToAddModList_r16 := utils.NewSequence[*SL_ResourcePoolConfig_r16]([]*SL_ResourcePoolConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofTXPool_r16}, false)
		fn_Sl_PoolToAddModList_r16 := func() *SL_ResourcePoolConfig_r16 {
			return new(SL_ResourcePoolConfig_r16)
		}
		if err = tmp_Sl_PoolToAddModList_r16.Decode(r, fn_Sl_PoolToAddModList_r16); err != nil {
			return utils.WrapError("Decode Sl_PoolToAddModList_r16", err)
		}
		ie.Sl_PoolToAddModList_r16 = []SL_ResourcePoolConfig_r16{}
		for _, i := range tmp_Sl_PoolToAddModList_r16.Value {
			ie.Sl_PoolToAddModList_r16 = append(ie.Sl_PoolToAddModList_r16, *i)
		}
	}
	return nil
}
