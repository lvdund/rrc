// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-InterUE-CoordinationScheme1-r17 (line 27354).

var sLInterUECoordinationScheme1R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-IUC-Explicit-r17", Optional: true},
		{Name: "sl-IUC-Condition-r17", Optional: true},
		{Name: "sl-Condition1-A-2-r17", Optional: true},
		{Name: "sl-ThresholdRSRP-Condition1-B-1-Option1List-r17", Optional: true},
		{Name: "sl-ThresholdRSRP-Condition1-B-1-Option2List-r17", Optional: true},
		{Name: "sl-ContainerCoordInfo-r17", Optional: true},
		{Name: "sl-ContainerRequest-r17", Optional: true},
		{Name: "sl-TriggerConditionCoordInfo-r17", Optional: true},
		{Name: "sl-TriggerConditionRequest-r17", Optional: true},
		{Name: "sl-PriorityCoordInfoExplicit-r17", Optional: true},
		{Name: "sl-PriorityCoordInfoCondition-r17", Optional: true},
		{Name: "sl-PriorityRequest-r17", Optional: true},
		{Name: "sl-PriorityPreferredResourceSet-r17", Optional: true},
		{Name: "sl-MaxSlotOffsetTRIV-r17", Optional: true},
		{Name: "sl-NumSubCH-PreferredResourceSet-r17", Optional: true},
		{Name: "sl-ReservedPeriodPreferredResourceSet-r17", Optional: true},
		{Name: "sl-DetermineResourceType-r17", Optional: true},
	},
}

const (
	SL_InterUE_CoordinationScheme1_r17_Sl_IUC_Explicit_r17_Enabled  = 0
	SL_InterUE_CoordinationScheme1_r17_Sl_IUC_Explicit_r17_Disabled = 1
)

var sLInterUECoordinationScheme1R17SlIUCExplicitR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_InterUE_CoordinationScheme1_r17_Sl_IUC_Explicit_r17_Enabled, SL_InterUE_CoordinationScheme1_r17_Sl_IUC_Explicit_r17_Disabled},
}

const (
	SL_InterUE_CoordinationScheme1_r17_Sl_IUC_Condition_r17_Enabled  = 0
	SL_InterUE_CoordinationScheme1_r17_Sl_IUC_Condition_r17_Disabled = 1
)

var sLInterUECoordinationScheme1R17SlIUCConditionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_InterUE_CoordinationScheme1_r17_Sl_IUC_Condition_r17_Enabled, SL_InterUE_CoordinationScheme1_r17_Sl_IUC_Condition_r17_Disabled},
}

const (
	SL_InterUE_CoordinationScheme1_r17_Sl_Condition1_A_2_r17_Disabled = 0
)

var sLInterUECoordinationScheme1R17SlCondition1A2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_InterUE_CoordinationScheme1_r17_Sl_Condition1_A_2_r17_Disabled},
}

var sLInterUECoordinationScheme1R17SlThresholdRSRPCondition1B1Option1ListR17Constraints = per.SizeRange(1, 8)

var sLInterUECoordinationScheme1R17SlThresholdRSRPCondition1B1Option2ListR17Constraints = per.SizeRange(1, 8)

const (
	SL_InterUE_CoordinationScheme1_r17_Sl_ContainerCoordInfo_r17_Enabled  = 0
	SL_InterUE_CoordinationScheme1_r17_Sl_ContainerCoordInfo_r17_Disabled = 1
)

var sLInterUECoordinationScheme1R17SlContainerCoordInfoR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_InterUE_CoordinationScheme1_r17_Sl_ContainerCoordInfo_r17_Enabled, SL_InterUE_CoordinationScheme1_r17_Sl_ContainerCoordInfo_r17_Disabled},
}

const (
	SL_InterUE_CoordinationScheme1_r17_Sl_ContainerRequest_r17_Enabled  = 0
	SL_InterUE_CoordinationScheme1_r17_Sl_ContainerRequest_r17_Disabled = 1
)

var sLInterUECoordinationScheme1R17SlContainerRequestR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_InterUE_CoordinationScheme1_r17_Sl_ContainerRequest_r17_Enabled, SL_InterUE_CoordinationScheme1_r17_Sl_ContainerRequest_r17_Disabled},
}

var sLInterUECoordinationScheme1R17SlTriggerConditionCoordInfoR17Constraints = per.Constrained(0, 1)

