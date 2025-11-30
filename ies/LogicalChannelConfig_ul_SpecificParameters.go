package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type LogicalChannelConfig_ul_SpecificParameters struct {
	Priority                           int64                                                                    `lb:1,ub:16,madatory`
	PrioritisedBitRate                 LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate            `madatory`
	BucketSizeDuration                 LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration            `madatory`
	AllowedServingCells                []ServCellIndex                                                          `lb:1,ub:maxNrofServingCells_1,optional`
	AllowedSCS_List                    []SubcarrierSpacing                                                      `lb:1,ub:maxSCSs,optional`
	MaxPUSCH_Duration                  *LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration            `optional`
	ConfiguredGrantType1Allowed        *LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed  `optional`
	LogicalChannelGroup                *int64                                                                   `lb:0,ub:maxLCG_ID,optional`
	SchedulingRequestID                *SchedulingRequestId                                                     `optional`
	LogicalChannelSR_Mask              bool                                                                     `madatory`
	LogicalChannelSR_DelayTimerApplied bool                                                                     `madatory`
	BitRateQueryProhibitTimer          *LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer    `optional`
	AllowedCG_List_r16                 []ConfiguredGrantConfigIndexMAC_r16                                      `lb:0,ub:maxNrofConfiguredGrantConfigMAC_1_r16,optional`
	AllowedPHY_PriorityIndex_r16       *LogicalChannelConfig_ul_SpecificParameters_allowedPHY_PriorityIndex_r16 `optional`
	LogicalChannelGroupIAB_Ext_r17     *int64                                                                   `lb:0,ub:maxLCG_ID_IAB_r17,optional`
	AllowedHARQ_mode_r17               *LogicalChannelConfig_ul_SpecificParameters_allowedHARQ_mode_r17         `optional`
}

func (ie *LogicalChannelConfig_ul_SpecificParameters) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.AllowedServingCells) > 0, len(ie.AllowedSCS_List) > 0, ie.MaxPUSCH_Duration != nil, ie.ConfiguredGrantType1Allowed != nil, ie.LogicalChannelGroup != nil, ie.SchedulingRequestID != nil, ie.BitRateQueryProhibitTimer != nil, len(ie.AllowedCG_List_r16) > 0, ie.AllowedPHY_PriorityIndex_r16 != nil, ie.LogicalChannelGroupIAB_Ext_r17 != nil, ie.AllowedHARQ_mode_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.Priority, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger Priority", err)
	}
	if err = ie.PrioritisedBitRate.Encode(w); err != nil {
		return utils.WrapError("Encode PrioritisedBitRate", err)
	}
	if err = ie.BucketSizeDuration.Encode(w); err != nil {
		return utils.WrapError("Encode BucketSizeDuration", err)
	}
	if len(ie.AllowedServingCells) > 0 {
		tmp_AllowedServingCells := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		for _, i := range ie.AllowedServingCells {
			tmp_AllowedServingCells.Value = append(tmp_AllowedServingCells.Value, &i)
		}
		if err = tmp_AllowedServingCells.Encode(w); err != nil {
			return utils.WrapError("Encode AllowedServingCells", err)
		}
	}
	if len(ie.AllowedSCS_List) > 0 {
		tmp_AllowedSCS_List := utils.NewSequence[*SubcarrierSpacing]([]*SubcarrierSpacing{}, aper.Constraint{Lb: 1, Ub: maxSCSs}, false)
		for _, i := range ie.AllowedSCS_List {
			tmp_AllowedSCS_List.Value = append(tmp_AllowedSCS_List.Value, &i)
		}
		if err = tmp_AllowedSCS_List.Encode(w); err != nil {
			return utils.WrapError("Encode AllowedSCS_List", err)
		}
	}
	if ie.MaxPUSCH_Duration != nil {
		if err = ie.MaxPUSCH_Duration.Encode(w); err != nil {
			return utils.WrapError("Encode MaxPUSCH_Duration", err)
		}
	}
	if ie.ConfiguredGrantType1Allowed != nil {
		if err = ie.ConfiguredGrantType1Allowed.Encode(w); err != nil {
			return utils.WrapError("Encode ConfiguredGrantType1Allowed", err)
		}
	}
	if ie.LogicalChannelGroup != nil {
		if err = w.WriteInteger(*ie.LogicalChannelGroup, &aper.Constraint{Lb: 0, Ub: maxLCG_ID}, false); err != nil {
			return utils.WrapError("Encode LogicalChannelGroup", err)
		}
	}
	if ie.SchedulingRequestID != nil {
		if err = ie.SchedulingRequestID.Encode(w); err != nil {
			return utils.WrapError("Encode SchedulingRequestID", err)
		}
	}
	if err = w.WriteBoolean(ie.LogicalChannelSR_Mask); err != nil {
		return utils.WrapError("WriteBoolean LogicalChannelSR_Mask", err)
	}
	if err = w.WriteBoolean(ie.LogicalChannelSR_DelayTimerApplied); err != nil {
		return utils.WrapError("WriteBoolean LogicalChannelSR_DelayTimerApplied", err)
	}
	if ie.BitRateQueryProhibitTimer != nil {
		if err = ie.BitRateQueryProhibitTimer.Encode(w); err != nil {
			return utils.WrapError("Encode BitRateQueryProhibitTimer", err)
		}
	}
	if len(ie.AllowedCG_List_r16) > 0 {
		tmp_AllowedCG_List_r16 := utils.NewSequence[*ConfiguredGrantConfigIndexMAC_r16]([]*ConfiguredGrantConfigIndexMAC_r16{}, aper.Constraint{Lb: 0, Ub: maxNrofConfiguredGrantConfigMAC_1_r16}, false)
		for _, i := range ie.AllowedCG_List_r16 {
			tmp_AllowedCG_List_r16.Value = append(tmp_AllowedCG_List_r16.Value, &i)
		}
		if err = tmp_AllowedCG_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AllowedCG_List_r16", err)
		}
	}
	if ie.AllowedPHY_PriorityIndex_r16 != nil {
		if err = ie.AllowedPHY_PriorityIndex_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AllowedPHY_PriorityIndex_r16", err)
		}
	}
	if ie.LogicalChannelGroupIAB_Ext_r17 != nil {
		if err = w.WriteInteger(*ie.LogicalChannelGroupIAB_Ext_r17, &aper.Constraint{Lb: 0, Ub: maxLCG_ID_IAB_r17}, false); err != nil {
			return utils.WrapError("Encode LogicalChannelGroupIAB_Ext_r17", err)
		}
	}
	if ie.AllowedHARQ_mode_r17 != nil {
		if err = ie.AllowedHARQ_mode_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AllowedHARQ_mode_r17", err)
		}
	}
	return nil
}

