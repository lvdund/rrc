package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_InterUE_CoordinationScheme1_r17 struct {
	Sl_IUC_Explicit_r17                             *SL_InterUE_CoordinationScheme1_r17_sl_IUC_Explicit_r17          `optional`
	Sl_IUC_Condition_r17                            *SL_InterUE_CoordinationScheme1_r17_sl_IUC_Condition_r17         `optional`
	Sl_Condition1_A_2_r17                           *SL_InterUE_CoordinationScheme1_r17_sl_Condition1_A_2_r17        `optional`
	Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 []SL_ThresholdRSRP_Condition1_B_1_r17                            `lb:1,ub:8,optional`
	Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 []SL_ThresholdRSRP_Condition1_B_1_r17                            `lb:1,ub:8,optional`
	Sl_ContainerCoordInfo_r17                       *SL_InterUE_CoordinationScheme1_r17_sl_ContainerCoordInfo_r17    `optional`
	Sl_ContainerRequest_r17                         *SL_InterUE_CoordinationScheme1_r17_sl_ContainerRequest_r17      `optional`
	Sl_TriggerConditionCoordInfo_r17                *int64                                                           `lb:0,ub:1,optional`
	Sl_TriggerConditionRequest_r17                  *int64                                                           `lb:0,ub:1,optional`
	Sl_PriorityCoordInfoExplicit_r17                *int64                                                           `lb:1,ub:8,optional`
	Sl_PriorityCoordInfoCondition_r17               *int64                                                           `lb:1,ub:8,optional`
	Sl_PriorityRequest_r17                          *int64                                                           `lb:1,ub:8,optional`
	Sl_PriorityPreferredResourceSet_r17             *int64                                                           `lb:1,ub:8,optional`
	Sl_MaxSlotOffsetTRIV_r17                        *int64                                                           `lb:1,ub:8000,optional`
	Sl_NumSubCH_PreferredResourceSet_r17            *int64                                                           `lb:1,ub:27,optional`
	Sl_ReservedPeriodPreferredResourceSet_r17       *int64                                                           `lb:1,ub:16,optional`
	Sl_DetermineResourceType_r17                    *SL_InterUE_CoordinationScheme1_r17_sl_DetermineResourceType_r17 `optional`
}

