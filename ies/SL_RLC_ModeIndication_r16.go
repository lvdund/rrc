package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RLC_ModeIndication_r16 struct {
	Sl_Mode_r16         SL_RLC_ModeIndication_r16_sl_Mode_r16 `madatory`
	Sl_QoS_InfoList_r16 []SL_QoS_Info_r16                     `lb:1,ub:maxNrofSL_QFIsPerDest_r16,madatory`
}

func (ie *SL_RLC_ModeIndication_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Sl_Mode_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_Mode_r16", err)
	}
	tmp_Sl_QoS_InfoList_r16 := utils.NewSequence[*SL_QoS_Info_r16]([]*SL_QoS_Info_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_QFIsPerDest_r16}, false)
	for _, i := range ie.Sl_QoS_InfoList_r16 {
		tmp_Sl_QoS_InfoList_r16.Value = append(tmp_Sl_QoS_InfoList_r16.Value, &i)
	}
	if err = tmp_Sl_QoS_InfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_QoS_InfoList_r16", err)
	}
	return nil
}

func (ie *SL_RLC_ModeIndication_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Sl_Mode_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_Mode_r16", err)
	}
	tmp_Sl_QoS_InfoList_r16 := utils.NewSequence[*SL_QoS_Info_r16]([]*SL_QoS_Info_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_QFIsPerDest_r16}, false)
	fn_Sl_QoS_InfoList_r16 := func() *SL_QoS_Info_r16 {
		return new(SL_QoS_Info_r16)
	}
	if err = tmp_Sl_QoS_InfoList_r16.Decode(r, fn_Sl_QoS_InfoList_r16); err != nil {
		return utils.WrapError("Decode Sl_QoS_InfoList_r16", err)
	}
	ie.Sl_QoS_InfoList_r16 = []SL_QoS_Info_r16{}
	for _, i := range tmp_Sl_QoS_InfoList_r16.Value {
		ie.Sl_QoS_InfoList_r16 = append(ie.Sl_QoS_InfoList_r16, *i)
	}
	return nil
}
