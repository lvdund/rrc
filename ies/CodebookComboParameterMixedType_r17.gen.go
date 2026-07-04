// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookComboParameterMixedType-r17 (line 18619).

var codebookComboParameterMixedTypeR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "type1SP-feType2PS-null-r17", Optional: true},
		{Name: "type1SP-feType2PS-M2R1-null-r17", Optional: true},
		{Name: "type1SP-feType2PS-M2R2-null-r17", Optional: true},
		{Name: "type1SP-Type2-feType2-PS-M1-r17", Optional: true},
		{Name: "type1SP-Type2-feType2-PS-M2R1-r17", Optional: true},
		{Name: "type1SP-eType2R1-feType2-PS-M1-r17", Optional: true},
		{Name: "type1SP-eType2R1-feType2-PS-M2R1-r17", Optional: true},
		{Name: "type1MP-feType2PS-null-r17", Optional: true},
		{Name: "type1MP-feType2PS-M2R1-null-r17", Optional: true},
		{Name: "type1MP-feType2PS-M2R2-null-r17", Optional: true},
		{Name: "type1MP-Type2-feType2-PS-M1-r17", Optional: true},
		{Name: "type1MP-Type2-feType2-PS-M2R1-r17", Optional: true},
		{Name: "type1MP-eType2R1-feType2-PS-M1-r17", Optional: true},
		{Name: "type1MP-eType2R1-feType2-PS-M2R1-r17", Optional: true},
	},
}

var codebookComboParameterMixedTypeR17Type1SPFeType2PSNullR17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMixedTypeR17Type1SPFeType2PSM2R1NullR17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMixedTypeR17Type1SPFeType2PSM2R2NullR17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMixedTypeR17Type1SPType2FeType2PSM1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMixedTypeR17Type1SPType2FeType2PSM2R1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMixedTypeR17Type1SPEType2R1FeType2PSM1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMixedTypeR17Type1SPEType2R1FeType2PSM2R1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMixedTypeR17Type1MPFeType2PSNullR17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMixedTypeR17Type1MPFeType2PSM2R1NullR17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMixedTypeR17Type1MPFeType2PSM2R2NullR17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMixedTypeR17Type1MPType2FeType2PSM1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMixedTypeR17Type1MPType2FeType2PSM2R1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMixedTypeR17Type1MPEType2R1FeType2PSM1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMixedTypeR17Type1MPEType2R1FeType2PSM2R1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

type CodebookComboParameterMixedType_r17 struct {
	Type1SP_FeType2PS_Null_r17           []int64
	Type1SP_FeType2PS_M2R1_Null_r17      []int64
	Type1SP_FeType2PS_M2R2_Null_r17      []int64
	Type1SP_Type2_FeType2_PS_M1_r17      []int64
	Type1SP_Type2_FeType2_PS_M2R1_r17    []int64
	Type1SP_EType2R1_FeType2_PS_M1_r17   []int64
	Type1SP_EType2R1_FeType2_PS_M2R1_r17 []int64
	Type1MP_FeType2PS_Null_r17           []int64
	Type1MP_FeType2PS_M2R1_Null_r17      []int64
	Type1MP_FeType2PS_M2R2_Null_r17      []int64
	Type1MP_Type2_FeType2_PS_M1_r17      []int64
	Type1MP_Type2_FeType2_PS_M2R1_r17    []int64
	Type1MP_EType2R1_FeType2_PS_M1_r17   []int64
	Type1MP_EType2R1_FeType2_PS_M2R1_r17 []int64
}

