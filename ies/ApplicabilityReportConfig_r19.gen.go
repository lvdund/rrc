// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ApplicabilityReportConfig-r19 (line 26537).

var applicabilityReportConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "applicabilityConfigToAddModList-r19", Optional: true},
		{Name: "applicabilityConfigToReleaseList-r19", Optional: true},
		{Name: "disableApplicabilityCSI-ReportConfig-r19", Optional: true},
	},
}

var applicabilityReportConfigR19ApplicabilityConfigToAddModListR19Constraints = per.SizeRange(1, common.MaxNrofServingCells)

var applicabilityReportConfigR19ApplicabilityConfigToReleaseListR19Constraints = per.SizeRange(1, common.MaxNrofServingCells)

const (
	ApplicabilityReportConfig_r19_DisableApplicabilityCSI_ReportConfig_r19_True = 0
)

var applicabilityReportConfigR19DisableApplicabilityCSIReportConfigR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ApplicabilityReportConfig_r19_DisableApplicabilityCSI_ReportConfig_r19_True},
}

type ApplicabilityReportConfig_r19 struct {
	ApplicabilityConfigToAddModList_r19      []ApplicabilityConfig_r19
	ApplicabilityConfigToReleaseList_r19     []ServCellIndex
	DisableApplicabilityCSI_ReportConfig_r19 *int64
}

func (ie *ApplicabilityReportConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(applicabilityReportConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ApplicabilityConfigToAddModList_r19 != nil, ie.ApplicabilityConfigToReleaseList_r19 != nil, ie.DisableApplicabilityCSI_ReportConfig_r19 != nil}); err != nil {
		return err
	}

	// 3. applicabilityConfigToAddModList-r19: sequence-of
	{
		if ie.ApplicabilityConfigToAddModList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(applicabilityReportConfigR19ApplicabilityConfigToAddModListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.ApplicabilityConfigToAddModList_r19))); err != nil {
				return err
			}
			for i := range ie.ApplicabilityConfigToAddModList_r19 {
				if err := ie.ApplicabilityConfigToAddModList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. applicabilityConfigToReleaseList-r19: sequence-of
	{
		if ie.ApplicabilityConfigToReleaseList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(applicabilityReportConfigR19ApplicabilityConfigToReleaseListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.ApplicabilityConfigToReleaseList_r19))); err != nil {
				return err
			}
			for i := range ie.ApplicabilityConfigToReleaseList_r19 {
				if err := ie.ApplicabilityConfigToReleaseList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. disableApplicabilityCSI-ReportConfig-r19: enumerated
	{
		if ie.DisableApplicabilityCSI_ReportConfig_r19 != nil {
			if err := e.EncodeEnumerated(*ie.DisableApplicabilityCSI_ReportConfig_r19, applicabilityReportConfigR19DisableApplicabilityCSIReportConfigR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ApplicabilityReportConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(applicabilityReportConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. applicabilityConfigToAddModList-r19: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(applicabilityReportConfigR19ApplicabilityConfigToAddModListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ApplicabilityConfigToAddModList_r19 = make([]ApplicabilityConfig_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ApplicabilityConfigToAddModList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. applicabilityConfigToReleaseList-r19: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(applicabilityReportConfigR19ApplicabilityConfigToReleaseListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ApplicabilityConfigToReleaseList_r19 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ApplicabilityConfigToReleaseList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. disableApplicabilityCSI-ReportConfig-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(applicabilityReportConfigR19DisableApplicabilityCSIReportConfigR19Constraints)
			if err != nil {
				return err
			}
			ie.DisableApplicabilityCSI_ReportConfig_r19 = &idx
		}
	}

	return nil
}
