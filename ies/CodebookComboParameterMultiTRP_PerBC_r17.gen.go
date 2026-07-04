// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookComboParameterMultiTRP-PerBC-r17 (line 18808).

var codebookComboParameterMultiTRPPerBCR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nCJT-null-null", Optional: true},
		{Name: "nCJT1SP-null-null", Optional: true},
		{Name: "nCJT-Type2-null-r16", Optional: true},
		{Name: "nCJT-Type2PS-null-r16", Optional: true},
		{Name: "nCJT-eType2R1-null-r16", Optional: true},
		{Name: "nCJT-eType2R2-null-r16", Optional: true},
		{Name: "nCJT-eType2R1PS-null-r16", Optional: true},
		{Name: "nCJT-eType2R2PS-null-r16", Optional: true},
		{Name: "nCJT-Type2-Type2PS-r16", Optional: true},
		{Name: "nCJT1SP-Type2-null-r16", Optional: true},
		{Name: "nCJT1SP-Type2PS-null-r16", Optional: true},
		{Name: "nCJT1SP-eType2R1-null-r16", Optional: true},
		{Name: "nCJT1SP-eType2R2-null-r16", Optional: true},
		{Name: "nCJT1SP-eType2R1PS-null-r16", Optional: true},
		{Name: "nCJT1SP-eType2R2PS-null-r16", Optional: true},
		{Name: "nCJT1SP-Type2-Type2PS-r16", Optional: true},
		{Name: "nCJT-feType2PS-null-r17", Optional: true},
		{Name: "nCJT-feType2PS-M2R1-null-r17", Optional: true},
		{Name: "nCJT-feType2PS-M2R2-null-r17", Optional: true},
		{Name: "nCJT-Type2-feType2-PS-M1-r17", Optional: true},
		{Name: "nCJT-Type2-feType2-PS-M2R1-r17", Optional: true},
		{Name: "nCJT-eType2R1-feType2-PS-M1-r17", Optional: true},
		{Name: "nCJT-eType2R1-feType2-PS-M2R1-r17", Optional: true},
		{Name: "nCJT1SP-feType2PS-null-r17", Optional: true},
		{Name: "nCJT1SP-feType2PS-M2R1-null-r17", Optional: true},
		{Name: "nCJT1SP-feType2PS-M2R2-null-r17", Optional: true},
		{Name: "nCJT1SP-Type2-feType2-PS-M1-r17", Optional: true},
		{Name: "nCJT1SP-Type2-feType2-PS-M2R1-r17", Optional: true},
		{Name: "nCJT1SP-eType2R1-feType2-PS-M1-r17", Optional: true},
		{Name: "nCJT1SP-eType2R1-feType2-PS-M2R1-r17", Optional: true},
	},
}

var codebookComboParameterMultiTRPPerBCR17NCJTNullNullConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPNullNullConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJTType2NullR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJTType2PSNullR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJTEType2R1NullR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJTEType2R2NullR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJTEType2R1PSNullR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJTEType2R2PSNullR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJTType2Type2PSR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPType2NullR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPType2PSNullR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R1NullR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R2NullR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R1PSNullR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R2PSNullR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPType2Type2PSR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJTFeType2PSNullR17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJTFeType2PSM2R1NullR17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJTFeType2PSM2R2NullR17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJTType2FeType2PSM1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJTType2FeType2PSM2R1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJTEType2R1FeType2PSM1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJTEType2R1FeType2PSM2R1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPFeType2PSNullR17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPFeType2PSM2R1NullR17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPFeType2PSM2R2NullR17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPType2FeType2PSM1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPType2FeType2PSM2R1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R1FeType2PSM1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R1FeType2PSM2R1R17Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