func (ie *LogicalChannelConfig_ul_SpecificParameters) Decode(r *aper.AperReader) error {
	var err error
	var AllowedServingCellsPresent bool
	if AllowedServingCellsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AllowedSCS_ListPresent bool
	if AllowedSCS_ListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxPUSCH_DurationPresent bool
	if MaxPUSCH_DurationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ConfiguredGrantType1AllowedPresent bool
	if ConfiguredGrantType1AllowedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LogicalChannelGroupPresent bool
	if LogicalChannelGroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SchedulingRequestIDPresent bool
	if SchedulingRequestIDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var BitRateQueryProhibitTimerPresent bool
	if BitRateQueryProhibitTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AllowedCG_List_r16Present bool
	if AllowedCG_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AllowedPHY_PriorityIndex_r16Present bool
	if AllowedPHY_PriorityIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LogicalChannelGroupIAB_Ext_r17Present bool
	if LogicalChannelGroupIAB_Ext_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AllowedHARQ_mode_r17Present bool
	if AllowedHARQ_mode_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_Priority int64
	if tmp_int_Priority, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger Priority", err)
	}
	ie.Priority = tmp_int_Priority
	if err = ie.PrioritisedBitRate.Decode(r); err != nil {
		return utils.WrapError("Decode PrioritisedBitRate", err)
	}
	if err = ie.BucketSizeDuration.Decode(r); err != nil {
		return utils.WrapError("Decode BucketSizeDuration", err)
	}
	if AllowedServingCellsPresent {
		tmp_AllowedServingCells := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		fn_AllowedServingCells := func() *ServCellIndex {
			return new(ServCellIndex)
		}
		if err = tmp_AllowedServingCells.Decode(r, fn_AllowedServingCells); err != nil {
			return utils.WrapError("Decode AllowedServingCells", err)
		}
		ie.AllowedServingCells = []ServCellIndex{}
		for _, i := range tmp_AllowedServingCells.Value {
			ie.AllowedServingCells = append(ie.AllowedServingCells, *i)
		}
	}
	if AllowedSCS_ListPresent {
		tmp_AllowedSCS_List := utils.NewSequence[*SubcarrierSpacing]([]*SubcarrierSpacing{}, aper.Constraint{Lb: 1, Ub: maxSCSs}, false)
		fn_AllowedSCS_List := func() *SubcarrierSpacing {
			return new(SubcarrierSpacing)
		}
		if err = tmp_AllowedSCS_List.Decode(r, fn_AllowedSCS_List); err != nil {
			return utils.WrapError("Decode AllowedSCS_List", err)
		}
		ie.AllowedSCS_List = []SubcarrierSpacing{}
		for _, i := range tmp_AllowedSCS_List.Value {
			ie.AllowedSCS_List = append(ie.AllowedSCS_List, *i)
		}
	}
	if MaxPUSCH_DurationPresent {
		ie.MaxPUSCH_Duration = new(LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration)
		if err = ie.MaxPUSCH_Duration.Decode(r); err != nil {
			return utils.WrapError("Decode MaxPUSCH_Duration", err)
		}
	}
	if ConfiguredGrantType1AllowedPresent {
		ie.ConfiguredGrantType1Allowed = new(LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed)
		if err = ie.ConfiguredGrantType1Allowed.Decode(r); err != nil {
			return utils.WrapError("Decode ConfiguredGrantType1Allowed", err)
		}
	}
	if LogicalChannelGroupPresent {
		var tmp_int_LogicalChannelGroup int64
		if tmp_int_LogicalChannelGroup, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxLCG_ID}, false); err != nil {
			return utils.WrapError("Decode LogicalChannelGroup", err)
		}
		ie.LogicalChannelGroup = &tmp_int_LogicalChannelGroup
	}
	if SchedulingRequestIDPresent {
		ie.SchedulingRequestID = new(SchedulingRequestId)
		if err = ie.SchedulingRequestID.Decode(r); err != nil {
			return utils.WrapError("Decode SchedulingRequestID", err)
		}
	}
	var tmp_bool_LogicalChannelSR_Mask bool
	if tmp_bool_LogicalChannelSR_Mask, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean LogicalChannelSR_Mask", err)
	}
	ie.LogicalChannelSR_Mask = tmp_bool_LogicalChannelSR_Mask
	var tmp_bool_LogicalChannelSR_DelayTimerApplied bool
	if tmp_bool_LogicalChannelSR_DelayTimerApplied, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean LogicalChannelSR_DelayTimerApplied", err)
	}
	ie.LogicalChannelSR_DelayTimerApplied = tmp_bool_LogicalChannelSR_DelayTimerApplied
	if BitRateQueryProhibitTimerPresent {
		ie.BitRateQueryProhibitTimer = new(LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer)
		if err = ie.BitRateQueryProhibitTimer.Decode(r); err != nil {
			return utils.WrapError("Decode BitRateQueryProhibitTimer", err)
		}
	}
	if AllowedCG_List_r16Present {
		tmp_AllowedCG_List_r16 := utils.NewSequence[*ConfiguredGrantConfigIndexMAC_r16]([]*ConfiguredGrantConfigIndexMAC_r16{}, aper.Constraint{Lb: 0, Ub: maxNrofConfiguredGrantConfigMAC_1_r16}, false)
		fn_AllowedCG_List_r16 := func() *ConfiguredGrantConfigIndexMAC_r16 {
			return new(ConfiguredGrantConfigIndexMAC_r16)
		}
		if err = tmp_AllowedCG_List_r16.Decode(r, fn_AllowedCG_List_r16); err != nil {
			return utils.WrapError("Decode AllowedCG_List_r16", err)
		}
		ie.AllowedCG_List_r16 = []ConfiguredGrantConfigIndexMAC_r16{}
		for _, i := range tmp_AllowedCG_List_r16.Value {
			ie.AllowedCG_List_r16 = append(ie.AllowedCG_List_r16, *i)
		}
	}
	if AllowedPHY_PriorityIndex_r16Present {
		ie.AllowedPHY_PriorityIndex_r16 = new(LogicalChannelConfig_ul_SpecificParameters_allowedPHY_PriorityIndex_r16)
		if err = ie.AllowedPHY_PriorityIndex_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AllowedPHY_PriorityIndex_r16", err)
		}
	}
	if LogicalChannelGroupIAB_Ext_r17Present {
		var tmp_int_LogicalChannelGroupIAB_Ext_r17 int64
		if tmp_int_LogicalChannelGroupIAB_Ext_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxLCG_ID_IAB_r17}, false); err != nil {
			return utils.WrapError("Decode LogicalChannelGroupIAB_Ext_r17", err)
		}
		ie.LogicalChannelGroupIAB_Ext_r17 = &tmp_int_LogicalChannelGroupIAB_Ext_r17
	}
	if AllowedHARQ_mode_r17Present {
		ie.AllowedHARQ_mode_r17 = new(LogicalChannelConfig_ul_SpecificParameters_allowedHARQ_mode_r17)
		if err = ie.AllowedHARQ_mode_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AllowedHARQ_mode_r17", err)
		}
	}
	return nil
}
