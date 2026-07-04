// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParameterseType2Ext-r19 (line 19136).

var codebookParameterseType2ExtR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "eType2-64PortExt-r19"},
		{Name: "eType2-48PortExt-r19", Optional: true},
		{Name: "eType2-128PortExt-r19", Optional: true},
		{Name: "eType2R2Ext-r19", Optional: true},
		{Name: "eType2ExtPC7-8-r19", Optional: true},
		{Name: "eType2R3R4Ext-r19", Optional: true},
	},
}

var codebookParameterseType2ExtR19EType2R2ExtR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParameterseType2Ext_r19_EType2ExtPC7_8_r19_Supported = 0
)

var codebookParameterseType2ExtR19EType2ExtPC78R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2Ext_r19_EType2ExtPC7_8_r19_Supported},
}

var codebookParameterseType2ExtR19EType2R3R4ExtR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParameterseType2ExtR19EType264PortExtR19SupportedCSIRSResourceExtListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParameterseType2Ext_r19_EType2_64PortExt_r19_ProcessingCapability_r19_Cap1 = 0
	CodebookParameterseType2Ext_r19_EType2_64PortExt_r19_ProcessingCapability_r19_Cap2 = 1
)

var codebookParameterseType2ExtR19EType264PortExtR19ProcessingCapabilityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2Ext_r19_EType2_64PortExt_r19_ProcessingCapability_r19_Cap1, CodebookParameterseType2Ext_r19_EType2_64PortExt_r19_ProcessingCapability_r19_Cap2},
}

const (
	CodebookParameterseType2Ext_r19_EType2_64PortExt_r19_MaxNumberResource_r19_N2 = 0
	CodebookParameterseType2Ext_r19_EType2_64PortExt_r19_MaxNumberResource_r19_N4 = 1
)

var codebookParameterseType2ExtR19EType264PortExtR19MaxNumberResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2Ext_r19_EType2_64PortExt_r19_MaxNumberResource_r19_N2, CodebookParameterseType2Ext_r19_EType2_64PortExt_r19_MaxNumberResource_r19_N4},
}

var codebookParameterseType2ExtR19EType264PortExtR19SupportedCSIRSResourceListPerCCR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParameterseType2ExtR19EType248PortExtR19SupportedCSIRSResourceExtListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParameterseType2Ext_r19_EType2_48PortExt_r19_ProcessingCapability_r19_Cap1 = 0
	CodebookParameterseType2Ext_r19_EType2_48PortExt_r19_ProcessingCapability_r19_Cap2 = 1
)

var codebookParameterseType2ExtR19EType248PortExtR19ProcessingCapabilityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2Ext_r19_EType2_48PortExt_r19_ProcessingCapability_r19_Cap1, CodebookParameterseType2Ext_r19_EType2_48PortExt_r19_ProcessingCapability_r19_Cap2},
}

const (
	CodebookParameterseType2Ext_r19_EType2_48PortExt_r19_MaxNumberResource_r19_N2 = 0
	CodebookParameterseType2Ext_r19_EType2_48PortExt_r19_MaxNumberResource_r19_N3 = 1
)

var codebookParameterseType2ExtR19EType248PortExtR19MaxNumberResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2Ext_r19_EType2_48PortExt_r19_MaxNumberResource_r19_N2, CodebookParameterseType2Ext_r19_EType2_48PortExt_r19_MaxNumberResource_r19_N3},
}

var codebookParameterseType2ExtR19EType248PortExtR19SupportedCSIRSResourceListPerCCR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParameterseType2ExtR19EType2128PortExtR19SupportedCSIRSResourceExtListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParameterseType2Ext_r19_EType2_128PortExt_r19_ProcessingCapability_r19_Cap1 = 0
	CodebookParameterseType2Ext_r19_EType2_128PortExt_r19_ProcessingCapability_r19_Cap2 = 1
)

