// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-PosRRC-InactiveEnhancedConfig-r18 (line 1454).

var sRSPosRRCInactiveEnhancedConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-PosRRC-InactiveAggBW-ConfigList-r18", Optional: true},
		{Name: "srs-PosTx-Hopping-r18", Optional: true},
		{Name: "srs-PosRRC-InactiveValidityAreaPreConfigList-r18", Optional: true},
		{Name: "srs-PosRRC-InactiveValidityAreaNonPreConfig-r18", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sRS_PosRRC_InactiveEnhancedConfig_r18SrsPosRRCInactiveAggBWConfigListR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveAggBW_ConfigList_r18_Release = 0
	SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveAggBW_ConfigList_r18_Setup   = 1
)

var sRS_PosRRC_InactiveEnhancedConfig_r18SrsPosTxHoppingR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosTx_Hopping_r18_Release = 0
	SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosTx_Hopping_r18_Setup   = 1
)

var sRS_PosRRC_InactiveEnhancedConfig_r18SrsPosRRCInactiveValidityAreaPreConfigListR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveValidityAreaPreConfigList_r18_Release = 0
	SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveValidityAreaPreConfigList_r18_Setup   = 1
)

var sRS_PosRRC_InactiveEnhancedConfig_r18SrsPosRRCInactiveValidityAreaNonPreConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18_Release = 0
	SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18_Setup   = 1
)

var sRSPosRRCInactiveEnhancedConfigR18ExtSrsPosRRCInactiveAggBWAdditionalCarriersR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SRS_PosRRC_InactiveEnhancedConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18_Release = 0
	SRS_PosRRC_InactiveEnhancedConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18_Setup   = 1
)

type SRS_PosRRC_InactiveEnhancedConfig_r18 struct {
	Srs_PosRRC_InactiveAggBW_ConfigList_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_PosRRC_InactiveAggBW_ConfigList_r18
	}
	Srs_PosTx_Hopping_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_PosTx_Hopping_r18
	}
	Srs_PosRRC_InactiveValidityAreaPreConfigList_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_PosRRC_InactiveValidityAreaPreConfigList_r18
	}
	Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_PosRRC_InactiveValidityAreaConfig_r18
	}
	Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_PosRRC_InactiveAggBW_AdditionalCarriers_r18
	}
}

