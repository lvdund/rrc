// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-UE-SelectedConfigRP-r16 (line 28058).

var sLUESelectedConfigRPR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-CBR-PriorityTxConfigList-r16", Optional: true},
		{Name: "sl-Thres-RSRP-List-r16", Optional: true},
		{Name: "sl-MultiReserveResource-r16", Optional: true},
		{Name: "sl-MaxNumPerReserve-r16", Optional: true},
		{Name: "sl-SensingWindow-r16", Optional: true},
		{Name: "sl-SelectionWindowList-r16", Optional: true},
		{Name: "sl-ResourceReservePeriodList-r16", Optional: true},
		{Name: "sl-RS-ForSensing-r16"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	SL_UE_SelectedConfigRP_r16_Sl_MultiReserveResource_r16_Enabled = 0
)

var sLUESelectedConfigRPR16SlMultiReserveResourceR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_UE_SelectedConfigRP_r16_Sl_MultiReserveResource_r16_Enabled},
}

const (
	SL_UE_SelectedConfigRP_r16_Sl_MaxNumPerReserve_r16_N2 = 0
	SL_UE_SelectedConfigRP_r16_Sl_MaxNumPerReserve_r16_N3 = 1
)

var sLUESelectedConfigRPR16SlMaxNumPerReserveR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_UE_SelectedConfigRP_r16_Sl_MaxNumPerReserve_r16_N2, SL_UE_SelectedConfigRP_r16_Sl_MaxNumPerReserve_r16_N3},
}

const (
	SL_UE_SelectedConfigRP_r16_Sl_SensingWindow_r16_Ms100  = 0
	SL_UE_SelectedConfigRP_r16_Sl_SensingWindow_r16_Ms1100 = 1
)

var sLUESelectedConfigRPR16SlSensingWindowR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_UE_SelectedConfigRP_r16_Sl_SensingWindow_r16_Ms100, SL_UE_SelectedConfigRP_r16_Sl_SensingWindow_r16_Ms1100},
}

var sLUESelectedConfigRPR16SlResourceReservePeriodListR16Constraints = per.SizeRange(1, 16)

const (
	SL_UE_SelectedConfigRP_r16_Sl_RS_ForSensing_r16_Pscch = 0
	SL_UE_SelectedConfigRP_r16_Sl_RS_ForSensing_r16_Pssch = 1
)

var sLUESelectedConfigRPR16SlRSForSensingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_UE_SelectedConfigRP_r16_Sl_RS_ForSensing_r16_Pscch, SL_UE_SelectedConfigRP_r16_Sl_RS_ForSensing_r16_Pssch},
}

type SL_UE_SelectedConfigRP_r16 struct {
	Sl_CBR_PriorityTxConfigList_r16     *SL_CBR_PriorityTxConfigList_r16
	Sl_Thres_RSRP_List_r16              *SL_Thres_RSRP_List_r16
	Sl_MultiReserveResource_r16         *int64
	Sl_MaxNumPerReserve_r16             *int64
	Sl_SensingWindow_r16                *int64
	Sl_SelectionWindowList_r16          *SL_SelectionWindowList_r16
	Sl_ResourceReservePeriodList_r16    []SL_ResourceReservePeriod_r16
	Sl_RS_ForSensing_r16                int64
	Sl_CBR_PriorityTxConfigList_v1650   *SL_CBR_PriorityTxConfigList_v1650
	Sl_NRPSSCH_EUTRA_ThresRSRP_List_r18 *SL_Thres_RSRP_List_r16
	Sl_NRPSFCH_EUTRA_ThresRSRP_List_r18 *SL_Thres_RSRP_List_r16
}

