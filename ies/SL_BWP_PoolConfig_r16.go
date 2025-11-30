package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_BWP_PoolConfig_r16 struct {
	Sl_RxPool_r16               []SL_ResourcePool_r16      `lb:1,ub:maxNrofRXPool_r16,optional`
	Sl_TxPoolSelectedNormal_r16 *SL_TxPoolDedicated_r16    `optional`
	Sl_TxPoolScheduling_r16     *SL_TxPoolDedicated_r16    `optional`
	Sl_TxPoolExceptional_r16    *SL_ResourcePoolConfig_r16 `optional`
}

func (ie *SL_BWP_PoolConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Sl_RxPool_r16) > 0, ie.Sl_TxPoolSelectedNormal_r16 != nil, ie.Sl_TxPoolScheduling_r16 != nil, ie.Sl_TxPoolExceptional_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Sl_RxPool_r16) > 0 {
		tmp_Sl_RxPool_r16 := utils.NewSequence[*SL_ResourcePool_r16]([]*SL_ResourcePool_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofRXPool_r16}, false)
		for _, i := range ie.Sl_RxPool_r16 {
			tmp_Sl_RxPool_r16.Value = append(tmp_Sl_RxPool_r16.Value, &i)
		}
		if err = tmp_Sl_RxPool_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RxPool_r16", err)
		}
	}
	if ie.Sl_TxPoolSelectedNormal_r16 != nil {
		if err = ie.Sl_TxPoolSelectedNormal_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TxPoolSelectedNormal_r16", err)
		}
	}
	if ie.Sl_TxPoolScheduling_r16 != nil {
		if err = ie.Sl_TxPoolScheduling_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TxPoolScheduling_r16", err)
		}
	}
	if ie.Sl_TxPoolExceptional_r16 != nil {
		if err = ie.Sl_TxPoolExceptional_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TxPoolExceptional_r16", err)
		}
	}
	return nil
}

func (ie *SL_BWP_PoolConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_RxPool_r16Present bool
	if Sl_RxPool_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TxPoolSelectedNormal_r16Present bool
	if Sl_TxPoolSelectedNormal_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TxPoolScheduling_r16Present bool
	if Sl_TxPoolScheduling_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TxPoolExceptional_r16Present bool
	if Sl_TxPoolExceptional_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_RxPool_r16Present {
		tmp_Sl_RxPool_r16 := utils.NewSequence[*SL_ResourcePool_r16]([]*SL_ResourcePool_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofRXPool_r16}, false)
		fn_Sl_RxPool_r16 := func() *SL_ResourcePool_r16 {
			return new(SL_ResourcePool_r16)
		}
		if err = tmp_Sl_RxPool_r16.Decode(r, fn_Sl_RxPool_r16); err != nil {
			return utils.WrapError("Decode Sl_RxPool_r16", err)
		}
		ie.Sl_RxPool_r16 = []SL_ResourcePool_r16{}
		for _, i := range tmp_Sl_RxPool_r16.Value {
			ie.Sl_RxPool_r16 = append(ie.Sl_RxPool_r16, *i)
		}
	}
	if Sl_TxPoolSelectedNormal_r16Present {
		ie.Sl_TxPoolSelectedNormal_r16 = new(SL_TxPoolDedicated_r16)
		if err = ie.Sl_TxPoolSelectedNormal_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TxPoolSelectedNormal_r16", err)
		}
	}
	if Sl_TxPoolScheduling_r16Present {
		ie.Sl_TxPoolScheduling_r16 = new(SL_TxPoolDedicated_r16)
		if err = ie.Sl_TxPoolScheduling_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TxPoolScheduling_r16", err)
		}
	}
	if Sl_TxPoolExceptional_r16Present {
		ie.Sl_TxPoolExceptional_r16 = new(SL_ResourcePoolConfig_r16)
		if err = ie.Sl_TxPoolExceptional_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TxPoolExceptional_r16", err)
		}
	}
	return nil
}
