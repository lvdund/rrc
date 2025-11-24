package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_LogicalChannelConfig_r16 struct {
	Sl_Priority_r16                           int64                                                           `lb:1,ub:8,madatory`
	Sl_PrioritisedBitRate_r16                 SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16           `madatory`
	Sl_BucketSizeDuration_r16                 SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16           `madatory`
	Sl_ConfiguredGrantType1Allowed_r16        *SL_LogicalChannelConfig_r16_sl_ConfiguredGrantType1Allowed_r16 `optional`
	Sl_HARQ_FeedbackEnabled_r16               *SL_LogicalChannelConfig_r16_sl_HARQ_FeedbackEnabled_r16        `optional`
	Sl_AllowedCG_List_r16                     []SL_ConfigIndexCG_r16                                          `lb:0,ub:maxNrofCG_SL_1_r16,optional`
	Sl_AllowedSCS_List_r16                    []SubcarrierSpacing                                             `lb:1,ub:maxSCSs,optional`
	Sl_MaxPUSCH_Duration_r16                  *SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16           `optional`
	Sl_LogicalChannelGroup_r16                *int64                                                          `lb:0,ub:maxLCG_ID,optional`
	Sl_SchedulingRequestId_r16                *SchedulingRequestId                                            `optional`
	Sl_LogicalChannelSR_DelayTimerApplied_r16 *bool                                                           `optional`
}

func (ie *SL_LogicalChannelConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_ConfiguredGrantType1Allowed_r16 != nil, ie.Sl_HARQ_FeedbackEnabled_r16 != nil, len(ie.Sl_AllowedCG_List_r16) > 0, len(ie.Sl_AllowedSCS_List_r16) > 0, ie.Sl_MaxPUSCH_Duration_r16 != nil, ie.Sl_LogicalChannelGroup_r16 != nil, ie.Sl_SchedulingRequestId_r16 != nil, ie.Sl_LogicalChannelSR_DelayTimerApplied_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.Sl_Priority_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_Priority_r16", err)
	}
	if err = ie.Sl_PrioritisedBitRate_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_PrioritisedBitRate_r16", err)
	}
	if err = ie.Sl_BucketSizeDuration_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_BucketSizeDuration_r16", err)
	}
	if ie.Sl_ConfiguredGrantType1Allowed_r16 != nil {
		if err = ie.Sl_ConfiguredGrantType1Allowed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ConfiguredGrantType1Allowed_r16", err)
		}
	}
	if ie.Sl_HARQ_FeedbackEnabled_r16 != nil {
		if err = ie.Sl_HARQ_FeedbackEnabled_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_HARQ_FeedbackEnabled_r16", err)
		}
	}
	if len(ie.Sl_AllowedCG_List_r16) > 0 {
		tmp_Sl_AllowedCG_List_r16 := utils.NewSequence[*SL_ConfigIndexCG_r16]([]*SL_ConfigIndexCG_r16{}, uper.Constraint{Lb: 0, Ub: maxNrofCG_SL_1_r16}, false)
		for _, i := range ie.Sl_AllowedCG_List_r16 {
			tmp_Sl_AllowedCG_List_r16.Value = append(tmp_Sl_AllowedCG_List_r16.Value, &i)
		}
		if err = tmp_Sl_AllowedCG_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_AllowedCG_List_r16", err)
		}
	}
	if len(ie.Sl_AllowedSCS_List_r16) > 0 {
		tmp_Sl_AllowedSCS_List_r16 := utils.NewSequence[*SubcarrierSpacing]([]*SubcarrierSpacing{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
		for _, i := range ie.Sl_AllowedSCS_List_r16 {
			tmp_Sl_AllowedSCS_List_r16.Value = append(tmp_Sl_AllowedSCS_List_r16.Value, &i)
		}
		if err = tmp_Sl_AllowedSCS_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_AllowedSCS_List_r16", err)
		}
	}
	if ie.Sl_MaxPUSCH_Duration_r16 != nil {
		if err = ie.Sl_MaxPUSCH_Duration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MaxPUSCH_Duration_r16", err)
		}
	}
	if ie.Sl_LogicalChannelGroup_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_LogicalChannelGroup_r16, &uper.Constraint{Lb: 0, Ub: maxLCG_ID}, false); err != nil {
			return utils.WrapError("Encode Sl_LogicalChannelGroup_r16", err)
		}
	}
	if ie.Sl_SchedulingRequestId_r16 != nil {
		if err = ie.Sl_SchedulingRequestId_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SchedulingRequestId_r16", err)
		}
	}
	if ie.Sl_LogicalChannelSR_DelayTimerApplied_r16 != nil {
		if err = w.WriteBoolean(*ie.Sl_LogicalChannelSR_DelayTimerApplied_r16); err != nil {
			return utils.WrapError("Encode Sl_LogicalChannelSR_DelayTimerApplied_r16", err)
		}
	}
	return nil
}

