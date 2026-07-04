// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasConfig (line 9095).

var measConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measObjectToRemoveList", Optional: true},
		{Name: "measObjectToAddModList", Optional: true},
		{Name: "reportConfigToRemoveList", Optional: true},
		{Name: "reportConfigToAddModList", Optional: true},
		{Name: "measIdToRemoveList", Optional: true},
		{Name: "measIdToAddModList", Optional: true},
		{Name: "s-MeasureConfig", Optional: true},
		{Name: "quantityConfig", Optional: true},
		{Name: "measGapConfig", Optional: true},
		{Name: "measGapSharingConfig", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var measConfigSMeasureConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-RSRP"},
		{Name: "csi-RSRP"},
	},
}

const (
	MeasConfig_S_MeasureConfig_Ssb_RSRP = 0
	MeasConfig_S_MeasureConfig_Csi_RSRP = 1
)

const (
	MeasConfig_Ext_InterFrequencyConfig_NoGap_r16_True = 0
)

var measConfigExtInterFrequencyConfigNoGapR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasConfig_Ext_InterFrequencyConfig_NoGap_r16_True},
}

var measConfigExtEffectiveMeasWindowConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MeasConfig_Ext_EffectiveMeasWindowConfig_r18_Release = 0
	MeasConfig_Ext_EffectiveMeasWindowConfig_r18_Setup   = 1
)

var measConfigExtFbsConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MeasConfig_Ext_Fbs_Config_r19_Release = 0
	MeasConfig_Ext_Fbs_Config_r19_Setup   = 1
)

var measConfigExtSpeedStateParsR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MeasConfig_Ext_SpeedStatePars_r19_Release = 0
	MeasConfig_Ext_SpeedStatePars_r19_Setup   = 1
)

var measConfigExtCssfConfigR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cssf-MeasMO-List-r19", Optional: true},
	},
}

type MeasConfig struct {
	MeasObjectToRemoveList   *MeasObjectToRemoveList
	MeasObjectToAddModList   *MeasObjectToAddModList
	ReportConfigToRemoveList *ReportConfigToRemoveList
	ReportConfigToAddModList *ReportConfigToAddModList
	MeasIdToRemoveList       *MeasIdToRemoveList
	MeasIdToAddModList       *MeasIdToAddModList
	S_MeasureConfig          *struct {
		Choice   int
		Ssb_RSRP *RSRP_Range
		Csi_RSRP *RSRP_Range
	}
	QuantityConfig                 *QuantityConfig
	MeasGapConfig                  *MeasGapConfig
	MeasGapSharingConfig           *MeasGapSharingConfig
	InterFrequencyConfig_NoGap_r16 *int64
	EffectiveMeasWindowConfig_r18  *struct {
		Choice  int
		Release *struct{}
		Setup   *MeasWindowConfig_r18
	}
	Cssf_Config_r19 *struct{ Cssf_MeasMO_List_r19 *CSSF_MeasMO_List_r19 }
	Fbs_Config_r19  *struct {
		Choice  int
		Release *struct{}
		Setup   *FBS_Config_r19
	}
	SpeedStatePars_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *SpeedStatePars_r19
	}
}

