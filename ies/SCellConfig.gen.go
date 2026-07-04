// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SCellConfig (line 5730).

var sCellConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sCellIndex"},
		{Name: "sCellConfigCommon", Optional: true},
		{Name: "sCellConfigDedicated", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
	},
}

const (
	SCellConfig_Ext_SCellState_r16_Activated = 0
)

var sCellConfigExtSCellStateR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SCellConfig_Ext_SCellState_r16_Activated},
}

const (
	SCellConfig_Ext_SecondaryDRX_GroupConfig_r16_True = 0
)

var sCellConfigExtSecondaryDRXGroupConfigR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SCellConfig_Ext_SecondaryDRX_GroupConfig_r16_True},
}

var sCellConfigPreConfGapStatusR17Constraints = per.FixedSize(common.MaxNrofGapId_r17)

var sCellConfigExtSCellSIB20R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SCellConfig_Ext_SCellSIB20_r17_Release = 0
	SCellConfig_Ext_SCellSIB20_r17_Setup   = 1
)

var sCellConfigExtPlmnIdentityInfoListR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SCellConfig_Ext_Plmn_IdentityInfoList_r17_Release = 0
	SCellConfig_Ext_Plmn_IdentityInfoList_r17_Setup   = 1
)

var sCellConfigExtNpnIdentityInfoListR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SCellConfig_Ext_Npn_IdentityInfoList_r17_Release = 0
	SCellConfig_Ext_Npn_IdentityInfoList_r17_Setup   = 1
)

var sCellConfigExtOdSsbR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SCellConfig_Ext_Od_Ssb_r19_Release = 0
	SCellConfig_Ext_Od_Ssb_r19_Setup   = 1
)

var sCellConfigExtAdaptSSBConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SCellConfig_Ext_Adapt_SSB_Config_r19_Release = 0
	SCellConfig_Ext_Adapt_SSB_Config_r19_Setup   = 1
)

type SCellConfig struct {
	SCellIndex                       SCellIndex
	SCellConfigCommon                *ServingCellConfigCommon
	SCellConfigDedicated             *ServingCellConfig
	Smtc                             *SSB_MTC
	SCellState_r16                   *int64
	SecondaryDRX_GroupConfig_r16     *int64
	PreConfGapStatus_r17             *per.BitString
	GoodServingCellEvaluationBFD_r17 *GoodServingCellEvaluation_r17
	SCellSIB20_r17                   *struct {
		Choice  int
		Release *struct{}
		Setup   *SCellSIB20_r17
	}
	Plmn_IdentityInfoList_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *PLMN_IdentityInfoList
	}
	Npn_IdentityInfoList_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *NPN_IdentityInfoList_r16
	}
	Od_Ssb_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *OD_SSB_r19
	}
	Adapt_SSB_Config_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *Adapt_SSB_Config_r19
	}
}

