// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParametersfeType2Ext-r19 (line 19173).

var codebookParametersfeType2ExtR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "feType2-64PortExt-r19"},
		{Name: "feType2-48PortExt-r19", Optional: true},
		{Name: "feType2-M2R1Ext-r19", Optional: true},
		{Name: "feType2-M2R2Ext-r19", Optional: true},
		{Name: "feType2-R3R4Ext-r19", Optional: true},
	},
}

var codebookParametersfeType2ExtR19FeType2M2R1ExtR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersfeType2ExtR19FeType2M2R2ExtR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersfeType2Ext_r19_FeType2_R3R4Ext_r19_Supported = 0
)

var codebookParametersfeType2ExtR19FeType2R3R4ExtR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfeType2Ext_r19_FeType2_R3R4Ext_r19_Supported},
}

var codebookParametersfeType2ExtR19FeType264PortExtR19SupportedCSIRSResourceExtListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersfeType2Ext_r19_FeType2_64PortExt_r19_ProcessingCapability_r19_Cap1 = 0
	CodebookParametersfeType2Ext_r19_FeType2_64PortExt_r19_ProcessingCapability_r19_Cap2 = 1
)

var codebookParametersfeType2ExtR19FeType264PortExtR19ProcessingCapabilityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfeType2Ext_r19_FeType2_64PortExt_r19_ProcessingCapability_r19_Cap1, CodebookParametersfeType2Ext_r19_FeType2_64PortExt_r19_ProcessingCapability_r19_Cap2},
}

const (
	CodebookParametersfeType2Ext_r19_FeType2_64PortExt_r19_MaxNumberResource_r19_N2 = 0
	CodebookParametersfeType2Ext_r19_FeType2_64PortExt_r19_MaxNumberResource_r19_N4 = 1
)

var codebookParametersfeType2ExtR19FeType264PortExtR19MaxNumberResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfeType2Ext_r19_FeType2_64PortExt_r19_MaxNumberResource_r19_N2, CodebookParametersfeType2Ext_r19_FeType2_64PortExt_r19_MaxNumberResource_r19_N4},
}

var codebookParametersfeType2ExtR19FeType264PortExtR19SupportedCSIRSResourceListPerCCR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersfeType2ExtR19FeType248PortExtR19SupportedCSIRSResourceExtListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersfeType2Ext_r19_FeType2_48PortExt_r19_ProcessingCapability_r19_Cap1 = 0
	CodebookParametersfeType2Ext_r19_FeType2_48PortExt_r19_ProcessingCapability_r19_Cap2 = 1
)

var codebookParametersfeType2ExtR19FeType248PortExtR19ProcessingCapabilityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfeType2Ext_r19_FeType2_48PortExt_r19_ProcessingCapability_r19_Cap1, CodebookParametersfeType2Ext_r19_FeType2_48PortExt_r19_ProcessingCapability_r19_Cap2},
}

const (
	CodebookParametersfeType2Ext_r19_FeType2_48PortExt_r19_MaxNumberResource_r19_N2 = 0
	CodebookParametersfeType2Ext_r19_FeType2_48PortExt_r19_MaxNumberResource_r19_N3 = 1
)

var codebookParametersfeType2ExtR19FeType248PortExtR19MaxNumberResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersfeType2Ext_r19_FeType2_48PortExt_r19_MaxNumberResource_r19_N2, CodebookParametersfeType2Ext_r19_FeType2_48PortExt_r19_MaxNumberResource_r19_N3},
}

var codebookParametersfeType2ExtR19FeType248PortExtR19SupportedCSIRSResourceListPerCCR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

type CodebookParametersfeType2Ext_r19 struct {
	FeType2_64PortExt_r19 struct {
		SupportedCSI_RS_ResourceExtList_r19   []int64
		ProcessingCapability_r19              int64
		MaxNumberResource_r19                 int64
		SupportedCSI_RS_ResourceListPerCC_r19 []int64
	}
	FeType2_48PortExt_r19 *struct {
		SupportedCSI_RS_ResourceExtList_r19   []int64
		ProcessingCapability_r19              int64
		MaxNumberResource_r19                 int64
		SupportedCSI_RS_ResourceListPerCC_r19 []int64
	}
	FeType2_M2R1Ext_r19 []int64
	FeType2_M2R2Ext_r19 []int64
	FeType2_R3R4Ext_r19 *int64
}