func (ie *MeasConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.InterFrequencyConfig_NoGap_r16 != nil
	hasExtGroup1 := ie.EffectiveMeasWindowConfig_r18 != nil
	hasExtGroup2 := ie.Cssf_Config_r19 != nil || ie.Fbs_Config_r19 != nil || ie.SpeedStatePars_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasObjectToRemoveList != nil, ie.MeasObjectToAddModList != nil, ie.ReportConfigToRemoveList != nil, ie.ReportConfigToAddModList != nil, ie.MeasIdToRemoveList != nil, ie.MeasIdToAddModList != nil, ie.S_MeasureConfig != nil, ie.QuantityConfig != nil, ie.MeasGapConfig != nil, ie.MeasGapSharingConfig != nil}); err != nil {
		return err
	}

	// 3. measObjectToRemoveList: ref
	{
		if ie.MeasObjectToRemoveList != nil {
			if err := ie.MeasObjectToRemoveList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. measObjectToAddModList: ref
	{
		if ie.MeasObjectToAddModList != nil {
			if err := ie.MeasObjectToAddModList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. reportConfigToRemoveList: ref
	{
		if ie.ReportConfigToRemoveList != nil {
			if err := ie.ReportConfigToRemoveList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. reportConfigToAddModList: ref
	{
		if ie.ReportConfigToAddModList != nil {
			if err := ie.ReportConfigToAddModList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. measIdToRemoveList: ref
	{
		if ie.MeasIdToRemoveList != nil {
			if err := ie.MeasIdToRemoveList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. measIdToAddModList: ref
	{
		if ie.MeasIdToAddModList != nil {
			if err := ie.MeasIdToAddModList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. s-MeasureConfig: choice
	{
		if ie.S_MeasureConfig != nil {
			choiceEnc := e.NewChoiceEncoder(measConfigSMeasureConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.S_MeasureConfig).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.S_MeasureConfig).Choice {
			case MeasConfig_S_MeasureConfig_Ssb_RSRP:
				if err := (*ie.S_MeasureConfig).Ssb_RSRP.Encode(e); err != nil {
					return err
				}
			case MeasConfig_S_MeasureConfig_Csi_RSRP:
				if err := (*ie.S_MeasureConfig).Csi_RSRP.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.S_MeasureConfig).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 10. quantityConfig: ref
	{
		if ie.QuantityConfig != nil {
			if err := ie.QuantityConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. measGapConfig: ref
	{
		if ie.MeasGapConfig != nil {
			if err := ie.MeasGapConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 12. measGapSharingConfig: ref
	{
		if ie.MeasGapSharingConfig != nil {
			if err := ie.MeasGapSharingConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "interFrequencyConfig-NoGap-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.InterFrequencyConfig_NoGap_r16 != nil}); err != nil {
				return err
			}

			if ie.InterFrequencyConfig_NoGap_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.InterFrequencyConfig_NoGap_r16, measConfigExtInterFrequencyConfigNoGapR16Constraints); err != nil {
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
					{Name: "effectiveMeasWindowConfig-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EffectiveMeasWindowConfig_r18 != nil}); err != nil {
				return err
			}

			if ie.EffectiveMeasWindowConfig_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(measConfigExtEffectiveMeasWindowConfigR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.EffectiveMeasWindowConfig_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.EffectiveMeasWindowConfig_r18).Choice {
				case MeasConfig_Ext_EffectiveMeasWindowConfig_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MeasConfig_Ext_EffectiveMeasWindowConfig_r18_Setup:
					if err := (*ie.EffectiveMeasWindowConfig_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "cssf-Config-r19", Optional: true},
					{Name: "fbs-Config-r19", Optional: true},
					{Name: "speedStatePars-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cssf_Config_r19 != nil, ie.Fbs_Config_r19 != nil, ie.SpeedStatePars_r19 != nil}); err != nil {
				return err
			}

			if ie.Cssf_Config_r19 != nil {
				c := ie.Cssf_Config_r19
				measConfigExtCssfConfigR19Seq := ex.NewSequenceEncoder(measConfigExtCssfConfigR19Constraints)
				if err := measConfigExtCssfConfigR19Seq.EncodePreamble([]bool{c.Cssf_MeasMO_List_r19 != nil}); err != nil {
					return err
				}
				if c.Cssf_MeasMO_List_r19 != nil {
					if err := c.Cssf_MeasMO_List_r19.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Fbs_Config_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(measConfigExtFbsConfigR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Fbs_Config_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Fbs_Config_r19).Choice {
				case MeasConfig_Ext_Fbs_Config_r19_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MeasConfig_Ext_Fbs_Config_r19_Setup:
					if err := (*ie.Fbs_Config_r19).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SpeedStatePars_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(measConfigExtSpeedStateParsR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.SpeedStatePars_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.SpeedStatePars_r19).Choice {
				case MeasConfig_Ext_SpeedStatePars_r19_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MeasConfig_Ext_SpeedStatePars_r19_Setup:
					if err := (*ie.SpeedStatePars_r19).Setup.Encode(ex); err != nil {
						return err
					}
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

func (ie *MeasConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. measObjectToRemoveList: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasObjectToRemoveList = new(MeasObjectToRemoveList)
			if err := ie.MeasObjectToRemoveList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. measObjectToAddModList: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MeasObjectToAddModList = new(MeasObjectToAddModList)
			if err := ie.MeasObjectToAddModList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. reportConfigToRemoveList: ref
	{
		if seq.IsComponentPresent(2) {
			ie.ReportConfigToRemoveList = new(ReportConfigToRemoveList)
			if err := ie.ReportConfigToRemoveList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. reportConfigToAddModList: ref
	{
		if seq.IsComponentPresent(3) {
			ie.ReportConfigToAddModList = new(ReportConfigToAddModList)
			if err := ie.ReportConfigToAddModList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. measIdToRemoveList: ref
	{
		if seq.IsComponentPresent(4) {
			ie.MeasIdToRemoveList = new(MeasIdToRemoveList)
			if err := ie.MeasIdToRemoveList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. measIdToAddModList: ref
	{
		if seq.IsComponentPresent(5) {
			ie.MeasIdToAddModList = new(MeasIdToAddModList)
			if err := ie.MeasIdToAddModList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. s-MeasureConfig: choice
	{
		if seq.IsComponentPresent(6) {
			ie.S_MeasureConfig = &struct {
				Choice   int
				Ssb_RSRP *RSRP_Range
				Csi_RSRP *RSRP_Range
			}{}
			choiceDec := d.NewChoiceDecoder(measConfigSMeasureConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.S_MeasureConfig).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case MeasConfig_S_MeasureConfig_Ssb_RSRP:
				(*ie.S_MeasureConfig).Ssb_RSRP = new(RSRP_Range)
				if err := (*ie.S_MeasureConfig).Ssb_RSRP.Decode(d); err != nil {
					return err
				}
			case MeasConfig_S_MeasureConfig_Csi_RSRP:
				(*ie.S_MeasureConfig).Csi_RSRP = new(RSRP_Range)
				if err := (*ie.S_MeasureConfig).Csi_RSRP.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. quantityConfig: ref
	{
		if seq.IsComponentPresent(7) {
			ie.QuantityConfig = new(QuantityConfig)
			if err := ie.QuantityConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. measGapConfig: ref
	{
		if seq.IsComponentPresent(8) {
			ie.MeasGapConfig = new(MeasGapConfig)
			if err := ie.MeasGapConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 12. measGapSharingConfig: ref
	{
		if seq.IsComponentPresent(9) {
			ie.MeasGapSharingConfig = new(MeasGapSharingConfig)
			if err := ie.MeasGapSharingConfig.Decode(d); err != nil {
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
				{Name: "interFrequencyConfig-NoGap-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measConfigExtInterFrequencyConfigNoGapR16Constraints)
			if err != nil {
				return err
			}
			ie.InterFrequencyConfig_NoGap_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "effectiveMeasWindowConfig-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.EffectiveMeasWindowConfig_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MeasWindowConfig_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(measConfigExtEffectiveMeasWindowConfigR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.EffectiveMeasWindowConfig_r18).Choice = int(index)
			switch index {
			case MeasConfig_Ext_EffectiveMeasWindowConfig_r18_Release:
				(*ie.EffectiveMeasWindowConfig_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MeasConfig_Ext_EffectiveMeasWindowConfig_r18_Setup:
				(*ie.EffectiveMeasWindowConfig_r18).Setup = new(MeasWindowConfig_r18)
				if err := (*ie.EffectiveMeasWindowConfig_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "cssf-Config-r19", Optional: true},
				{Name: "fbs-Config-r19", Optional: true},
				{Name: "speedStatePars-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Cssf_Config_r19 = &struct{ Cssf_MeasMO_List_r19 *CSSF_MeasMO_List_r19 }{}
			c := ie.Cssf_Config_r19
			measConfigExtCssfConfigR19Seq := dx.NewSequenceDecoder(measConfigExtCssfConfigR19Constraints)
			if err := measConfigExtCssfConfigR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if measConfigExtCssfConfigR19Seq.IsComponentPresent(0) {
				c.Cssf_MeasMO_List_r19 = new(CSSF_MeasMO_List_r19)
				if err := (*c.Cssf_MeasMO_List_r19).Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Fbs_Config_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *FBS_Config_r19
			}{}
			choiceDec := dx.NewChoiceDecoder(measConfigExtFbsConfigR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Fbs_Config_r19).Choice = int(index)
			switch index {
			case MeasConfig_Ext_Fbs_Config_r19_Release:
				(*ie.Fbs_Config_r19).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MeasConfig_Ext_Fbs_Config_r19_Setup:
				(*ie.Fbs_Config_r19).Setup = new(FBS_Config_r19)
				if err := (*ie.Fbs_Config_r19).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SpeedStatePars_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SpeedStatePars_r19
			}{}
			choiceDec := dx.NewChoiceDecoder(measConfigExtSpeedStateParsR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SpeedStatePars_r19).Choice = int(index)
			switch index {
			case MeasConfig_Ext_SpeedStatePars_r19_Release:
				(*ie.SpeedStatePars_r19).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MeasConfig_Ext_SpeedStatePars_r19_Setup:
				(*ie.SpeedStatePars_r19).Setup = new(SpeedStatePars_r19)
				if err := (*ie.SpeedStatePars_r19).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
