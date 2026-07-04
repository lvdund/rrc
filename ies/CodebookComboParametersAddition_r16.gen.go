// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookComboParametersAddition-r16 (line 18560).

var codebookComboParametersAdditionR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "type1SP-Type2-null-r16", Optional: true},
		{Name: "type1SP-Type2PS-null-r16", Optional: true},
		{Name: "type1SP-eType2R1-null-r16", Optional: true},
		{Name: "type1SP-eType2R2-null-r16", Optional: true},
		{Name: "type1SP-eType2R1PS-null-r16", Optional: true},
		{Name: "type1SP-eType2R2PS-null-r16", Optional: true},
		{Name: "type1SP-Type2-Type2PS-r16", Optional: true},
		{Name: "type1MP-Type2-null-r16", Optional: true},
		{Name: "type1MP-Type2PS-null-r16", Optional: true},
		{Name: "type1MP-eType2R1-null-r16", Optional: true},
		{Name: "type1MP-eType2R2-null-r16", Optional: true},
		{Name: "type1MP-eType2R1PS-null-r16", Optional: true},
		{Name: "type1MP-eType2R2PS-null-r16", Optional: true},
		{Name: "type1MP-Type2-Type2PS-r16", Optional: true},
	},
}

var codebookComboParametersAdditionR16Type1SPType2NullR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersAdditionR16Type1SPType2PSNullR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersAdditionR16Type1SPEType2R1NullR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersAdditionR16Type1SPEType2R2NullR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersAdditionR16Type1SPEType2R1PSNullR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersAdditionR16Type1SPEType2R2PSNullR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersAdditionR16Type1SPType2Type2PSR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersAdditionR16Type1MPType2NullR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersAdditionR16Type1MPType2PSNullR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersAdditionR16Type1MPEType2R1NullR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersAdditionR16Type1MPEType2R2NullR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersAdditionR16Type1MPEType2R1PSNullR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersAdditionR16Type1MPEType2R2PSNullR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookComboParametersAdditionR16Type1MPType2Type2PSR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

type CodebookComboParametersAddition_r16 struct {
	Type1SP_Type2_Null_r16      *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
	Type1SP_Type2PS_Null_r16    *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
	Type1SP_EType2R1_Null_r16   *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
	Type1SP_EType2R2_Null_r16   *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
	Type1SP_EType2R1PS_Null_r16 *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
	Type1SP_EType2R2PS_Null_r16 *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
	Type1SP_Type2_Type2PS_r16   *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
	Type1MP_Type2_Null_r16      *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
	Type1MP_Type2PS_Null_r16    *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
	Type1MP_EType2R1_Null_r16   *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
	Type1MP_EType2R2_Null_r16   *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
	Type1MP_EType2R1PS_Null_r16 *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
	Type1MP_EType2R2PS_Null_r16 *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
	Type1MP_Type2_Type2PS_r16   *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
}

