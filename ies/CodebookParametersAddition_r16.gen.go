// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParametersAddition-r16 (line 18525).

var codebookParametersAdditionR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "etype2-r16", Optional: true},
		{Name: "etype2-PS-r16", Optional: true},
	},
}

var codebookParametersAdditionR16Etype2R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "etype2R1-r16"},
		{Name: "etype2R2-r16", Optional: true},
		{Name: "paramComb7-8-r16", Optional: true},
		{Name: "rank3-4-r16", Optional: true},
		{Name: "amplitudeSubsetRestriction-r16", Optional: true},
	},
}

var codebookParametersAdditionR16Etype2R16Etype2R1R16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersAdditionR16Etype2R16Etype2R2R16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersAddition_r16_Etype2_r16_ParamComb7_8_r16_Supported = 0
)

var codebookParametersAdditionR16Etype2R16ParamComb78R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersAddition_r16_Etype2_r16_ParamComb7_8_r16_Supported},
}

const (
	CodebookParametersAddition_r16_Etype2_r16_Rank3_4_r16_Supported = 0
)

var codebookParametersAdditionR16Etype2R16Rank34R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersAddition_r16_Etype2_r16_Rank3_4_r16_Supported},
}

const (
	CodebookParametersAddition_r16_Etype2_r16_AmplitudeSubsetRestriction_r16_Supported = 0
)

var codebookParametersAdditionR16Etype2R16AmplitudeSubsetRestrictionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersAddition_r16_Etype2_r16_AmplitudeSubsetRestriction_r16_Supported},
}

var codebookParametersAdditionR16Etype2PSR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "etype2R1-PortSelection-r16"},
		{Name: "etype2R2-PortSelection-r16", Optional: true},
		{Name: "rank3-4-r16", Optional: true},
	},
}

var codebookParametersAdditionR16Etype2PSR16Etype2R1PortSelectionR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersAdditionR16Etype2PSR16Etype2R2PortSelectionR16SupportedCSIRSResourceListAddR16Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersAddition_r16_Etype2_PS_r16_Rank3_4_r16_Supported = 0
)

var codebookParametersAdditionR16Etype2PSR16Rank34R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersAddition_r16_Etype2_PS_r16_Rank3_4_r16_Supported},
}

type CodebookParametersAddition_r16 struct {
	Etype2_r16 *struct {
		Etype2R1_r16                   struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
		Etype2R2_r16                   *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
		ParamComb7_8_r16               *int64
		Rank3_4_r16                    *int64
		AmplitudeSubsetRestriction_r16 *int64
	}
	Etype2_PS_r16 *struct {
		Etype2R1_PortSelection_r16 struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
		Etype2R2_PortSelection_r16 *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
		Rank3_4_r16                *int64
	}
}

