package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_BWP_DiscPoolConfigCommon_r17 struct {
	Sl_DiscRxPool_r17         []SL_ResourcePool_r16       `lb:1,ub:maxNrofRXPool_r16,optional`
	Sl_DiscTxPoolSelected_r17 []SL_ResourcePoolConfig_r16 `lb:1,ub:maxNrofTXPool_r16,optional`
}

func (ie *SL_BWP_DiscPoolConfigCommon_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Sl_DiscRxPool_r17) > 0, len(ie.Sl_DiscTxPoolSelected_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Sl_DiscRxPool_r17) > 0 {
		tmp_Sl_DiscRxPool_r17 := utils.NewSequence[*SL_ResourcePool_r16]([]*SL_ResourcePool_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofRXPool_r16}, false)
		for _, i := range ie.Sl_DiscRxPool_r17 {
			tmp_Sl_DiscRxPool_r17.Value = append(tmp_Sl_DiscRxPool_r17.Value, &i)
		}
		if err = tmp_Sl_DiscRxPool_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DiscRxPool_r17", err)
		}
	}
	if len(ie.Sl_DiscTxPoolSelected_r17) > 0 {
		tmp_Sl_DiscTxPoolSelected_r17 := utils.NewSequence[*SL_ResourcePoolConfig_r16]([]*SL_ResourcePoolConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofTXPool_r16}, false)
		for _, i := range ie.Sl_DiscTxPoolSelected_r17 {
			tmp_Sl_DiscTxPoolSelected_r17.Value = append(tmp_Sl_DiscTxPoolSelected_r17.Value, &i)
		}
		if err = tmp_Sl_DiscTxPoolSelected_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DiscTxPoolSelected_r17", err)
		}
	}
	return nil
}

func (ie *SL_BWP_DiscPoolConfigCommon_r17) Decode(r *uper.UperReader) error {
	var err error
	var Sl_DiscRxPool_r17Present bool
	if Sl_DiscRxPool_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_DiscTxPoolSelected_r17Present bool
	if Sl_DiscTxPoolSelected_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_DiscRxPool_r17Present {
		tmp_Sl_DiscRxPool_r17 := utils.NewSequence[*SL_ResourcePool_r16]([]*SL_ResourcePool_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofRXPool_r16}, false)
		fn_Sl_DiscRxPool_r17 := func() *SL_ResourcePool_r16 {
			return new(SL_ResourcePool_r16)
		}
		if err = tmp_Sl_DiscRxPool_r17.Decode(r, fn_Sl_DiscRxPool_r17); err != nil {
			return utils.WrapError("Decode Sl_DiscRxPool_r17", err)
		}
		ie.Sl_DiscRxPool_r17 = []SL_ResourcePool_r16{}
		for _, i := range tmp_Sl_DiscRxPool_r17.Value {
			ie.Sl_DiscRxPool_r17 = append(ie.Sl_DiscRxPool_r17, *i)
		}
	}
	if Sl_DiscTxPoolSelected_r17Present {
		tmp_Sl_DiscTxPoolSelected_r17 := utils.NewSequence[*SL_ResourcePoolConfig_r16]([]*SL_ResourcePoolConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofTXPool_r16}, false)
		fn_Sl_DiscTxPoolSelected_r17 := func() *SL_ResourcePoolConfig_r16 {
			return new(SL_ResourcePoolConfig_r16)
		}
		if err = tmp_Sl_DiscTxPoolSelected_r17.Decode(r, fn_Sl_DiscTxPoolSelected_r17); err != nil {
			return utils.WrapError("Decode Sl_DiscTxPoolSelected_r17", err)
		}
		ie.Sl_DiscTxPoolSelected_r17 = []SL_ResourcePoolConfig_r16{}
		for _, i := range tmp_Sl_DiscTxPoolSelected_r17.Value {
			ie.Sl_DiscTxPoolSelected_r17 = append(ie.Sl_DiscTxPoolSelected_r17, *i)
		}
	}
	return nil
}
