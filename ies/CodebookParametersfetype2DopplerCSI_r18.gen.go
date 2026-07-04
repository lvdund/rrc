// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParametersfetype2DopplerCSI-r18 (line 18911).

var codebookParametersfetype2DopplerCSIR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "feType2Doppler-r18"},
		{Name: "maxNumberAperiodicCSI-RS-Resource-r18", Optional: true},
		{Name: "feType2DopplerM2R1-r18", Optional: true},
		{Name: "feType2DopplerR2-r18", Optional: true},
		{Name: "feType2DopplerL-N4D1-r18", Optional: true},
		{Name: "feType2DopplerR3R4-r18", Optional: true},
	},
}

const (
	CodebookParametersfetype2DopplerCSI_r18_MaxNumberAperiodicCSI_RS_Resource_r18_N4  = 0
	CodebookParametersfetype2DopplerCSI_r18_MaxNumberAperiodicCSI_RS_Resource_r18_N8  = 1
	CodebookParametersfetype2DopplerCSI_r18_MaxNumberAperiodicCSI_RS_Resource_r18_N12 = 2
)

var codebookParametersfetype2DopplerCSIR18MaxNumberAperiodicCSIRSResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfetype2DopplerCSI_r18_MaxNumberAperiodicCSI_RS_Resource_r18_N4, CodebookParametersfetype2DopplerCSI_r18_MaxNumberAperiodicCSI_RS_Resource_r18_N8, CodebookParametersfetype2DopplerCSI_r18_MaxNumberAperiodicCSI_RS_Resource_r18_N12},
}

var codebookParametersfetype2DopplerCSIR18FeType2DopplerM2R1R18Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersfetype2DopplerCSIR18FeType2DopplerR2R18Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersfetype2DopplerCSI_r18_FeType2DopplerL_N4D1_r18_Supported = 0
)

var codebookParametersfetype2DopplerCSIR18FeType2DopplerLN4D1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfetype2DopplerCSI_r18_FeType2DopplerL_N4D1_r18_Supported},
}

const (
	CodebookParametersfetype2DopplerCSI_r18_FeType2DopplerR3R4_r18_Supported = 0
)

var codebookParametersfetype2DopplerCSIR18FeType2DopplerR3R4R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfetype2DopplerCSI_r18_FeType2DopplerR3R4_r18_Supported},
}

var codebookParametersfetype2DopplerCSIR18FeType2DopplerR18SupportedCSIRSResourceListR18Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersfetype2DopplerCSI_r18_FeType2Doppler_r18_Scalingfactor_r18_N1 = 0
	CodebookParametersfetype2DopplerCSI_r18_FeType2Doppler_r18_Scalingfactor_r18_N2 = 1
	CodebookParametersfetype2DopplerCSI_r18_FeType2Doppler_r18_Scalingfactor_r18_N4 = 2
)

var codebookParametersfetype2DopplerCSIR18FeType2DopplerR18ScalingfactorR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfetype2DopplerCSI_r18_FeType2Doppler_r18_Scalingfactor_r18_N1, CodebookParametersfetype2DopplerCSI_r18_FeType2Doppler_r18_Scalingfactor_r18_N2, CodebookParametersfetype2DopplerCSI_r18_FeType2Doppler_r18_Scalingfactor_r18_N4},
}

type CodebookParametersfetype2DopplerCSI_r18 struct {
	FeType2Doppler_r18 struct {
		SupportedCSI_RS_ResourceList_r18 []int64
		ValueY_A_CSI_RS_r18              int64
		Scalingfactor_r18                int64
	}
	MaxNumberAperiodicCSI_RS_Resource_r18 *int64
	FeType2DopplerM2R1_r18                []int64
	FeType2DopplerR2_r18                  []int64
	FeType2DopplerL_N4D1_r18              *int64
	FeType2DopplerR3R4_r18                *int64
}

