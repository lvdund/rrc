// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParametersfetype2CJT-r18 (line 18969).

var codebookParametersfetype2CJTR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "feType2CJT-r18"},
		{Name: "feType2CJT-FD-IO-r18", Optional: true},
		{Name: "feType2CJT-FD-FO-r18", Optional: true},
		{Name: "feType2CJT-M2R1-r18", Optional: true},
		{Name: "feType2CJT-R2-r18", Optional: true},
		{Name: "feType2CJT-2NN1N2-r18", Optional: true},
		{Name: "feType2CJT-Rank3Rank4-r18", Optional: true},
		{Name: "feType2CJT-NN-r18", Optional: true},
		{Name: "feType2CJT-NL-r18", Optional: true},
		{Name: "feType2CJT-Unequal-r18", Optional: true},
	},
}

var codebookParametersfetype2CJTR18FeType2CJTFDIOR18Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersfetype2CJT_r18_FeType2CJT_FD_FO_r18_Supported = 0
)

var codebookParametersfetype2CJTR18FeType2CJTFDFOR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfetype2CJT_r18_FeType2CJT_FD_FO_r18_Supported},
}

var codebookParametersfetype2CJTR18FeType2CJTM2R1R18Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersfetype2CJTR18FeType2CJTR2R18Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersfetype2CJT_r18_FeType2CJT_2NN1N2_r18_N64  = 0
	CodebookParametersfetype2CJT_r18_FeType2CJT_2NN1N2_r18_N96  = 1
	CodebookParametersfetype2CJT_r18_FeType2CJT_2NN1N2_r18_N128 = 2
)

var codebookParametersfetype2CJTR18FeType2CJT2NN1N2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfetype2CJT_r18_FeType2CJT_2NN1N2_r18_N64, CodebookParametersfetype2CJT_r18_FeType2CJT_2NN1N2_r18_N96, CodebookParametersfetype2CJT_r18_FeType2CJT_2NN1N2_r18_N128},
}

const (
	CodebookParametersfetype2CJT_r18_FeType2CJT_Rank3Rank4_r18_Supported = 0
)

var codebookParametersfetype2CJTR18FeType2CJTRank3Rank4R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfetype2CJT_r18_FeType2CJT_Rank3Rank4_r18_Supported},
}

const (
	CodebookParametersfetype2CJT_r18_FeType2CJT_NN_r18_Supported = 0
)

var codebookParametersfetype2CJTR18FeType2CJTNNR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfetype2CJT_r18_FeType2CJT_NN_r18_Supported},
}

const (
	CodebookParametersfetype2CJT_r18_FeType2CJT_NL_r18_N2 = 0
	CodebookParametersfetype2CJT_r18_FeType2CJT_NL_r18_N4 = 1
)

var codebookParametersfetype2CJTR18FeType2CJTNLR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfetype2CJT_r18_FeType2CJT_NL_r18_N2, CodebookParametersfetype2CJT_r18_FeType2CJT_NL_r18_N4},
}

const (
	CodebookParametersfetype2CJT_r18_FeType2CJT_Unequal_r18_Supported = 0
)

var codebookParametersfetype2CJTR18FeType2CJTUnequalR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfetype2CJT_r18_FeType2CJT_Unequal_r18_Supported},
}

var codebookParametersfetype2CJTR18FeType2CJTR18SupportedCSIRSResourceListR18Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersfetype2CJT_r18_FeType2CJT_r18_Scalingfactor_r18_N1     = 0
	CodebookParametersfetype2CJT_r18_FeType2CJT_r18_Scalingfactor_r18_N1dot5 = 1
	CodebookParametersfetype2CJT_r18_FeType2CJT_r18_Scalingfactor_r18_N2     = 2
)

var codebookParametersfetype2CJTR18FeType2CJTR18ScalingfactorR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfetype2CJT_r18_FeType2CJT_r18_Scalingfactor_r18_N1, CodebookParametersfetype2CJT_r18_FeType2CJT_r18_Scalingfactor_r18_N1dot5, CodebookParametersfetype2CJT_r18_FeType2CJT_r18_Scalingfactor_r18_N2},
}

type CodebookParametersfetype2CJT_r18 struct {
	FeType2CJT_r18 struct {
		SupportedCSI_RS_ResourceList_r18     []int64
		Scalingfactor_r18                    int64
		MaxNumberNZP_CSI_RS_MultiTRP_CJT_r18 int64
	}
	FeType2CJT_FD_IO_r18      []int64
	FeType2CJT_FD_FO_r18      *int64
	FeType2CJT_M2R1_r18       []int64
	FeType2CJT_R2_r18         []int64
	FeType2CJT_2NN1N2_r18     *int64
	FeType2CJT_Rank3Rank4_r18 *int64
	FeType2CJT_NN_r18         *int64
	FeType2CJT_NL_r18         *int64
	FeType2CJT_Unequal_r18    *int64
}