func (ie *SL_InterUE_CoordinationScheme1_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_IUC_Explicit_r17 != nil, ie.Sl_IUC_Condition_r17 != nil, ie.Sl_Condition1_A_2_r17 != nil, len(ie.Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17) > 0, len(ie.Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17) > 0, ie.Sl_ContainerCoordInfo_r17 != nil, ie.Sl_ContainerRequest_r17 != nil, ie.Sl_TriggerConditionCoordInfo_r17 != nil, ie.Sl_TriggerConditionRequest_r17 != nil, ie.Sl_PriorityCoordInfoExplicit_r17 != nil, ie.Sl_PriorityCoordInfoCondition_r17 != nil, ie.Sl_PriorityRequest_r17 != nil, ie.Sl_PriorityPreferredResourceSet_r17 != nil, ie.Sl_MaxSlotOffsetTRIV_r17 != nil, ie.Sl_NumSubCH_PreferredResourceSet_r17 != nil, ie.Sl_ReservedPeriodPreferredResourceSet_r17 != nil, ie.Sl_DetermineResourceType_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_IUC_Explicit_r17 != nil {
		if err = ie.Sl_IUC_Explicit_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_IUC_Explicit_r17", err)
		}
	}
	if ie.Sl_IUC_Condition_r17 != nil {
		if err = ie.Sl_IUC_Condition_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_IUC_Condition_r17", err)
		}
	}
	if ie.Sl_Condition1_A_2_r17 != nil {
		if err = ie.Sl_Condition1_A_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_Condition1_A_2_r17", err)
		}
	}
	if len(ie.Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17) > 0 {
		tmp_Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 := utils.NewSequence[*SL_ThresholdRSRP_Condition1_B_1_r17]([]*SL_ThresholdRSRP_Condition1_B_1_r17{}, aper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 {
			tmp_Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17.Value = append(tmp_Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17.Value, &i)
		}
		if err = tmp_Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17", err)
		}
	}
	if len(ie.Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17) > 0 {
		tmp_Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 := utils.NewSequence[*SL_ThresholdRSRP_Condition1_B_1_r17]([]*SL_ThresholdRSRP_Condition1_B_1_r17{}, aper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 {
			tmp_Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17.Value = append(tmp_Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17.Value, &i)
		}
		if err = tmp_Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17", err)
		}
	}
	if ie.Sl_ContainerCoordInfo_r17 != nil {
		if err = ie.Sl_ContainerCoordInfo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ContainerCoordInfo_r17", err)
		}
	}
	if ie.Sl_ContainerRequest_r17 != nil {
		if err = ie.Sl_ContainerRequest_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ContainerRequest_r17", err)
		}
	}
	if ie.Sl_TriggerConditionCoordInfo_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_TriggerConditionCoordInfo_r17, &aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode Sl_TriggerConditionCoordInfo_r17", err)
		}
	}
	if ie.Sl_TriggerConditionRequest_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_TriggerConditionRequest_r17, &aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode Sl_TriggerConditionRequest_r17", err)
		}
	}
	if ie.Sl_PriorityCoordInfoExplicit_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_PriorityCoordInfoExplicit_r17, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Sl_PriorityCoordInfoExplicit_r17", err)
		}
	}
	if ie.Sl_PriorityCoordInfoCondition_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_PriorityCoordInfoCondition_r17, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Sl_PriorityCoordInfoCondition_r17", err)
		}
	}
	if ie.Sl_PriorityRequest_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_PriorityRequest_r17, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Sl_PriorityRequest_r17", err)
		}
	}
	if ie.Sl_PriorityPreferredResourceSet_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_PriorityPreferredResourceSet_r17, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Sl_PriorityPreferredResourceSet_r17", err)
		}
	}
	if ie.Sl_MaxSlotOffsetTRIV_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_MaxSlotOffsetTRIV_r17, &aper.Constraint{Lb: 1, Ub: 8000}, false); err != nil {
			return utils.WrapError("Encode Sl_MaxSlotOffsetTRIV_r17", err)
		}
	}
	if ie.Sl_NumSubCH_PreferredResourceSet_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_NumSubCH_PreferredResourceSet_r17, &aper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
			return utils.WrapError("Encode Sl_NumSubCH_PreferredResourceSet_r17", err)
		}
	}
	if ie.Sl_ReservedPeriodPreferredResourceSet_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_ReservedPeriodPreferredResourceSet_r17, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Sl_ReservedPeriodPreferredResourceSet_r17", err)
		}
	}
	if ie.Sl_DetermineResourceType_r17 != nil {
		if err = ie.Sl_DetermineResourceType_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DetermineResourceType_r17", err)
		}
	}
	return nil
}

