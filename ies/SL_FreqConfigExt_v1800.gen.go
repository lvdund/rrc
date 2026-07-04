// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-FreqConfigExt-v1800 (line 27272).

var sLFreqConfigExtV1800Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "absenceOfAnyOtherTechnology-r18", Optional: true},
		{Name: "sl-FreqSelectionConfigList-r18", Optional: true},
		{Name: "sl-SyncTxDisabled-r18", Optional: true},
		{Name: "sl-EnergyDetectionConfig-r18", Optional: true},
		{Name: "ue-ToUE-COT-SharingED-Threshold-r18", Optional: true},
		{Name: "harq-ACK-FeedbackRatioforCW-AdjustmentGC-Option2-r18", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	SL_FreqConfigExt_v1800_AbsenceOfAnyOtherTechnology_r18_True = 0
)

var sLFreqConfigExtV1800AbsenceOfAnyOtherTechnologyR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_FreqConfigExt_v1800_AbsenceOfAnyOtherTechnology_r18_True},
}

var sLFreqConfigExtV1800SlFreqSelectionConfigListR18Constraints = per.SizeRange(1, 8)

const (
	SL_FreqConfigExt_v1800_Sl_SyncTxDisabled_r18_True = 0
)

var sLFreqConfigExtV1800SlSyncTxDisabledR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_FreqConfigExt_v1800_Sl_SyncTxDisabled_r18_True},
}

var sL_FreqConfigExt_v1800SlEnergyDetectionConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-MaxEnergyDetectionThreshold-r18"},
		{Name: "sl-EnergyDetectionThresholdOffset-r18"},
	},
}

const (
	SL_FreqConfigExt_v1800_Sl_EnergyDetectionConfig_r18_Sl_MaxEnergyDetectionThreshold_r18    = 0
	SL_FreqConfigExt_v1800_Sl_EnergyDetectionConfig_r18_Sl_EnergyDetectionThresholdOffset_r18 = 1
)

var sLFreqConfigExtV1800UeToUECOTSharingEDThresholdR18Constraints = per.Constrained(-85, -52)

var sLFreqConfigExtV1800HarqACKFeedbackRatioforCWAdjustmentGCOption2R18Constraints = per.Constrained(10, 100)

type SL_FreqConfigExt_v1800 struct {
	AbsenceOfAnyOtherTechnology_r18 *int64
	Sl_FreqSelectionConfigList_r18  []SL_FreqSelectionConfig_r18
	Sl_SyncTxDisabled_r18           *int64
	Sl_EnergyDetectionConfig_r18    *struct {
		Choice                                int
		Sl_MaxEnergyDetectionThreshold_r18    *int64
		Sl_EnergyDetectionThresholdOffset_r18 *int64
	}
	Ue_ToUE_COT_SharingED_Threshold_r18                  *int64
	Harq_ACK_FeedbackRatioforCW_AdjustmentGC_Option2_r18 *int64
	AdditionalSpectrumEmission_v1860                     *AdditionalSpectrumEmission_v1760
}