var sLInterUECoordinationScheme1R17SlTriggerConditionRequestR17Constraints = per.Constrained(0, 1)

var sLInterUECoordinationScheme1R17SlPriorityCoordInfoExplicitR17Constraints = per.Constrained(1, 8)

var sLInterUECoordinationScheme1R17SlPriorityCoordInfoConditionR17Constraints = per.Constrained(1, 8)

var sLInterUECoordinationScheme1R17SlPriorityRequestR17Constraints = per.Constrained(1, 8)

var sLInterUECoordinationScheme1R17SlPriorityPreferredResourceSetR17Constraints = per.Constrained(1, 8)

var sLInterUECoordinationScheme1R17SlMaxSlotOffsetTRIVR17Constraints = per.Constrained(1, 8000)

var sLInterUECoordinationScheme1R17SlNumSubCHPreferredResourceSetR17Constraints = per.Constrained(1, 27)

var sLInterUECoordinationScheme1R17SlReservedPeriodPreferredResourceSetR17Constraints = per.Constrained(1, 16)

const (
	SL_InterUE_CoordinationScheme1_r17_Sl_DetermineResourceType_r17_Uea = 0
	SL_InterUE_CoordinationScheme1_r17_Sl_DetermineResourceType_r17_Ueb = 1
)

var sLInterUECoordinationScheme1R17SlDetermineResourceTypeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_InterUE_CoordinationScheme1_r17_Sl_DetermineResourceType_r17_Uea, SL_InterUE_CoordinationScheme1_r17_Sl_DetermineResourceType_r17_Ueb},
}

type SL_InterUE_CoordinationScheme1_r17 struct {
	Sl_IUC_Explicit_r17                             *int64
	Sl_IUC_Condition_r17                            *int64
	Sl_Condition1_A_2_r17                           *int64
	Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 []SL_ThresholdRSRP_Condition1_B_1_r17
	Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 []SL_ThresholdRSRP_Condition1_B_1_r17
	Sl_ContainerCoordInfo_r17                       *int64
	Sl_ContainerRequest_r17                         *int64
	Sl_TriggerConditionCoordInfo_r17                *int64
	Sl_TriggerConditionRequest_r17                  *int64
	Sl_PriorityCoordInfoExplicit_r17                *int64
	Sl_PriorityCoordInfoCondition_r17               *int64
	Sl_PriorityRequest_r17                          *int64
	Sl_PriorityPreferredResourceSet_r17             *int64
	Sl_MaxSlotOffsetTRIV_r17                        *int64
	Sl_NumSubCH_PreferredResourceSet_r17            *int64
	Sl_ReservedPeriodPreferredResourceSet_r17       *int64
	Sl_DetermineResourceType_r17                    *int64
}