func (ie *CodebookParametersfeType2Ext_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParametersfeType2ExtR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FeType2_48PortExt_r19 != nil, ie.FeType2_M2R1Ext_r19 != nil, ie.FeType2_M2R2Ext_r19 != nil, ie.FeType2_R3R4Ext_r19 != nil}); err != nil {
		return err
	}

	// 2. feType2-64PortExt-r19: sequence
	{
		{
			c := &ie.FeType2_64PortExt_r19
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersfeType2ExtR19FeType264PortExtR19SupportedCSIRSResourceExtListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceExtList_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceExtList_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceExtList_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.ProcessingCapability_r19, codebookParametersfeType2ExtR19FeType264PortExtR19ProcessingCapabilityR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberResource_r19, codebookParametersfeType2ExtR19FeType264PortExtR19MaxNumberResourceR19Constraints); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersfeType2ExtR19FeType264PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
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

	// 3. feType2-48PortExt-r19: sequence
	{
		if ie.FeType2_48PortExt_r19 != nil {
			c := ie.FeType2_48PortExt_r19
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersfeType2ExtR19FeType248PortExtR19SupportedCSIRSResourceExtListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceExtList_r19))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceExtList_r19 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceExtList_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.ProcessingCapability_r19, codebookParametersfeType2ExtR19FeType248PortExtR19ProcessingCapabilityR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberResource_r19, codebookParametersfeType2ExtR19FeType248PortExtR19MaxNumberResourceR19Constraints); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersfeType2ExtR19FeType248PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
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

	// 4. feType2-M2R1Ext-r19: sequence-of
	{
		if ie.FeType2_M2R1Ext_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersfeType2ExtR19FeType2M2R1ExtR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeType2_M2R1Ext_r19))); err != nil {
				return err
			}
			for i := range ie.FeType2_M2R1Ext_r19 {
				if err := e.EncodeInteger(ie.FeType2_M2R1Ext_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 5. feType2-M2R2Ext-r19: sequence-of
	{
		if ie.FeType2_M2R2Ext_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersfeType2ExtR19FeType2M2R2ExtR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeType2_M2R2Ext_r19))); err != nil {
				return err
			}
			for i := range ie.FeType2_M2R2Ext_r19 {
				if err := e.EncodeInteger(ie.FeType2_M2R2Ext_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 6. feType2-R3R4Ext-r19: enumerated
	{
		if ie.FeType2_R3R4Ext_r19 != nil {
			if err := e.EncodeEnumerated(*ie.FeType2_R3R4Ext_r19, codebookParametersfeType2ExtR19FeType2R3R4ExtR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CodebookParametersfeType2Ext_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParametersfeType2ExtR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. feType2-64PortExt-r19: sequence
	{
		{
			c := &ie.FeType2_64PortExt_r19
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersfeType2ExtR19FeType264PortExtR19SupportedCSIRSResourceExtListR19Constraints)
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
				v, err := d.DecodeEnumerated(codebookParametersfeType2ExtR19FeType264PortExtR19ProcessingCapabilityR19Constraints)
				if err != nil {
					return err
				}
				c.ProcessingCapability_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParametersfeType2ExtR19FeType264PortExtR19MaxNumberResourceR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberResource_r19 = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersfeType2ExtR19FeType264PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
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

	// 3. feType2-48PortExt-r19: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.FeType2_48PortExt_r19 = &struct {
				SupportedCSI_RS_ResourceExtList_r19   []int64
				ProcessingCapability_r19              int64
				MaxNumberResource_r19                 int64
				SupportedCSI_RS_ResourceListPerCC_r19 []int64
			}{}
			c := ie.FeType2_48PortExt_r19
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersfeType2ExtR19FeType248PortExtR19SupportedCSIRSResourceExtListR19Constraints)
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
				v, err := d.DecodeEnumerated(codebookParametersfeType2ExtR19FeType248PortExtR19ProcessingCapabilityR19Constraints)
				if err != nil {
					return err
				}
				c.ProcessingCapability_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(codebookParametersfeType2ExtR19FeType248PortExtR19MaxNumberResourceR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberResource_r19 = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersfeType2ExtR19FeType248PortExtR19SupportedCSIRSResourceListPerCCR19Constraints)
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

	// 4. feType2-M2R1Ext-r19: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersfeType2ExtR19FeType2M2R1ExtR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeType2_M2R1Ext_r19 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.FeType2_M2R1Ext_r19[i] = v
			}
		}
	}

	// 5. feType2-M2R2Ext-r19: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersfeType2ExtR19FeType2M2R2ExtR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeType2_M2R2Ext_r19 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.FeType2_M2R2Ext_r19[i] = v
			}
		}
	}

	// 6. feType2-R3R4Ext-r19: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(codebookParametersfeType2ExtR19FeType2R3R4ExtR19Constraints)
			if err != nil {
				return err
			}
			ie.FeType2_R3R4Ext_r19 = &idx
		}
	}

	return nil
}
