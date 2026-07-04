// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDSCH-ServingCellConfig (line 11480).

var pDSCHServingCellConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "codeBlockGroupTransmission", Optional: true},
		{Name: "xOverhead", Optional: true},
		{Name: "nrofHARQ-ProcessesForPDSCH", Optional: true},
		{Name: "pucch-Cell", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var pDSCH_ServingCellConfigCodeBlockGroupTransmissionConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_ServingCellConfig_CodeBlockGroupTransmission_Release = 0
	PDSCH_ServingCellConfig_CodeBlockGroupTransmission_Setup   = 1
)

const (
	PDSCH_ServingCellConfig_XOverhead_XOh6  = 0
	PDSCH_ServingCellConfig_XOverhead_XOh12 = 1
	PDSCH_ServingCellConfig_XOverhead_XOh18 = 2
)

var pDSCHServingCellConfigXOverheadConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_ServingCellConfig_XOverhead_XOh6, PDSCH_ServingCellConfig_XOverhead_XOh12, PDSCH_ServingCellConfig_XOverhead_XOh18},
}

const (
	PDSCH_ServingCellConfig_NrofHARQ_ProcessesForPDSCH_N2  = 0
	PDSCH_ServingCellConfig_NrofHARQ_ProcessesForPDSCH_N4  = 1
	PDSCH_ServingCellConfig_NrofHARQ_ProcessesForPDSCH_N6  = 2
	PDSCH_ServingCellConfig_NrofHARQ_ProcessesForPDSCH_N10 = 3
	PDSCH_ServingCellConfig_NrofHARQ_ProcessesForPDSCH_N12 = 4
	PDSCH_ServingCellConfig_NrofHARQ_ProcessesForPDSCH_N16 = 5
)

var pDSCHServingCellConfigNrofHARQProcessesForPDSCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_ServingCellConfig_NrofHARQ_ProcessesForPDSCH_N2, PDSCH_ServingCellConfig_NrofHARQ_ProcessesForPDSCH_N4, PDSCH_ServingCellConfig_NrofHARQ_ProcessesForPDSCH_N6, PDSCH_ServingCellConfig_NrofHARQ_ProcessesForPDSCH_N10, PDSCH_ServingCellConfig_NrofHARQ_ProcessesForPDSCH_N12, PDSCH_ServingCellConfig_NrofHARQ_ProcessesForPDSCH_N16},
}

var pDSCHServingCellConfigMaxMIMOLayersConstraints = per.Constrained(1, 8)

var pDSCHServingCellConfigExtPdschCodeBlockGroupTransmissionListR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_ServingCellConfig_Ext_Pdsch_CodeBlockGroupTransmissionList_r16_Release = 0
	PDSCH_ServingCellConfig_Ext_Pdsch_CodeBlockGroupTransmissionList_r16_Setup   = 1
)

var pDSCHServingCellConfigExtDownlinkHARQFeedbackDisabledR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_ServingCellConfig_Ext_DownlinkHARQ_FeedbackDisabled_r17_Release = 0
	PDSCH_ServingCellConfig_Ext_DownlinkHARQ_FeedbackDisabled_r17_Setup   = 1
)

const (
	PDSCH_ServingCellConfig_Ext_NrofHARQ_ProcessesForPDSCH_v1700_N32 = 0
)

var pDSCHServingCellConfigExtNrofHARQProcessesForPDSCHV1700Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_ServingCellConfig_Ext_NrofHARQ_ProcessesForPDSCH_v1700_N32},
}

type PDSCH_ServingCellConfig struct {
	CodeBlockGroupTransmission *struct {
		Choice  int
		Release *struct{}
		Setup   *PDSCH_CodeBlockGroupTransmission
	}
	XOverhead                                *int64
	NrofHARQ_ProcessesForPDSCH               *int64
	Pucch_Cell                               *ServCellIndex
	MaxMIMO_Layers                           *int64
	ProcessingType2Enabled                   *bool
	Pdsch_CodeBlockGroupTransmissionList_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *PDSCH_CodeBlockGroupTransmissionList_r16
	}
	DownlinkHARQ_FeedbackDisabled_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *DownlinkHARQ_FeedbackDisabled_r17
	}
	NrofHARQ_ProcessesForPDSCH_v1700 *int64
}