var codebookParameterseType2ExtR19EType2128PortExtR19ProcessingCapabilityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParameterseType2Ext_r19_EType2_128PortExt_r19_ProcessingCapability_r19_Cap1, CodebookParameterseType2Ext_r19_EType2_128PortExt_r19_ProcessingCapability_r19_Cap2},
}

var codebookParameterseType2ExtR19EType2128PortExtR19SupportedCSIRSResourceListPerCCR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

type CodebookParameterseType2Ext_r19 struct {
	EType2_64PortExt_r19 struct {
		SupportedCSI_RS_ResourceExtList_r19   []int64
		ProcessingCapability_r19              int64
		MaxNumberResource_r19                 int64
		SupportedCSI_RS_ResourceListPerCC_r19 []int64
	}
	EType2_48PortExt_r19 *struct {
		SupportedCSI_RS_ResourceExtList_r19   []int64
		ProcessingCapability_r19              int64
		MaxNumberResource_r19                 int64
		SupportedCSI_RS_ResourceListPerCC_r19 []int64
	}
	EType2_128PortExt_r19 *struct {
		SupportedCSI_RS_ResourceExtList_r19   []int64
		ProcessingCapability_r19              int64
		SupportedCSI_RS_ResourceListPerCC_r19 []int64
	}
	EType2R2Ext_r19    []int64
	EType2ExtPC7_8_r19 *int64
	EType2R3R4Ext_r19  []int64
}

