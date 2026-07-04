// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ApplicabilityReport-r19 (line 5000).

var applicabilityReportR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "applicabilityCellId-r19"},
		{Name: "applicabilityInfoReportList-r19", Optional: true},
	},
}

var applicabilityReportR19ApplicabilityInfoReportListR19Constraints = per.SizeRange(1, common.MaxNrofApplicabilityReports_r19)

type ApplicabilityReport_r19 struct {
	ApplicabilityCellId_r19         ServCellIndex
	ApplicabilityInfoReportList_r19 []ApplicabilityInfoReport_r19
}

func (ie *ApplicabilityReport_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(applicabilityReportR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ApplicabilityInfoReportList_r19 != nil}); err != nil {
		return err
	}

	// 3. applicabilityCellId-r19: ref
	{
		if err := ie.ApplicabilityCellId_r19.Encode(e); err != nil {
			return err
		}
	}

	// 4. applicabilityInfoReportList-r19: sequence-of
	{
		if ie.ApplicabilityInfoReportList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(applicabilityReportR19ApplicabilityInfoReportListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.ApplicabilityInfoReportList_r19))); err != nil {
				return err
			}
			for i := range ie.ApplicabilityInfoReportList_r19 {
				if err := ie.ApplicabilityInfoReportList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *ApplicabilityReport_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(applicabilityReportR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. applicabilityCellId-r19: ref
	{
		if err := ie.ApplicabilityCellId_r19.Decode(d); err != nil {
			return err
		}
	}

	// 4. applicabilityInfoReportList-r19: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(applicabilityReportR19ApplicabilityInfoReportListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ApplicabilityInfoReportList_r19 = make([]ApplicabilityInfoReport_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ApplicabilityInfoReportList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