func (ie *PDSCH_ServingCellConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDSCHServingCellConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MaxMIMO_Layers != nil || ie.ProcessingType2Enabled != nil
	hasExtGroup1 := ie.Pdsch_CodeBlockGroupTransmissionList_r16 != nil
	hasExtGroup2 := ie.DownlinkHARQ_FeedbackDisabled_r17 != nil || ie.NrofHARQ_ProcessesForPDSCH_v1700 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CodeBlockGroupTransmission != nil, ie.XOverhead != nil, ie.NrofHARQ_ProcessesForPDSCH != nil, ie.Pucch_Cell != nil}); err != nil {
		return err
	}

	// 3. codeBlockGroupTransmission: choice
	{
		if ie.CodeBlockGroupTransmission != nil {
			choiceEnc := e.NewChoiceEncoder(pDSCH_ServingCellConfigCodeBlockGroupTransmissionConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.CodeBlockGroupTransmission).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.CodeBlockGroupTransmission).Choice {
			case PDSCH_ServingCellConfig_CodeBlockGroupTransmission_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PDSCH_ServingCellConfig_CodeBlockGroupTransmission_Setup:
				if err := (*ie.CodeBlockGroupTransmission).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.CodeBlockGroupTransmission).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. xOverhead: enumerated
	{
		if ie.XOverhead != nil {
			if err := e.EncodeEnumerated(*ie.XOverhead, pDSCHServingCellConfigXOverheadConstraints); err != nil {
				return err
			}
		}
	}

	// 5. nrofHARQ-ProcessesForPDSCH: enumerated
	{
		if ie.NrofHARQ_ProcessesForPDSCH != nil {
			if err := e.EncodeEnumerated(*ie.NrofHARQ_ProcessesForPDSCH, pDSCHServingCellConfigNrofHARQProcessesForPDSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 6. pucch-Cell: ref
	{
		if ie.Pucch_Cell != nil {
			if err := ie.Pucch_Cell.Encode(e); err != nil {
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
					{Name: "maxMIMO-Layers", Optional: true},
					{Name: "processingType2Enabled", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxMIMO_Layers != nil, ie.ProcessingType2Enabled != nil}); err != nil {
				return err
			}

			if ie.MaxMIMO_Layers != nil {
				if err := ex.EncodeInteger(*ie.MaxMIMO_Layers, pDSCHServingCellConfigMaxMIMOLayersConstraints); err != nil {
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
					{Name: "pdsch-CodeBlockGroupTransmissionList-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdsch_CodeBlockGroupTransmissionList_r16 != nil}); err != nil {
				return err
			}

			if ie.Pdsch_CodeBlockGroupTransmissionList_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHServingCellConfigExtPdschCodeBlockGroupTransmissionListR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_CodeBlockGroupTransmissionList_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdsch_CodeBlockGroupTransmissionList_r16).Choice {
				case PDSCH_ServingCellConfig_Ext_Pdsch_CodeBlockGroupTransmissionList_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_ServingCellConfig_Ext_Pdsch_CodeBlockGroupTransmissionList_r16_Setup:
					if err := (*ie.Pdsch_CodeBlockGroupTransmissionList_r16).Setup.Encode(ex); err != nil {
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
					{Name: "downlinkHARQ-FeedbackDisabled-r17", Optional: true},
					{Name: "nrofHARQ-ProcessesForPDSCH-v1700", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.DownlinkHARQ_FeedbackDisabled_r17 != nil, ie.NrofHARQ_ProcessesForPDSCH_v1700 != nil}); err != nil {
				return err
			}

			if ie.DownlinkHARQ_FeedbackDisabled_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHServingCellConfigExtDownlinkHARQFeedbackDisabledR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.DownlinkHARQ_FeedbackDisabled_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.DownlinkHARQ_FeedbackDisabled_r17).Choice {
				case PDSCH_ServingCellConfig_Ext_DownlinkHARQ_FeedbackDisabled_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_ServingCellConfig_Ext_DownlinkHARQ_FeedbackDisabled_r17_Setup:
					if err := (*ie.DownlinkHARQ_FeedbackDisabled_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.NrofHARQ_ProcessesForPDSCH_v1700 != nil {
				if err := ex.EncodeEnumerated(*ie.NrofHARQ_ProcessesForPDSCH_v1700, pDSCHServingCellConfigExtNrofHARQProcessesForPDSCHV1700Constraints); err != nil {
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

func (ie *PDSCH_ServingCellConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDSCHServingCellConfigConstraints)

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
				Setup   *PDSCH_CodeBlockGroupTransmission
			}{}
			choiceDec := d.NewChoiceDecoder(pDSCH_ServingCellConfigCodeBlockGroupTransmissionConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.CodeBlockGroupTransmission).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PDSCH_ServingCellConfig_CodeBlockGroupTransmission_Release:
				(*ie.CodeBlockGroupTransmission).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_ServingCellConfig_CodeBlockGroupTransmission_Setup:
				(*ie.CodeBlockGroupTransmission).Setup = new(PDSCH_CodeBlockGroupTransmission)
				if err := (*ie.CodeBlockGroupTransmission).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. xOverhead: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(pDSCHServingCellConfigXOverheadConstraints)
			if err != nil {
				return err
			}
			ie.XOverhead = &idx
		}
	}

	// 5. nrofHARQ-ProcessesForPDSCH: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(pDSCHServingCellConfigNrofHARQProcessesForPDSCHConstraints)
			if err != nil {
				return err
			}
			ie.NrofHARQ_ProcessesForPDSCH = &idx
		}
	}

	// 6. pucch-Cell: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Pucch_Cell = new(ServCellIndex)
			if err := ie.Pucch_Cell.Decode(d); err != nil {
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
				{Name: "maxMIMO-Layers", Optional: true},
				{Name: "processingType2Enabled", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(pDSCHServingCellConfigMaxMIMOLayersConstraints)
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
				{Name: "pdsch-CodeBlockGroupTransmissionList-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Pdsch_CodeBlockGroupTransmissionList_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDSCH_CodeBlockGroupTransmissionList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHServingCellConfigExtPdschCodeBlockGroupTransmissionListR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_CodeBlockGroupTransmissionList_r16).Choice = int(index)
			switch index {
			case PDSCH_ServingCellConfig_Ext_Pdsch_CodeBlockGroupTransmissionList_r16_Release:
				(*ie.Pdsch_CodeBlockGroupTransmissionList_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_ServingCellConfig_Ext_Pdsch_CodeBlockGroupTransmissionList_r16_Setup:
				(*ie.Pdsch_CodeBlockGroupTransmissionList_r16).Setup = new(PDSCH_CodeBlockGroupTransmissionList_r16)
				if err := (*ie.Pdsch_CodeBlockGroupTransmissionList_r16).Setup.Decode(dx); err != nil {
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
				{Name: "downlinkHARQ-FeedbackDisabled-r17", Optional: true},
				{Name: "nrofHARQ-ProcessesForPDSCH-v1700", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.DownlinkHARQ_FeedbackDisabled_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DownlinkHARQ_FeedbackDisabled_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHServingCellConfigExtDownlinkHARQFeedbackDisabledR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.DownlinkHARQ_FeedbackDisabled_r17).Choice = int(index)
			switch index {
			case PDSCH_ServingCellConfig_Ext_DownlinkHARQ_FeedbackDisabled_r17_Release:
				(*ie.DownlinkHARQ_FeedbackDisabled_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_ServingCellConfig_Ext_DownlinkHARQ_FeedbackDisabled_r17_Setup:
				(*ie.DownlinkHARQ_FeedbackDisabled_r17).Setup = new(DownlinkHARQ_FeedbackDisabled_r17)
				if err := (*ie.DownlinkHARQ_FeedbackDisabled_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pDSCHServingCellConfigExtNrofHARQProcessesForPDSCHV1700Constraints)
			if err != nil {
				return err
			}
			ie.NrofHARQ_ProcessesForPDSCH_v1700 = &v
		}
	}

	return nil
}