func (ie *CodebookParameterseType2Ext_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParameterseType2ExtR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.EType2_48PortExt_r19 != nil, ie.EType2_128PortExt_r19 != nil, ie.EType2R2Ext_r19 != nil, ie.EType2ExtPC7_8_r19 != nil, ie.EType2R3R4Ext_r19 != nil}); err != nil {
		return err
	}

	// 2. eType2-64PortExt-r19: sequence
	{
		{
			c := &ie.EType2_64PortExt_r19
			{
				seqOf := e.NewSequenceOfEncoder(codebookParameterseType2ExtR19EType264PortExtR19SupportedCSIRSResourceExtListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceExtList_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceExtList_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceExtList_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.ProcessingCapability_r19, codebookParameterseType2ExtR19EType264PortExtR19ProcessingCapabilityR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberResource_r19, codebookParameterseType2ExtR19EType264PortExtR19MaxNumberResourceR19Constraints); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(codebookParameterseType2ExtR19EType264PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
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

	// 3. eType2-48PortExt-r19: sequence
	{
		if ie.EType2_48PortExt_r19 != nil {
			c := ie.EType2_48PortExt_r19
			{
				seqOf := e.NewSequenceOfEncoder(codebookParameterseType2ExtR19EType248PortExtR19SupportedCSIRSResourceExtListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceExtList_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceExtList_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceExtList_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.ProcessingCapability_r19, codebookParameterseType2ExtR19EType248PortExtR19ProcessingCapabilityR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberResource_r19, codebookParameterseType2ExtR19EType248PortExtR19MaxNumberResourceR19Constraints); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(codebookParameterseType2ExtR19EType248PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
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

	// 4. eType2-128PortExt-r19: sequence
	{
		if ie.EType2_128PortExt_r19 != nil {
			c := ie.EType2_128PortExt_r19
			{
				seqOf := e.NewSequenceOfEncoder(codebookParameterseType2ExtR19EType2128PortExtR19SupportedCSIRSResourceExtListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceExtList_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceExtList_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceExtList_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.ProcessingCapability_r19, codebookParameterseType2ExtR19EType2128PortExtR19ProcessingCapabilityR19Constraints); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(codebookParameterseType2ExtR19EType2128PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
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

	// 5. eType2R2Ext-r19: sequence-of
	{
		if ie.EType2R2Ext_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParameterseType2ExtR19EType2R2ExtR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.EType2R2Ext_r19))); err != nil {
				return err
			}
			for i := range ie.EType2R2Ext_r19 {
				if err := e.EncodeInteger(ie.EType2R2Ext_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 6. eType2ExtPC7-8-r19: enumerated
	{
		if ie.EType2ExtPC7_8_r19 != nil {
			if err := e.EncodeEnumerated(*ie.EType2ExtPC7_8_r19, codebookParameterseType2ExtR19EType2ExtPC78R19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. eType2R3R4Ext-r19: sequence-of
	{
		if ie.EType2R3R4Ext_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParameterseType2ExtR19EType2R3R4ExtR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.EType2R3R4Ext_r19))); err != nil {
				return err
			}
			for i := range ie.EType2R3R4Ext_r19 {
				if err := e.EncodeInteger(ie.EType2R3R4Ext_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CodebookParameterseType2Ext_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParameterseType2ExtR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. eType2-64PortExt-r19: sequence
	{
		{
			c := &ie.EType2_64PortExt_r19
			{
				seqOf := d.NewSequenceOfDecoder(codebookParameterseType2ExtR19EType264PortExtR19SupportedCSIRSResourceExtListR19Constraints)
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
				v, err := d.DecodeEnumerated(codebookParameterseType2ExtR19EType264PortExtR19ProcessingCapabilityR19Constraints)
				if err != nil {
					return err
				}
				c.ProcessingCapability_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParameterseType2ExtR19EType264PortExtR19MaxNumberResourceR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberResource_r19 = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(codebookParameterseType2ExtR19EType264PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
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

	// 3. eType2-48PortExt-r19: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.EType2_48PortExt_r19 = &struct {
				SupportedCSI_RS_ResourceExtList_r19   []int64
				ProcessingCapability_r19              int64
				MaxNumberResource_r19                 int64
				SupportedCSI_RS_ResourceListPerCC_r19 []int64
			}{}
			c := ie.EType2_48PortExt_r19
			{
				seqOf := d.NewSequenceOfDecoder(codebookParameterseType2ExtR19EType248PortExtR19SupportedCSIRSResourceExtListR19Constraints)
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
				v, err := d.DecodeEnumerated(codebookParameterseType2ExtR19EType248PortExtR19ProcessingCapabilityR19Constraints)
				if err != nil {
					return err
				}
				c.ProcessingCapability_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParameterseType2ExtR19EType248PortExtR19MaxNumberResourceR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberResource_r19 = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(codebookParameterseType2ExtR19EType248PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
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

	// 4. eType2-128PortExt-r19: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.EType2_128PortExt_r19 = &struct {
				SupportedCSI_RS_ResourceExtList_r19   []int64
				ProcessingCapability_r19              int64
				SupportedCSI_RS_ResourceListPerCC_r19 []int64
			}{}
			c := ie.EType2_128PortExt_r19
			{
				seqOf := d.NewSequenceOfDecoder(codebookParameterseType2ExtR19EType2128PortExtR19SupportedCSIRSResourceExtListR19Constraints)
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
				v, err := d.DecodeEnumerated(codebookParameterseType2ExtR19EType2128PortExtR19ProcessingCapabilityR19Constraints)
				if err != nil {
					return err
				}
				c.ProcessingCapability_r19 = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(codebookParameterseType2ExtR19EType2128PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
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

	// 5. eType2R2Ext-r19: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(codebookParameterseType2ExtR19EType2R2ExtR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.EType2R2Ext_r19 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.EType2R2Ext_r19[i] = v
			}
		}
	}

	// 6. eType2ExtPC7-8-r19: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(codebookParameterseType2ExtR19EType2ExtPC78R19Constraints)
			if err != nil {
				return err
			}
			ie.EType2ExtPC7_8_r19 = &idx
		}
	}

	// 7. eType2R3R4Ext-r19: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(codebookParameterseType2ExtR19EType2R3R4ExtR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.EType2R3R4Ext_r19 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.EType2R3R4Ext_r19[i] = v
			}
		}
	}

	return nil
}