func (ie *CodebookComboParametersAddition_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookComboParametersAdditionR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Type1SP_Type2_Null_r16 != nil, ie.Type1SP_Type2PS_Null_r16 != nil, ie.Type1SP_EType2R1_Null_r16 != nil, ie.Type1SP_EType2R2_Null_r16 != nil, ie.Type1SP_EType2R1PS_Null_r16 != nil, ie.Type1SP_EType2R2PS_Null_r16 != nil, ie.Type1SP_Type2_Type2PS_r16 != nil, ie.Type1MP_Type2_Null_r16 != nil, ie.Type1MP_Type2PS_Null_r16 != nil, ie.Type1MP_EType2R1_Null_r16 != nil, ie.Type1MP_EType2R2_Null_r16 != nil, ie.Type1MP_EType2R1PS_Null_r16 != nil, ie.Type1MP_EType2R2PS_Null_r16 != nil, ie.Type1MP_Type2_Type2PS_r16 != nil}); err != nil {
		return err
	}

	// 2. type1SP-Type2-null-r16: sequence
	{
		if ie.Type1SP_Type2_Null_r16 != nil {
			c := ie.Type1SP_Type2_Null_r16
			{
				seqOf := e.NewSequenceOfEncoder(codebookComboParametersAdditionR16Type1SPType2NullR16SupportedCSIRSResourceListAddR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListAdd_r16))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListAdd_r16 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListAdd_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 3. type1SP-Type2PS-null-r16: sequence
	{
		if ie.Type1SP_Type2PS_Null_r16 != nil {
			c := ie.Type1SP_Type2PS_Null_r16
			{
				seqOf := e.NewSequenceOfEncoder(codebookComboParametersAdditionR16Type1SPType2PSNullR16SupportedCSIRSResourceListAddR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListAdd_r16))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListAdd_r16 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListAdd_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 4. type1SP-eType2R1-null-r16: sequence
	{
		if ie.Type1SP_EType2R1_Null_r16 != nil {
			c := ie.Type1SP_EType2R1_Null_r16
			{
				seqOf := e.NewSequenceOfEncoder(codebookComboParametersAdditionR16Type1SPEType2R1NullR16SupportedCSIRSResourceListAddR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListAdd_r16))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListAdd_r16 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListAdd_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 5. type1SP-eType2R2-null-r16: sequence
	{
		if ie.Type1SP_EType2R2_Null_r16 != nil {
			c := ie.Type1SP_EType2R2_Null_r16
			{
				seqOf := e.NewSequenceOfEncoder(codebookComboParametersAdditionR16Type1SPEType2R2NullR16SupportedCSIRSResourceListAddR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListAdd_r16))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListAdd_r16 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListAdd_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 6. type1SP-eType2R1PS-null-r16: sequence
	{
		if ie.Type1SP_EType2R1PS_Null_r16 != nil {
			c := ie.Type1SP_EType2R1PS_Null_r16
			{
				seqOf := e.NewSequenceOfEncoder(codebookComboParametersAdditionR16Type1SPEType2R1PSNullR16SupportedCSIRSResourceListAddR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListAdd_r16))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListAdd_r16 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListAdd_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 7. type1SP-eType2R2PS-null-r16: sequence
	{
		if ie.Type1SP_EType2R2PS_Null_r16 != nil {
			c := ie.Type1SP_EType2R2PS_Null_r16
			{
				seqOf := e.NewSequenceOfEncoder(codebookComboParametersAdditionR16Type1SPEType2R2PSNullR16SupportedCSIRSResourceListAddR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListAdd_r16))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListAdd_r16 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListAdd_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 8. type1SP-Type2-Type2PS-r16: sequence
	{
		if ie.Type1SP_Type2_Type2PS_r16 != nil {
			c := ie.Type1SP_Type2_Type2PS_r16
			{
				seqOf := e.NewSequenceOfEncoder(codebookComboParametersAdditionR16Type1SPType2Type2PSR16SupportedCSIRSResourceListAddR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListAdd_r16))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListAdd_r16 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListAdd_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 9. type1MP-Type2-null-r16: sequence
	{
		if ie.Type1MP_Type2_Null_r16 != nil {
			c := ie.Type1MP_Type2_Null_r16
			{
				seqOf := e.NewSequenceOfEncoder(codebookComboParametersAdditionR16Type1MPType2NullR16SupportedCSIRSResourceListAddR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListAdd_r16))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListAdd_r16 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListAdd_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 10. type1MP-Type2PS-null-r16: sequence
	{
		if ie.Type1MP_Type2PS_Null_r16 != nil {
			c := ie.Type1MP_Type2PS_Null_r16
			{
				seqOf := e.NewSequenceOfEncoder(codebookComboParametersAdditionR16Type1MPType2PSNullR16SupportedCSIRSResourceListAddR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListAdd_r16))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListAdd_r16 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListAdd_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 11. type1MP-eType2R1-null-r16: sequence
	{
		if ie.Type1MP_EType2R1_Null_r16 != nil {
			c := ie.Type1MP_EType2R1_Null_r16
			{
				seqOf := e.NewSequenceOfEncoder(codebookComboParametersAdditionR16Type1MPEType2R1NullR16SupportedCSIRSResourceListAddR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListAdd_r16))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListAdd_r16 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListAdd_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 12. type1MP-eType2R2-null-r16: sequence
	{
		if ie.Type1MP_EType2R2_Null_r16 != nil {
			c := ie.Type1MP_EType2R2_Null_r16
			{
				seqOf := e.NewSequenceOfEncoder(codebookComboParametersAdditionR16Type1MPEType2R2NullR16SupportedCSIRSResourceListAddR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListAdd_r16))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListAdd_r16 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListAdd_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 13. type1MP-eType2R1PS-null-r16: sequence
	{
		if ie.Type1MP_EType2R1PS_Null_r16 != nil {
			c := ie.Type1MP_EType2R1PS_Null_r16
			{
				seqOf := e.NewSequenceOfEncoder(codebookComboParametersAdditionR16Type1MPEType2R1PSNullR16SupportedCSIRSResourceListAddR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListAdd_r16))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListAdd_r16 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListAdd_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 14. type1MP-eType2R2PS-null-r16: sequence
	{
		if ie.Type1MP_EType2R2PS_Null_r16 != nil {
			c := ie.Type1MP_EType2R2PS_Null_r16
			{
				seqOf := e.NewSequenceOfEncoder(codebookComboParametersAdditionR16Type1MPEType2R2PSNullR16SupportedCSIRSResourceListAddR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListAdd_r16))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListAdd_r16 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListAdd_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 15. type1MP-Type2-Type2PS-r16: sequence
	{
		if ie.Type1MP_Type2_Type2PS_r16 != nil {
			c := ie.Type1MP_Type2_Type2PS_r16
			{
				seqOf := e.NewSequenceOfEncoder(codebookComboParametersAdditionR16Type1MPType2Type2PSR16SupportedCSIRSResourceListAddR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListAdd_r16))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListAdd_r16 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListAdd_r16[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (ie *CodebookComboParametersAddition_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookComboParametersAdditionR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. type1SP-Type2-null-r16: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Type1SP_Type2_Null_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
			c := ie.Type1SP_Type2_Null_r16
			{
				seqOf := d.NewSequenceOfDecoder(codebookComboParametersAdditionR16Type1SPType2NullR16SupportedCSIRSResourceListAddR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListAdd_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListAdd_r16[i] = v
				}
			}
		}
	}

	// 3. type1SP-Type2PS-null-r16: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Type1SP_Type2PS_Null_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
			c := ie.Type1SP_Type2PS_Null_r16
			{
				seqOf := d.NewSequenceOfDecoder(codebookComboParametersAdditionR16Type1SPType2PSNullR16SupportedCSIRSResourceListAddR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListAdd_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListAdd_r16[i] = v
				}
			}
		}
	}

	// 4. type1SP-eType2R1-null-r16: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.Type1SP_EType2R1_Null_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
			c := ie.Type1SP_EType2R1_Null_r16
			{
				seqOf := d.NewSequenceOfDecoder(codebookComboParametersAdditionR16Type1SPEType2R1NullR16SupportedCSIRSResourceListAddR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListAdd_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListAdd_r16[i] = v
				}
			}
		}
	}

	// 5. type1SP-eType2R2-null-r16: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.Type1SP_EType2R2_Null_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
			c := ie.Type1SP_EType2R2_Null_r16
			{
				seqOf := d.NewSequenceOfDecoder(codebookComboParametersAdditionR16Type1SPEType2R2NullR16SupportedCSIRSResourceListAddR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListAdd_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListAdd_r16[i] = v
				}
			}
		}
	}

	// 6. type1SP-eType2R1PS-null-r16: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.Type1SP_EType2R1PS_Null_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
			c := ie.Type1SP_EType2R1PS_Null_r16
			{
				seqOf := d.NewSequenceOfDecoder(codebookComboParametersAdditionR16Type1SPEType2R1PSNullR16SupportedCSIRSResourceListAddR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListAdd_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListAdd_r16[i] = v
				}
			}
		}
	}

	// 7. type1SP-eType2R2PS-null-r16: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.Type1SP_EType2R2PS_Null_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
			c := ie.Type1SP_EType2R2PS_Null_r16
			{
				seqOf := d.NewSequenceOfDecoder(codebookComboParametersAdditionR16Type1SPEType2R2PSNullR16SupportedCSIRSResourceListAddR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListAdd_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListAdd_r16[i] = v
				}
			}
		}
	}

	// 8. type1SP-Type2-Type2PS-r16: sequence
	{
		if seq.IsComponentPresent(6) {
			ie.Type1SP_Type2_Type2PS_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
			c := ie.Type1SP_Type2_Type2PS_r16
			{
				seqOf := d.NewSequenceOfDecoder(codebookComboParametersAdditionR16Type1SPType2Type2PSR16SupportedCSIRSResourceListAddR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListAdd_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListAdd_r16[i] = v
				}
			}
		}
	}

	// 9. type1MP-Type2-null-r16: sequence
	{
		if seq.IsComponentPresent(7) {
			ie.Type1MP_Type2_Null_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
			c := ie.Type1MP_Type2_Null_r16
			{
				seqOf := d.NewSequenceOfDecoder(codebookComboParametersAdditionR16Type1MPType2NullR16SupportedCSIRSResourceListAddR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListAdd_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListAdd_r16[i] = v
				}
			}
		}
	}

	// 10. type1MP-Type2PS-null-r16: sequence
	{
		if seq.IsComponentPresent(8) {
			ie.Type1MP_Type2PS_Null_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
			c := ie.Type1MP_Type2PS_Null_r16
			{
				seqOf := d.NewSequenceOfDecoder(codebookComboParametersAdditionR16Type1MPType2PSNullR16SupportedCSIRSResourceListAddR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListAdd_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListAdd_r16[i] = v
				}
			}
		}
	}

	// 11. type1MP-eType2R1-null-r16: sequence
	{
		if seq.IsComponentPresent(9) {
			ie.Type1MP_EType2R1_Null_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
			c := ie.Type1MP_EType2R1_Null_r16
			{
				seqOf := d.NewSequenceOfDecoder(codebookComboParametersAdditionR16Type1MPEType2R1NullR16SupportedCSIRSResourceListAddR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListAdd_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListAdd_r16[i] = v
				}
			}
		}
	}

	// 12. type1MP-eType2R2-null-r16: sequence
	{
		if seq.IsComponentPresent(10) {
			ie.Type1MP_EType2R2_Null_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
			c := ie.Type1MP_EType2R2_Null_r16
			{
				seqOf := d.NewSequenceOfDecoder(codebookComboParametersAdditionR16Type1MPEType2R2NullR16SupportedCSIRSResourceListAddR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListAdd_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListAdd_r16[i] = v
				}
			}
		}
	}

	// 13. type1MP-eType2R1PS-null-r16: sequence
	{
		if seq.IsComponentPresent(11) {
			ie.Type1MP_EType2R1PS_Null_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
			c := ie.Type1MP_EType2R1PS_Null_r16
			{
				seqOf := d.NewSequenceOfDecoder(codebookComboParametersAdditionR16Type1MPEType2R1PSNullR16SupportedCSIRSResourceListAddR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListAdd_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListAdd_r16[i] = v
				}
			}
		}
	}

	// 14. type1MP-eType2R2PS-null-r16: sequence
	{
		if seq.IsComponentPresent(12) {
			ie.Type1MP_EType2R2PS_Null_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
			c := ie.Type1MP_EType2R2PS_Null_r16
			{
				seqOf := d.NewSequenceOfDecoder(codebookComboParametersAdditionR16Type1MPEType2R2PSNullR16SupportedCSIRSResourceListAddR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListAdd_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListAdd_r16[i] = v
				}
			}
		}
	}

	// 15. type1MP-Type2-Type2PS-r16: sequence
	{
		if seq.IsComponentPresent(13) {
			ie.Type1MP_Type2_Type2PS_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
			c := ie.Type1MP_Type2_Type2PS_r16
			{
				seqOf := d.NewSequenceOfDecoder(codebookComboParametersAdditionR16Type1MPType2Type2PSR16SupportedCSIRSResourceListAddR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListAdd_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListAdd_r16[i] = v
				}
			}
		}
	}

	return nil
}
