// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUSCH-ServingCellConfig (line 12664).

var pUSCHServingCellConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "codeBlockGroupTransmission", Optional: true},
		{Name: "rateMatching", Optional: true},
		{Name: "xOverhead", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var pUSCH_ServingCellConfigCodeBlockGroupTransmissionConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_ServingCellConfig_CodeBlockGroupTransmission_Release = 0
	PUSCH_ServingCellConfig_CodeBlockGroupTransmission_Setup   = 1
)

const (
	PUSCH_ServingCellConfig_RateMatching_LimitedBufferRM = 0
)

var pUSCHServingCellConfigRateMatchingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_ServingCellConfig_RateMatching_LimitedBufferRM},
}

const (
	PUSCH_ServingCellConfig_XOverhead_Xoh6  = 0
	PUSCH_ServingCellConfig_XOverhead_Xoh12 = 1
	PUSCH_ServingCellConfig_XOverhead_Xoh18 = 2
)

var pUSCHServingCellConfigXOverheadConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_ServingCellConfig_XOverhead_Xoh6, PUSCH_ServingCellConfig_XOverhead_Xoh12, PUSCH_ServingCellConfig_XOverhead_Xoh18},
}

var pUSCHServingCellConfigMaxMIMOLayersConstraints = per.Constrained(1, 4)

var pUSCHServingCellConfigExtMaxMIMOLayersDCI02R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_ServingCellConfig_Ext_MaxMIMO_LayersDCI_0_2_r16_Release = 0
	PUSCH_ServingCellConfig_Ext_MaxMIMO_LayersDCI_0_2_r16_Setup   = 1
)

const (
	PUSCH_ServingCellConfig_Ext_NrofHARQ_ProcessesForPUSCH_r17_N32 = 0
)

var pUSCHServingCellConfigExtNrofHARQProcessesForPUSCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_ServingCellConfig_Ext_NrofHARQ_ProcessesForPUSCH_r17_N32},
}

var pUSCHServingCellConfigExtUplinkHARQModeR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_ServingCellConfig_Ext_UplinkHARQ_Mode_r17_Release = 0
	PUSCH_ServingCellConfig_Ext_UplinkHARQ_Mode_r17_Setup   = 1
)

var pUSCHServingCellConfigMaxMIMOLayersV1810Constraints = per.Constrained(5, 8)

var pUSCHServingCellConfigMaxMIMOLayersforSDMR18Constraints = per.Constrained(1, 2)

var pUSCHServingCellConfigMaxMIMOLayersforSDMDCI02R18Constraints = per.Constrained(1, 2)

var pUSCHServingCellConfigMaxMIMOLayersforSFNR18Constraints = per.Constrained(1, 2)

var pUSCHServingCellConfigMaxMIMOLayersforSFNDCI02R18Constraints = per.Constrained(1, 2)

type PUSCH_ServingCellConfig struct {
	CodeBlockGroupTransmission *struct {
		Choice  int
		Release *struct{}
		Setup   *PUSCH_CodeBlockGroupTransmission
	}
	RateMatching              *int64
	XOverhead                 *int64
	MaxMIMO_Layers            *int64
	ProcessingType2Enabled    *bool
	MaxMIMO_LayersDCI_0_2_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *MaxMIMO_LayersDCI_0_2_r16
	}
	NrofHARQ_ProcessesForPUSCH_r17 *int64
	UplinkHARQ_Mode_r17            *struct {
		Choice  int
		Release *struct{}
		Setup   *UplinkHARQ_Mode_r17
	}
	MaxMIMO_Layers_v1810             *int64
	MaxMIMO_LayersforSDM_r18         *int64
	MaxMIMO_LayersforSDM_DCI_0_2_r18 *int64
	MaxMIMO_LayersforSFN_r18         *int64
	MaxMIMO_LayersforSFN_DCI_0_2_r18 *int64
}

