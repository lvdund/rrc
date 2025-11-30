package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_FreqConfigCommon_r16 struct {
	Sl_SCS_SpecificCarrierList_r16 []SCS_SpecificCarrier                               `lb:1,ub:maxSCSs,madatory`
	Sl_AbsoluteFrequencyPointA_r16 ARFCN_ValueNR                                       `madatory`
	Sl_AbsoluteFrequencySSB_r16    *ARFCN_ValueNR                                      `optional`
	FrequencyShift7p5khzSL_r16     *SL_FreqConfigCommon_r16_frequencyShift7p5khzSL_r16 `optional`
	ValueN_r16                     int64                                               `lb:-1,ub:1,madatory`
	Sl_BWP_List_r16                []SL_BWP_ConfigCommon_r16                           `lb:1,ub:maxNrofSL_BWPs_r16,optional`
	Sl_SyncPriority_r16            *SL_FreqConfigCommon_r16_sl_SyncPriority_r16        `optional`
	Sl_NbAsSync_r16                *bool                                               `optional`
	Sl_SyncConfigList_r16          *SL_SyncConfigList_r16                              `optional`
}

func (ie *SL_FreqConfigCommon_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_AbsoluteFrequencySSB_r16 != nil, ie.FrequencyShift7p5khzSL_r16 != nil, len(ie.Sl_BWP_List_r16) > 0, ie.Sl_SyncPriority_r16 != nil, ie.Sl_NbAsSync_r16 != nil, ie.Sl_SyncConfigList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_Sl_SCS_SpecificCarrierList_r16 := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, aper.Constraint{Lb: 1, Ub: maxSCSs}, false)
	for _, i := range ie.Sl_SCS_SpecificCarrierList_r16 {
		tmp_Sl_SCS_SpecificCarrierList_r16.Value = append(tmp_Sl_SCS_SpecificCarrierList_r16.Value, &i)
	}
	if err = tmp_Sl_SCS_SpecificCarrierList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_SCS_SpecificCarrierList_r16", err)
	}
	if err = ie.Sl_AbsoluteFrequencyPointA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_AbsoluteFrequencyPointA_r16", err)
	}
	if ie.Sl_AbsoluteFrequencySSB_r16 != nil {
		if err = ie.Sl_AbsoluteFrequencySSB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_AbsoluteFrequencySSB_r16", err)
		}
	}
	if ie.FrequencyShift7p5khzSL_r16 != nil {
		if err = ie.FrequencyShift7p5khzSL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyShift7p5khzSL_r16", err)
		}
	}
	if err = w.WriteInteger(ie.ValueN_r16, &aper.Constraint{Lb: -1, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteInteger ValueN_r16", err)
	}
	if len(ie.Sl_BWP_List_r16) > 0 {
		tmp_Sl_BWP_List_r16 := utils.NewSequence[*SL_BWP_ConfigCommon_r16]([]*SL_BWP_ConfigCommon_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_BWPs_r16}, false)
		for _, i := range ie.Sl_BWP_List_r16 {
			tmp_Sl_BWP_List_r16.Value = append(tmp_Sl_BWP_List_r16.Value, &i)
		}
		if err = tmp_Sl_BWP_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_BWP_List_r16", err)
		}
	}
	if ie.Sl_SyncPriority_r16 != nil {
		if err = ie.Sl_SyncPriority_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SyncPriority_r16", err)
		}
	}
	if ie.Sl_NbAsSync_r16 != nil {
		if err = w.WriteBoolean(*ie.Sl_NbAsSync_r16); err != nil {
			return utils.WrapError("Encode Sl_NbAsSync_r16", err)
		}
	}
	if ie.Sl_SyncConfigList_r16 != nil {
		if err = ie.Sl_SyncConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SyncConfigList_r16", err)
		}
	}
	return nil
}

