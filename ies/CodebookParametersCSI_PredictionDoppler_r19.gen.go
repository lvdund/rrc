// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParametersCSI-PredictionDoppler-r19 (line 19293).

var codebookParametersCSIPredictionDopplerR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberOfAperiodic-CSI-RS-Resource-r19", Optional: true},
		{Name: "eType2DopplerR2-r19", Optional: true},
		{Name: "eType2DopplerX1-r19", Optional: true},
		{Name: "eType2DopplerX2-r19", Optional: true},
		{Name: "eType2DopplerL-N4D1-r19", Optional: true},
		{Name: "eType2DopplerL6-r19", Optional: true},
		{Name: "eType2DopplerR3R4-r19", Optional: true},
		{Name: "codebookComboParameterMixedTypePrediction-r19", Optional: true},
	},
}

const (
	CodebookParametersCSI_PredictionDoppler_r19_MaxNumberOfAperiodic_CSI_RS_Resource_r19_N4  = 0
	CodebookParametersCSI_PredictionDoppler_r19_MaxNumberOfAperiodic_CSI_RS_Resource_r19_N8  = 1
	CodebookParametersCSI_PredictionDoppler_r19_MaxNumberOfAperiodic_CSI_RS_Resource_r19_N12 = 2
)

var codebookParametersCSIPredictionDopplerR19MaxNumberOfAperiodicCSIRSResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersCSI_PredictionDoppler_r19_MaxNumberOfAperiodic_CSI_RS_Resource_r19_N4, CodebookParametersCSI_PredictionDoppler_r19_MaxNumberOfAperiodic_CSI_RS_Resource_r19_N8, CodebookParametersCSI_PredictionDoppler_r19_MaxNumberOfAperiodic_CSI_RS_Resource_r19_N12},
}

var codebookParametersCSIPredictionDopplerR19EType2DopplerR2R19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersCSI_PredictionDoppler_r19_EType2DopplerX1_r19_Supported = 0
)

var codebookParametersCSIPredictionDopplerR19EType2DopplerX1R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersCSI_PredictionDoppler_r19_EType2DopplerX1_r19_Supported},
}

const (
	CodebookParametersCSI_PredictionDoppler_r19_EType2DopplerX2_r19_Supported = 0
)

var codebookParametersCSIPredictionDopplerR19EType2DopplerX2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersCSI_PredictionDoppler_r19_EType2DopplerX2_r19_Supported},
}

const (
	CodebookParametersCSI_PredictionDoppler_r19_EType2DopplerL_N4D1_r19_Supported = 0
)

var codebookParametersCSIPredictionDopplerR19EType2DopplerLN4D1R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersCSI_PredictionDoppler_r19_EType2DopplerL_N4D1_r19_Supported},
}

const (
	CodebookParametersCSI_PredictionDoppler_r19_EType2DopplerL6_r19_Supported = 0
)

var codebookParametersCSIPredictionDopplerR19EType2DopplerL6R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersCSI_PredictionDoppler_r19_EType2DopplerL6_r19_Supported},
}

const (
	CodebookParametersCSI_PredictionDoppler_r19_EType2DopplerR3R4_r19_Supported = 0
)

var codebookParametersCSIPredictionDopplerR19EType2DopplerR3R4R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersCSI_PredictionDoppler_r19_EType2DopplerR3R4_r19_Supported},
}

var codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "type1SP-Type1SP-N4-r19", Optional: true},
		{Name: "type1SP-eType2SP-r19", Optional: true},
		{Name: "type1SP-eType2SP-N4-r19", Optional: true},
		{Name: "type1SP-N4-eType2SP-r19", Optional: true},
		{Name: "type1SP-N4-eType2SP-N4-r19", Optional: true},
		{Name: "eType2SP-eType2SP-N4-r19", Optional: true},
	},
}

var codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPType1SPN4R19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPEType2SPR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPEType2SPN4R19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPN4EType2SPR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPN4EType2SPN4R19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19EType2SPEType2SPN4R19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

type CodebookParametersCSI_PredictionDoppler_r19 struct {
	MaxNumberOfAperiodic_CSI_RS_Resource_r19      *int64
	EType2DopplerR2_r19                           []int64
	EType2DopplerX1_r19                           *int64
	EType2DopplerX2_r19                           *int64
	EType2DopplerL_N4D1_r19                       *int64
	EType2DopplerL6_r19                           *int64
	EType2DopplerR3R4_r19                         *int64
	CodebookComboParameterMixedTypePrediction_r19 *struct {
		Type1SP_Type1SP_N4_r19     []int64
		Type1SP_EType2SP_r19       []int64
		Type1SP_EType2SP_N4_r19    []int64
		Type1SP_N4_EType2SP_r19    []int64
		Type1SP_N4_EType2SP_N4_r19 []int64
		EType2SP_EType2SP_N4_r19   []int64
	}
}