func (ie *SRS_PosRRC_InactiveEnhancedConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSPosRRCInactiveEnhancedConfigR18Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_PosRRC_InactiveAggBW_ConfigList_r18 != nil, ie.Srs_PosTx_Hopping_r18 != nil, ie.Srs_PosRRC_InactiveValidityAreaPreConfigList_r18 != nil, ie.Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18 != nil}); err != nil {
		return err
	}

	// 3. srs-PosRRC-InactiveAggBW-ConfigList-r18: choice
	{
		if ie.Srs_PosRRC_InactiveAggBW_ConfigList_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(sRS_PosRRC_InactiveEnhancedConfig_r18SrsPosRRCInactiveAggBWConfigListR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Srs_PosRRC_InactiveAggBW_ConfigList_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Srs_PosRRC_InactiveAggBW_ConfigList_r18).Choice {
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveAggBW_ConfigList_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveAggBW_ConfigList_r18_Setup:
				if err := (*ie.Srs_PosRRC_InactiveAggBW_ConfigList_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Srs_PosRRC_InactiveAggBW_ConfigList_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. srs-PosTx-Hopping-r18: choice
	{
		if ie.Srs_PosTx_Hopping_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(sRS_PosRRC_InactiveEnhancedConfig_r18SrsPosTxHoppingR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Srs_PosTx_Hopping_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Srs_PosTx_Hopping_r18).Choice {
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosTx_Hopping_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosTx_Hopping_r18_Setup:
				if err := (*ie.Srs_PosTx_Hopping_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Srs_PosTx_Hopping_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. srs-PosRRC-InactiveValidityAreaPreConfigList-r18: choice
	{
		if ie.Srs_PosRRC_InactiveValidityAreaPreConfigList_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(sRS_PosRRC_InactiveEnhancedConfig_r18SrsPosRRCInactiveValidityAreaPreConfigListR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Srs_PosRRC_InactiveValidityAreaPreConfigList_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Srs_PosRRC_InactiveValidityAreaPreConfigList_r18).Choice {
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveValidityAreaPreConfigList_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveValidityAreaPreConfigList_r18_Setup:
				if err := (*ie.Srs_PosRRC_InactiveValidityAreaPreConfigList_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Srs_PosRRC_InactiveValidityAreaPreConfigList_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. srs-PosRRC-InactiveValidityAreaNonPreConfig-r18: choice
	{
		if ie.Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(sRS_PosRRC_InactiveEnhancedConfig_r18SrsPosRRCInactiveValidityAreaNonPreConfigR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18).Choice {
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18_Setup:
				if err := (*ie.Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18).Choice), Constraint: "index out of range"}
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
					{Name: "srs-PosRRC-InactiveAggBW-AdditionalCarriers-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18 != nil}); err != nil {
				return err
			}

			if ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(sRSPosRRCInactiveEnhancedConfigR18ExtSrsPosRRCInactiveAggBWAdditionalCarriersR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18).Choice {
				case SRS_PosRRC_InactiveEnhancedConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SRS_PosRRC_InactiveEnhancedConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18_Setup:
					if err := (*ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18).Setup.Encode(ex); err != nil {
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

func (ie *SRS_PosRRC_InactiveEnhancedConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSPosRRCInactiveEnhancedConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. srs-PosRRC-InactiveAggBW-ConfigList-r18: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Srs_PosRRC_InactiveAggBW_ConfigList_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_PosRRC_InactiveAggBW_ConfigList_r18
			}{}
			choiceDec := d.NewChoiceDecoder(sRS_PosRRC_InactiveEnhancedConfig_r18SrsPosRRCInactiveAggBWConfigListR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_PosRRC_InactiveAggBW_ConfigList_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveAggBW_ConfigList_r18_Release:
				(*ie.Srs_PosRRC_InactiveAggBW_ConfigList_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveAggBW_ConfigList_r18_Setup:
				(*ie.Srs_PosRRC_InactiveAggBW_ConfigList_r18).Setup = new(SRS_PosRRC_InactiveAggBW_ConfigList_r18)
				if err := (*ie.Srs_PosRRC_InactiveAggBW_ConfigList_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. srs-PosTx-Hopping-r18: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Srs_PosTx_Hopping_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_PosTx_Hopping_r18
			}{}
			choiceDec := d.NewChoiceDecoder(sRS_PosRRC_InactiveEnhancedConfig_r18SrsPosTxHoppingR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_PosTx_Hopping_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosTx_Hopping_r18_Release:
				(*ie.Srs_PosTx_Hopping_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosTx_Hopping_r18_Setup:
				(*ie.Srs_PosTx_Hopping_r18).Setup = new(SRS_PosTx_Hopping_r18)
				if err := (*ie.Srs_PosTx_Hopping_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. srs-PosRRC-InactiveValidityAreaPreConfigList-r18: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Srs_PosRRC_InactiveValidityAreaPreConfigList_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_PosRRC_InactiveValidityAreaPreConfigList_r18
			}{}
			choiceDec := d.NewChoiceDecoder(sRS_PosRRC_InactiveEnhancedConfig_r18SrsPosRRCInactiveValidityAreaPreConfigListR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_PosRRC_InactiveValidityAreaPreConfigList_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveValidityAreaPreConfigList_r18_Release:
				(*ie.Srs_PosRRC_InactiveValidityAreaPreConfigList_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveValidityAreaPreConfigList_r18_Setup:
				(*ie.Srs_PosRRC_InactiveValidityAreaPreConfigList_r18).Setup = new(SRS_PosRRC_InactiveValidityAreaPreConfigList_r18)
				if err := (*ie.Srs_PosRRC_InactiveValidityAreaPreConfigList_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. srs-PosRRC-InactiveValidityAreaNonPreConfig-r18: choice
	{
		if seq.IsComponentPresent(3) {
			ie.Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_PosRRC_InactiveValidityAreaConfig_r18
			}{}
			choiceDec := d.NewChoiceDecoder(sRS_PosRRC_InactiveEnhancedConfig_r18SrsPosRRCInactiveValidityAreaNonPreConfigR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18_Release:
				(*ie.Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18_Setup:
				(*ie.Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18).Setup = new(SRS_PosRRC_InactiveValidityAreaConfig_r18)
				if err := (*ie.Srs_PosRRC_InactiveValidityAreaNonPreConfig_r18).Setup.Decode(d); err != nil {
					return err
				}
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
				{Name: "srs-PosRRC-InactiveAggBW-AdditionalCarriers-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_PosRRC_InactiveAggBW_AdditionalCarriers_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(sRSPosRRCInactiveEnhancedConfigR18ExtSrsPosRRCInactiveAggBWAdditionalCarriersR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18).Choice = int(index)
			switch index {
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18_Release:
				(*ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SRS_PosRRC_InactiveEnhancedConfig_r18_Ext_Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18_Setup:
				(*ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18).Setup = new(SRS_PosRRC_InactiveAggBW_AdditionalCarriers_r18)
				if err := (*ie.Srs_PosRRC_InactiveAggBW_AdditionalCarriers_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