func (ie *SCellConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sCellConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Smtc != nil
	hasExtGroup1 := ie.SCellState_r16 != nil || ie.SecondaryDRX_GroupConfig_r16 != nil
	hasExtGroup2 := ie.PreConfGapStatus_r17 != nil || ie.GoodServingCellEvaluationBFD_r17 != nil || ie.SCellSIB20_r17 != nil
	hasExtGroup3 := ie.Plmn_IdentityInfoList_r17 != nil || ie.Npn_IdentityInfoList_r17 != nil
	hasExtGroup4 := ie.Od_Ssb_r19 != nil || ie.Adapt_SSB_Config_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SCellConfigCommon != nil, ie.SCellConfigDedicated != nil}); err != nil {
		return err
	}

	// 3. sCellIndex: ref
	{
		if err := ie.SCellIndex.Encode(e); err != nil {
			return err
		}
	}

	// 4. sCellConfigCommon: ref
	{
		if ie.SCellConfigCommon != nil {
			if err := ie.SCellConfigCommon.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sCellConfigDedicated: ref
	{
		if ie.SCellConfigDedicated != nil {
			if err := ie.SCellConfigDedicated.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "smtc", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Smtc != nil}); err != nil {
				return err
			}

			if ie.Smtc != nil {
				if err := ie.Smtc.Encode(ex); err != nil {
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
					{Name: "sCellState-r16", Optional: true},
					{Name: "secondaryDRX-GroupConfig-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SCellState_r16 != nil, ie.SecondaryDRX_GroupConfig_r16 != nil}); err != nil {
				return err
			}

			if ie.SCellState_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SCellState_r16, sCellConfigExtSCellStateR16Constraints); err != nil {
					return err
				}
			}

			if ie.SecondaryDRX_GroupConfig_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SecondaryDRX_GroupConfig_r16, sCellConfigExtSecondaryDRXGroupConfigR16Constraints); err != nil {
					return err
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
					{Name: "preConfGapStatus-r17", Optional: true},
					{Name: "goodServingCellEvaluationBFD-r17", Optional: true},
					{Name: "sCellSIB20-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PreConfGapStatus_r17 != nil, ie.GoodServingCellEvaluationBFD_r17 != nil, ie.SCellSIB20_r17 != nil}); err != nil {
				return err
			}

			if ie.PreConfGapStatus_r17 != nil {
				if err := ex.EncodeBitString(*ie.PreConfGapStatus_r17, sCellConfigPreConfGapStatusR17Constraints); err != nil {
					return err
				}
			}

			if ie.GoodServingCellEvaluationBFD_r17 != nil {
				if err := ie.GoodServingCellEvaluationBFD_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SCellSIB20_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(sCellConfigExtSCellSIB20R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.SCellSIB20_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.SCellSIB20_r17).Choice {
				case SCellConfig_Ext_SCellSIB20_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SCellConfig_Ext_SCellSIB20_r17_Setup:
					if err := (*ie.SCellSIB20_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "plmn-IdentityInfoList-r17", Optional: true},
					{Name: "npn-IdentityInfoList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Plmn_IdentityInfoList_r17 != nil, ie.Npn_IdentityInfoList_r17 != nil}); err != nil {
				return err
			}

			if ie.Plmn_IdentityInfoList_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(sCellConfigExtPlmnIdentityInfoListR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Plmn_IdentityInfoList_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Plmn_IdentityInfoList_r17).Choice {
				case SCellConfig_Ext_Plmn_IdentityInfoList_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SCellConfig_Ext_Plmn_IdentityInfoList_r17_Setup:
					if err := (*ie.Plmn_IdentityInfoList_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Npn_IdentityInfoList_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(sCellConfigExtNpnIdentityInfoListR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Npn_IdentityInfoList_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Npn_IdentityInfoList_r17).Choice {
				case SCellConfig_Ext_Npn_IdentityInfoList_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SCellConfig_Ext_Npn_IdentityInfoList_r17_Setup:
					if err := (*ie.Npn_IdentityInfoList_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 4:
		if hasExtGroup4 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "od-ssb-r19", Optional: true},
					{Name: "adapt-SSB-Config-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Od_Ssb_r19 != nil, ie.Adapt_SSB_Config_r19 != nil}); err != nil {
				return err
			}

			if ie.Od_Ssb_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(sCellConfigExtOdSsbR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Od_Ssb_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Od_Ssb_r19).Choice {
				case SCellConfig_Ext_Od_Ssb_r19_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SCellConfig_Ext_Od_Ssb_r19_Setup:
					if err := (*ie.Od_Ssb_r19).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Adapt_SSB_Config_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(sCellConfigExtAdaptSSBConfigR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Adapt_SSB_Config_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Adapt_SSB_Config_r19).Choice {
				case SCellConfig_Ext_Adapt_SSB_Config_r19_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SCellConfig_Ext_Adapt_SSB_Config_r19_Setup:
					if err := (*ie.Adapt_SSB_Config_r19).Setup.Encode(ex); err != nil {
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

func (ie *SCellConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sCellConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sCellIndex: ref
	{
		if err := ie.SCellIndex.Decode(d); err != nil {
			return err
		}
	}

	// 4. sCellConfigCommon: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SCellConfigCommon = new(ServingCellConfigCommon)
			if err := ie.SCellConfigCommon.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sCellConfigDedicated: ref
	{
		if seq.IsComponentPresent(2) {
			ie.SCellConfigDedicated = new(ServingCellConfig)
			if err := ie.SCellConfigDedicated.Decode(d); err != nil {
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
				{Name: "smtc", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Smtc = new(SSB_MTC)
			if err := ie.Smtc.Decode(dx); err != nil {
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
				{Name: "sCellState-r16", Optional: true},
				{Name: "secondaryDRX-GroupConfig-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sCellConfigExtSCellStateR16Constraints)
			if err != nil {
				return err
			}
			ie.SCellState_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(sCellConfigExtSecondaryDRXGroupConfigR16Constraints)
			if err != nil {
				return err
			}
			ie.SecondaryDRX_GroupConfig_r16 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "preConfGapStatus-r17", Optional: true},
				{Name: "goodServingCellEvaluationBFD-r17", Optional: true},
				{Name: "sCellSIB20-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeBitString(sCellConfigPreConfGapStatusR17Constraints)
			if err != nil {
				return err
			}
			ie.PreConfGapStatus_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.GoodServingCellEvaluationBFD_r17 = new(GoodServingCellEvaluation_r17)
			if err := ie.GoodServingCellEvaluationBFD_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SCellSIB20_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SCellSIB20_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(sCellConfigExtSCellSIB20R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SCellSIB20_r17).Choice = int(index)
			switch index {
			case SCellConfig_Ext_SCellSIB20_r17_Release:
				(*ie.SCellSIB20_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SCellConfig_Ext_SCellSIB20_r17_Setup:
				(*ie.SCellSIB20_r17).Setup = new(SCellSIB20_r17)
				if err := (*ie.SCellSIB20_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "plmn-IdentityInfoList-r17", Optional: true},
				{Name: "npn-IdentityInfoList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Plmn_IdentityInfoList_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PLMN_IdentityInfoList
			}{}
			choiceDec := dx.NewChoiceDecoder(sCellConfigExtPlmnIdentityInfoListR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Plmn_IdentityInfoList_r17).Choice = int(index)
			switch index {
			case SCellConfig_Ext_Plmn_IdentityInfoList_r17_Release:
				(*ie.Plmn_IdentityInfoList_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SCellConfig_Ext_Plmn_IdentityInfoList_r17_Setup:
				(*ie.Plmn_IdentityInfoList_r17).Setup = new(PLMN_IdentityInfoList)
				if err := (*ie.Plmn_IdentityInfoList_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Npn_IdentityInfoList_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *NPN_IdentityInfoList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(sCellConfigExtNpnIdentityInfoListR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Npn_IdentityInfoList_r17).Choice = int(index)
			switch index {
			case SCellConfig_Ext_Npn_IdentityInfoList_r17_Release:
				(*ie.Npn_IdentityInfoList_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SCellConfig_Ext_Npn_IdentityInfoList_r17_Setup:
				(*ie.Npn_IdentityInfoList_r17).Setup = new(NPN_IdentityInfoList_r16)
				if err := (*ie.Npn_IdentityInfoList_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "od-ssb-r19", Optional: true},
				{Name: "adapt-SSB-Config-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Od_Ssb_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *OD_SSB_r19
			}{}
			choiceDec := dx.NewChoiceDecoder(sCellConfigExtOdSsbR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Od_Ssb_r19).Choice = int(index)
			switch index {
			case SCellConfig_Ext_Od_Ssb_r19_Release:
				(*ie.Od_Ssb_r19).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SCellConfig_Ext_Od_Ssb_r19_Setup:
				(*ie.Od_Ssb_r19).Setup = new(OD_SSB_r19)
				if err := (*ie.Od_Ssb_r19).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Adapt_SSB_Config_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *Adapt_SSB_Config_r19
			}{}
			choiceDec := dx.NewChoiceDecoder(sCellConfigExtAdaptSSBConfigR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Adapt_SSB_Config_r19).Choice = int(index)
			switch index {
			case SCellConfig_Ext_Adapt_SSB_Config_r19_Release:
				(*ie.Adapt_SSB_Config_r19).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SCellConfig_Ext_Adapt_SSB_Config_r19_Setup:
				(*ie.Adapt_SSB_Config_r19).Setup = new(Adapt_SSB_Config_r19)
				if err := (*ie.Adapt_SSB_Config_r19).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
