// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookComboParametersCJT-r18 (line 19001).

var codebookComboParametersCJTR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cjt-Type1SP-eType2R1-null", Optional: true},
		{Name: "cjt-Type1SP-eType2R2-null", Optional: true},
		{Name: "cjt-Type1SP-feType2R1M1-null", Optional: true},
		{Name: "cjt-Type1SP-feType2R1M2-null", Optional: true},
		{Name: "cjt-Type1SP-feType2R2M2-null", Optional: true},
		{Name: "cjt-Type1MP-eType2R1-null", Optional: true},
		{Name: "cjt-Type1MP-eType2R2-null", Optional: true},
		{Name: "cjt-Type1MP-feType2R1M1-null", Optional: true},
		{Name: "cjt-Type1MP-feType2R1M2-null", Optional: true},
		{Name: "cjt-Type1MP-feType2R2M2-null", Optional: true},
	},
}

var codebookComboParametersCJTR18CjtType1SPEType2R1NullConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersCJTR18CjtType1SPEType2R2NullConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersCJTR18CjtType1SPFeType2R1M1NullConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersCJTR18CjtType1SPFeType2R1M2NullConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersCJTR18CjtType1SPFeType2R2M2NullConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersCJTR18CjtType1MPEType2R1NullConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersCJTR18CjtType1MPEType2R2NullConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersCJTR18CjtType1MPFeType2R1M1NullConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersCJTR18CjtType1MPFeType2R1M2NullConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersCJTR18CjtType1MPFeType2R2M2NullConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

type CodebookComboParametersCJT_r18 struct {
	Cjt_Type1SP_EType2R1_Null    []int64
	Cjt_Type1SP_EType2R2_Null    []int64
	Cjt_Type1SP_FeType2R1M1_Null []int64
	Cjt_Type1SP_FeType2R1M2_Null []int64
	Cjt_Type1SP_FeType2R2M2_Null []int64
	Cjt_Type1MP_EType2R1_Null    []int64
	Cjt_Type1MP_EType2R2_Null    []int64
	Cjt_Type1MP_FeType2R1M1_Null []int64
	Cjt_Type1MP_FeType2R1M2_Null []int64
	Cjt_Type1MP_FeType2R2M2_Null []int64
}