func (ie *SL_FreqConfigCommon_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_AbsoluteFrequencySSB_r16Present bool
	if Sl_AbsoluteFrequencySSB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var FrequencyShift7p5khzSL_r16Present bool
	if FrequencyShift7p5khzSL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_BWP_List_r16Present bool
	if Sl_BWP_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SyncPriority_r16Present bool
	if Sl_SyncPriority_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_NbAsSync_r16Present bool
	if Sl_NbAsSync_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SyncConfigList_r16Present bool
	if Sl_SyncConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_Sl_SCS_SpecificCarrierList_r16 := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, aper.Constraint{Lb: 1, Ub: maxSCSs}, false)
	fn_Sl_SCS_SpecificCarrierList_r16 := func() *SCS_SpecificCarrier {
		return new(SCS_SpecificCarrier)
	}
	if err = tmp_Sl_SCS_SpecificCarrierList_r16.Decode(r, fn_Sl_SCS_SpecificCarrierList_r16); err != nil {
		return utils.WrapError("Decode Sl_SCS_SpecificCarrierList_r16", err)
	}
	ie.Sl_SCS_SpecificCarrierList_r16 = []SCS_SpecificCarrier{}
	for _, i := range tmp_Sl_SCS_SpecificCarrierList_r16.Value {
		ie.Sl_SCS_SpecificCarrierList_r16 = append(ie.Sl_SCS_SpecificCarrierList_r16, *i)
	}
	if err = ie.Sl_AbsoluteFrequencyPointA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_AbsoluteFrequencyPointA_r16", err)
	}
	if Sl_AbsoluteFrequencySSB_r16Present {
		ie.Sl_AbsoluteFrequencySSB_r16 = new(ARFCN_ValueNR)
		if err = ie.Sl_AbsoluteFrequencySSB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_AbsoluteFrequencySSB_r16", err)
		}
	}
	if FrequencyShift7p5khzSL_r16Present {
		ie.FrequencyShift7p5khzSL_r16 = new(SL_FreqConfigCommon_r16_frequencyShift7p5khzSL_r16)
		if err = ie.FrequencyShift7p5khzSL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyShift7p5khzSL_r16", err)
		}
	}
	var tmp_int_ValueN_r16 int64
	if tmp_int_ValueN_r16, err = r.ReadInteger(&aper.Constraint{Lb: -1, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadInteger ValueN_r16", err)
	}
	ie.ValueN_r16 = tmp_int_ValueN_r16
	if Sl_BWP_List_r16Present {
		tmp_Sl_BWP_List_r16 := utils.NewSequence[*SL_BWP_ConfigCommon_r16]([]*SL_BWP_ConfigCommon_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_BWPs_r16}, false)
		fn_Sl_BWP_List_r16 := func() *SL_BWP_ConfigCommon_r16 {
			return new(SL_BWP_ConfigCommon_r16)
		}
		if err = tmp_Sl_BWP_List_r16.Decode(r, fn_Sl_BWP_List_r16); err != nil {
			return utils.WrapError("Decode Sl_BWP_List_r16", err)
		}
		ie.Sl_BWP_List_r16 = []SL_BWP_ConfigCommon_r16{}
		for _, i := range tmp_Sl_BWP_List_r16.Value {
			ie.Sl_BWP_List_r16 = append(ie.Sl_BWP_List_r16, *i)
		}
	}
	if Sl_SyncPriority_r16Present {
		ie.Sl_SyncPriority_r16 = new(SL_FreqConfigCommon_r16_sl_SyncPriority_r16)
		if err = ie.Sl_SyncPriority_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SyncPriority_r16", err)
		}
	}
	if Sl_NbAsSync_r16Present {
		var tmp_bool_Sl_NbAsSync_r16 bool
		if tmp_bool_Sl_NbAsSync_r16, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode Sl_NbAsSync_r16", err)
		}
		ie.Sl_NbAsSync_r16 = &tmp_bool_Sl_NbAsSync_r16
	}
	if Sl_SyncConfigList_r16Present {
		ie.Sl_SyncConfigList_r16 = new(SL_SyncConfigList_r16)
		if err = ie.Sl_SyncConfigList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SyncConfigList_r16", err)
		}
	}
	return nil
}
