// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParametersHybridBF-Type1SP-r19 (line 19278).

var codebookParametersHybridBFType1SPR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberCRI-Report-r19"},
		{Name: "supportedCSI-RS-ResourceHybridList-r19"},
		{Name: "maxValueKs-r19"},
	},
}

var codebookParametersHybridBFType1SPR19MaxNumberCRIReportR19Constraints = per.Constrained(1, 4)

var codebookParametersHybridBFType1SPR19SupportedCSIRSResourceHybridListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var codebookParametersHybridBFType1SPR19MaxValueKsR19Constraints = per.Constrained(2, 8)

type CodebookParametersHybridBF_Type1SP_r19 struct {
	MaxNumberCRI_Report_r19                int64
	SupportedCSI_RS_ResourceHybridList_r19 []int64
	MaxValueKs_r19                         int64
}

func (ie *CodebookParametersHybridBF_Type1SP_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParametersHybridBFType1SPR19Constraints)
	_ = seq

	// 1. maxNumberCRI-Report-r19: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberCRI_Report_r19, codebookParametersHybridBFType1SPR19MaxNumberCRIReportR19Constraints); err != nil {
			return err
		}
	}

	// 2. supportedCSI-RS-ResourceHybridList-r19: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(codebookParametersHybridBFType1SPR19SupportedCSIRSResourceHybridListR19Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.SupportedCSI_RS_ResourceHybridList_r19))); err != nil {
			return err
		}
		for i := range ie.SupportedCSI_RS_ResourceHybridList_r19 {
			if err := e.EncodeInteger(ie.SupportedCSI_RS_ResourceHybridList_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
				return err
			}
		}
	}

	// 3. maxValueKs-r19: integer
	{
		if err := e.EncodeInteger(ie.MaxValueKs_r19, codebookParametersHybridBFType1SPR19MaxValueKsR19Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CodebookParametersHybridBF_Type1SP_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParametersHybridBFType1SPR19Constraints)
	_ = seq

	// 1. maxNumberCRI-Report-r19: integer
	{
		v0, err := d.DecodeInteger(codebookParametersHybridBFType1SPR19MaxNumberCRIReportR19Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberCRI_Report_r19 = v0
	}

	// 2. supportedCSI-RS-ResourceHybridList-r19: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(codebookParametersHybridBFType1SPR19SupportedCSIRSResourceHybridListR19Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.SupportedCSI_RS_ResourceHybridList_r19 = make([]int64, n)
		for i := int64(0); i < n; i++ {
			v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
			if err != nil {
				return err
			}
			ie.SupportedCSI_RS_ResourceHybridList_r19[i] = v
		}
	}

	// 3. maxValueKs-r19: integer
	{
		v2, err := d.DecodeInteger(codebookParametersHybridBFType1SPR19MaxValueKsR19Constraints)
		if err != nil {
			return err
		}
		ie.MaxValueKs_r19 = v2
	}

	return nil
}
