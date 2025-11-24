package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_DRX_GC_BC_QoS_r17 struct {
	Sl_DRX_GC_BC_MappedQoS_FlowList_r17 []SL_QoS_Profile_r16                                  `lb:1,ub:maxNrofSL_QFIs_r16,optional`
	Sl_DRX_GC_BC_OnDurationTimer_r17    SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_BC_OnDurationTimer_r17 `lb:1,ub:31,madatory`
	Sl_DRX_GC_InactivityTimer_r17       SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17    `madatory`
	Sl_DRX_GC_BC_Cycle_r17              SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_BC_Cycle_r17           `madatory`
}

func (ie *SL_DRX_GC_BC_QoS_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Sl_DRX_GC_BC_MappedQoS_FlowList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Sl_DRX_GC_BC_MappedQoS_FlowList_r17) > 0 {
		tmp_Sl_DRX_GC_BC_MappedQoS_FlowList_r17 := utils.NewSequence[*SL_QoS_Profile_r16]([]*SL_QoS_Profile_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_QFIs_r16}, false)
		for _, i := range ie.Sl_DRX_GC_BC_MappedQoS_FlowList_r17 {
			tmp_Sl_DRX_GC_BC_MappedQoS_FlowList_r17.Value = append(tmp_Sl_DRX_GC_BC_MappedQoS_FlowList_r17.Value, &i)
		}
		if err = tmp_Sl_DRX_GC_BC_MappedQoS_FlowList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DRX_GC_BC_MappedQoS_FlowList_r17", err)
		}
	}
	if err = ie.Sl_DRX_GC_BC_OnDurationTimer_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_DRX_GC_BC_OnDurationTimer_r17", err)
	}
	if err = ie.Sl_DRX_GC_InactivityTimer_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_DRX_GC_InactivityTimer_r17", err)
	}
	if err = ie.Sl_DRX_GC_BC_Cycle_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_DRX_GC_BC_Cycle_r17", err)
	}
	return nil
}

func (ie *SL_DRX_GC_BC_QoS_r17) Decode(r *uper.UperReader) error {
	var err error
	var Sl_DRX_GC_BC_MappedQoS_FlowList_r17Present bool
	if Sl_DRX_GC_BC_MappedQoS_FlowList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_DRX_GC_BC_MappedQoS_FlowList_r17Present {
		tmp_Sl_DRX_GC_BC_MappedQoS_FlowList_r17 := utils.NewSequence[*SL_QoS_Profile_r16]([]*SL_QoS_Profile_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_QFIs_r16}, false)
		fn_Sl_DRX_GC_BC_MappedQoS_FlowList_r17 := func() *SL_QoS_Profile_r16 {
			return new(SL_QoS_Profile_r16)
		}
		if err = tmp_Sl_DRX_GC_BC_MappedQoS_FlowList_r17.Decode(r, fn_Sl_DRX_GC_BC_MappedQoS_FlowList_r17); err != nil {
			return utils.WrapError("Decode Sl_DRX_GC_BC_MappedQoS_FlowList_r17", err)
		}
		ie.Sl_DRX_GC_BC_MappedQoS_FlowList_r17 = []SL_QoS_Profile_r16{}
		for _, i := range tmp_Sl_DRX_GC_BC_MappedQoS_FlowList_r17.Value {
			ie.Sl_DRX_GC_BC_MappedQoS_FlowList_r17 = append(ie.Sl_DRX_GC_BC_MappedQoS_FlowList_r17, *i)
		}
	}
	if err = ie.Sl_DRX_GC_BC_OnDurationTimer_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_DRX_GC_BC_OnDurationTimer_r17", err)
	}
	if err = ie.Sl_DRX_GC_InactivityTimer_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_DRX_GC_InactivityTimer_r17", err)
	}
	if err = ie.Sl_DRX_GC_BC_Cycle_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_DRX_GC_BC_Cycle_r17", err)
	}
	return nil
}