func (ie *CodebookParametersfetype2DopplerCSI_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParametersfetype2DopplerCSIR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxNumberAperiodicCSI_RS_Resource_r18 != nil, ie.FeType2DopplerM2R1_r18 != nil, ie.FeType2DopplerR2_r18 != nil, ie.FeType2DopplerL_N4D1_r18 != nil, ie.FeType2DopplerR3R4_r18 != nil}); err != nil {
		return err
	}

	// 2. feType2Doppler-r18: sequence
	{
		{
			c := &ie.FeType2Doppler_r18
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersfetype2DopplerCSIR18FeType2DopplerR18SupportedCSIRSResourceListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceList_r18))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceList_r18 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceList_r18[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeInteger(c.ValueY_A_CSI_RS_r18, per.Constrained(1, 3)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Scalingfactor_r18, codebookParametersfetype2DopplerCSIR18FeType2DopplerR18ScalingfactorR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. maxNumberAperiodicCSI-RS-Resource-r18: enumerated
	{
		if ie.MaxNumberAperiodicCSI_RS_Resource_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumberAperiodicCSI_RS_Resource_r18, codebookParametersfetype2DopplerCSIR18MaxNumberAperiodicCSIRSResourceR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. feType2DopplerM2R1-r18: sequence-of
	{
		if ie.FeType2DopplerM2R1_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersfetype2DopplerCSIR18FeType2DopplerM2R1R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeType2DopplerM2R1_r18))); err != nil {
				return err
			}
			for i := range ie.FeType2DopplerM2R1_r18 {
				if err := e.EncodeInteger(ie.FeType2DopplerM2R1_r18[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 5. feType2DopplerR2-r18: sequence-of
	{
		if ie.FeType2DopplerR2_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersfetype2DopplerCSIR18FeType2DopplerR2R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeType2DopplerR2_r18))); err != nil {
				return err
			}
			for i := range ie.FeType2DopplerR2_r18 {
				if err := e.EncodeInteger(ie.FeType2DopplerR2_r18[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 6. feType2DopplerL-N4D1-r18: enumerated
	{
		if ie.FeType2DopplerL_N4D1_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FeType2DopplerL_N4D1_r18, codebookParametersfetype2DopplerCSIR18FeType2DopplerLN4D1R18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. feType2DopplerR3R4-r18: enumerated
	{
		if ie.FeType2DopplerR3R4_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FeType2DopplerR3R4_r18, codebookParametersfetype2DopplerCSIR18FeType2DopplerR3R4R18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CodebookParametersfetype2DopplerCSI_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParametersfetype2DopplerCSIR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. feType2Doppler-r18: sequence
	{
		{
			c := &ie.FeType2Doppler_r18
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersfetype2DopplerCSIR18FeType2DopplerR18SupportedCSIRSResourceListR18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceList_r18 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceList_r18[i] = v
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.ValueY_A_CSI_RS_r18 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParametersfetype2DopplerCSIR18FeType2DopplerR18ScalingfactorR18Constraints)
				if err != nil {
					return err
				}
				c.Scalingfactor_r18 = v
			}
		}
	}

	// 3. maxNumberAperiodicCSI-RS-Resource-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(codebookParametersfetype2DopplerCSIR18MaxNumberAperiodicCSIRSResourceR18Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberAperiodicCSI_RS_Resource_r18 = &idx
		}
	}

	// 4. feType2DopplerM2R1-r18: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersfetype2DopplerCSIR18FeType2DopplerM2R1R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeType2DopplerM2R1_r18 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.FeType2DopplerM2R1_r18[i] = v
			}
		}
	}

	// 5. feType2DopplerR2-r18: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersfetype2DopplerCSIR18FeType2DopplerR2R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeType2DopplerR2_r18 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.FeType2DopplerR2_r18[i] = v
			}
		}
	}

	// 6. feType2DopplerL-N4D1-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(codebookParametersfetype2DopplerCSIR18FeType2DopplerLN4D1R18Constraints)
			if err != nil {
				return err
			}
			ie.FeType2DopplerL_N4D1_r18 = &idx
		}
	}

	// 7. feType2DopplerR3R4-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(codebookParametersfetype2DopplerCSIR18FeType2DopplerR3R4R18Constraints)
			if err != nil {
				return err
			}
			ie.FeType2DopplerR3R4_r18 = &idx
		}
	}

	return nil
}
