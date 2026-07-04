// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParametersType1SP-SchemeA-r19 (line 19040).

var codebookParametersType1SPSchemeAR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "enhType1SP64PortsSchemeA-r19"},
		{Name: "enhType1SP48PortsSchemeA-r19", Optional: true},
		{Name: "enhType1SP128PortsSchemeA-r19", Optional: true},
	},
}

var codebookParametersType1SPSchemeAR19EnhType1SP64PortsSchemeAR19SupportedCSIRSResourceExtListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersType1SP_SchemeA_r19_EnhType1SP64PortsSchemeA_r19_MaxNumberResource_r19_N2 = 0
	CodebookParametersType1SP_SchemeA_r19_EnhType1SP64PortsSchemeA_r19_MaxNumberResource_r19_N4 = 1
)

var codebookParametersType1SPSchemeAR19EnhType1SP64PortsSchemeAR19MaxNumberResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersType1SP_SchemeA_r19_EnhType1SP64PortsSchemeA_r19_MaxNumberResource_r19_N2, CodebookParametersType1SP_SchemeA_r19_EnhType1SP64PortsSchemeA_r19_MaxNumberResource_r19_N4},
}

const (
	CodebookParametersType1SP_SchemeA_r19_EnhType1SP64PortsSchemeA_r19_ProcessingCapability_r19_Cap1 = 0
	CodebookParametersType1SP_SchemeA_r19_EnhType1SP64PortsSchemeA_r19_ProcessingCapability_r19_Cap2 = 1
)

var codebookParametersType1SPSchemeAR19EnhType1SP64PortsSchemeAR19ProcessingCapabilityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersType1SP_SchemeA_r19_EnhType1SP64PortsSchemeA_r19_ProcessingCapability_r19_Cap1, CodebookParametersType1SP_SchemeA_r19_EnhType1SP64PortsSchemeA_r19_ProcessingCapability_r19_Cap2},
}

var codebookParametersType1SPSchemeAR19EnhType1SP64PortsSchemeAR19SupportedCSIRSResourceListPerCCR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersType1SPSchemeAR19EnhType1SP48PortsSchemeAR19SupportedCSIRSResourceExtListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersType1SP_SchemeA_r19_EnhType1SP48PortsSchemeA_r19_ProcessingCapability_r19_Cap1 = 0
	CodebookParametersType1SP_SchemeA_r19_EnhType1SP48PortsSchemeA_r19_ProcessingCapability_r19_Cap2 = 1
)

var codebookParametersType1SPSchemeAR19EnhType1SP48PortsSchemeAR19ProcessingCapabilityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersType1SP_SchemeA_r19_EnhType1SP48PortsSchemeA_r19_ProcessingCapability_r19_Cap1, CodebookParametersType1SP_SchemeA_r19_EnhType1SP48PortsSchemeA_r19_ProcessingCapability_r19_Cap2},
}

var codebookParametersType1SPSchemeAR19EnhType1SP48PortsSchemeAR19SupportedCSIRSResourceListPerCCR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersType1SPSchemeAR19EnhType1SP128PortsSchemeAR19SupportedCSIRSResourceExtListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersType1SP_SchemeA_r19_EnhType1SP128PortsSchemeA_r19_ProcessingCapability_r19_Cap1 = 0
	CodebookParametersType1SP_SchemeA_r19_EnhType1SP128PortsSchemeA_r19_ProcessingCapability_r19_Cap2 = 1
)

var codebookParametersType1SPSchemeAR19EnhType1SP128PortsSchemeAR19ProcessingCapabilityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersType1SP_SchemeA_r19_EnhType1SP128PortsSchemeA_r19_ProcessingCapability_r19_Cap1, CodebookParametersType1SP_SchemeA_r19_EnhType1SP128PortsSchemeA_r19_ProcessingCapability_r19_Cap2},
}

var codebookParametersType1SPSchemeAR19EnhType1SP128PortsSchemeAR19SupportedCSIRSResourceListPerCCR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

type CodebookParametersType1SP_SchemeA_r19 struct {
	EnhType1SP64PortsSchemeA_r19 struct {
		SupportedCSI_RS_ResourceExtList_r19   []int64
		MaxRank_r19                           int64
		MaxNumberResource_r19                 int64
		ProcessingCapability_r19              int64
		SupportedCSI_RS_ResourceListPerCC_r19 []int64
	}
	EnhType1SP48PortsSchemeA_r19 *struct {
		SupportedCSI_RS_ResourceExtList_r19   []int64
		MaxRank_r19                           int64
		MaxNumberResource_r19                 int64
		ProcessingCapability_r19              int64
		SupportedCSI_RS_ResourceListPerCC_r19 []int64
	}
	EnhType1SP128PortsSchemeA_r19 *struct {
		SupportedCSI_RS_ResourceExtList_r19   []int64
		MaxRank_r19                           int64
		ProcessingCapability_r19              int64
		SupportedCSI_RS_ResourceListPerCC_r19 []int64
	}
}