func (ie *SL_FreqConfigExt_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLFreqConfigExtV1800Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.AdditionalSpectrumEmission_v1860 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AbsenceOfAnyOtherTechnology_r18 != nil, ie.Sl_FreqSelectionConfigList_r18 != nil, ie.Sl_SyncTxDisabled_r18 != nil, ie.Sl_EnergyDetectionConfig_r18 != nil, ie.Ue_ToUE_COT_SharingED_Threshold_r18 != nil, ie.Harq_ACK_FeedbackRatioforCW_AdjustmentGC_Option2_r18 != nil}); err != nil {
		return err
	}

	// 3. absenceOfAnyOtherTechnology-r18: enumerated
	{
		if ie.AbsenceOfAnyOtherTechnology_r18 != nil {
			if err := e.EncodeEnumerated(*ie.AbsenceOfAnyOtherTechnology_r18, sLFreqConfigExtV1800AbsenceOfAnyOtherTechnologyR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-FreqSelectionConfigList-r18: sequence-of
	{
		if ie.Sl_FreqSelectionConfigList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLFreqConfigExtV1800SlFreqSelectionConfigListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_FreqSelectionConfigList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_FreqSelectionConfigList_r18 {
				if err := ie.Sl_FreqSelectionConfigList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-SyncTxDisabled-r18: enumerated
	{
		if ie.Sl_SyncTxDisabled_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_SyncTxDisabled_r18, sLFreqConfigExtV1800SlSyncTxDisabledR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-EnergyDetectionConfig-r18: choice
	{
		if ie.Sl_EnergyDetectionConfig_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_FreqConfigExt_v1800SlEnergyDetectionConfigR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_EnergyDetectionConfig_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_EnergyDetectionConfig_r18).Choice {
			case SL_FreqConfigExt_v1800_Sl_EnergyDetectionConfig_r18_Sl_MaxEnergyDetectionThreshold_r18:
				if err := e.EncodeInteger((*(*ie.Sl_EnergyDetectionConfig_r18).Sl_MaxEnergyDetectionThreshold_r18), per.Constrained(-85, -52)); err != nil {
					return err
				}
			case SL_FreqConfigExt_v1800_Sl_EnergyDetectionConfig_r18_Sl_EnergyDetectionThresholdOffset_r18:
				if err := e.EncodeInteger((*(*ie.Sl_EnergyDetectionConfig_r18).Sl_EnergyDetectionThresholdOffset_r18), per.Constrained(-13, 20)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_EnergyDetectionConfig_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 7. ue-ToUE-COT-SharingED-Threshold-r18: integer
	{
		if ie.Ue_ToUE_COT_SharingED_Threshold_r18 != nil {
			if err := e.EncodeInteger(*ie.Ue_ToUE_COT_SharingED_Threshold_r18, sLFreqConfigExtV1800UeToUECOTSharingEDThresholdR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. harq-ACK-FeedbackRatioforCW-AdjustmentGC-Option2-r18: integer
	{
		if ie.Harq_ACK_FeedbackRatioforCW_AdjustmentGC_Option2_r18 != nil {
			if err := e.EncodeInteger(*ie.Harq_ACK_FeedbackRatioforCW_AdjustmentGC_Option2_r18, sLFreqConfigExtV1800HarqACKFeedbackRatioforCWAdjustmentGCOption2R18Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "additionalSpectrumEmission-v1860", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AdditionalSpectrumEmission_v1860 != nil}); err != nil {
				return err
			}

			if ie.AdditionalSpectrumEmission_v1860 != nil {
				if err := ie.AdditionalSpectrumEmission_v1860.Encode(ex); err != nil {
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

func (ie *SL_FreqConfigExt_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLFreqConfigExtV1800Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. absenceOfAnyOtherTechnology-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sLFreqConfigExtV1800AbsenceOfAnyOtherTechnologyR18Constraints)
			if err != nil {
				return err
			}
			ie.AbsenceOfAnyOtherTechnology_r18 = &idx
		}
	}

	// 4. sl-FreqSelectionConfigList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLFreqConfigExtV1800SlFreqSelectionConfigListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_FreqSelectionConfigList_r18 = make([]SL_FreqSelectionConfig_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_FreqSelectionConfigList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-SyncTxDisabled-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLFreqConfigExtV1800SlSyncTxDisabledR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SyncTxDisabled_r18 = &idx
		}
	}

	// 6. sl-EnergyDetectionConfig-r18: choice
	{
		if seq.IsComponentPresent(3) {
			ie.Sl_EnergyDetectionConfig_r18 = &struct {
				Choice                                int
				Sl_MaxEnergyDetectionThreshold_r18    *int64
				Sl_EnergyDetectionThresholdOffset_r18 *int64
			}{}
			choiceDec := d.NewChoiceDecoder(sL_FreqConfigExt_v1800SlEnergyDetectionConfigR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_EnergyDetectionConfig_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_FreqConfigExt_v1800_Sl_EnergyDetectionConfig_r18_Sl_MaxEnergyDetectionThreshold_r18:
				(*ie.Sl_EnergyDetectionConfig_r18).Sl_MaxEnergyDetectionThreshold_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(-85, -52))
				if err != nil {
					return err
				}
				(*(*ie.Sl_EnergyDetectionConfig_r18).Sl_MaxEnergyDetectionThreshold_r18) = v
			case SL_FreqConfigExt_v1800_Sl_EnergyDetectionConfig_r18_Sl_EnergyDetectionThresholdOffset_r18:
				(*ie.Sl_EnergyDetectionConfig_r18).Sl_EnergyDetectionThresholdOffset_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(-13, 20))
				if err != nil {
					return err
				}
				(*(*ie.Sl_EnergyDetectionConfig_r18).Sl_EnergyDetectionThresholdOffset_r18) = v
			}
		}
	}

	// 7. ue-ToUE-COT-SharingED-Threshold-r18: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(sLFreqConfigExtV1800UeToUECOTSharingEDThresholdR18Constraints)
			if err != nil {
				return err
			}
			ie.Ue_ToUE_COT_SharingED_Threshold_r18 = &v
		}
	}

	// 8. harq-ACK-FeedbackRatioforCW-AdjustmentGC-Option2-r18: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(sLFreqConfigExtV1800HarqACKFeedbackRatioforCWAdjustmentGCOption2R18Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ACK_FeedbackRatioforCW_AdjustmentGC_Option2_r18 = &v
		}
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
				{Name: "additionalSpectrumEmission-v1860", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.AdditionalSpectrumEmission_v1860 = new(AdditionalSpectrumEmission_v1760)
			if err := ie.AdditionalSpectrumEmission_v1860.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