func (ie *CodebookComboParametersCJT_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookComboParametersCJTR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Cjt_Type1SP_EType2R1_Null != nil, ie.Cjt_Type1SP_EType2R2_Null != nil, ie.Cjt_Type1SP_FeType2R1M1_Null != nil, ie.Cjt_Type1SP_FeType2R1M2_Null != nil, ie.Cjt_Type1SP_FeType2R2M2_Null != nil, ie.Cjt_Type1MP_EType2R1_Null != nil, ie.Cjt_Type1MP_EType2R2_Null != nil, ie.Cjt_Type1MP_FeType2R1M1_Null != nil, ie.Cjt_Type1MP_FeType2R1M2_Null != nil, ie.Cjt_Type1MP_FeType2R2M2_Null != nil}); err != nil {
		return err
	}

	// 2. cjt-Type1SP-eType2R1-null: sequence-of
	{
		if ie.Cjt_Type1SP_EType2R1_Null != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParametersCJTR18CjtType1SPEType2R1NullConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cjt_Type1SP_EType2R1_Null))); err != nil {
				return err
			}
			for i := range ie.Cjt_Type1SP_EType2R1_Null {
				if err := e.EncodeInteger(ie.Cjt_Type1SP_EType2R1_Null[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 3. cjt-Type1SP-eType2R2-null: sequence-of
	{
		if ie.Cjt_Type1SP_EType2R2_Null != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParametersCJTR18CjtType1SPEType2R2NullConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cjt_Type1SP_EType2R2_Null))); err != nil {
				return err
			}
			for i := range ie.Cjt_Type1SP_EType2R2_Null {
				if err := e.EncodeInteger(ie.Cjt_Type1SP_EType2R2_Null[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 4. cjt-Type1SP-feType2R1M1-null: sequence-of
	{
		if ie.Cjt_Type1SP_FeType2R1M1_Null != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParametersCJTR18CjtType1SPFeType2R1M1NullConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cjt_Type1SP_FeType2R1M1_Null))); err != nil {
				return err
			}
			for i := range ie.Cjt_Type1SP_FeType2R1M1_Null {
				if err := e.EncodeInteger(ie.Cjt_Type1SP_FeType2R1M1_Null[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 5. cjt-Type1SP-feType2R1M2-null: sequence-of
	{
		if ie.Cjt_Type1SP_FeType2R1M2_Null != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParametersCJTR18CjtType1SPFeType2R1M2NullConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cjt_Type1SP_FeType2R1M2_Null))); err != nil {
				return err
			}
			for i := range ie.Cjt_Type1SP_FeType2R1M2_Null {
				if err := e.EncodeInteger(ie.Cjt_Type1SP_FeType2R1M2_Null[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 6. cjt-Type1SP-feType2R2M2-null: sequence-of
	{
		if ie.Cjt_Type1SP_FeType2R2M2_Null != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParametersCJTR18CjtType1SPFeType2R2M2NullConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cjt_Type1SP_FeType2R2M2_Null))); err != nil {
				return err
			}
			for i := range ie.Cjt_Type1SP_FeType2R2M2_Null {
				if err := e.EncodeInteger(ie.Cjt_Type1SP_FeType2R2M2_Null[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 7. cjt-Type1MP-eType2R1-null: sequence-of
	{
		if ie.Cjt_Type1MP_EType2R1_Null != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParametersCJTR18CjtType1MPEType2R1NullConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cjt_Type1MP_EType2R1_Null))); err != nil {
				return err
			}
			for i := range ie.Cjt_Type1MP_EType2R1_Null {
				if err := e.EncodeInteger(ie.Cjt_Type1MP_EType2R1_Null[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 8. cjt-Type1MP-eType2R2-null: sequence-of
	{
		if ie.Cjt_Type1MP_EType2R2_Null != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParametersCJTR18CjtType1MPEType2R2NullConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cjt_Type1MP_EType2R2_Null))); err != nil {
				return err
			}
			for i := range ie.Cjt_Type1MP_EType2R2_Null {
				if err := e.EncodeInteger(ie.Cjt_Type1MP_EType2R2_Null[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 9. cjt-Type1MP-feType2R1M1-null: sequence-of
	{
		if ie.Cjt_Type1MP_FeType2R1M1_Null != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParametersCJTR18CjtType1MPFeType2R1M1NullConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cjt_Type1MP_FeType2R1M1_Null))); err != nil {
				return err
			}
			for i := range ie.Cjt_Type1MP_FeType2R1M1_Null {
				if err := e.EncodeInteger(ie.Cjt_Type1MP_FeType2R1M1_Null[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 10. cjt-Type1MP-feType2R1M2-null: sequence-of
	{
		if ie.Cjt_Type1MP_FeType2R1M2_Null != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParametersCJTR18CjtType1MPFeType2R1M2NullConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cjt_Type1MP_FeType2R1M2_Null))); err != nil {
				return err
			}
			for i := range ie.Cjt_Type1MP_FeType2R1M2_Null {
				if err := e.EncodeInteger(ie.Cjt_Type1MP_FeType2R1M2_Null[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 11. cjt-Type1MP-feType2R2M2-null: sequence-of
	{
		if ie.Cjt_Type1MP_FeType2R2M2_Null != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParametersCJTR18CjtType1MPFeType2R2M2NullConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cjt_Type1MP_FeType2R2M2_Null))); err != nil {
				return err
			}
			for i := range ie.Cjt_Type1MP_FeType2R2M2_Null {
				if err := e.EncodeInteger(ie.Cjt_Type1MP_FeType2R2M2_Null[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CodebookComboParametersCJT_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookComboParametersCJTR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. cjt-Type1SP-eType2R1-null: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParametersCJTR18CjtType1SPEType2R1NullConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cjt_Type1SP_EType2R1_Null = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Cjt_Type1SP_EType2R1_Null[i] = v
			}
		}
	}

	// 3. cjt-Type1SP-eType2R2-null: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParametersCJTR18CjtType1SPEType2R2NullConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cjt_Type1SP_EType2R2_Null = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Cjt_Type1SP_EType2R2_Null[i] = v
			}
		}
	}

	// 4. cjt-Type1SP-feType2R1M1-null: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParametersCJTR18CjtType1SPFeType2R1M1NullConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cjt_Type1SP_FeType2R1M1_Null = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Cjt_Type1SP_FeType2R1M1_Null[i] = v
			}
		}
	}

	// 5. cjt-Type1SP-feType2R1M2-null: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParametersCJTR18CjtType1SPFeType2R1M2NullConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cjt_Type1SP_FeType2R1M2_Null = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Cjt_Type1SP_FeType2R1M2_Null[i] = v
			}
		}
	}

	// 6. cjt-Type1SP-feType2R2M2-null: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParametersCJTR18CjtType1SPFeType2R2M2NullConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cjt_Type1SP_FeType2R2M2_Null = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Cjt_Type1SP_FeType2R2M2_Null[i] = v
			}
		}
	}

	// 7. cjt-Type1MP-eType2R1-null: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParametersCJTR18CjtType1MPEType2R1NullConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cjt_Type1MP_EType2R1_Null = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Cjt_Type1MP_EType2R1_Null[i] = v
			}
		}
	}

	// 8. cjt-Type1MP-eType2R2-null: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParametersCJTR18CjtType1MPEType2R2NullConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cjt_Type1MP_EType2R2_Null = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Cjt_Type1MP_EType2R2_Null[i] = v
			}
		}
	}

	// 9. cjt-Type1MP-feType2R1M1-null: sequence-of
	{
		if seq.IsComponentPresent(7) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParametersCJTR18CjtType1MPFeType2R1M1NullConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cjt_Type1MP_FeType2R1M1_Null = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Cjt_Type1MP_FeType2R1M1_Null[i] = v
			}
		}
	}

	// 10. cjt-Type1MP-feType2R1M2-null: sequence-of
	{
		if seq.IsComponentPresent(8) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParametersCJTR18CjtType1MPFeType2R1M2NullConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cjt_Type1MP_FeType2R1M2_Null = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Cjt_Type1MP_FeType2R1M2_Null[i] = v
			}
		}
	}

	// 11. cjt-Type1MP-feType2R2M2-null: sequence-of
	{
		if seq.IsComponentPresent(9) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParametersCJTR18CjtType1MPFeType2R2M2NullConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cjt_Type1MP_FeType2R2M2_Null = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Cjt_Type1MP_FeType2R2M2_Null[i] = v
			}
		}
	}

	return nil
}
