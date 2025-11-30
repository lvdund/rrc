package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_MappedQoS_FlowsListDedicated_r16 struct {
	Sl_MappedQoS_FlowsToAddList_r16     []SL_QoS_FlowIdentity_r16 `lb:1,ub:maxNrofSL_QFIs_r16,optional`
	Sl_MappedQoS_FlowsToReleaseList_r16 []SL_QoS_FlowIdentity_r16 `lb:1,ub:maxNrofSL_QFIs_r16,optional`
}

func (ie *SL_MappedQoS_FlowsListDedicated_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Sl_MappedQoS_FlowsToAddList_r16) > 0, len(ie.Sl_MappedQoS_FlowsToReleaseList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Sl_MappedQoS_FlowsToAddList_r16) > 0 {
		tmp_Sl_MappedQoS_FlowsToAddList_r16 := utils.NewSequence[*SL_QoS_FlowIdentity_r16]([]*SL_QoS_FlowIdentity_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_QFIs_r16}, false)
		for _, i := range ie.Sl_MappedQoS_FlowsToAddList_r16 {
			tmp_Sl_MappedQoS_FlowsToAddList_r16.Value = append(tmp_Sl_MappedQoS_FlowsToAddList_r16.Value, &i)
		}
		if err = tmp_Sl_MappedQoS_FlowsToAddList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MappedQoS_FlowsToAddList_r16", err)
		}
	}
	if len(ie.Sl_MappedQoS_FlowsToReleaseList_r16) > 0 {
		tmp_Sl_MappedQoS_FlowsToReleaseList_r16 := utils.NewSequence[*SL_QoS_FlowIdentity_r16]([]*SL_QoS_FlowIdentity_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_QFIs_r16}, false)
		for _, i := range ie.Sl_MappedQoS_FlowsToReleaseList_r16 {
			tmp_Sl_MappedQoS_FlowsToReleaseList_r16.Value = append(tmp_Sl_MappedQoS_FlowsToReleaseList_r16.Value, &i)
		}
		if err = tmp_Sl_MappedQoS_FlowsToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MappedQoS_FlowsToReleaseList_r16", err)
		}
	}
	return nil
}

func (ie *SL_MappedQoS_FlowsListDedicated_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_MappedQoS_FlowsToAddList_r16Present bool
	if Sl_MappedQoS_FlowsToAddList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MappedQoS_FlowsToReleaseList_r16Present bool
	if Sl_MappedQoS_FlowsToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_MappedQoS_FlowsToAddList_r16Present {
		tmp_Sl_MappedQoS_FlowsToAddList_r16 := utils.NewSequence[*SL_QoS_FlowIdentity_r16]([]*SL_QoS_FlowIdentity_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_QFIs_r16}, false)
		fn_Sl_MappedQoS_FlowsToAddList_r16 := func() *SL_QoS_FlowIdentity_r16 {
			return new(SL_QoS_FlowIdentity_r16)
		}
		if err = tmp_Sl_MappedQoS_FlowsToAddList_r16.Decode(r, fn_Sl_MappedQoS_FlowsToAddList_r16); err != nil {
			return utils.WrapError("Decode Sl_MappedQoS_FlowsToAddList_r16", err)
		}
		ie.Sl_MappedQoS_FlowsToAddList_r16 = []SL_QoS_FlowIdentity_r16{}
		for _, i := range tmp_Sl_MappedQoS_FlowsToAddList_r16.Value {
			ie.Sl_MappedQoS_FlowsToAddList_r16 = append(ie.Sl_MappedQoS_FlowsToAddList_r16, *i)
		}
	}
	if Sl_MappedQoS_FlowsToReleaseList_r16Present {
		tmp_Sl_MappedQoS_FlowsToReleaseList_r16 := utils.NewSequence[*SL_QoS_FlowIdentity_r16]([]*SL_QoS_FlowIdentity_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_QFIs_r16}, false)
		fn_Sl_MappedQoS_FlowsToReleaseList_r16 := func() *SL_QoS_FlowIdentity_r16 {
			return new(SL_QoS_FlowIdentity_r16)
		}
		if err = tmp_Sl_MappedQoS_FlowsToReleaseList_r16.Decode(r, fn_Sl_MappedQoS_FlowsToReleaseList_r16); err != nil {
			return utils.WrapError("Decode Sl_MappedQoS_FlowsToReleaseList_r16", err)
		}
		ie.Sl_MappedQoS_FlowsToReleaseList_r16 = []SL_QoS_FlowIdentity_r16{}
		for _, i := range tmp_Sl_MappedQoS_FlowsToReleaseList_r16.Value {
			ie.Sl_MappedQoS_FlowsToReleaseList_r16 = append(ie.Sl_MappedQoS_FlowsToReleaseList_r16, *i)
		}
	}
	return nil
}
