// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ApplicabilityInfoReport-r19 (line 5006).

var applicabilityInfoReportR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "applicabilityInfoReportId-r19", Optional: true},
		{Name: "applicabilityStatus-r19"},
		{Name: "releaseConfigurationPreference-r19", Optional: true},
	},
}

var applicabilityInfoReport_r19ApplicabilityInfoReportIdR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "csi-ReportConfigId-r19"},
		{Name: "applicabilitySetId-r19"},
	},
}

const (
	ApplicabilityInfoReport_r19_ApplicabilityInfoReportId_r19_Csi_ReportConfigId_r19 = 0
	ApplicabilityInfoReport_r19_ApplicabilityInfoReportId_r19_ApplicabilitySetId_r19 = 1
)

const (
	ApplicabilityInfoReport_r19_ApplicabilityStatus_r19_Applicable   = 0
	ApplicabilityInfoReport_r19_ApplicabilityStatus_r19_Inapplicable = 1
)

var applicabilityInfoReportR19ApplicabilityStatusR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ApplicabilityInfoReport_r19_ApplicabilityStatus_r19_Applicable, ApplicabilityInfoReport_r19_ApplicabilityStatus_r19_Inapplicable},
}

const (
	ApplicabilityInfoReport_r19_ReleaseConfigurationPreference_r19_True = 0
)

var applicabilityInfoReportR19ReleaseConfigurationPreferenceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ApplicabilityInfoReport_r19_ReleaseConfigurationPreference_r19_True},
}

type ApplicabilityInfoReport_r19 struct {
	ApplicabilityInfoReportId_r19 *struct {
		Choice                 int
		Csi_ReportConfigId_r19 *CSI_ReportConfigId
		ApplicabilitySetId_r19 *ApplicabilitySetConfigId_r19
	}
	ApplicabilityStatus_r19            int64
	ReleaseConfigurationPreference_r19 *int64
}

func (ie *ApplicabilityInfoReport_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(applicabilityInfoReportR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ApplicabilityInfoReportId_r19 != nil, ie.ReleaseConfigurationPreference_r19 != nil}); err != nil {
		return err
	}

	// 3. applicabilityInfoReportId-r19: choice
	{
		if ie.ApplicabilityInfoReportId_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(applicabilityInfoReport_r19ApplicabilityInfoReportIdR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.ApplicabilityInfoReportId_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.ApplicabilityInfoReportId_r19).Choice {
			case ApplicabilityInfoReport_r19_ApplicabilityInfoReportId_r19_Csi_ReportConfigId_r19:
				if err := (*ie.ApplicabilityInfoReportId_r19).Csi_ReportConfigId_r19.Encode(e); err != nil {
					return err
				}
			case ApplicabilityInfoReport_r19_ApplicabilityInfoReportId_r19_ApplicabilitySetId_r19:
				if err := (*ie.ApplicabilityInfoReportId_r19).ApplicabilitySetId_r19.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.ApplicabilityInfoReportId_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. applicabilityStatus-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.ApplicabilityStatus_r19, applicabilityInfoReportR19ApplicabilityStatusR19Constraints); err != nil {
			return err
		}
	}

	// 5. releaseConfigurationPreference-r19: enumerated
	{
		if ie.ReleaseConfigurationPreference_r19 != nil {
			if err := e.EncodeEnumerated(*ie.ReleaseConfigurationPreference_r19, applicabilityInfoReportR19ReleaseConfigurationPreferenceR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ApplicabilityInfoReport_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(applicabilityInfoReportR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. applicabilityInfoReportId-r19: choice
	{
		if seq.IsComponentPresent(0) {
			ie.ApplicabilityInfoReportId_r19 = &struct {
				Choice                 int
				Csi_ReportConfigId_r19 *CSI_ReportConfigId
				ApplicabilitySetId_r19 *ApplicabilitySetConfigId_r19
			}{}
			choiceDec := d.NewChoiceDecoder(applicabilityInfoReport_r19ApplicabilityInfoReportIdR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ApplicabilityInfoReportId_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case ApplicabilityInfoReport_r19_ApplicabilityInfoReportId_r19_Csi_ReportConfigId_r19:
				(*ie.ApplicabilityInfoReportId_r19).Csi_ReportConfigId_r19 = new(CSI_ReportConfigId)
				if err := (*ie.ApplicabilityInfoReportId_r19).Csi_ReportConfigId_r19.Decode(d); err != nil {
					return err
				}
			case ApplicabilityInfoReport_r19_ApplicabilityInfoReportId_r19_ApplicabilitySetId_r19:
				(*ie.ApplicabilityInfoReportId_r19).ApplicabilitySetId_r19 = new(ApplicabilitySetConfigId_r19)
				if err := (*ie.ApplicabilityInfoReportId_r19).ApplicabilitySetId_r19.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. applicabilityStatus-r19: enumerated
	{
		v1, err := d.DecodeEnumerated(applicabilityInfoReportR19ApplicabilityStatusR19Constraints)
		if err != nil {
			return err
		}
		ie.ApplicabilityStatus_r19 = v1
	}

	// 5. releaseConfigurationPreference-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(applicabilityInfoReportR19ReleaseConfigurationPreferenceR19Constraints)
			if err != nil {
				return err
			}
			ie.ReleaseConfigurationPreference_r19 = &idx
		}
	}

	return nil
}