func (ie *CodebookComboParameterMixedType_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookComboParameterMixedTypeR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Type1SP_FeType2PS_Null_r17 != nil, ie.Type1SP_FeType2PS_M2R1_Null_r17 != nil, ie.Type1SP_FeType2PS_M2R2_Null_r17 != nil, ie.Type1SP_Type2_FeType2_PS_M1_r17 != nil, ie.Type1SP_Type2_FeType2_PS_M2R1_r17 != nil, ie.Type1SP_EType2R1_FeType2_PS_M1_r17 != nil, ie.Type1SP_EType2R1_FeType2_PS_M2R1_r17 != nil, ie.Type1MP_FeType2PS_Null_r17 != nil, ie.Type1MP_FeType2PS_M2R1_Null_r17 != nil, ie.Type1MP_FeType2PS_M2R2_Null_r17 != nil, ie.Type1MP_Type2_FeType2_PS_M1_r17 != nil, ie.Type1MP_Type2_FeType2_PS_M2R1_r17 != nil, ie.Type1MP_EType2R1_FeType2_PS_M1_r17 != nil, ie.Type1MP_EType2R1_FeType2_PS_M2R1_r17 != nil}); err != nil {
		return err
	}

	// 2. type1SP-feType2PS-null-r17: sequence-of
	{
		if ie.Type1SP_FeType2PS_Null_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMixedTypeR17Type1SPFeType2PSNullR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Type1SP_FeType2PS_Null_r17))); err != nil {
				return err
			}
			for i := range ie.Type1SP_FeType2PS_Null_r17 {
				if err := e.EncodeInteger(ie.Type1SP_FeType2PS_Null_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 3. type1SP-feType2PS-M2R1-null-r17: sequence-of
	{
		if ie.Type1SP_FeType2PS_M2R1_Null_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMixedTypeR17Type1SPFeType2PSM2R1NullR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Type1SP_FeType2PS_M2R1_Null_r17))); err != nil {
				return err
			}
			for i := range ie.Type1SP_FeType2PS_M2R1_Null_r17 {
				if err := e.EncodeInteger(ie.Type1SP_FeType2PS_M2R1_Null_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 4. type1SP-feType2PS-M2R2-null-r17: sequence-of
	{
		if ie.Type1SP_FeType2PS_M2R2_Null_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMixedTypeR17Type1SPFeType2PSM2R2NullR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Type1SP_FeType2PS_M2R2_Null_r17))); err != nil {
				return err
			}
			for i := range ie.Type1SP_FeType2PS_M2R2_Null_r17 {
				if err := e.EncodeInteger(ie.Type1SP_FeType2PS_M2R2_Null_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 5. type1SP-Type2-feType2-PS-M1-r17: sequence-of
	{
		if ie.Type1SP_Type2_FeType2_PS_M1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMixedTypeR17Type1SPType2FeType2PSM1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Type1SP_Type2_FeType2_PS_M1_r17))); err != nil {
				return err
			}
			for i := range ie.Type1SP_Type2_FeType2_PS_M1_r17 {
				if err := e.EncodeInteger(ie.Type1SP_Type2_FeType2_PS_M1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 6. type1SP-Type2-feType2-PS-M2R1-r17: sequence-of
	{
		if ie.Type1SP_Type2_FeType2_PS_M2R1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMixedTypeR17Type1SPType2FeType2PSM2R1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Type1SP_Type2_FeType2_PS_M2R1_r17))); err != nil {
				return err
			}
			for i := range ie.Type1SP_Type2_FeType2_PS_M2R1_r17 {
				if err := e.EncodeInteger(ie.Type1SP_Type2_FeType2_PS_M2R1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 7. type1SP-eType2R1-feType2-PS-M1-r17: sequence-of
	{
		if ie.Type1SP_EType2R1_FeType2_PS_M1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMixedTypeR17Type1SPEType2R1FeType2PSM1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Type1SP_EType2R1_FeType2_PS_M1_r17))); err != nil {
				return err
			}
			for i := range ie.Type1SP_EType2R1_FeType2_PS_M1_r17 {
				if err := e.EncodeInteger(ie.Type1SP_EType2R1_FeType2_PS_M1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 8. type1SP-eType2R1-feType2-PS-M2R1-r17: sequence-of
	{
		if ie.Type1SP_EType2R1_FeType2_PS_M2R1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMixedTypeR17Type1SPEType2R1FeType2PSM2R1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Type1SP_EType2R1_FeType2_PS_M2R1_r17))); err != nil {
				return err
			}
			for i := range ie.Type1SP_EType2R1_FeType2_PS_M2R1_r17 {
				if err := e.EncodeInteger(ie.Type1SP_EType2R1_FeType2_PS_M2R1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 9. type1MP-feType2PS-null-r17: sequence-of
	{
		if ie.Type1MP_FeType2PS_Null_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMixedTypeR17Type1MPFeType2PSNullR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Type1MP_FeType2PS_Null_r17))); err != nil {
				return err
			}
			for i := range ie.Type1MP_FeType2PS_Null_r17 {
				if err := e.EncodeInteger(ie.Type1MP_FeType2PS_Null_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 10. type1MP-feType2PS-M2R1-null-r17: sequence-of
	{
		if ie.Type1MP_FeType2PS_M2R1_Null_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMixedTypeR17Type1MPFeType2PSM2R1NullR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Type1MP_FeType2PS_M2R1_Null_r17))); err != nil {
				return err
			}
			for i := range ie.Type1MP_FeType2PS_M2R1_Null_r17 {
				if err := e.EncodeInteger(ie.Type1MP_FeType2PS_M2R1_Null_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 11. type1MP-feType2PS-M2R2-null-r17: sequence-of
	{
		if ie.Type1MP_FeType2PS_M2R2_Null_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMixedTypeR17Type1MPFeType2PSM2R2NullR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Type1MP_FeType2PS_M2R2_Null_r17))); err != nil {
				return err
			}
			for i := range ie.Type1MP_FeType2PS_M2R2_Null_r17 {
				if err := e.EncodeInteger(ie.Type1MP_FeType2PS_M2R2_Null_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 12. type1MP-Type2-feType2-PS-M1-r17: sequence-of
	{
		if ie.Type1MP_Type2_FeType2_PS_M1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMixedTypeR17Type1MPType2FeType2PSM1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Type1MP_Type2_FeType2_PS_M1_r17))); err != nil {
				return err
			}
			for i := range ie.Type1MP_Type2_FeType2_PS_M1_r17 {
				if err := e.EncodeInteger(ie.Type1MP_Type2_FeType2_PS_M1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 13. type1MP-Type2-feType2-PS-M2R1-r17: sequence-of
	{
		if ie.Type1MP_Type2_FeType2_PS_M2R1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMixedTypeR17Type1MPType2FeType2PSM2R1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Type1MP_Type2_FeType2_PS_M2R1_r17))); err != nil {
				return err
			}
			for i := range ie.Type1MP_Type2_FeType2_PS_M2R1_r17 {
				if err := e.EncodeInteger(ie.Type1MP_Type2_FeType2_PS_M2R1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 14. type1MP-eType2R1-feType2-PS-M1-r17: sequence-of
	{
		if ie.Type1MP_EType2R1_FeType2_PS_M1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMixedTypeR17Type1MPEType2R1FeType2PSM1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Type1MP_EType2R1_FeType2_PS_M1_r17))); err != nil {
				return err
			}
			for i := range ie.Type1MP_EType2R1_FeType2_PS_M1_r17 {
				if err := e.EncodeInteger(ie.Type1MP_EType2R1_FeType2_PS_M1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 15. type1MP-eType2R1-feType2-PS-M2R1-r17: sequence-of
	{
		if ie.Type1MP_EType2R1_FeType2_PS_M2R1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMixedTypeR17Type1MPEType2R1FeType2PSM2R1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Type1MP_EType2R1_FeType2_PS_M2R1_r17))); err != nil {
				return err
			}
			for i := range ie.Type1MP_EType2R1_FeType2_PS_M2R1_r17 {
				if err := e.EncodeInteger(ie.Type1MP_EType2R1_FeType2_PS_M2R1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CodebookComboParameterMixedType_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookComboParameterMixedTypeR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. type1SP-feType2PS-null-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMixedTypeR17Type1SPFeType2PSNullR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Type1SP_FeType2PS_Null_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Type1SP_FeType2PS_Null_r17[i] = v
			}
		}
	}

	// 3. type1SP-feType2PS-M2R1-null-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMixedTypeR17Type1SPFeType2PSM2R1NullR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Type1SP_FeType2PS_M2R1_Null_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Type1SP_FeType2PS_M2R1_Null_r17[i] = v
			}
		}
	}

	// 4. type1SP-feType2PS-M2R2-null-r17: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMixedTypeR17Type1SPFeType2PSM2R2NullR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Type1SP_FeType2PS_M2R2_Null_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Type1SP_FeType2PS_M2R2_Null_r17[i] = v
			}
		}
	}

	// 5. type1SP-Type2-feType2-PS-M1-r17: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMixedTypeR17Type1SPType2FeType2PSM1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Type1SP_Type2_FeType2_PS_M1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Type1SP_Type2_FeType2_PS_M1_r17[i] = v
			}
		}
	}

	// 6. type1SP-Type2-feType2-PS-M2R1-r17: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMixedTypeR17Type1SPType2FeType2PSM2R1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Type1SP_Type2_FeType2_PS_M2R1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Type1SP_Type2_FeType2_PS_M2R1_r17[i] = v
			}
		}
	}

	// 7. type1SP-eType2R1-feType2-PS-M1-r17: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMixedTypeR17Type1SPEType2R1FeType2PSM1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Type1SP_EType2R1_FeType2_PS_M1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Type1SP_EType2R1_FeType2_PS_M1_r17[i] = v
			}
		}
	}

	// 8. type1SP-eType2R1-feType2-PS-M2R1-r17: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMixedTypeR17Type1SPEType2R1FeType2PSM2R1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Type1SP_EType2R1_FeType2_PS_M2R1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Type1SP_EType2R1_FeType2_PS_M2R1_r17[i] = v
			}
		}
	}

	// 9. type1MP-feType2PS-null-r17: sequence-of
	{
		if seq.IsComponentPresent(7) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMixedTypeR17Type1MPFeType2PSNullR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Type1MP_FeType2PS_Null_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Type1MP_FeType2PS_Null_r17[i] = v
			}
		}
	}

	// 10. type1MP-feType2PS-M2R1-null-r17: sequence-of
	{
		if seq.IsComponentPresent(8) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMixedTypeR17Type1MPFeType2PSM2R1NullR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Type1MP_FeType2PS_M2R1_Null_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Type1MP_FeType2PS_M2R1_Null_r17[i] = v
			}
		}
	}

	// 11. type1MP-feType2PS-M2R2-null-r17: sequence-of
	{
		if seq.IsComponentPresent(9) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMixedTypeR17Type1MPFeType2PSM2R2NullR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Type1MP_FeType2PS_M2R2_Null_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Type1MP_FeType2PS_M2R2_Null_r17[i] = v
			}
		}
	}

	// 12. type1MP-Type2-feType2-PS-M1-r17: sequence-of
	{
		if seq.IsComponentPresent(10) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMixedTypeR17Type1MPType2FeType2PSM1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Type1MP_Type2_FeType2_PS_M1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Type1MP_Type2_FeType2_PS_M1_r17[i] = v
			}
		}
	}

	// 13. type1MP-Type2-feType2-PS-M2R1-r17: sequence-of
	{
		if seq.IsComponentPresent(11) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMixedTypeR17Type1MPType2FeType2PSM2R1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Type1MP_Type2_FeType2_PS_M2R1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Type1MP_Type2_FeType2_PS_M2R1_r17[i] = v
			}
		}
	}

	// 14. type1MP-eType2R1-feType2-PS-M1-r17: sequence-of
	{
		if seq.IsComponentPresent(12) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMixedTypeR17Type1MPEType2R1FeType2PSM1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Type1MP_EType2R1_FeType2_PS_M1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Type1MP_EType2R1_FeType2_PS_M1_r17[i] = v
			}
		}
	}

	// 15. type1MP-eType2R1-feType2-PS-M2R1-r17: sequence-of
	{
		if seq.IsComponentPresent(13) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMixedTypeR17Type1MPEType2R1FeType2PSM2R1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Type1MP_EType2R1_FeType2_PS_M2R1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.Type1MP_EType2R1_FeType2_PS_M2R1_r17[i] = v
			}
		}
	}

	return nil
}