func (ie *CodebookParametersType1SP_SchemeA_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParametersType1SPSchemeAR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.EnhType1SP48PortsSchemeA_r19 != nil, ie.EnhType1SP128PortsSchemeA_r19 != nil}); err != nil {
		return err
	}

	// 2. enhType1SP64PortsSchemeA-r19: sequence
	{
		{
			c := &ie.EnhType1SP64PortsSchemeA_r19
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersType1SPSchemeAR19EnhType1SP64PortsSchemeAR19SupportedCSIRSResourceExtListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceExtList_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceExtList_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceExtList_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeInteger(c.MaxRank_r19, per.Constrained(4, 8)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberResource_r19, codebookParametersType1SPSchemeAR19EnhType1SP64PortsSchemeAR19MaxNumberResourceR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.ProcessingCapability_r19, codebookParametersType1SPSchemeAR19EnhType1SP64PortsSchemeAR19ProcessingCapabilityR19Constraints); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersType1SPSchemeAR19EnhType1SP64PortsSchemeAR19SupportedCSIRSResourceListPerCCR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListPerCC_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListPerCC_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListPerCC_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 3. enhType1SP48PortsSchemeA-r19: sequence
	{
		if ie.EnhType1SP48PortsSchemeA_r19 != nil {
			c := ie.EnhType1SP48PortsSchemeA_r19
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersType1SPSchemeAR19EnhType1SP48PortsSchemeAR19SupportedCSIRSResourceExtListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceExtList_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceExtList_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceExtList_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeInteger(c.MaxRank_r19, per.Constrained(4, 8)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberResource_r19, per.Constrained(2, 3)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.ProcessingCapability_r19, codebookParametersType1SPSchemeAR19EnhType1SP48PortsSchemeAR19ProcessingCapabilityR19Constraints); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersType1SPSchemeAR19EnhType1SP48PortsSchemeAR19SupportedCSIRSResourceListPerCCR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListPerCC_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListPerCC_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListPerCC_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 4. enhType1SP128PortsSchemeA-r19: sequence
	{
		if ie.EnhType1SP128PortsSchemeA_r19 != nil {
			c := ie.EnhType1SP128PortsSchemeA_r19
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersType1SPSchemeAR19EnhType1SP128PortsSchemeAR19SupportedCSIRSResourceExtListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceExtList_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceExtList_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceExtList_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeInteger(c.MaxRank_r19, per.Constrained(4, 8)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.ProcessingCapability_r19, codebookParametersType1SPSchemeAR19EnhType1SP128PortsSchemeAR19ProcessingCapabilityR19Constraints); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersType1SPSchemeAR19EnhType1SP128PortsSchemeAR19SupportedCSIRSResourceListPerCCR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceListPerCC_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceListPerCC_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceListPerCC_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (ie *CodebookParametersType1SP_SchemeA_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParametersType1SPSchemeAR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. enhType1SP64PortsSchemeA-r19: sequence
	{
		{
			c := &ie.EnhType1SP64PortsSchemeA_r19
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersType1SPSchemeAR19EnhType1SP64PortsSchemeAR19SupportedCSIRSResourceExtListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceExtList_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceExtList_r19[i] = v
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(4, 8))
				if err != nil {
					return err
				}
				c.MaxRank_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParametersType1SPSchemeAR19EnhType1SP64PortsSchemeAR19MaxNumberResourceR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberResource_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParametersType1SPSchemeAR19EnhType1SP64PortsSchemeAR19ProcessingCapabilityR19Constraints)
				if err != nil {
					return err
				}
				c.ProcessingCapability_r19 = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersType1SPSchemeAR19EnhType1SP64PortsSchemeAR19SupportedCSIRSResourceListPerCCR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListPerCC_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListPerCC_r19[i] = v
				}
			}
		}
	}

	// 3. enhType1SP48PortsSchemeA-r19: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.EnhType1SP48PortsSchemeA_r19 = &struct {
				SupportedCSI_RS_ResourceExtList_r19   []int64
				MaxRank_r19                           int64
				MaxNumberResource_r19                 int64
				ProcessingCapability_r19              int64
				SupportedCSI_RS_ResourceListPerCC_r19 []int64
			}{}
			c := ie.EnhType1SP48PortsSchemeA_r19
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersType1SPSchemeAR19EnhType1SP48PortsSchemeAR19SupportedCSIRSResourceExtListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceExtList_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceExtList_r19[i] = v
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(4, 8))
				if err != nil {
					return err
				}
				c.MaxRank_r19 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(2, 3))
				if err != nil {
					return err
				}
				c.MaxNumberResource_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParametersType1SPSchemeAR19EnhType1SP48PortsSchemeAR19ProcessingCapabilityR19Constraints)
				if err != nil {
					return err
				}
				c.ProcessingCapability_r19 = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersType1SPSchemeAR19EnhType1SP48PortsSchemeAR19SupportedCSIRSResourceListPerCCR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListPerCC_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListPerCC_r19[i] = v
				}
			}
		}
	}

	// 4. enhType1SP128PortsSchemeA-r19: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.EnhType1SP128PortsSchemeA_r19 = &struct {
				SupportedCSI_RS_ResourceExtList_r19   []int64
				MaxRank_r19                           int64
				ProcessingCapability_r19              int64
				SupportedCSI_RS_ResourceListPerCC_r19 []int64
			}{}
			c := ie.EnhType1SP128PortsSchemeA_r19
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersType1SPSchemeAR19EnhType1SP128PortsSchemeAR19SupportedCSIRSResourceExtListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceExtList_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceExtList_r19[i] = v
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(4, 8))
				if err != nil {
					return err
				}
				c.MaxRank_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParametersType1SPSchemeAR19EnhType1SP128PortsSchemeAR19ProcessingCapabilityR19Constraints)
				if err != nil {
					return err
				}
				c.ProcessingCapability_r19 = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersType1SPSchemeAR19EnhType1SP128PortsSchemeAR19SupportedCSIRSResourceListPerCCR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceListPerCC_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceListPerCC_r19[i] = v
				}
			}
		}
	}

	return nil
}
