package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RxInterestedGC_BC_Dest_r17 struct {
	Sl_RxInterestedQoS_InfoList_r17 []SL_QoS_Info_r16          `lb:1,ub:maxNrofSL_QFIsPerDest_r16,madatory`
	Sl_DestinationIdentity_r16      SL_DestinationIdentity_r16 `madatory`
}

func (ie *SL_RxInterestedGC_BC_Dest_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp_Sl_RxInterestedQoS_InfoList_r17 := utils.NewSequence[*SL_QoS_Info_r16]([]*SL_QoS_Info_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_QFIsPerDest_r16}, false)
	for _, i := range ie.Sl_RxInterestedQoS_InfoList_r17 {
		tmp_Sl_RxInterestedQoS_InfoList_r17.Value = append(tmp_Sl_RxInterestedQoS_InfoList_r17.Value, &i)
	}
	if err = tmp_Sl_RxInterestedQoS_InfoList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_RxInterestedQoS_InfoList_r17", err)
	}
	if err = ie.Sl_DestinationIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_DestinationIdentity_r16", err)
	}
	return nil
}

func (ie *SL_RxInterestedGC_BC_Dest_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp_Sl_RxInterestedQoS_InfoList_r17 := utils.NewSequence[*SL_QoS_Info_r16]([]*SL_QoS_Info_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_QFIsPerDest_r16}, false)
	fn_Sl_RxInterestedQoS_InfoList_r17 := func() *SL_QoS_Info_r16 {
		return new(SL_QoS_Info_r16)
	}
	if err = tmp_Sl_RxInterestedQoS_InfoList_r17.Decode(r, fn_Sl_RxInterestedQoS_InfoList_r17); err != nil {
		return utils.WrapError("Decode Sl_RxInterestedQoS_InfoList_r17", err)
	}
	ie.Sl_RxInterestedQoS_InfoList_r17 = []SL_QoS_Info_r16{}
	for _, i := range tmp_Sl_RxInterestedQoS_InfoList_r17.Value {
		ie.Sl_RxInterestedQoS_InfoList_r17 = append(ie.Sl_RxInterestedQoS_InfoList_r17, *i)
	}
	if err = ie.Sl_DestinationIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_DestinationIdentity_r16", err)
	}
	return nil
}