func (ie *SL_InterUE_CoordinationScheme1_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLInterUECoordinationScheme1R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_IUC_Explicit_r17 != nil, ie.Sl_IUC_Condition_r17 != nil, ie.Sl_Condition1_A_2_r17 != nil, ie.Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 != nil, ie.Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 != nil, ie.Sl_ContainerCoordInfo_r17 != nil, ie.Sl_ContainerRequest_r17 != nil, ie.Sl_TriggerConditionCoordInfo_r17 != nil, ie.Sl_TriggerConditionRequest_r17 != nil, ie.Sl_PriorityCoordInfoExplicit_r17 != nil, ie.Sl_PriorityCoordInfoCondition_r17 != nil, ie.Sl_PriorityRequest_r17 != nil, ie.Sl_PriorityPreferredResourceSet_r17 != nil, ie.Sl_MaxSlotOffsetTRIV_r17 != nil, ie.Sl_NumSubCH_PreferredResourceSet_r17 != nil, ie.Sl_ReservedPeriodPreferredResourceSet_r17 != nil, ie.Sl_DetermineResourceType_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-IUC-Explicit-r17: enumerated
	{
		if ie.Sl_IUC_Explicit_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_IUC_Explicit_r17, sLInterUECoordinationScheme1R17SlIUCExplicitR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-IUC-Condition-r17: enumerated
	{
		if ie.Sl_IUC_Condition_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_IUC_Condition_r17, sLInterUECoordinationScheme1R17SlIUCConditionR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-Condition1-A-2-r17: enumerated
	{
		if ie.Sl_Condition1_A_2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_Condition1_A_2_r17, sLInterUECoordinationScheme1R17SlCondition1A2R17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-ThresholdRSRP-Condition1-B-1-Option1List-r17: sequence-of
	{
		if ie.Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sLInterUECoordinationScheme1R17SlThresholdRSRPCondition1B1Option1ListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17))); err != nil {
				return err
			}
			for i := range ie.Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 {
				if err := ie.Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-ThresholdRSRP-Condition1-B-1-Option2List-r17: sequence-of
	{
		if ie.Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sLInterUECoordinationScheme1R17SlThresholdRSRPCondition1B1Option2ListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17))); err != nil {
				return err
			}
			for i := range ie.Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 {
				if err := ie.Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. sl-ContainerCoordInfo-r17: enumerated
	{
		if ie.Sl_ContainerCoordInfo_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_ContainerCoordInfo_r17, sLInterUECoordinationScheme1R17SlContainerCoordInfoR17Constraints); err != nil {
				return err
			}
		}
	}

	// 9. sl-ContainerRequest-r17: enumerated
	{
		if ie.Sl_ContainerRequest_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_ContainerRequest_r17, sLInterUECoordinationScheme1R17SlContainerRequestR17Constraints); err != nil {
				return err
			}
		}
	}

	// 10. sl-TriggerConditionCoordInfo-r17: integer
	{
		if ie.Sl_TriggerConditionCoordInfo_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_TriggerConditionCoordInfo_r17, sLInterUECoordinationScheme1R17SlTriggerConditionCoordInfoR17Constraints); err != nil {
				return err
			}
		}
	}

	// 11. sl-TriggerConditionRequest-r17: integer
	{
		if ie.Sl_TriggerConditionRequest_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_TriggerConditionRequest_r17, sLInterUECoordinationScheme1R17SlTriggerConditionRequestR17Constraints); err != nil {
				return err
			}
		}
	}

	// 12. sl-PriorityCoordInfoExplicit-r17: integer
	{
		if ie.Sl_PriorityCoordInfoExplicit_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_PriorityCoordInfoExplicit_r17, sLInterUECoordinationScheme1R17SlPriorityCoordInfoExplicitR17Constraints); err != nil {
				return err
			}
		}
	}

	// 13. sl-PriorityCoordInfoCondition-r17: integer
	{
		if ie.Sl_PriorityCoordInfoCondition_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_PriorityCoordInfoCondition_r17, sLInterUECoordinationScheme1R17SlPriorityCoordInfoConditionR17Constraints); err != nil {
				return err
			}
		}
	}

	// 14. sl-PriorityRequest-r17: integer
	{
		if ie.Sl_PriorityRequest_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_PriorityRequest_r17, sLInterUECoordinationScheme1R17SlPriorityRequestR17Constraints); err != nil {
				return err
			}
		}
	}

	// 15. sl-PriorityPreferredResourceSet-r17: integer
	{
		if ie.Sl_PriorityPreferredResourceSet_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_PriorityPreferredResourceSet_r17, sLInterUECoordinationScheme1R17SlPriorityPreferredResourceSetR17Constraints); err != nil {
				return err
			}
		}
	}

	// 16. sl-MaxSlotOffsetTRIV-r17: integer
	{
		if ie.Sl_MaxSlotOffsetTRIV_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_MaxSlotOffsetTRIV_r17, sLInterUECoordinationScheme1R17SlMaxSlotOffsetTRIVR17Constraints); err != nil {
				return err
			}
		}
	}

	// 17. sl-NumSubCH-PreferredResourceSet-r17: integer
	{
		if ie.Sl_NumSubCH_PreferredResourceSet_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_NumSubCH_PreferredResourceSet_r17, sLInterUECoordinationScheme1R17SlNumSubCHPreferredResourceSetR17Constraints); err != nil {
				return err
			}
		}
	}

	// 18. sl-ReservedPeriodPreferredResourceSet-r17: integer
	{
		if ie.Sl_ReservedPeriodPreferredResourceSet_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_ReservedPeriodPreferredResourceSet_r17, sLInterUECoordinationScheme1R17SlReservedPeriodPreferredResourceSetR17Constraints); err != nil {
				return err
			}
		}
	}

	// 19. sl-DetermineResourceType-r17: enumerated
	{
		if ie.Sl_DetermineResourceType_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_DetermineResourceType_r17, sLInterUECoordinationScheme1R17SlDetermineResourceTypeR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_InterUE_CoordinationScheme1_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLInterUECoordinationScheme1R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-IUC-Explicit-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sLInterUECoordinationScheme1R17SlIUCExplicitR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_IUC_Explicit_r17 = &idx
		}
	}

	// 4. sl-IUC-Condition-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLInterUECoordinationScheme1R17SlIUCConditionR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_IUC_Condition_r17 = &idx
		}
	}

	// 5. sl-Condition1-A-2-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLInterUECoordinationScheme1R17SlCondition1A2R17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_Condition1_A_2_r17 = &idx
		}
	}

	// 6. sl-ThresholdRSRP-Condition1-B-1-Option1List-r17: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(sLInterUECoordinationScheme1R17SlThresholdRSRPCondition1B1Option1ListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 = make([]SL_ThresholdRSRP_Condition1_B_1_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_ThresholdRSRP_Condition1_B_1_Option1List_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-ThresholdRSRP-Condition1-B-1-Option2List-r17: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(sLInterUECoordinationScheme1R17SlThresholdRSRPCondition1B1Option2ListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 = make([]SL_ThresholdRSRP_Condition1_B_1_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_ThresholdRSRP_Condition1_B_1_Option2List_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. sl-ContainerCoordInfo-r17: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(sLInterUECoordinationScheme1R17SlContainerCoordInfoR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ContainerCoordInfo_r17 = &idx
		}
	}

	// 9. sl-ContainerRequest-r17: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(sLInterUECoordinationScheme1R17SlContainerRequestR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ContainerRequest_r17 = &idx
		}
	}

	// 10. sl-TriggerConditionCoordInfo-r17: integer
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeInteger(sLInterUECoordinationScheme1R17SlTriggerConditionCoordInfoR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TriggerConditionCoordInfo_r17 = &v
		}
	}

	// 11. sl-TriggerConditionRequest-r17: integer
	{
		if seq.IsComponentPresent(8) {
			v, err := d.DecodeInteger(sLInterUECoordinationScheme1R17SlTriggerConditionRequestR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TriggerConditionRequest_r17 = &v
		}
	}

	// 12. sl-PriorityCoordInfoExplicit-r17: integer
	{
		if seq.IsComponentPresent(9) {
			v, err := d.DecodeInteger(sLInterUECoordinationScheme1R17SlPriorityCoordInfoExplicitR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PriorityCoordInfoExplicit_r17 = &v
		}
	}

	// 13. sl-PriorityCoordInfoCondition-r17: integer
	{
		if seq.IsComponentPresent(10) {
			v, err := d.DecodeInteger(sLInterUECoordinationScheme1R17SlPriorityCoordInfoConditionR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PriorityCoordInfoCondition_r17 = &v
		}
	}

	// 14. sl-PriorityRequest-r17: integer
	{
		if seq.IsComponentPresent(11) {
			v, err := d.DecodeInteger(sLInterUECoordinationScheme1R17SlPriorityRequestR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PriorityRequest_r17 = &v
		}
	}

	// 15. sl-PriorityPreferredResourceSet-r17: integer
	{
		if seq.IsComponentPresent(12) {
			v, err := d.DecodeInteger(sLInterUECoordinationScheme1R17SlPriorityPreferredResourceSetR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PriorityPreferredResourceSet_r17 = &v
		}
	}

	// 16. sl-MaxSlotOffsetTRIV-r17: integer
	{
		if seq.IsComponentPresent(13) {
			v, err := d.DecodeInteger(sLInterUECoordinationScheme1R17SlMaxSlotOffsetTRIVR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MaxSlotOffsetTRIV_r17 = &v
		}
	}

	// 17. sl-NumSubCH-PreferredResourceSet-r17: integer
	{
		if seq.IsComponentPresent(14) {
			v, err := d.DecodeInteger(sLInterUECoordinationScheme1R17SlNumSubCHPreferredResourceSetR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumSubCH_PreferredResourceSet_r17 = &v
		}
	}

	// 18. sl-ReservedPeriodPreferredResourceSet-r17: integer
	{
		if seq.IsComponentPresent(15) {
			v, err := d.DecodeInteger(sLInterUECoordinationScheme1R17SlReservedPeriodPreferredResourceSetR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ReservedPeriodPreferredResourceSet_r17 = &v
		}
	}

	// 19. sl-DetermineResourceType-r17: enumerated
	{
		if seq.IsComponentPresent(16) {
			idx, err := d.DecodeEnumerated(sLInterUECoordinationScheme1R17SlDetermineResourceTypeR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DetermineResourceType_r17 = &idx
		}
	}

	return nil
}
