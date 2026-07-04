// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-PosRRC-InactiveValidityAreaConfig-r18 (line 1480).

var sRSPosRRCInactiveValidityAreaConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-PosConfigValidityArea-r18"},
		{Name: "srs-PosConfigNUL-r18", Optional: true},
		{Name: "srs-PosConfigSUL-r18", Optional: true},
		{Name: "bwp-NUL-r18", Optional: true},
		{Name: "bwp-SUL-r18", Optional: true},
		{Name: "areaValidityTA-Config-r18", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var sRSPosRRCInactiveValidityAreaConfigR18SrsPosConfigValidityAreaR18Constraints = per.SizeRange(1, common.MaxNrOfCellsInVA_r18)

var sRSPosRRCInactiveValidityAreaConfigR18ExtSrsPosConfigValidityAreaExtV1830Constraints = per.SizeRange(1, common.MaxNrOfCellsInVA_Ext_r18)

var sRSPosRRCInactiveValidityAreaConfigR18ExtSrsPosRRCInactiveAggBWAdditionalCarriersPerVAR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SRS_PosRRC_InactiveValidityAreaConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18_Release = 0
	SRS_PosRRC_InactiveValidityAreaConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18_Setup   = 1
)

var sRSPosRRCInactiveValidityAreaConfigR18ExtSrsPosRRCInactiveAggBWConfigListPerVAR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SRS_PosRRC_InactiveValidityAreaConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18_Release = 0
	SRS_PosRRC_InactiveValidityAreaConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18_Setup   = 1
)

type SRS_PosRRC_InactiveValidityAreaConfig_r18 struct {
	Srs_PosConfigValidityArea_r18                        []CellIdentity
	Srs_PosConfigNUL_r18                                 *SRS_PosConfig_r17
	Srs_PosConfigSUL_r18                                 *SRS_PosConfig_r17
	Bwp_NUL_r18                                          *BWP
	Bwp_SUL_r18                                          *BWP
	AreaValidityTA_Config_r18                            *AreaValidityTA_Config_r18
	Srs_PosConfigValidityAreaExt_v1830                   []CellIdentity
	Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_PosRRC_InactiveAggBW_AdditionalCarriers_r18
	}
	Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_PosRRC_InactiveAggBW_ConfigList_r18
	}
}