type CodebookComboParameterMultiTRP_PerBC_r17 struct {
	NCJT_Null_Null                       []int64
	NCJT1SP_Null_Null                    []int64
	NCJT_Type2_Null_r16                  []int64
	NCJT_Type2PS_Null_r16                []int64
	NCJT_EType2R1_Null_r16               []int64
	NCJT_EType2R2_Null_r16               []int64
	NCJT_EType2R1PS_Null_r16             []int64
	NCJT_EType2R2PS_Null_r16             []int64
	NCJT_Type2_Type2PS_r16               []int64
	NCJT1SP_Type2_Null_r16               []int64
	NCJT1SP_Type2PS_Null_r16             []int64
	NCJT1SP_EType2R1_Null_r16            []int64
	NCJT1SP_EType2R2_Null_r16            []int64
	NCJT1SP_EType2R1PS_Null_r16          []int64
	NCJT1SP_EType2R2PS_Null_r16          []int64
	NCJT1SP_Type2_Type2PS_r16            []int64
	NCJT_FeType2PS_Null_r17              []int64
	NCJT_FeType2PS_M2R1_Null_r17         []int64
	NCJT_FeType2PS_M2R2_Null_r17         []int64
	NCJT_Type2_FeType2_PS_M1_r17         []int64
	NCJT_Type2_FeType2_PS_M2R1_r17       []int64
	NCJT_EType2R1_FeType2_PS_M1_r17      []int64
	NCJT_EType2R1_FeType2_PS_M2R1_r17    []int64
	NCJT1SP_FeType2PS_Null_r17           []int64
	NCJT1SP_FeType2PS_M2R1_Null_r17      []int64
	NCJT1SP_FeType2PS_M2R2_Null_r17      []int64
	NCJT1SP_Type2_FeType2_PS_M1_r17      []int64
	NCJT1SP_Type2_FeType2_PS_M2R1_r17    []int64
	NCJT1SP_EType2R1_FeType2_PS_M1_r17   []int64
	NCJT1SP_EType2R1_FeType2_PS_M2R1_r17 []int64
}

