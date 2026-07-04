// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-FreqConfigCommon-r16 (line 27291).

var sLFreqConfigCommonR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-SCS-SpecificCarrierList-r16"},
		{Name: "sl-AbsoluteFrequencyPointA-r16"},
		{Name: "sl-AbsoluteFrequencySSB-r16", Optional: true},
		{Name: "frequencyShift7p5khzSL-r16", Optional: true},
		{Name: "valueN-r16"},
		{Name: "sl-BWP-List-r16", Optional: true},
		{Name: "sl-SyncPriority-r16", Optional: true},
		{Name: "sl-NbAsSync-r16", Optional: true},
		{Name: "sl-SyncConfigList-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var sLFreqConfigCommonR16SlSCSSpecificCarrierListR16Constraints = per.SizeRange(1, common.MaxSCSs)

const (
	SL_FreqConfigCommon_r16_FrequencyShift7p5khzSL_r16_True = 0
)

var sLFreqConfigCommonR16FrequencyShift7p5khzSLR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_FreqConfigCommon_r16_FrequencyShift7p5khzSL_r16_True},
}

var sLFreqConfigCommonR16ValueNR16Constraints = per.Constrained(-1, 1)

var sLFreqConfigCommonR16SlBWPListR16Constraints = per.SizeRange(1, common.MaxNrofSL_BWPs_r16)

const (
	SL_FreqConfigCommon_r16_Sl_SyncPriority_r16_Gnss   = 0
	SL_FreqConfigCommon_r16_Sl_SyncPriority_r16_GnbEnb = 1
)

var sLFreqConfigCommonR16SlSyncPriorityR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_FreqConfigCommon_r16_Sl_SyncPriority_r16_Gnss, SL_FreqConfigCommon_r16_Sl_SyncPriority_r16_GnbEnb},
}

var sLFreqConfigCommonR16ExtSlPosBWPListR18Constraints = per.SizeRange(1, common.MaxNrofSL_BWPs_r16)

var sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "absenceOfAnyOtherTechnology-r18", Optional: true},
		{Name: "sl-FreqSelectionConfigList-r18", Optional: true},
		{Name: "sl-SyncTxDisabled-r18", Optional: true},
		{Name: "sl-EnergyDetectionConfig-r18", Optional: true},
		{Name: "ue-ToUE-COT-SharingED-Threshold-r18", Optional: true},
		{Name: "harq-ACK-FeedbackRatioforCW-AdjustmentGC-Option2-r18", Optional: true},
	},
}

const (
	SL_FreqConfigCommon_r16_Ext_Sl_UnlicensedFreqConfigCommon_r18_AbsenceOfAnyOtherTechnology_r18_True = 0
)

var sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18AbsenceOfAnyOtherTechnologyR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_FreqConfigCommon_r16_Ext_Sl_UnlicensedFreqConfigCommon_r18_AbsenceOfAnyOtherTechnology_r18_True},
}

var sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18SlFreqSelectionConfigListR18Constraints = per.SizeRange(1, 8)

const (
	SL_FreqConfigCommon_r16_Ext_Sl_UnlicensedFreqConfigCommon_r18_Sl_SyncTxDisabled_r18_True = 0
)

var sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18SlSyncTxDisabledR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_FreqConfigCommon_r16_Ext_Sl_UnlicensedFreqConfigCommon_r18_Sl_SyncTxDisabled_r18_True},
}

var sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18SlEnergyDetectionConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-MaxEnergyDetectionThreshold-r18"},
		{Name: "sl-EnergyDetectionThresholdOffset-r18"},
	},
}

const (
	SL_FreqConfigCommon_r16_Ext_Sl_UnlicensedFreqConfigCommon_r18_Sl_EnergyDetectionConfig_r18_Sl_MaxEnergyDetectionThreshold_r18    = 0
	SL_FreqConfigCommon_r16_Ext_Sl_UnlicensedFreqConfigCommon_r18_Sl_EnergyDetectionConfig_r18_Sl_EnergyDetectionThresholdOffset_r18 = 1
)

type SL_FreqConfigCommon_r16 struct {
	Sl_SCS_SpecificCarrierList_r16    []SCS_SpecificCarrier
	Sl_AbsoluteFrequencyPointA_r16    ARFCN_ValueNR
	Sl_AbsoluteFrequencySSB_r16       *ARFCN_ValueNR
	FrequencyShift7p5khzSL_r16        *int64
	ValueN_r16                        int64
	Sl_BWP_List_r16                   []SL_BWP_ConfigCommon_r16
	Sl_SyncPriority_r16               *int64
	Sl_NbAsSync_r16                   *bool
	Sl_SyncConfigList_r16             *SL_SyncConfigList_r16
	Sl_UnlicensedFreqConfigCommon_r18 *struct {
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
	}
	Sl_PosBWP_List_r18               []SL_PosBWP_ConfigCommon_r18
	AdditionalSpectrumEmission_v1860 *AdditionalSpectrumEmission_v1760
}