func (ie *SL_UE_SelectedConfigRP_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLUESelectedConfigRPR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_CBR_PriorityTxConfigList_v1650 != nil
	hasExtGroup1 := ie.Sl_NRPSSCH_EUTRA_ThresRSRP_List_r18 != nil || ie.Sl_NRPSFCH_EUTRA_ThresRSRP_List_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_CBR_PriorityTxConfigList_r16 != nil, ie.Sl_Thres_RSRP_List_r16 != nil, ie.Sl_MultiReserveResource_r16 != nil, ie.Sl_MaxNumPerReserve_r16 != nil, ie.Sl_SensingWindow_r16 != nil, ie.Sl_SelectionWindowList_r16 != nil, ie.Sl_ResourceReservePeriodList_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-CBR-PriorityTxConfigList-r16: ref
	{
		if ie.Sl_CBR_PriorityTxConfigList_r16 != nil {
			if err := ie.Sl_CBR_PriorityTxConfigList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-Thres-RSRP-List-r16: ref
	{
		if ie.Sl_Thres_RSRP_List_r16 != nil {
			if err := ie.Sl_Thres_RSRP_List_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-MultiReserveResource-r16: enumerated
	{
		if ie.Sl_MultiReserveResource_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_MultiReserveResource_r16, sLUESelectedConfigRPR16SlMultiReserveResourceR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-MaxNumPerReserve-r16: enumerated
	{
		if ie.Sl_MaxNumPerReserve_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_MaxNumPerReserve_r16, sLUESelectedConfigRPR16SlMaxNumPerReserveR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sl-SensingWindow-r16: enumerated
	{
		if ie.Sl_SensingWindow_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_SensingWindow_r16, sLUESelectedConfigRPR16SlSensingWindowR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. sl-SelectionWindowList-r16: ref
	{
		if ie.Sl_SelectionWindowList_r16 != nil {
			if err := ie.Sl_SelectionWindowList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. sl-ResourceReservePeriodList-r16: sequence-of
	{
		if ie.Sl_ResourceReservePeriodList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLUESelectedConfigRPR16SlResourceReservePeriodListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_ResourceReservePeriodList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_ResourceReservePeriodList_r16 {
				if err := ie.Sl_ResourceReservePeriodList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 10. sl-RS-ForSensing-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_RS_ForSensing_r16, sLUESelectedConfigRPR16SlRSForSensingR16Constraints); err != nil {
			return err
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-CBR-PriorityTxConfigList-v1650", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_CBR_PriorityTxConfigList_v1650 != nil}); err != nil {
				return err
			}

			if ie.Sl_CBR_PriorityTxConfigList_v1650 != nil {
				if err := ie.Sl_CBR_PriorityTxConfigList_v1650.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-NRPSSCH-EUTRA-ThresRSRP-List-r18", Optional: true},
					{Name: "sl-NRPSFCH-EUTRA-ThresRSRP-List-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_NRPSSCH_EUTRA_ThresRSRP_List_r18 != nil, ie.Sl_NRPSFCH_EUTRA_ThresRSRP_List_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_NRPSSCH_EUTRA_ThresRSRP_List_r18 != nil {
				if err := ie.Sl_NRPSSCH_EUTRA_ThresRSRP_List_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_NRPSFCH_EUTRA_ThresRSRP_List_r18 != nil {
				if err := ie.Sl_NRPSFCH_EUTRA_ThresRSRP_List_r18.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_UE_SelectedConfigRP_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLUESelectedConfigRPR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-CBR-PriorityTxConfigList-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_CBR_PriorityTxConfigList_r16 = new(SL_CBR_PriorityTxConfigList_r16)
			if err := ie.Sl_CBR_PriorityTxConfigList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-Thres-RSRP-List-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_Thres_RSRP_List_r16 = new(SL_Thres_RSRP_List_r16)
			if err := ie.Sl_Thres_RSRP_List_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-MultiReserveResource-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLUESelectedConfigRPR16SlMultiReserveResourceR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MultiReserveResource_r16 = &idx
		}
	}

	// 6. sl-MaxNumPerReserve-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sLUESelectedConfigRPR16SlMaxNumPerReserveR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MaxNumPerReserve_r16 = &idx
		}
	}

	// 7. sl-SensingWindow-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(sLUESelectedConfigRPR16SlSensingWindowR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SensingWindow_r16 = &idx
		}
	}

	// 8. sl-SelectionWindowList-r16: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Sl_SelectionWindowList_r16 = new(SL_SelectionWindowList_r16)
			if err := ie.Sl_SelectionWindowList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. sl-ResourceReservePeriodList-r16: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(sLUESelectedConfigRPR16SlResourceReservePeriodListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_ResourceReservePeriodList_r16 = make([]SL_ResourceReservePeriod_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_ResourceReservePeriodList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. sl-RS-ForSensing-r16: enumerated
	{
		v7, err := d.DecodeEnumerated(sLUESelectedConfigRPR16SlRSForSensingR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_RS_ForSensing_r16 = v7
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sl-CBR-PriorityTxConfigList-v1650", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_CBR_PriorityTxConfigList_v1650 = new(SL_CBR_PriorityTxConfigList_v1650)
			if err := ie.Sl_CBR_PriorityTxConfigList_v1650.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sl-NRPSSCH-EUTRA-ThresRSRP-List-r18", Optional: true},
				{Name: "sl-NRPSFCH-EUTRA-ThresRSRP-List-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_NRPSSCH_EUTRA_ThresRSRP_List_r18 = new(SL_Thres_RSRP_List_r16)
			if err := ie.Sl_NRPSSCH_EUTRA_ThresRSRP_List_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sl_NRPSFCH_EUTRA_ThresRSRP_List_r18 = new(SL_Thres_RSRP_List_r16)
			if err := ie.Sl_NRPSFCH_EUTRA_ThresRSRP_List_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