func (ie *PUSCH_ServingCellConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHServingCellConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MaxMIMO_Layers != nil || ie.ProcessingType2Enabled != nil
	hasExtGroup1 := ie.MaxMIMO_LayersDCI_0_2_r16 != nil
	hasExtGroup2 := ie.NrofHARQ_ProcessesForPUSCH_r17 != nil || ie.UplinkHARQ_Mode_r17 != nil
	hasExtGroup3 := ie.MaxMIMO_Layers_v1810 != nil || ie.MaxMIMO_LayersforSDM_r18 != nil || ie.MaxMIMO_LayersforSDM_DCI_0_2_r18 != nil || ie.MaxMIMO_LayersforSFN_r18 != nil || ie.MaxMIMO_LayersforSFN_DCI_0_2_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CodeBlockGroupTransmission != nil, ie.RateMatching != nil, ie.XOverhead != nil}); err != nil {
		return err
	}

	// 3. codeBlockGroupTransmission: choice
	{
		if ie.CodeBlockGroupTransmission != nil {
			choiceEnc := e.NewChoiceEncoder(pUSCH_ServingCellConfigCodeBlockGroupTransmissionConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.CodeBlockGroupTransmission).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.CodeBlockGroupTransmission).Choice {
			case PUSCH_ServingCellConfig_CodeBlockGroupTransmission_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PUSCH_ServingCellConfig_CodeBlockGroupTransmission_Setup:
				if err := (*ie.CodeBlockGroupTransmission).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.CodeBlockGroupTransmission).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. rateMatching: enumerated
	{
		if ie.RateMatching != nil {
			if err := e.EncodeEnumerated(*ie.RateMatching, pUSCHServingCellConfigRateMatchingConstraints); err != nil {
				return err
			}
		}
	}

	// 5. xOverhead: enumerated
	{
		if ie.XOverhead != nil {
			if err := e.EncodeEnumerated(*ie.XOverhead, pUSCHServingCellConfigXOverheadConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "maxMIMO-Layers", Optional: true},
					{Name: "processingType2Enabled", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxMIMO_Layers != nil, ie.ProcessingType2Enabled != nil}); err != nil {
				return err
			}

			if ie.MaxMIMO_Layers != nil {
				if err := ex.EncodeInteger(*ie.MaxMIMO_Layers, pUSCHServingCellConfigMaxMIMOLayersConstraints); err != nil {
					return err
				}
			}

			if ie.ProcessingType2Enabled != nil {
				if err := ex.EncodeBoolean(*ie.ProcessingType2Enabled); err != nil {
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
					{Name: "maxMIMO-LayersDCI-0-2-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxMIMO_LayersDCI_0_2_r16 != nil}); err != nil {
				return err
			}

			if ie.MaxMIMO_LayersDCI_0_2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHServingCellConfigExtMaxMIMOLayersDCI02R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.MaxMIMO_LayersDCI_0_2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.MaxMIMO_LayersDCI_0_2_r16).Choice {
				case PUSCH_ServingCellConfig_Ext_MaxMIMO_LayersDCI_0_2_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_ServingCellConfig_Ext_MaxMIMO_LayersDCI_0_2_r16_Setup:
					if err := (*ie.MaxMIMO_LayersDCI_0_2_r16).Setup.Encode(ex); err != nil {
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
					{Name: "nrofHARQ-ProcessesForPUSCH-r17", Optional: true},
					{Name: "uplinkHARQ-mode-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.NrofHARQ_ProcessesForPUSCH_r17 != nil, ie.UplinkHARQ_Mode_r17 != nil}); err != nil {
				return err
			}

			if ie.NrofHARQ_ProcessesForPUSCH_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.NrofHARQ_ProcessesForPUSCH_r17, pUSCHServingCellConfigExtNrofHARQProcessesForPUSCHR17Constraints); err != nil {
					return err
				}
			}

			if ie.UplinkHARQ_Mode_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHServingCellConfigExtUplinkHARQModeR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.UplinkHARQ_Mode_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.UplinkHARQ_Mode_r17).Choice {
				case PUSCH_ServingCellConfig_Ext_UplinkHARQ_Mode_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_ServingCellConfig_Ext_UplinkHARQ_Mode_r17_Setup:
					if err := (*ie.UplinkHARQ_Mode_r17).Setup.Encode(ex); err != nil {
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
					{Name: "maxMIMO-Layers-v1810", Optional: true},
					{Name: "maxMIMO-LayersforSDM-r18", Optional: true},
					{Name: "maxMIMO-LayersforSDM-DCI-0-2-r18", Optional: true},
					{Name: "maxMIMO-LayersforSFN-r18", Optional: true},
					{Name: "maxMIMO-LayersforSFN-DCI-0-2-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxMIMO_Layers_v1810 != nil, ie.MaxMIMO_LayersforSDM_r18 != nil, ie.MaxMIMO_LayersforSDM_DCI_0_2_r18 != nil, ie.MaxMIMO_LayersforSFN_r18 != nil, ie.MaxMIMO_LayersforSFN_DCI_0_2_r18 != nil}); err != nil {
				return err
			}

			if ie.MaxMIMO_Layers_v1810 != nil {
				if err := ex.EncodeInteger(*ie.MaxMIMO_Layers_v1810, pUSCHServingCellConfigMaxMIMOLayersV1810Constraints); err != nil {
					return err
				}
			}

			if ie.MaxMIMO_LayersforSDM_r18 != nil {
				if err := ex.EncodeInteger(*ie.MaxMIMO_LayersforSDM_r18, pUSCHServingCellConfigMaxMIMOLayersforSDMR18Constraints); err != nil {
					return err
				}
			}

			if ie.MaxMIMO_LayersforSDM_DCI_0_2_r18 != nil {
				if err := ex.EncodeInteger(*ie.MaxMIMO_LayersforSDM_DCI_0_2_r18, pUSCHServingCellConfigMaxMIMOLayersforSDMDCI02R18Constraints); err != nil {
					return err
				}
			}

			if ie.MaxMIMO_LayersforSFN_r18 != nil {
				if err := ex.EncodeInteger(*ie.MaxMIMO_LayersforSFN_r18, pUSCHServingCellConfigMaxMIMOLayersforSFNR18Constraints); err != nil {
					return err
				}
			}

			if ie.MaxMIMO_LayersforSFN_DCI_0_2_r18 != nil {
				if err := ex.EncodeInteger(*ie.MaxMIMO_LayersforSFN_DCI_0_2_r18, pUSCHServingCellConfigMaxMIMOLayersforSFNDCI02R18Constraints); err != nil {
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

func (ie *PUSCH_ServingCellConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHServingCellConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. codeBlockGroupTransmission: choice
	{
		if seq.IsComponentPresent(0) {
			ie.CodeBlockGroupTransmission = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUSCH_CodeBlockGroupTransmission
			}{}
			choiceDec := d.NewChoiceDecoder(pUSCH_ServingCellConfigCodeBlockGroupTransmissionConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.CodeBlockGroupTransmission).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PUSCH_ServingCellConfig_CodeBlockGroupTransmission_Release:
				(*ie.CodeBlockGroupTransmission).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_ServingCellConfig_CodeBlockGroupTransmission_Setup:
				(*ie.CodeBlockGroupTransmission).Setup = new(PUSCH_CodeBlockGroupTransmission)
				if err := (*ie.CodeBlockGroupTransmission).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. rateMatching: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(pUSCHServingCellConfigRateMatchingConstraints)
			if err != nil {
				return err
			}
			ie.RateMatching = &idx
		}
	}

	// 5. xOverhead: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(pUSCHServingCellConfigXOverheadConstraints)
			if err != nil {
				return err
			}
			ie.XOverhead = &idx
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
				{Name: "maxMIMO-Layers", Optional: true},
				{Name: "processingType2Enabled", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(pUSCHServingCellConfigMaxMIMOLayersConstraints)
			if err != nil {
				return err
			}
			ie.MaxMIMO_Layers = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.ProcessingType2Enabled = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "maxMIMO-LayersDCI-0-2-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.MaxMIMO_LayersDCI_0_2_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MaxMIMO_LayersDCI_0_2_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHServingCellConfigExtMaxMIMOLayersDCI02R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MaxMIMO_LayersDCI_0_2_r16).Choice = int(index)
			switch index {
			case PUSCH_ServingCellConfig_Ext_MaxMIMO_LayersDCI_0_2_r16_Release:
				(*ie.MaxMIMO_LayersDCI_0_2_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_ServingCellConfig_Ext_MaxMIMO_LayersDCI_0_2_r16_Setup:
				(*ie.MaxMIMO_LayersDCI_0_2_r16).Setup = new(MaxMIMO_LayersDCI_0_2_r16)
				if err := (*ie.MaxMIMO_LayersDCI_0_2_r16).Setup.Decode(dx); err != nil {
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
				{Name: "nrofHARQ-ProcessesForPUSCH-r17", Optional: true},
				{Name: "uplinkHARQ-mode-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pUSCHServingCellConfigExtNrofHARQProcessesForPUSCHR17Constraints)
			if err != nil {
				return err
			}
			ie.NrofHARQ_ProcessesForPUSCH_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.UplinkHARQ_Mode_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UplinkHARQ_Mode_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHServingCellConfigExtUplinkHARQModeR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.UplinkHARQ_Mode_r17).Choice = int(index)
			switch index {
			case PUSCH_ServingCellConfig_Ext_UplinkHARQ_Mode_r17_Release:
				(*ie.UplinkHARQ_Mode_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_ServingCellConfig_Ext_UplinkHARQ_Mode_r17_Setup:
				(*ie.UplinkHARQ_Mode_r17).Setup = new(UplinkHARQ_Mode_r17)
				if err := (*ie.UplinkHARQ_Mode_r17).Setup.Decode(dx); err != nil {
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
				{Name: "maxMIMO-Layers-v1810", Optional: true},
				{Name: "maxMIMO-LayersforSDM-r18", Optional: true},
				{Name: "maxMIMO-LayersforSDM-DCI-0-2-r18", Optional: true},
				{Name: "maxMIMO-LayersforSFN-r18", Optional: true},
				{Name: "maxMIMO-LayersforSFN-DCI-0-2-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(pUSCHServingCellConfigMaxMIMOLayersV1810Constraints)
			if err != nil {
				return err
			}
			ie.MaxMIMO_Layers_v1810 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(pUSCHServingCellConfigMaxMIMOLayersforSDMR18Constraints)
			if err != nil {
				return err
			}
			ie.MaxMIMO_LayersforSDM_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(pUSCHServingCellConfigMaxMIMOLayersforSDMDCI02R18Constraints)
			if err != nil {
				return err
			}
			ie.MaxMIMO_LayersforSDM_DCI_0_2_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeInteger(pUSCHServingCellConfigMaxMIMOLayersforSFNR18Constraints)
			if err != nil {
				return err
			}
			ie.MaxMIMO_LayersforSFN_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeInteger(pUSCHServingCellConfigMaxMIMOLayersforSFNDCI02R18Constraints)
			if err != nil {
				return err
			}
			ie.MaxMIMO_LayersforSFN_DCI_0_2_r18 = &v
		}
	}

	return nil
}