func (ie *CodebookParametersAddition_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParametersAdditionR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Etype2_r16 != nil, ie.Etype2_PS_r16 != nil}); err != nil {
		return err
	}

	// 2. etype2-r16: sequence
	{
		if ie.Etype2_r16 != nil {
			c := ie.Etype2_r16
			codebookParametersAdditionR16Etype2R16Seq := e.NewSequenceEncoder(codebookParametersAdditionR16Etype2R16Constraints)
			if err := codebookParametersAdditionR16Etype2R16Seq.EncodePreamble([]bool{c.Etype2R2_r16 != nil, c.ParamComb7_8_r16 != nil, c.Rank3_4_r16 != nil, c.AmplitudeSubsetRestriction_r16 != nil}); err != nil {
				return err
			}
			{
				c := &c.Etype2R1_r16
				{
					seqOf := e.NewSequenceOfEncoder(codebookParametersAdditionR16Etype2R16Etype2R1R16SupportedCSIRSResourceListAddR16Constraints)
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
			if c.Etype2R2_r16 != nil {
				c := (*c.Etype2R2_r16)
				{
					seqOf := e.NewSequenceOfEncoder(codebookParametersAdditionR16Etype2R16Etype2R2R16SupportedCSIRSResourceListAddR16Constraints)
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
			if c.ParamComb7_8_r16 != nil {
				if err := e.EncodeEnumerated((*c.ParamComb7_8_r16), codebookParametersAdditionR16Etype2R16ParamComb78R16Constraints); err != nil {
					return err
				}
			}
			if c.Rank3_4_r16 != nil {
				if err := e.EncodeEnumerated((*c.Rank3_4_r16), codebookParametersAdditionR16Etype2R16Rank34R16Constraints); err != nil {
					return err
				}
			}
			if c.AmplitudeSubsetRestriction_r16 != nil {
				if err := e.EncodeEnumerated((*c.AmplitudeSubsetRestriction_r16), codebookParametersAdditionR16Etype2R16AmplitudeSubsetRestrictionR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 3. etype2-PS-r16: sequence
	{
		if ie.Etype2_PS_r16 != nil {
			c := ie.Etype2_PS_r16
			codebookParametersAdditionR16Etype2PSR16Seq := e.NewSequenceEncoder(codebookParametersAdditionR16Etype2PSR16Constraints)
			if err := codebookParametersAdditionR16Etype2PSR16Seq.EncodePreamble([]bool{c.Etype2R2_PortSelection_r16 != nil, c.Rank3_4_r16 != nil}); err != nil {
				return err
			}
			{
				c := &c.Etype2R1_PortSelection_r16
				{
					seqOf := e.NewSequenceOfEncoder(codebookParametersAdditionR16Etype2PSR16Etype2R1PortSelectionR16SupportedCSIRSResourceListAddR16Constraints)
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
			if c.Etype2R2_PortSelection_r16 != nil {
				c := (*c.Etype2R2_PortSelection_r16)
				{
					seqOf := e.NewSequenceOfEncoder(codebookParametersAdditionR16Etype2PSR16Etype2R2PortSelectionR16SupportedCSIRSResourceListAddR16Constraints)
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
			if c.Rank3_4_r16 != nil {
				if err := e.EncodeEnumerated((*c.Rank3_4_r16), codebookParametersAdditionR16Etype2PSR16Rank34R16Constraints); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CodebookParametersAddition_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParametersAdditionR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. etype2-r16: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Etype2_r16 = &struct {
				Etype2R1_r16                   struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
				Etype2R2_r16                   *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
				ParamComb7_8_r16               *int64
				Rank3_4_r16                    *int64
				AmplitudeSubsetRestriction_r16 *int64
			}{}
			c := ie.Etype2_r16
			codebookParametersAdditionR16Etype2R16Seq := d.NewSequenceDecoder(codebookParametersAdditionR16Etype2R16Constraints)
			if err := codebookParametersAdditionR16Etype2R16Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				c := &c.Etype2R1_r16
				{
					seqOf := d.NewSequenceOfDecoder(codebookParametersAdditionR16Etype2R16Etype2R1R16SupportedCSIRSResourceListAddR16Constraints)
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
			if codebookParametersAdditionR16Etype2R16Seq.IsComponentPresent(1) {
				c.Etype2R2_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
				c := (*c.Etype2R2_r16)
				{
					seqOf := d.NewSequenceOfDecoder(codebookParametersAdditionR16Etype2R16Etype2R2R16SupportedCSIRSResourceListAddR16Constraints)
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
			if codebookParametersAdditionR16Etype2R16Seq.IsComponentPresent(2) {
				c.ParamComb7_8_r16 = new(int64)
				v, err := d.DecodeEnumerated(codebookParametersAdditionR16Etype2R16ParamComb78R16Constraints)
				if err != nil {
					return err
				}
				(*c.ParamComb7_8_r16) = v
			}
			if codebookParametersAdditionR16Etype2R16Seq.IsComponentPresent(3) {
				c.Rank3_4_r16 = new(int64)
				v, err := d.DecodeEnumerated(codebookParametersAdditionR16Etype2R16Rank34R16Constraints)
				if err != nil {
					return err
				}
				(*c.Rank3_4_r16) = v
			}
			if codebookParametersAdditionR16Etype2R16Seq.IsComponentPresent(4) {
				c.AmplitudeSubsetRestriction_r16 = new(int64)
				v, err := d.DecodeEnumerated(codebookParametersAdditionR16Etype2R16AmplitudeSubsetRestrictionR16Constraints)
				if err != nil {
					return err
				}
				(*c.AmplitudeSubsetRestriction_r16) = v
			}
		}
	}

	// 3. etype2-PS-r16: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Etype2_PS_r16 = &struct {
				Etype2R1_PortSelection_r16 struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
				Etype2R2_PortSelection_r16 *struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }
				Rank3_4_r16                *int64
			}{}
			c := ie.Etype2_PS_r16
			codebookParametersAdditionR16Etype2PSR16Seq := d.NewSequenceDecoder(codebookParametersAdditionR16Etype2PSR16Constraints)
			if err := codebookParametersAdditionR16Etype2PSR16Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				c := &c.Etype2R1_PortSelection_r16
				{
					seqOf := d.NewSequenceOfDecoder(codebookParametersAdditionR16Etype2PSR16Etype2R1PortSelectionR16SupportedCSIRSResourceListAddR16Constraints)
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
			if codebookParametersAdditionR16Etype2PSR16Seq.IsComponentPresent(1) {
				c.Etype2R2_PortSelection_r16 = &struct{ SupportedCSI_RS_ResourceListAdd_r16 []int64 }{}
				c := (*c.Etype2R2_PortSelection_r16)
				{
					seqOf := d.NewSequenceOfDecoder(codebookParametersAdditionR16Etype2PSR16Etype2R2PortSelectionR16SupportedCSIRSResourceListAddR16Constraints)
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
			if codebookParametersAdditionR16Etype2PSR16Seq.IsComponentPresent(2) {
				c.Rank3_4_r16 = new(int64)
				v, err := d.DecodeEnumerated(codebookParametersAdditionR16Etype2PSR16Rank34R16Constraints)
				if err != nil {
					return err
				}
				(*c.Rank3_4_r16) = v
			}
		}
	}

	return nil
}