func (ie *SL_FreqConfigCommon_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLFreqConfigCommonR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_UnlicensedFreqConfigCommon_r18 != nil || ie.Sl_PosBWP_List_r18 != nil
	hasExtGroup1 := ie.AdditionalSpectrumEmission_v1860 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_AbsoluteFrequencySSB_r16 != nil, ie.FrequencyShift7p5khzSL_r16 != nil, ie.Sl_BWP_List_r16 != nil, ie.Sl_SyncPriority_r16 != nil, ie.Sl_NbAsSync_r16 != nil, ie.Sl_SyncConfigList_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-SCS-SpecificCarrierList-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sLFreqConfigCommonR16SlSCSSpecificCarrierListR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sl_SCS_SpecificCarrierList_r16))); err != nil {
			return err
		}
		for i := range ie.Sl_SCS_SpecificCarrierList_r16 {
			if err := ie.Sl_SCS_SpecificCarrierList_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-AbsoluteFrequencyPointA-r16: ref
	{
		if err := ie.Sl_AbsoluteFrequencyPointA_r16.Encode(e); err != nil {
			return err
		}
	}

	// 5. sl-AbsoluteFrequencySSB-r16: ref
	{
		if ie.Sl_AbsoluteFrequencySSB_r16 != nil {
			if err := ie.Sl_AbsoluteFrequencySSB_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. frequencyShift7p5khzSL-r16: enumerated
	{
		if ie.FrequencyShift7p5khzSL_r16 != nil {
			if err := e.EncodeEnumerated(*ie.FrequencyShift7p5khzSL_r16, sLFreqConfigCommonR16FrequencyShift7p5khzSLR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. valueN-r16: integer
	{
		if err := e.EncodeInteger(ie.ValueN_r16, sLFreqConfigCommonR16ValueNR16Constraints); err != nil {
			return err
		}
	}

	// 8. sl-BWP-List-r16: sequence-of
	{
		if ie.Sl_BWP_List_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLFreqConfigCommonR16SlBWPListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_BWP_List_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_BWP_List_r16 {
				if err := ie.Sl_BWP_List_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 9. sl-SyncPriority-r16: enumerated
	{
		if ie.Sl_SyncPriority_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_SyncPriority_r16, sLFreqConfigCommonR16SlSyncPriorityR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. sl-NbAsSync-r16: boolean
	{
		if ie.Sl_NbAsSync_r16 != nil {
			if err := e.EncodeBoolean(*ie.Sl_NbAsSync_r16); err != nil {
				return err
			}
		}
	}

	// 11. sl-SyncConfigList-r16: ref
	{
		if ie.Sl_SyncConfigList_r16 != nil {
			if err := ie.Sl_SyncConfigList_r16.Encode(e); err != nil {
				return err
			}
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
					{Name: "sl-UnlicensedFreqConfigCommon-r18", Optional: true},
					{Name: "sl-PosBWP-List-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_UnlicensedFreqConfigCommon_r18 != nil, ie.Sl_PosBWP_List_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_UnlicensedFreqConfigCommon_r18 != nil {
				c := ie.Sl_UnlicensedFreqConfigCommon_r18
				sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18Seq := ex.NewSequenceEncoder(sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18Constraints)
				if err := sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18Seq.EncodePreamble([]bool{c.AbsenceOfAnyOtherTechnology_r18 != nil, c.Sl_FreqSelectionConfigList_r18 != nil, c.Sl_SyncTxDisabled_r18 != nil, c.Sl_EnergyDetectionConfig_r18 != nil, c.Ue_ToUE_COT_SharingED_Threshold_r18 != nil, c.Harq_ACK_FeedbackRatioforCW_AdjustmentGC_Option2_r18 != nil}); err != nil {
					return err
				}
				if c.AbsenceOfAnyOtherTechnology_r18 != nil {
					if err := ex.EncodeEnumerated((*c.AbsenceOfAnyOtherTechnology_r18), sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18AbsenceOfAnyOtherTechnologyR18Constraints); err != nil {
						return err
					}
				}
				if c.Sl_FreqSelectionConfigList_r18 != nil {
					seqOf := ex.NewSequenceOfEncoder(sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18SlFreqSelectionConfigListR18Constraints)
					if err := seqOf.EncodeLength(int64(len(c.Sl_FreqSelectionConfigList_r18))); err != nil {
						return err
					}
					for i := range c.Sl_FreqSelectionConfigList_r18 {
						if err := c.Sl_FreqSelectionConfigList_r18[i].Encode(ex); err != nil {
							return err
						}
					}
				}
				if c.Sl_SyncTxDisabled_r18 != nil {
					if err := ex.EncodeEnumerated((*c.Sl_SyncTxDisabled_r18), sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18SlSyncTxDisabledR18Constraints); err != nil {
						return err
					}
				}
				if c.Sl_EnergyDetectionConfig_r18 != nil {
					choiceEnc := ex.NewChoiceEncoder(sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18SlEnergyDetectionConfigR18Constraints)
					if err := choiceEnc.EncodeChoice(int64((*c.Sl_EnergyDetectionConfig_r18).Choice), false, nil); err != nil {
						return err
					}
					switch (*c.Sl_EnergyDetectionConfig_r18).Choice {
					case SL_FreqConfigCommon_r16_Ext_Sl_UnlicensedFreqConfigCommon_r18_Sl_EnergyDetectionConfig_r18_Sl_MaxEnergyDetectionThreshold_r18:
						if err := ex.EncodeInteger((*(*c.Sl_EnergyDetectionConfig_r18).Sl_MaxEnergyDetectionThreshold_r18), per.Constrained(-85, -52)); err != nil {
							return err
						}
					case SL_FreqConfigCommon_r16_Ext_Sl_UnlicensedFreqConfigCommon_r18_Sl_EnergyDetectionConfig_r18_Sl_EnergyDetectionThresholdOffset_r18:
						if err := ex.EncodeInteger((*(*c.Sl_EnergyDetectionConfig_r18).Sl_EnergyDetectionThresholdOffset_r18), per.Constrained(-13, 20)); err != nil {
							return err
						}
					}
				}
				if c.Ue_ToUE_COT_SharingED_Threshold_r18 != nil {
					if err := ex.EncodeInteger((*c.Ue_ToUE_COT_SharingED_Threshold_r18), per.Constrained(-85, -52)); err != nil {
						return err
					}
				}
				if c.Harq_ACK_FeedbackRatioforCW_AdjustmentGC_Option2_r18 != nil {
					if err := ex.EncodeInteger((*c.Harq_ACK_FeedbackRatioforCW_AdjustmentGC_Option2_r18), per.Constrained(10, 100)); err != nil {
						return err
					}
				}
			}

			if ie.Sl_PosBWP_List_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLFreqConfigCommonR16ExtSlPosBWPListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_PosBWP_List_r18))); err != nil {
					return err
				}
				for i := range ie.Sl_PosBWP_List_r18 {
					if err := ie.Sl_PosBWP_List_r18[i].Encode(ex); err != nil {
						return err
					}
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

func (ie *SL_FreqConfigCommon_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLFreqConfigCommonR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-SCS-SpecificCarrierList-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sLFreqConfigCommonR16SlSCSSpecificCarrierListR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sl_SCS_SpecificCarrierList_r16 = make([]SCS_SpecificCarrier, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Sl_SCS_SpecificCarrierList_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-AbsoluteFrequencyPointA-r16: ref
	{
		if err := ie.Sl_AbsoluteFrequencyPointA_r16.Decode(d); err != nil {
			return err
		}
	}

	// 5. sl-AbsoluteFrequencySSB-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_AbsoluteFrequencySSB_r16 = new(ARFCN_ValueNR)
			if err := ie.Sl_AbsoluteFrequencySSB_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. frequencyShift7p5khzSL-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sLFreqConfigCommonR16FrequencyShift7p5khzSLR16Constraints)
			if err != nil {
				return err
			}
			ie.FrequencyShift7p5khzSL_r16 = &idx
		}
	}

	// 7. valueN-r16: integer
	{
		v4, err := d.DecodeInteger(sLFreqConfigCommonR16ValueNR16Constraints)
		if err != nil {
			return err
		}
		ie.ValueN_r16 = v4
	}

	// 8. sl-BWP-List-r16: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(sLFreqConfigCommonR16SlBWPListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_BWP_List_r16 = make([]SL_BWP_ConfigCommon_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_BWP_List_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. sl-SyncPriority-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(sLFreqConfigCommonR16SlSyncPriorityR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SyncPriority_r16 = &idx
		}
	}

	// 10. sl-NbAsSync-r16: boolean
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.Sl_NbAsSync_r16 = &v
		}
	}

	// 11. sl-SyncConfigList-r16: ref
	{
		if seq.IsComponentPresent(8) {
			ie.Sl_SyncConfigList_r16 = new(SL_SyncConfigList_r16)
			if err := ie.Sl_SyncConfigList_r16.Decode(d); err != nil {
				return err
			}
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
				{Name: "sl-UnlicensedFreqConfigCommon-r18", Optional: true},
				{Name: "sl-PosBWP-List-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_UnlicensedFreqConfigCommon_r18 = &struct {
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
			}{}
			c := ie.Sl_UnlicensedFreqConfigCommon_r18
			sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18Seq := dx.NewSequenceDecoder(sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18Constraints)
			if err := sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18Seq.IsComponentPresent(0) {
				c.AbsenceOfAnyOtherTechnology_r18 = new(int64)
				v, err := dx.DecodeEnumerated(sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18AbsenceOfAnyOtherTechnologyR18Constraints)
				if err != nil {
					return err
				}
				(*c.AbsenceOfAnyOtherTechnology_r18) = v
			}
			if sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18Seq.IsComponentPresent(1) {
				seqOf := dx.NewSequenceOfDecoder(sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18SlFreqSelectionConfigListR18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Sl_FreqSelectionConfigList_r18 = make([]SL_FreqSelectionConfig_r18, n)
				for i := int64(0); i < n; i++ {
					if err := c.Sl_FreqSelectionConfigList_r18[i].Decode(dx); err != nil {
						return err
					}
				}
			}
			if sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18Seq.IsComponentPresent(2) {
				c.Sl_SyncTxDisabled_r18 = new(int64)
				v, err := dx.DecodeEnumerated(sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18SlSyncTxDisabledR18Constraints)
				if err != nil {
					return err
				}
				(*c.Sl_SyncTxDisabled_r18) = v
			}
			if sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18Seq.IsComponentPresent(3) {
				c.Sl_EnergyDetectionConfig_r18 = &struct {
					Choice                                int
					Sl_MaxEnergyDetectionThreshold_r18    *int64
					Sl_EnergyDetectionThresholdOffset_r18 *int64
				}{}
				choiceDec := dx.NewChoiceDecoder(sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18SlEnergyDetectionConfigR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.Sl_EnergyDetectionConfig_r18).Choice = int(index)
				switch index {
				case SL_FreqConfigCommon_r16_Ext_Sl_UnlicensedFreqConfigCommon_r18_Sl_EnergyDetectionConfig_r18_Sl_MaxEnergyDetectionThreshold_r18:
					(*c.Sl_EnergyDetectionConfig_r18).Sl_MaxEnergyDetectionThreshold_r18 = new(int64)
					v, err := dx.DecodeInteger(per.Constrained(-85, -52))
					if err != nil {
						return err
					}
					(*(*c.Sl_EnergyDetectionConfig_r18).Sl_MaxEnergyDetectionThreshold_r18) = v
				case SL_FreqConfigCommon_r16_Ext_Sl_UnlicensedFreqConfigCommon_r18_Sl_EnergyDetectionConfig_r18_Sl_EnergyDetectionThresholdOffset_r18:
					(*c.Sl_EnergyDetectionConfig_r18).Sl_EnergyDetectionThresholdOffset_r18 = new(int64)
					v, err := dx.DecodeInteger(per.Constrained(-13, 20))
					if err != nil {
						return err
					}
					(*(*c.Sl_EnergyDetectionConfig_r18).Sl_EnergyDetectionThresholdOffset_r18) = v
				}
			}
			if sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18Seq.IsComponentPresent(4) {
				c.Ue_ToUE_COT_SharingED_Threshold_r18 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(-85, -52))
				if err != nil {
					return err
				}
				(*c.Ue_ToUE_COT_SharingED_Threshold_r18) = v
			}
			if sLFreqConfigCommonR16ExtSlUnlicensedFreqConfigCommonR18Seq.IsComponentPresent(5) {
				c.Harq_ACK_FeedbackRatioforCW_AdjustmentGC_Option2_r18 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(10, 100))
				if err != nil {
					return err
				}
				(*c.Harq_ACK_FeedbackRatioforCW_AdjustmentGC_Option2_r18) = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(sLFreqConfigCommonR16ExtSlPosBWPListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PosBWP_List_r18 = make([]SL_PosBWP_ConfigCommon_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PosBWP_List_r18[i].Decode(dx); err != nil {
					return err
				}
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
