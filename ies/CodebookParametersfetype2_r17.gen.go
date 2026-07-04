// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParametersfetype2-r17 (line 18606).

var codebookParametersfetype2R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fetype2basic-r17"},
		{Name: "fetype2R1-r17", Optional: true},
		{Name: "fetype2R2-r17", Optional: true},
		{Name: "fetype2Rank3Rank4-r17", Optional: true},
	},
}

var codebookParametersfetype2R17Fetype2basicR17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersfetype2R17Fetype2R1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r17)

var codebookParametersfetype2R17Fetype2R2R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r17)

const (
	CodebookParametersfetype2_r17_Fetype2Rank3Rank4_r17_Supported = 0
)

var codebookParametersfetype2R17Fetype2Rank3Rank4R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfetype2_r17_Fetype2Rank3Rank4_r17_Supported},
}

type CodebookParametersfetype2_r17 struct {
	Fetype2basic_r17      []int64
	Fetype2R1_r17         []int64
	Fetype2R2_r17         []int64
	Fetype2Rank3Rank4_r17 *int64
}

func (ie *CodebookParametersfetype2_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParametersfetype2R17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Fetype2R1_r17 != nil, ie.Fetype2R2_r17 != nil, ie.Fetype2Rank3Rank4_r17 != nil}); err != nil {
		return err
	}

	// 2. fetype2basic-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(codebookParametersfetype2R17Fetype2basicR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Fetype2basic_r17))); err != nil {
			return err
		}
		for i := range ie.Fetype2basic_r17 {
			if err := e.EncodeInteger(ie.Fetype2basic_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
				return err
			}
		}
	}

	// 3. fetype2R1-r17: sequence-of
	{
		if ie.Fetype2R1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersfetype2R17Fetype2R1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Fetype2R1_r17))); err != nil {
				return err
			}
			for i := range ie.Fetype2R1_r17 {
				if err := e.EncodeInteger(ie.Fetype2R1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 4. fetype2R2-r17: sequence-of
	{
		if ie.Fetype2R2_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersfetype2R17Fetype2R2R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Fetype2R2_r17))); err != nil {
				return err
			}
			for i := range ie.Fetype2R2_r17 {
				if err := e.EncodeInteger(ie.Fetype2R2_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 5. fetype2Rank3Rank4-r17: enumerated
	{
		if ie.Fetype2Rank3Rank4_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Fetype2Rank3Rank4_r17, codebookParametersfetype2R17Fetype2Rank3Rank4R17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CodebookParametersfetype2_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParametersfetype2R17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. fetype2basic-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(codebookParametersfetype2R17Fetype2basicR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Fetype2basic_r17 = make([]int64, n)
		for i := int64(0); i < n; i++ {
			v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
			if err != nil {
				return err
			}
			ie.Fetype2basic_r17[i] = v
		}
	}

	// 3. fetype2R1-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersfetype2R17Fetype2R1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Fetype2R1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Fetype2R1_r17[i] = v
			}
		}
	}

	// 4. fetype2R2-r17: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersfetype2R17Fetype2R2R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Fetype2R2_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Fetype2R2_r17[i] = v
			}
		}
	}

	// 5. fetype2Rank3Rank4-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(codebookParametersfetype2R17Fetype2Rank3Rank4R17Constraints)
			if err != nil {
				return err
			}
			ie.Fetype2Rank3Rank4_r17 = &idx
		}
	}

	return nil
}