func (ie *CodebookComboParameterMultiTRP_PerBC_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookComboParameterMultiTRPPerBCR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.NCJT_Null_Null != nil, ie.NCJT1SP_Null_Null != nil, ie.NCJT_Type2_Null_r16 != nil, ie.NCJT_Type2PS_Null_r16 != nil, ie.NCJT_EType2R1_Null_r16 != nil, ie.NCJT_EType2R2_Null_r16 != nil, ie.NCJT_EType2R1PS_Null_r16 != nil, ie.NCJT_EType2R2PS_Null_r16 != nil, ie.NCJT_Type2_Type2PS_r16 != nil, ie.NCJT1SP_Type2_Null_r16 != nil, ie.NCJT1SP_Type2PS_Null_r16 != nil, ie.NCJT1SP_EType2R1_Null_r16 != nil, ie.NCJT1SP_EType2R2_Null_r16 != nil, ie.NCJT1SP_EType2R1PS_Null_r16 != nil, ie.NCJT1SP_EType2R2PS_Null_r16 != nil, ie.NCJT1SP_Type2_Type2PS_r16 != nil, ie.NCJT_FeType2PS_Null_r17 != nil, ie.NCJT_FeType2PS_M2R1_Null_r17 != nil, ie.NCJT_FeType2PS_M2R2_Null_r17 != nil, ie.NCJT_Type2_FeType2_PS_M1_r17 != nil, ie.NCJT_Type2_FeType2_PS_M2R1_r17 != nil, ie.NCJT_EType2R1_FeType2_PS_M1_r17 != nil, ie.NCJT_EType2R1_FeType2_PS_M2R1_r17 != nil, ie.NCJT1SP_FeType2PS_Null_r17 != nil, ie.NCJT1SP_FeType2PS_M2R1_Null_r17 != nil, ie.NCJT1SP_FeType2PS_M2R2_Null_r17 != nil, ie.NCJT1SP_Type2_FeType2_PS_M1_r17 != nil, ie.NCJT1SP_Type2_FeType2_PS_M2R1_r17 != nil, ie.NCJT1SP_EType2R1_FeType2_PS_M1_r17 != nil, ie.NCJT1SP_EType2R1_FeType2_PS_M2R1_r17 != nil}); err != nil {
		return err
	}

	// 2. nCJT-null-null: sequence-of
	{
		if ie.NCJT_Null_Null != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTNullNullConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_Null_Null))); err != nil {
				return err
			}
			for i := range ie.NCJT_Null_Null {
				if err := e.EncodeInteger(ie.NCJT_Null_Null[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 3. nCJT1SP-null-null: sequence-of
	{
		if ie.NCJT1SP_Null_Null != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPNullNullConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_Null_Null))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_Null_Null {
				if err := e.EncodeInteger(ie.NCJT1SP_Null_Null[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 4. nCJT-Type2-null-r16: sequence-of
	{
		if ie.NCJT_Type2_Null_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTType2NullR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_Type2_Null_r16))); err != nil {
				return err
			}
			for i := range ie.NCJT_Type2_Null_r16 {
				if err := e.EncodeInteger(ie.NCJT_Type2_Null_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 5. nCJT-Type2PS-null-r16: sequence-of
	{
		if ie.NCJT_Type2PS_Null_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTType2PSNullR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_Type2PS_Null_r16))); err != nil {
				return err
			}
			for i := range ie.NCJT_Type2PS_Null_r16 {
				if err := e.EncodeInteger(ie.NCJT_Type2PS_Null_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 6. nCJT-eType2R1-null-r16: sequence-of
	{
		if ie.NCJT_EType2R1_Null_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTEType2R1NullR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_EType2R1_Null_r16))); err != nil {
				return err
			}
			for i := range ie.NCJT_EType2R1_Null_r16 {
				if err := e.EncodeInteger(ie.NCJT_EType2R1_Null_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 7. nCJT-eType2R2-null-r16: sequence-of
	{
		if ie.NCJT_EType2R2_Null_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTEType2R2NullR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_EType2R2_Null_r16))); err != nil {
				return err
			}
			for i := range ie.NCJT_EType2R2_Null_r16 {
				if err := e.EncodeInteger(ie.NCJT_EType2R2_Null_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 8. nCJT-eType2R1PS-null-r16: sequence-of
	{
		if ie.NCJT_EType2R1PS_Null_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTEType2R1PSNullR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_EType2R1PS_Null_r16))); err != nil {
				return err
			}
			for i := range ie.NCJT_EType2R1PS_Null_r16 {
				if err := e.EncodeInteger(ie.NCJT_EType2R1PS_Null_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 9. nCJT-eType2R2PS-null-r16: sequence-of
	{
		if ie.NCJT_EType2R2PS_Null_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTEType2R2PSNullR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_EType2R2PS_Null_r16))); err != nil {
				return err
			}
			for i := range ie.NCJT_EType2R2PS_Null_r16 {
				if err := e.EncodeInteger(ie.NCJT_EType2R2PS_Null_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 10. nCJT-Type2-Type2PS-r16: sequence-of
	{
		if ie.NCJT_Type2_Type2PS_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTType2Type2PSR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_Type2_Type2PS_r16))); err != nil {
				return err
			}
			for i := range ie.NCJT_Type2_Type2PS_r16 {
				if err := e.EncodeInteger(ie.NCJT_Type2_Type2PS_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 11. nCJT1SP-Type2-null-r16: sequence-of
	{
		if ie.NCJT1SP_Type2_Null_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPType2NullR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_Type2_Null_r16))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_Type2_Null_r16 {
				if err := e.EncodeInteger(ie.NCJT1SP_Type2_Null_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 12. nCJT1SP-Type2PS-null-r16: sequence-of
	{
		if ie.NCJT1SP_Type2PS_Null_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPType2PSNullR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_Type2PS_Null_r16))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_Type2PS_Null_r16 {
				if err := e.EncodeInteger(ie.NCJT1SP_Type2PS_Null_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 13. nCJT1SP-eType2R1-null-r16: sequence-of
	{
		if ie.NCJT1SP_EType2R1_Null_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R1NullR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_EType2R1_Null_r16))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_EType2R1_Null_r16 {
				if err := e.EncodeInteger(ie.NCJT1SP_EType2R1_Null_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 14. nCJT1SP-eType2R2-null-r16: sequence-of
	{
		if ie.NCJT1SP_EType2R2_Null_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R2NullR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_EType2R2_Null_r16))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_EType2R2_Null_r16 {
				if err := e.EncodeInteger(ie.NCJT1SP_EType2R2_Null_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 15. nCJT1SP-eType2R1PS-null-r16: sequence-of
	{
		if ie.NCJT1SP_EType2R1PS_Null_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R1PSNullR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_EType2R1PS_Null_r16))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_EType2R1PS_Null_r16 {
				if err := e.EncodeInteger(ie.NCJT1SP_EType2R1PS_Null_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 16. nCJT1SP-eType2R2PS-null-r16: sequence-of
	{
		if ie.NCJT1SP_EType2R2PS_Null_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R2PSNullR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_EType2R2PS_Null_r16))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_EType2R2PS_Null_r16 {
				if err := e.EncodeInteger(ie.NCJT1SP_EType2R2PS_Null_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 17. nCJT1SP-Type2-Type2PS-r16: sequence-of
	{
		if ie.NCJT1SP_Type2_Type2PS_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPType2Type2PSR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_Type2_Type2PS_r16))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_Type2_Type2PS_r16 {
				if err := e.EncodeInteger(ie.NCJT1SP_Type2_Type2PS_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 18. nCJT-feType2PS-null-r17: sequence-of
	{
		if ie.NCJT_FeType2PS_Null_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTFeType2PSNullR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_FeType2PS_Null_r17))); err != nil {
				return err
			}
			for i := range ie.NCJT_FeType2PS_Null_r17 {
				if err := e.EncodeInteger(ie.NCJT_FeType2PS_Null_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 19. nCJT-feType2PS-M2R1-null-r17: sequence-of
	{
		if ie.NCJT_FeType2PS_M2R1_Null_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTFeType2PSM2R1NullR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_FeType2PS_M2R1_Null_r17))); err != nil {
				return err
			}
			for i := range ie.NCJT_FeType2PS_M2R1_Null_r17 {
				if err := e.EncodeInteger(ie.NCJT_FeType2PS_M2R1_Null_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 20. nCJT-feType2PS-M2R2-null-r17: sequence-of
	{
		if ie.NCJT_FeType2PS_M2R2_Null_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTFeType2PSM2R2NullR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_FeType2PS_M2R2_Null_r17))); err != nil {
				return err
			}
			for i := range ie.NCJT_FeType2PS_M2R2_Null_r17 {
				if err := e.EncodeInteger(ie.NCJT_FeType2PS_M2R2_Null_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 21. nCJT-Type2-feType2-PS-M1-r17: sequence-of
	{
		if ie.NCJT_Type2_FeType2_PS_M1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTType2FeType2PSM1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_Type2_FeType2_PS_M1_r17))); err != nil {
				return err
			}
			for i := range ie.NCJT_Type2_FeType2_PS_M1_r17 {
				if err := e.EncodeInteger(ie.NCJT_Type2_FeType2_PS_M1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 22. nCJT-Type2-feType2-PS-M2R1-r17: sequence-of
	{
		if ie.NCJT_Type2_FeType2_PS_M2R1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTType2FeType2PSM2R1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_Type2_FeType2_PS_M2R1_r17))); err != nil {
				return err
			}
			for i := range ie.NCJT_Type2_FeType2_PS_M2R1_r17 {
				if err := e.EncodeInteger(ie.NCJT_Type2_FeType2_PS_M2R1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 23. nCJT-eType2R1-feType2-PS-M1-r17: sequence-of
	{
		if ie.NCJT_EType2R1_FeType2_PS_M1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTEType2R1FeType2PSM1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_EType2R1_FeType2_PS_M1_r17))); err != nil {
				return err
			}
			for i := range ie.NCJT_EType2R1_FeType2_PS_M1_r17 {
				if err := e.EncodeInteger(ie.NCJT_EType2R1_FeType2_PS_M1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 24. nCJT-eType2R1-feType2-PS-M2R1-r17: sequence-of
	{
		if ie.NCJT_EType2R1_FeType2_PS_M2R1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJTEType2R1FeType2PSM2R1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT_EType2R1_FeType2_PS_M2R1_r17))); err != nil {
				return err
			}
			for i := range ie.NCJT_EType2R1_FeType2_PS_M2R1_r17 {
				if err := e.EncodeInteger(ie.NCJT_EType2R1_FeType2_PS_M2R1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 25. nCJT1SP-feType2PS-null-r17: sequence-of
	{
		if ie.NCJT1SP_FeType2PS_Null_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPFeType2PSNullR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_FeType2PS_Null_r17))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_FeType2PS_Null_r17 {
				if err := e.EncodeInteger(ie.NCJT1SP_FeType2PS_Null_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 26. nCJT1SP-feType2PS-M2R1-null-r17: sequence-of
	{
		if ie.NCJT1SP_FeType2PS_M2R1_Null_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPFeType2PSM2R1NullR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_FeType2PS_M2R1_Null_r17))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_FeType2PS_M2R1_Null_r17 {
				if err := e.EncodeInteger(ie.NCJT1SP_FeType2PS_M2R1_Null_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 27. nCJT1SP-feType2PS-M2R2-null-r17: sequence-of
	{
		if ie.NCJT1SP_FeType2PS_M2R2_Null_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPFeType2PSM2R2NullR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_FeType2PS_M2R2_Null_r17))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_FeType2PS_M2R2_Null_r17 {
				if err := e.EncodeInteger(ie.NCJT1SP_FeType2PS_M2R2_Null_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 28. nCJT1SP-Type2-feType2-PS-M1-r17: sequence-of
	{
		if ie.NCJT1SP_Type2_FeType2_PS_M1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPType2FeType2PSM1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_Type2_FeType2_PS_M1_r17))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_Type2_FeType2_PS_M1_r17 {
				if err := e.EncodeInteger(ie.NCJT1SP_Type2_FeType2_PS_M1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 29. nCJT1SP-Type2-feType2-PS-M2R1-r17: sequence-of
	{
		if ie.NCJT1SP_Type2_FeType2_PS_M2R1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPType2FeType2PSM2R1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_Type2_FeType2_PS_M2R1_r17))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_Type2_FeType2_PS_M2R1_r17 {
				if err := e.EncodeInteger(ie.NCJT1SP_Type2_FeType2_PS_M2R1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 30. nCJT1SP-eType2R1-feType2-PS-M1-r17: sequence-of
	{
		if ie.NCJT1SP_EType2R1_FeType2_PS_M1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R1FeType2PSM1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_EType2R1_FeType2_PS_M1_r17))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_EType2R1_FeType2_PS_M1_r17 {
				if err := e.EncodeInteger(ie.NCJT1SP_EType2R1_FeType2_PS_M1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 31. nCJT1SP-eType2R1-feType2-PS-M2R1-r17: sequence-of
	{
		if ie.NCJT1SP_EType2R1_FeType2_PS_M2R1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R1FeType2PSM2R1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NCJT1SP_EType2R1_FeType2_PS_M2R1_r17))); err != nil {
				return err
			}
			for i := range ie.NCJT1SP_EType2R1_FeType2_PS_M2R1_r17 {
				if err := e.EncodeInteger(ie.NCJT1SP_EType2R1_FeType2_PS_M2R1_r17[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CodebookComboParameterMultiTRP_PerBC_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookComboParameterMultiTRPPerBCR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. nCJT-null-null: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTNullNullConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_Null_Null = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_Null_Null[i] = v
			}
		}
	}

	// 3. nCJT1SP-null-null: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPNullNullConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_Null_Null = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_Null_Null[i] = v
			}
		}
	}

	// 4. nCJT-Type2-null-r16: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTType2NullR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_Type2_Null_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_Type2_Null_r16[i] = v
			}
		}
	}

	// 5. nCJT-Type2PS-null-r16: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTType2PSNullR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_Type2PS_Null_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_Type2PS_Null_r16[i] = v
			}
		}
	}

	// 6. nCJT-eType2R1-null-r16: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTEType2R1NullR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_EType2R1_Null_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_EType2R1_Null_r16[i] = v
			}
		}
	}

	// 7. nCJT-eType2R2-null-r16: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTEType2R2NullR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_EType2R2_Null_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_EType2R2_Null_r16[i] = v
			}
		}
	}

	// 8. nCJT-eType2R1PS-null-r16: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTEType2R1PSNullR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_EType2R1PS_Null_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_EType2R1PS_Null_r16[i] = v
			}
		}
	}

	// 9. nCJT-eType2R2PS-null-r16: sequence-of
	{
		if seq.IsComponentPresent(7) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTEType2R2PSNullR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_EType2R2PS_Null_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_EType2R2PS_Null_r16[i] = v
			}
		}
	}

	// 10. nCJT-Type2-Type2PS-r16: sequence-of
	{
		if seq.IsComponentPresent(8) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTType2Type2PSR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_Type2_Type2PS_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_Type2_Type2PS_r16[i] = v
			}
		}
	}

	// 11. nCJT1SP-Type2-null-r16: sequence-of
	{
		if seq.IsComponentPresent(9) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPType2NullR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_Type2_Null_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_Type2_Null_r16[i] = v
			}
		}
	}

	// 12. nCJT1SP-Type2PS-null-r16: sequence-of
	{
		if seq.IsComponentPresent(10) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPType2PSNullR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_Type2PS_Null_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_Type2PS_Null_r16[i] = v
			}
		}
	}

	// 13. nCJT1SP-eType2R1-null-r16: sequence-of
	{
		if seq.IsComponentPresent(11) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R1NullR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_EType2R1_Null_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_EType2R1_Null_r16[i] = v
			}
		}
	}

	// 14. nCJT1SP-eType2R2-null-r16: sequence-of
	{
		if seq.IsComponentPresent(12) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R2NullR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_EType2R2_Null_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_EType2R2_Null_r16[i] = v
			}
		}
	}

	// 15. nCJT1SP-eType2R1PS-null-r16: sequence-of
	{
		if seq.IsComponentPresent(13) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R1PSNullR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_EType2R1PS_Null_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_EType2R1PS_Null_r16[i] = v
			}
		}
	}

	// 16. nCJT1SP-eType2R2PS-null-r16: sequence-of
	{
		if seq.IsComponentPresent(14) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R2PSNullR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_EType2R2PS_Null_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_EType2R2PS_Null_r16[i] = v
			}
		}
	}

	// 17. nCJT1SP-Type2-Type2PS-r16: sequence-of
	{
		if seq.IsComponentPresent(15) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPType2Type2PSR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_Type2_Type2PS_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_Type2_Type2PS_r16[i] = v
			}
		}
	}

	// 18. nCJT-feType2PS-null-r17: sequence-of
	{
		if seq.IsComponentPresent(16) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTFeType2PSNullR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_FeType2PS_Null_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_FeType2PS_Null_r17[i] = v
			}
		}
	}

	// 19. nCJT-feType2PS-M2R1-null-r17: sequence-of
	{
		if seq.IsComponentPresent(17) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTFeType2PSM2R1NullR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_FeType2PS_M2R1_Null_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_FeType2PS_M2R1_Null_r17[i] = v
			}
		}
	}

	// 20. nCJT-feType2PS-M2R2-null-r17: sequence-of
	{
		if seq.IsComponentPresent(18) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTFeType2PSM2R2NullR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_FeType2PS_M2R2_Null_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_FeType2PS_M2R2_Null_r17[i] = v
			}
		}
	}

	// 21. nCJT-Type2-feType2-PS-M1-r17: sequence-of
	{
		if seq.IsComponentPresent(19) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTType2FeType2PSM1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_Type2_FeType2_PS_M1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_Type2_FeType2_PS_M1_r17[i] = v
			}
		}
	}

	// 22. nCJT-Type2-feType2-PS-M2R1-r17: sequence-of
	{
		if seq.IsComponentPresent(20) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTType2FeType2PSM2R1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_Type2_FeType2_PS_M2R1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_Type2_FeType2_PS_M2R1_r17[i] = v
			}
		}
	}

	// 23. nCJT-eType2R1-feType2-PS-M1-r17: sequence-of
	{
		if seq.IsComponentPresent(21) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTEType2R1FeType2PSM1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_EType2R1_FeType2_PS_M1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_EType2R1_FeType2_PS_M1_r17[i] = v
			}
		}
	}

	// 24. nCJT-eType2R1-feType2-PS-M2R1-r17: sequence-of
	{
		if seq.IsComponentPresent(22) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJTEType2R1FeType2PSM2R1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT_EType2R1_FeType2_PS_M2R1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT_EType2R1_FeType2_PS_M2R1_r17[i] = v
			}
		}
	}

	// 25. nCJT1SP-feType2PS-null-r17: sequence-of
	{
		if seq.IsComponentPresent(23) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPFeType2PSNullR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_FeType2PS_Null_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_FeType2PS_Null_r17[i] = v
			}
		}
	}

	// 26. nCJT1SP-feType2PS-M2R1-null-r17: sequence-of
	{
		if seq.IsComponentPresent(24) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPFeType2PSM2R1NullR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_FeType2PS_M2R1_Null_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_FeType2PS_M2R1_Null_r17[i] = v
			}
		}
	}

	// 27. nCJT1SP-feType2PS-M2R2-null-r17: sequence-of
	{
		if seq.IsComponentPresent(25) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPFeType2PSM2R2NullR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_FeType2PS_M2R2_Null_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_FeType2PS_M2R2_Null_r17[i] = v
			}
		}
	}

	// 28. nCJT1SP-Type2-feType2-PS-M1-r17: sequence-of
	{
		if seq.IsComponentPresent(26) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPType2FeType2PSM1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_Type2_FeType2_PS_M1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_Type2_FeType2_PS_M1_r17[i] = v
			}
		}
	}

	// 29. nCJT1SP-Type2-feType2-PS-M2R1-r17: sequence-of
	{
		if seq.IsComponentPresent(27) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPType2FeType2PSM2R1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_Type2_FeType2_PS_M2R1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_Type2_FeType2_PS_M2R1_r17[i] = v
			}
		}
	}

	// 30. nCJT1SP-eType2R1-feType2-PS-M1-r17: sequence-of
	{
		if seq.IsComponentPresent(28) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R1FeType2PSM1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_EType2R1_FeType2_PS_M1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_EType2R1_FeType2_PS_M1_r17[i] = v
			}
		}
	}

	// 31. nCJT1SP-eType2R1-feType2-PS-M2R1-r17: sequence-of
	{
		if seq.IsComponentPresent(29) {
			seqOf := d.NewSequenceOfDecoder(codebookComboParameterMultiTRPPerBCR17NCJT1SPEType2R1FeType2PSM2R1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NCJT1SP_EType2R1_FeType2_PS_M2R1_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.NCJT1SP_EType2R1_FeType2_PS_M2R1_r17[i] = v
			}
		}
	}

	return nil
}