func (ie *CodebookParametersCSI_PredictionDoppler_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParametersCSIPredictionDopplerR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxNumberOfAperiodic_CSI_RS_Resource_r19 != nil, ie.EType2DopplerR2_r19 != nil, ie.EType2DopplerX1_r19 != nil, ie.EType2DopplerX2_r19 != nil, ie.EType2DopplerL_N4D1_r19 != nil, ie.EType2DopplerL6_r19 != nil, ie.EType2DopplerR3R4_r19 != nil, ie.CodebookComboParameterMixedTypePrediction_r19 != nil}); err != nil {
		return err
	}

	// 2. maxNumberOfAperiodic-CSI-RS-Resource-r19: enumerated
	{
		if ie.MaxNumberOfAperiodic_CSI_RS_Resource_r19 != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumberOfAperiodic_CSI_RS_Resource_r19, codebookParametersCSIPredictionDopplerR19MaxNumberOfAperiodicCSIRSResourceR19Constraints); err != nil {
				return err
			}
		}
	}

	// 3. eType2DopplerR2-r19: sequence-of
	{
		if ie.EType2DopplerR2_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersCSIPredictionDopplerR19EType2DopplerR2R19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.EType2DopplerR2_r19))); err != nil {
				return err
			}
			for i := range ie.EType2DopplerR2_r19 {
				if err := e.EncodeInteger(ie.EType2DopplerR2_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 4. eType2DopplerX1-r19: enumerated
	{
		if ie.EType2DopplerX1_r19 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerX1_r19, codebookParametersCSIPredictionDopplerR19EType2DopplerX1R19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. eType2DopplerX2-r19: enumerated
	{
		if ie.EType2DopplerX2_r19 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerX2_r19, codebookParametersCSIPredictionDopplerR19EType2DopplerX2R19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. eType2DopplerL-N4D1-r19: enumerated
	{
		if ie.EType2DopplerL_N4D1_r19 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerL_N4D1_r19, codebookParametersCSIPredictionDopplerR19EType2DopplerLN4D1R19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. eType2DopplerL6-r19: enumerated
	{
		if ie.EType2DopplerL6_r19 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerL6_r19, codebookParametersCSIPredictionDopplerR19EType2DopplerL6R19Constraints); err != nil {
				return err
			}
		}
	}

	// 8. eType2DopplerR3R4-r19: enumerated
	{
		if ie.EType2DopplerR3R4_r19 != nil {
			if err := e.EncodeEnumerated(*ie.EType2DopplerR3R4_r19, codebookParametersCSIPredictionDopplerR19EType2DopplerR3R4R19Constraints); err != nil {
				return err
			}
		}
	}

	// 9. codebookComboParameterMixedTypePrediction-r19: sequence
	{
		if ie.CodebookComboParameterMixedTypePrediction_r19 != nil {
			c := ie.CodebookComboParameterMixedTypePrediction_r19
			codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Seq := e.NewSequenceEncoder(codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Constraints)
			if err := codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Seq.EncodePreamble([]bool{c.Type1SP_Type1SP_N4_r19 != nil, c.Type1SP_EType2SP_r19 != nil, c.Type1SP_EType2SP_N4_r19 != nil, c.Type1SP_N4_EType2SP_r19 != nil, c.Type1SP_N4_EType2SP_N4_r19 != nil, c.EType2SP_EType2SP_N4_r19 != nil}); err != nil {
				return err
			}
			if c.Type1SP_Type1SP_N4_r19 != nil {
				seqOf := e.NewSequenceOfEncoder(codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPType1SPN4R19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.Type1SP_Type1SP_N4_r19))); err != nil {
					return err
				}
				for i := range c.Type1SP_Type1SP_N4_r19 {
					if err := e.EncodeInteger(c.Type1SP_Type1SP_N4_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if c.Type1SP_EType2SP_r19 != nil {
				seqOf := e.NewSequenceOfEncoder(codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPEType2SPR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.Type1SP_EType2SP_r19))); err != nil {
					return err
				}
				for i := range c.Type1SP_EType2SP_r19 {
					if err := e.EncodeInteger(c.Type1SP_EType2SP_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if c.Type1SP_EType2SP_N4_r19 != nil {
				seqOf := e.NewSequenceOfEncoder(codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPEType2SPN4R19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.Type1SP_EType2SP_N4_r19))); err != nil {
					return err
				}
				for i := range c.Type1SP_EType2SP_N4_r19 {
					if err := e.EncodeInteger(c.Type1SP_EType2SP_N4_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if c.Type1SP_N4_EType2SP_r19 != nil {
				seqOf := e.NewSequenceOfEncoder(codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPN4EType2SPR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.Type1SP_N4_EType2SP_r19))); err != nil {
					return err
				}
				for i := range c.Type1SP_N4_EType2SP_r19 {
					if err := e.EncodeInteger(c.Type1SP_N4_EType2SP_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if c.Type1SP_N4_EType2SP_N4_r19 != nil {
				seqOf := e.NewSequenceOfEncoder(codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPN4EType2SPN4R19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.Type1SP_N4_EType2SP_N4_r19))); err != nil {
					return err
				}
				for i := range c.Type1SP_N4_EType2SP_N4_r19 {
					if err := e.EncodeInteger(c.Type1SP_N4_EType2SP_N4_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if c.EType2SP_EType2SP_N4_r19 != nil {
				seqOf := e.NewSequenceOfEncoder(codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19EType2SPEType2SPN4R19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.EType2SP_EType2SP_N4_r19))); err != nil {
					return err
				}
				for i := range c.EType2SP_EType2SP_N4_r19 {
					if err := e.EncodeInteger(c.EType2SP_EType2SP_N4_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (ie *CodebookParametersCSI_PredictionDoppler_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParametersCSIPredictionDopplerR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. maxNumberOfAperiodic-CSI-RS-Resource-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(codebookParametersCSIPredictionDopplerR19MaxNumberOfAperiodicCSIRSResourceR19Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberOfAperiodic_CSI_RS_Resource_r19 = &idx
		}
	}

	// 3. eType2DopplerR2-r19: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersCSIPredictionDopplerR19EType2DopplerR2R19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.EType2DopplerR2_r19 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.EType2DopplerR2_r19[i] = v
			}
		}
	}

	// 4. eType2DopplerX1-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(codebookParametersCSIPredictionDopplerR19EType2DopplerX1R19Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerX1_r19 = &idx
		}
	}

	// 5. eType2DopplerX2-r19: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(codebookParametersCSIPredictionDopplerR19EType2DopplerX2R19Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerX2_r19 = &idx
		}
	}

	// 6. eType2DopplerL-N4D1-r19: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(codebookParametersCSIPredictionDopplerR19EType2DopplerLN4D1R19Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerL_N4D1_r19 = &idx
		}
	}

	// 7. eType2DopplerL6-r19: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(codebookParametersCSIPredictionDopplerR19EType2DopplerL6R19Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerL6_r19 = &idx
		}
	}

	// 8. eType2DopplerR3R4-r19: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(codebookParametersCSIPredictionDopplerR19EType2DopplerR3R4R19Constraints)
			if err != nil {
				return err
			}
			ie.EType2DopplerR3R4_r19 = &idx
		}
	}

	// 9. codebookComboParameterMixedTypePrediction-r19: sequence
	{
		if seq.IsComponentPresent(7) {
			ie.CodebookComboParameterMixedTypePrediction_r19 = &struct {
				Type1SP_Type1SP_N4_r19     []int64
				Type1SP_EType2SP_r19       []int64
				Type1SP_EType2SP_N4_r19    []int64
				Type1SP_N4_EType2SP_r19    []int64
				Type1SP_N4_EType2SP_N4_r19 []int64
				EType2SP_EType2SP_N4_r19   []int64
			}{}
			c := ie.CodebookComboParameterMixedTypePrediction_r19
			codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Seq := d.NewSequenceDecoder(codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Constraints)
			if err := codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Seq.IsComponentPresent(0) {
				seqOf := d.NewSequenceOfDecoder(codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPType1SPN4R19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Type1SP_Type1SP_N4_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.Type1SP_Type1SP_N4_r19[i] = v
				}
			}
			if codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Seq.IsComponentPresent(1) {
				seqOf := d.NewSequenceOfDecoder(codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPEType2SPR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Type1SP_EType2SP_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.Type1SP_EType2SP_r19[i] = v
				}
			}
			if codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Seq.IsComponentPresent(2) {
				seqOf := d.NewSequenceOfDecoder(codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPEType2SPN4R19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Type1SP_EType2SP_N4_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.Type1SP_EType2SP_N4_r19[i] = v
				}
			}
			if codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Seq.IsComponentPresent(3) {
				seqOf := d.NewSequenceOfDecoder(codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPN4EType2SPR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Type1SP_N4_EType2SP_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.Type1SP_N4_EType2SP_r19[i] = v
				}
			}
			if codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Seq.IsComponentPresent(4) {
				seqOf := d.NewSequenceOfDecoder(codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Type1SPN4EType2SPN4R19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Type1SP_N4_EType2SP_N4_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.Type1SP_N4_EType2SP_N4_r19[i] = v
				}
			}
			if codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19Seq.IsComponentPresent(5) {
				seqOf := d.NewSequenceOfDecoder(codebookParametersCSIPredictionDopplerR19CodebookComboParameterMixedTypePredictionR19EType2SPEType2SPN4R19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.EType2SP_EType2SP_N4_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.EType2SP_EType2SP_N4_r19[i] = v
				}
			}
		}
	}

	return nil
}