func (ie *SRS_PosRRC_InactiveValidityAreaConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSPosRRCInactiveValidityAreaConfigR18Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Srs_PosConfigValidityAreaExt_v1830 != nil
	hasExtGroup1 := ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18 != nil || ie.Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_PosConfigNUL_r18 != nil, ie.Srs_PosConfigSUL_r18 != nil, ie.Bwp_NUL_r18 != nil, ie.Bwp_SUL_r18 != nil, ie.AreaValidityTA_Config_r18 != nil}); err != nil {
		return err
	}

	// 3. srs-PosConfigValidityArea-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sRSPosRRCInactiveValidityAreaConfigR18SrsPosConfigValidityAreaR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Srs_PosConfigValidityArea_r18))); err != nil {
			return err
		}
		for i := range ie.Srs_PosConfigValidityArea_r18 {
			if err := ie.Srs_PosConfigValidityArea_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. srs-PosConfigNUL-r18: ref
	{
		if ie.Srs_PosConfigNUL_r18 != nil {
			if err := ie.Srs_PosConfigNUL_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. srs-PosConfigSUL-r18: ref
	{
		if ie.Srs_PosConfigSUL_r18 != nil {
			if err := ie.Srs_PosConfigSUL_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. bwp-NUL-r18: ref
	{
		if ie.Bwp_NUL_r18 != nil {
			if err := ie.Bwp_NUL_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. bwp-SUL-r18: ref
	{
		if ie.Bwp_SUL_r18 != nil {
			if err := ie.Bwp_SUL_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. areaValidityTA-Config-r18: ref
	{
		if ie.AreaValidityTA_Config_r18 != nil {
			if err := ie.AreaValidityTA_Config_r18.Encode(e); err != nil {
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
					{Name: "srs-PosConfigValidityAreaExt-v1830", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srs_PosConfigValidityAreaExt_v1830 != nil}); err != nil {
				return err
			}

			if ie.Srs_PosConfigValidityAreaExt_v1830 != nil {
				seqOf := ex.NewSequenceOfEncoder(sRSPosRRCInactiveValidityAreaConfigR18ExtSrsPosConfigValidityAreaExtV1830Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Srs_PosConfigValidityAreaExt_v1830))); err != nil {
					return err
				}
				for i := range ie.Srs_PosConfigValidityAreaExt_v1830 {
					if err := ie.Srs_PosConfigValidityAreaExt_v1830[i].Encode(ex); err != nil {
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
					{Name: "srs-PosRRC-InactiveAggBW-AdditionalCarriersPerVA-r18", Optional: true},
					{Name: "srs-PosRRC-InactiveAggBW-ConfigListPerVA-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18 != nil, ie.Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18 != nil}); err != nil {
				return err
			}

			if ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(sRSPosRRCInactiveValidityAreaConfigR18ExtSrsPosRRCInactiveAggBWAdditionalCarriersPerVAR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18).Choice {
				case SRS_PosRRC_InactiveValidityAreaConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SRS_PosRRC_InactiveValidityAreaConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18_Setup:
					if err := (*ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(sRSPosRRCInactiveValidityAreaConfigR18ExtSrsPosRRCInactiveAggBWConfigListPerVAR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18).Choice {
				case SRS_PosRRC_InactiveValidityAreaConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SRS_PosRRC_InactiveValidityAreaConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18_Setup:
					if err := (*ie.Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18).Setup.Encode(ex); err != nil {
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

func (ie *SRS_PosRRC_InactiveValidityAreaConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSPosRRCInactiveValidityAreaConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. srs-PosConfigValidityArea-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sRSPosRRCInactiveValidityAreaConfigR18SrsPosConfigValidityAreaR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Srs_PosConfigValidityArea_r18 = make([]CellIdentity, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Srs_PosConfigValidityArea_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. srs-PosConfigNUL-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Srs_PosConfigNUL_r18 = new(SRS_PosConfig_r17)
			if err := ie.Srs_PosConfigNUL_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. srs-PosConfigSUL-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Srs_PosConfigSUL_r18 = new(SRS_PosConfig_r17)
			if err := ie.Srs_PosConfigSUL_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. bwp-NUL-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Bwp_NUL_r18 = new(BWP)
			if err := ie.Bwp_NUL_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. bwp-SUL-r18: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Bwp_SUL_r18 = new(BWP)
			if err := ie.Bwp_SUL_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. areaValidityTA-Config-r18: ref
	{
		if seq.IsComponentPresent(5) {
			ie.AreaValidityTA_Config_r18 = new(AreaValidityTA_Config_r18)
			if err := ie.AreaValidityTA_Config_r18.Decode(d); err != nil {
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
				{Name: "srs-PosConfigValidityAreaExt-v1830", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(sRSPosRRCInactiveValidityAreaConfigR18ExtSrsPosConfigValidityAreaExtV1830Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_PosConfigValidityAreaExt_v1830 = make([]CellIdentity, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_PosConfigValidityAreaExt_v1830[i].Decode(dx); err != nil {
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
				{Name: "srs-PosRRC-InactiveAggBW-AdditionalCarriersPerVA-r18", Optional: true},
				{Name: "srs-PosRRC-InactiveAggBW-ConfigListPerVA-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_PosRRC_InactiveAggBW_AdditionalCarriers_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(sRSPosRRCInactiveValidityAreaConfigR18ExtSrsPosRRCInactiveAggBWAdditionalCarriersPerVAR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18).Choice = int(index)
			switch index {
			case SRS_PosRRC_InactiveValidityAreaConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18_Release:
				(*ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SRS_PosRRC_InactiveValidityAreaConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18_Setup:
				(*ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18).Setup = new(SRS_PosRRC_InactiveAggBW_AdditionalCarriers_r18)
				if err := (*ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriersPerVA_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_PosRRC_InactiveAggBW_ConfigList_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(sRSPosRRCInactiveValidityAreaConfigR18ExtSrsPosRRCInactiveAggBWConfigListPerVAR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18).Choice = int(index)
			switch index {
			case SRS_PosRRC_InactiveValidityAreaConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18_Release:
				(*ie.Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SRS_PosRRC_InactiveValidityAreaConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18_Setup:
				(*ie.Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18).Setup = new(SRS_PosRRC_InactiveAggBW_ConfigList_r18)
				if err := (*ie.Srs_PosRRC_InactiveAggBW_ConfigListPerVA_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