func (ie *SL_LogicalChannelConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_ConfiguredGrantType1Allowed_r16Present bool
	if Sl_ConfiguredGrantType1Allowed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_HARQ_FeedbackEnabled_r16Present bool
	if Sl_HARQ_FeedbackEnabled_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_AllowedCG_List_r16Present bool
	if Sl_AllowedCG_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_AllowedSCS_List_r16Present bool
	if Sl_AllowedSCS_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MaxPUSCH_Duration_r16Present bool
	if Sl_MaxPUSCH_Duration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_LogicalChannelGroup_r16Present bool
	if Sl_LogicalChannelGroup_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SchedulingRequestId_r16Present bool
	if Sl_SchedulingRequestId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_LogicalChannelSR_DelayTimerApplied_r16Present bool
	if Sl_LogicalChannelSR_DelayTimerApplied_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_Sl_Priority_r16 int64
	if tmp_int_Sl_Priority_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_Priority_r16", err)
	}
	ie.Sl_Priority_r16 = tmp_int_Sl_Priority_r16
	if err = ie.Sl_PrioritisedBitRate_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_PrioritisedBitRate_r16", err)
	}
	if err = ie.Sl_BucketSizeDuration_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_BucketSizeDuration_r16", err)
	}
	if Sl_ConfiguredGrantType1Allowed_r16Present {
		ie.Sl_ConfiguredGrantType1Allowed_r16 = new(SL_LogicalChannelConfig_r16_sl_ConfiguredGrantType1Allowed_r16)
		if err = ie.Sl_ConfiguredGrantType1Allowed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ConfiguredGrantType1Allowed_r16", err)
		}
	}
	if Sl_HARQ_FeedbackEnabled_r16Present {
		ie.Sl_HARQ_FeedbackEnabled_r16 = new(SL_LogicalChannelConfig_r16_sl_HARQ_FeedbackEnabled_r16)
		if err = ie.Sl_HARQ_FeedbackEnabled_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_HARQ_FeedbackEnabled_r16", err)
		}
	}
	if Sl_AllowedCG_List_r16Present {
		tmp_Sl_AllowedCG_List_r16 := utils.NewSequence[*SL_ConfigIndexCG_r16]([]*SL_ConfigIndexCG_r16{}, uper.Constraint{Lb: 0, Ub: maxNrofCG_SL_1_r16}, false)
		fn_Sl_AllowedCG_List_r16 := func() *SL_ConfigIndexCG_r16 {
			return new(SL_ConfigIndexCG_r16)
		}
		if err = tmp_Sl_AllowedCG_List_r16.Decode(r, fn_Sl_AllowedCG_List_r16); err != nil {
			return utils.WrapError("Decode Sl_AllowedCG_List_r16", err)
		}
		ie.Sl_AllowedCG_List_r16 = []SL_ConfigIndexCG_r16{}
		for _, i := range tmp_Sl_AllowedCG_List_r16.Value {
			ie.Sl_AllowedCG_List_r16 = append(ie.Sl_AllowedCG_List_r16, *i)
		}
	}
	if Sl_AllowedSCS_List_r16Present {
		tmp_Sl_AllowedSCS_List_r16 := utils.NewSequence[*SubcarrierSpacing]([]*SubcarrierSpacing{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
		fn_Sl_AllowedSCS_List_r16 := func() *SubcarrierSpacing {
			return new(SubcarrierSpacing)
		}
		if err = tmp_Sl_AllowedSCS_List_r16.Decode(r, fn_Sl_AllowedSCS_List_r16); err != nil {
			return utils.WrapError("Decode Sl_AllowedSCS_List_r16", err)
		}
		ie.Sl_AllowedSCS_List_r16 = []SubcarrierSpacing{}
		for _, i := range tmp_Sl_AllowedSCS_List_r16.Value {
			ie.Sl_AllowedSCS_List_r16 = append(ie.Sl_AllowedSCS_List_r16, *i)
		}
	}
	if Sl_MaxPUSCH_Duration_r16Present {
		ie.Sl_MaxPUSCH_Duration_r16 = new(SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16)
		if err = ie.Sl_MaxPUSCH_Duration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MaxPUSCH_Duration_r16", err)
		}
	}
	if Sl_LogicalChannelGroup_r16Present {
		var tmp_int_Sl_LogicalChannelGroup_r16 int64
		if tmp_int_Sl_LogicalChannelGroup_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxLCG_ID}, false); err != nil {
			return utils.WrapError("Decode Sl_LogicalChannelGroup_r16", err)
		}
		ie.Sl_LogicalChannelGroup_r16 = &tmp_int_Sl_LogicalChannelGroup_r16
	}
	if Sl_SchedulingRequestId_r16Present {
		ie.Sl_SchedulingRequestId_r16 = new(SchedulingRequestId)
		if err = ie.Sl_SchedulingRequestId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SchedulingRequestId_r16", err)
		}
	}
	if Sl_LogicalChannelSR_DelayTimerApplied_r16Present {
		var tmp_bool_Sl_LogicalChannelSR_DelayTimerApplied_r16 bool
		if tmp_bool_Sl_LogicalChannelSR_DelayTimerApplied_r16, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode Sl_LogicalChannelSR_DelayTimerApplied_r16", err)
		}
		ie.Sl_LogicalChannelSR_DelayTimerApplied_r16 = &tmp_bool_Sl_LogicalChannelSR_DelayTimerApplied_r16
	}
	return nil
}