func (ie *SL_InterUE_CoordinationScheme1_r17) Decode(r *aper.AperReader) error {
	var err error
	var Sl_IUC_Explicit_r17Present bool
	if Sl_IUC_Explicit_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_IUC_Condition_r17Present bool
	if Sl_IUC_Condition_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_Condition1_A_2_r17Present bool
	if Sl_Condition1_A_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17Present bool
	if Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17Present bool
	if Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ContainerCoordInfo_r17Present bool
	if Sl_ContainerCoordInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ContainerRequest_r17Present bool
	if Sl_ContainerRequest_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TriggerConditionCoordInfo_r17Present bool
	if Sl_TriggerConditionCoordInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TriggerConditionRequest_r17Present bool
	if Sl_TriggerConditionRequest_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PriorityCoordInfoExplicit_r17Present bool
	if Sl_PriorityCoordInfoExplicit_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PriorityCoordInfoCondition_r17Present bool
	if Sl_PriorityCoordInfoCondition_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PriorityRequest_r17Present bool
	if Sl_PriorityRequest_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PriorityPreferredResourceSet_r17Present bool
	if Sl_PriorityPreferredResourceSet_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MaxSlotOffsetTRIV_r17Present bool
	if Sl_MaxSlotOffsetTRIV_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_NumSubCH_PreferredResourceSet_r17Present bool
	if Sl_NumSubCH_PreferredResourceSet_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ReservedPeriodPreferredResourceSet_r17Present bool
	if Sl_ReservedPeriodPreferredResourceSet_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_DetermineResourceType_r17Present bool
	if Sl_DetermineResourceType_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_IUC_Explicit_r17Present {
		ie.Sl_IUC_Explicit_r17 = new(SL_InterUE_CoordinationScheme1_r17_sl_IUC_Explicit_r17)
		if err = ie.Sl_IUC_Explicit_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_IUC_Explicit_r17", err)
		}
	}
	if Sl_IUC_Condition_r17Present {
		ie.Sl_IUC_Condition_r17 = new(SL_InterUE_CoordinationScheme1_r17_sl_IUC_Condition_r17)
		if err = ie.Sl_IUC_Condition_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_IUC_Condition_r17", err)
		}
	}
	if Sl_Condition1_A_2_r17Present {
		ie.Sl_Condition1_A_2_r17 = new(SL_InterUE_CoordinationScheme1_r17_sl_Condition1_A_2_r17)
		if err = ie.Sl_Condition1_A_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_Condition1_A_2_r17", err)
		}
	}
	if Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17Present {
		tmp_Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 := utils.NewSequence[*SL_ThresholdRSRP_Condition1_B_1_r17]([]*SL_ThresholdRSRP_Condition1_B_1_r17{}, aper.Constraint{Lb: 1, Ub: 8}, false)
		fn_Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 := func() *SL_ThresholdRSRP_Condition1_B_1_r17 {
			return new(SL_ThresholdRSRP_Condition1_B_1_r17)
		}
		if err = tmp_Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17.Decode(r, fn_Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17); err != nil {
			return utils.WrapError("Decode Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17", err)
		}
		ie.Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 = []SL_ThresholdRSRP_Condition1_B_1_r17{}
		for _, i := range tmp_Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17.Value {
			ie.Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 = append(ie.Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17, *i)
		}
	}
	if Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17Present {
		tmp_Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 := utils.NewSequence[*SL_ThresholdRSRP_Condition1_B_1_r17]([]*SL_ThresholdRSRP_Condition1_B_1_r17{}, aper.Constraint{Lb: 1, Ub: 8}, false)
		fn_Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 := func() *SL_ThresholdRSRP_Condition1_B_1_r17 {
			return new(SL_ThresholdRSRP_Condition1_B_1_r17)
		}
		if err = tmp_Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17.Decode(r, fn_Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17); err != nil {
			return utils.WrapError("Decode Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17", err)
		}
		ie.Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 = []SL_ThresholdRSRP_Condition1_B_1_r17{}
		for _, i := range tmp_Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17.Value {
			ie.Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 = append(ie.Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17, *i)
		}
	}
	if Sl_ContainerCoordInfo_r17Present {
		ie.Sl_ContainerCoordInfo_r17 = new(SL_InterUE_CoordinationScheme1_r17_sl_ContainerCoordInfo_r17)
		if err = ie.Sl_ContainerCoordInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ContainerCoordInfo_r17", err)
		}
	}
	if Sl_ContainerRequest_r17Present {
		ie.Sl_ContainerRequest_r17 = new(SL_InterUE_CoordinationScheme1_r17_sl_ContainerRequest_r17)
		if err = ie.Sl_ContainerRequest_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ContainerRequest_r17", err)
		}
	}
	if Sl_TriggerConditionCoordInfo_r17Present {
		var tmp_int_Sl_TriggerConditionCoordInfo_r17 int64
		if tmp_int_Sl_TriggerConditionCoordInfo_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode Sl_TriggerConditionCoordInfo_r17", err)
		}
		ie.Sl_TriggerConditionCoordInfo_r17 = &tmp_int_Sl_TriggerConditionCoordInfo_r17
	}
	if Sl_TriggerConditionRequest_r17Present {
		var tmp_int_Sl_TriggerConditionRequest_r17 int64
		if tmp_int_Sl_TriggerConditionRequest_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode Sl_TriggerConditionRequest_r17", err)
		}
		ie.Sl_TriggerConditionRequest_r17 = &tmp_int_Sl_TriggerConditionRequest_r17
	}
	if Sl_PriorityCoordInfoExplicit_r17Present {
		var tmp_int_Sl_PriorityCoordInfoExplicit_r17 int64
		if tmp_int_Sl_PriorityCoordInfoExplicit_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Sl_PriorityCoordInfoExplicit_r17", err)
		}
		ie.Sl_PriorityCoordInfoExplicit_r17 = &tmp_int_Sl_PriorityCoordInfoExplicit_r17
	}
	if Sl_PriorityCoordInfoCondition_r17Present {
		var tmp_int_Sl_PriorityCoordInfoCondition_r17 int64
		if tmp_int_Sl_PriorityCoordInfoCondition_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Sl_PriorityCoordInfoCondition_r17", err)
		}
		ie.Sl_PriorityCoordInfoCondition_r17 = &tmp_int_Sl_PriorityCoordInfoCondition_r17
	}
	if Sl_PriorityRequest_r17Present {
		var tmp_int_Sl_PriorityRequest_r17 int64
		if tmp_int_Sl_PriorityRequest_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Sl_PriorityRequest_r17", err)
		}
		ie.Sl_PriorityRequest_r17 = &tmp_int_Sl_PriorityRequest_r17
	}
	if Sl_PriorityPreferredResourceSet_r17Present {
		var tmp_int_Sl_PriorityPreferredResourceSet_r17 int64
		if tmp_int_Sl_PriorityPreferredResourceSet_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Sl_PriorityPreferredResourceSet_r17", err)
		}
		ie.Sl_PriorityPreferredResourceSet_r17 = &tmp_int_Sl_PriorityPreferredResourceSet_r17
	}
	if Sl_MaxSlotOffsetTRIV_r17Present {
		var tmp_int_Sl_MaxSlotOffsetTRIV_r17 int64
		if tmp_int_Sl_MaxSlotOffsetTRIV_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8000}, false); err != nil {
			return utils.WrapError("Decode Sl_MaxSlotOffsetTRIV_r17", err)
		}
		ie.Sl_MaxSlotOffsetTRIV_r17 = &tmp_int_Sl_MaxSlotOffsetTRIV_r17
	}
	if Sl_NumSubCH_PreferredResourceSet_r17Present {
		var tmp_int_Sl_NumSubCH_PreferredResourceSet_r17 int64
		if tmp_int_Sl_NumSubCH_PreferredResourceSet_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
			return utils.WrapError("Decode Sl_NumSubCH_PreferredResourceSet_r17", err)
		}
		ie.Sl_NumSubCH_PreferredResourceSet_r17 = &tmp_int_Sl_NumSubCH_PreferredResourceSet_r17
	}
	if Sl_ReservedPeriodPreferredResourceSet_r17Present {
		var tmp_int_Sl_ReservedPeriodPreferredResourceSet_r17 int64
		if tmp_int_Sl_ReservedPeriodPreferredResourceSet_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Sl_ReservedPeriodPreferredResourceSet_r17", err)
		}
		ie.Sl_ReservedPeriodPreferredResourceSet_r17 = &tmp_int_Sl_ReservedPeriodPreferredResourceSet_r17
	}
	if Sl_DetermineResourceType_r17Present {
		ie.Sl_DetermineResourceType_r17 = new(SL_InterUE_CoordinationScheme1_r17_sl_DetermineResourceType_r17)
		if err = ie.Sl_DetermineResourceType_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_DetermineResourceType_r17", err)
		}
	}
	return nil
}