func (ie *CodebookParametersfetype2CJT_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParametersfetype2CJTR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FeType2CJT_FD_IO_r18 != nil, ie.FeType2CJT_FD_FO_r18 != nil, ie.FeType2CJT_M2R1_r18 != nil, ie.FeType2CJT_R2_r18 != nil, ie.FeType2CJT_2NN1N2_r18 != nil, ie.FeType2CJT_Rank3Rank4_r18 != nil, ie.FeType2CJT_NN_r18 != nil, ie.FeType2CJT_NL_r18 != nil, ie.FeType2CJT_Unequal_r18 != nil}); err != nil {
		return err
	}

	// 2. feType2CJT-r18: sequence
	{
		{
			c := &ie.FeType2CJT_r18
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersfetype2CJTR18FeType2CJTR18SupportedCSIRSResourceListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceList_r18))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceList_r18 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceList_r18[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.Scalingfactor_r18, codebookParametersfetype2CJTR18FeType2CJTR18ScalingfactorR18Constraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberNZP_CSI_RS_MultiTRP_CJT_r18, per.Constrained(2, 4)); err != nil {
				return err
			}
		}
	}

	// 3. feType2CJT-FD-IO-r18: sequence-of
	{
		if ie.FeType2CJT_FD_IO_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersfetype2CJTR18FeType2CJTFDIOR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeType2CJT_FD_IO_r18))); err != nil {
				return err
			}
			for i := range ie.FeType2CJT_FD_IO_r18 {
				if err := e.EncodeInteger(ie.FeType2CJT_FD_IO_r18[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 4. feType2CJT-FD-FO-r18: enumerated
	{
		if ie.FeType2CJT_FD_FO_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FeType2CJT_FD_FO_r18, codebookParametersfetype2CJTR18FeType2CJTFDFOR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. feType2CJT-M2R1-r18: sequence-of
	{
		if ie.FeType2CJT_M2R1_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersfetype2CJTR18FeType2CJTM2R1R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeType2CJT_M2R1_r18))); err != nil {
				return err
			}
			for i := range ie.FeType2CJT_M2R1_r18 {
				if err := e.EncodeInteger(ie.FeType2CJT_M2R1_r18[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 6. feType2CJT-R2-r18: sequence-of
	{
		if ie.FeType2CJT_R2_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersfetype2CJTR18FeType2CJTR2R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeType2CJT_R2_r18))); err != nil {
				return err
			}
			for i := range ie.FeType2CJT_R2_r18 {
				if err := e.EncodeInteger(ie.FeType2CJT_R2_r18[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 7. feType2CJT-2NN1N2-r18: enumerated
	{
		if ie.FeType2CJT_2NN1N2_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FeType2CJT_2NN1N2_r18, codebookParametersfetype2CJTR18FeType2CJT2NN1N2R18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. feType2CJT-Rank3Rank4-r18: enumerated
	{
		if ie.FeType2CJT_Rank3Rank4_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FeType2CJT_Rank3Rank4_r18, codebookParametersfetype2CJTR18FeType2CJTRank3Rank4R18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. feType2CJT-NN-r18: enumerated
	{
		if ie.FeType2CJT_NN_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FeType2CJT_NN_r18, codebookParametersfetype2CJTR18FeType2CJTNNR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. feType2CJT-NL-r18: enumerated
	{
		if ie.FeType2CJT_NL_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FeType2CJT_NL_r18, codebookParametersfetype2CJTR18FeType2CJTNLR18Constraints); err != nil {
				return err
			}
		}
	}

	// 11. feType2CJT-Unequal-r18: enumerated
	{
		if ie.FeType2CJT_Unequal_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FeType2CJT_Unequal_r18, codebookParametersfetype2CJTR18FeType2CJTUnequalR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CodebookParametersfetype2CJT_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParametersfetype2CJTR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. feType2CJT-r18: sequence
	{
		{
			c := &ie.FeType2CJT_r18
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersfetype2CJTR18FeType2CJTR18SupportedCSIRSResourceListR18Constraints)
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
				v, err := d.DecodeEnumerated(codebookParametersfetype2CJTR18FeType2CJTR18ScalingfactorR18Constraints)
				if err != nil {
					return err
				}
				c.Scalingfactor_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.MaxNumberNZP_CSI_RS_MultiTRP_CJT_r18 = v
			}
		}
	}

	// 3. feType2CJT-FD-IO-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersfetype2CJTR18FeType2CJTFDIOR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeType2CJT_FD_IO_r18 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.FeType2CJT_FD_IO_r18[i] = v
			}
		}
	}

	// 4. feType2CJT-FD-FO-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(codebookParametersfetype2CJTR18FeType2CJTFDFOR18Constraints)
			if err != nil {
				return err
			}
			ie.FeType2CJT_FD_FO_r18 = &idx
		}
	}

	// 5. feType2CJT-M2R1-r18: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersfetype2CJTR18FeType2CJTM2R1R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeType2CJT_M2R1_r18 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.FeType2CJT_M2R1_r18[i] = v
			}
		}
	}

	// 6. feType2CJT-R2-r18: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersfetype2CJTR18FeType2CJTR2R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeType2CJT_R2_r18 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.FeType2CJT_R2_r18[i] = v
			}
		}
	}

	// 7. feType2CJT-2NN1N2-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(codebookParametersfetype2CJTR18FeType2CJT2NN1N2R18Constraints)
			if err != nil {
				return err
			}
			ie.FeType2CJT_2NN1N2_r18 = &idx
		}
	}

	// 8. feType2CJT-Rank3Rank4-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(codebookParametersfetype2CJTR18FeType2CJTRank3Rank4R18Constraints)
			if err != nil {
				return err
			}
			ie.FeType2CJT_Rank3Rank4_r18 = &idx
		}
	}

	// 9. feType2CJT-NN-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(codebookParametersfetype2CJTR18FeType2CJTNNR18Constraints)
			if err != nil {
				return err
			}
			ie.FeType2CJT_NN_r18 = &idx
		}
	}

	// 10. feType2CJT-NL-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(codebookParametersfetype2CJTR18FeType2CJTNLR18Constraints)
			if err != nil {
				return err
			}
			ie.FeType2CJT_NL_r18 = &idx
		}
	}

	// 11. feType2CJT-Unequal-r18: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(codebookParametersfetype2CJTR18FeType2CJTUnequalR18Constraints)
			if err != nil {
				return err
			}
			ie.FeType2CJT_Unequal_r18 = &idx
		}
	}

	return nil
}
