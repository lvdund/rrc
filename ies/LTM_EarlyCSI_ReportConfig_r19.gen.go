// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LTM-EarlyCSI-ReportConfig-r19 (line 8894).

var lTMEarlyCSIReportConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ltm-ResourceForInterferenceMeasurements-r19", Optional: true},
		{Name: "ltm-CodebookConfig-r19", Optional: true},
		{Name: "reportQuantity-r19", Optional: true},
		{Name: "ltm-CQI-Table-r19", Optional: true},
	},
}

var lTM_EarlyCSI_ReportConfig_r19LtmCodebookConfigR19Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "twoToThirtyTwoPorts"},
		{Name: "moreThanThirtyTwoPorts"},
	},
}

const (
	LTM_EarlyCSI_ReportConfig_r19_Ltm_CodebookConfig_r19_TwoToThirtyTwoPorts    = 0
	LTM_EarlyCSI_ReportConfig_r19_Ltm_CodebookConfig_r19_MoreThanThirtyTwoPorts = 1
)

const (
	LTM_EarlyCSI_ReportConfig_r19_ReportQuantity_r19_Cri_RI_PMI_CQI = 0
)

var lTMEarlyCSIReportConfigR19ReportQuantityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_EarlyCSI_ReportConfig_r19_ReportQuantity_r19_Cri_RI_PMI_CQI},
}

type LTM_EarlyCSI_ReportConfig_r19 struct {
	Ltm_ResourceForInterferenceMeasurements_r19 *LTM_CSI_ResourceConfigId_r18
	Ltm_CodebookConfig_r19                      *struct {
		Choice                 int
		TwoToThirtyTwoPorts    *CodebookConfig
		MoreThanThirtyTwoPorts *CodebookConfig_r19
	}
	ReportQuantity_r19 *int64
	Ltm_CQI_Table_r19  *CQI_Table
}

func (ie *LTM_EarlyCSI_ReportConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMEarlyCSIReportConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ltm_ResourceForInterferenceMeasurements_r19 != nil, ie.Ltm_CodebookConfig_r19 != nil, ie.ReportQuantity_r19 != nil, ie.Ltm_CQI_Table_r19 != nil}); err != nil {
		return err
	}

	// 3. ltm-ResourceForInterferenceMeasurements-r19: ref
	{
		if ie.Ltm_ResourceForInterferenceMeasurements_r19 != nil {
			if err := ie.Ltm_ResourceForInterferenceMeasurements_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ltm-CodebookConfig-r19: choice
	{
		if ie.Ltm_CodebookConfig_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lTM_EarlyCSI_ReportConfig_r19LtmCodebookConfigR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Ltm_CodebookConfig_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Ltm_CodebookConfig_r19).Choice {
			case LTM_EarlyCSI_ReportConfig_r19_Ltm_CodebookConfig_r19_TwoToThirtyTwoPorts:
				if err := (*ie.Ltm_CodebookConfig_r19).TwoToThirtyTwoPorts.Encode(e); err != nil {
					return err
				}
			case LTM_EarlyCSI_ReportConfig_r19_Ltm_CodebookConfig_r19_MoreThanThirtyTwoPorts:
				if err := (*ie.Ltm_CodebookConfig_r19).MoreThanThirtyTwoPorts.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Ltm_CodebookConfig_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. reportQuantity-r19: enumerated
	{
		if ie.ReportQuantity_r19 != nil {
			if err := e.EncodeEnumerated(*ie.ReportQuantity_r19, lTMEarlyCSIReportConfigR19ReportQuantityR19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. ltm-CQI-Table-r19: ref
	{
		if ie.Ltm_CQI_Table_r19 != nil {
			if err := ie.Ltm_CQI_Table_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LTM_EarlyCSI_ReportConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMEarlyCSIReportConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ltm-ResourceForInterferenceMeasurements-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ltm_ResourceForInterferenceMeasurements_r19 = new(LTM_CSI_ResourceConfigId_r18)
			if err := ie.Ltm_ResourceForInterferenceMeasurements_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ltm-CodebookConfig-r19: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Ltm_CodebookConfig_r19 = &struct {
				Choice                 int
				TwoToThirtyTwoPorts    *CodebookConfig
				MoreThanThirtyTwoPorts *CodebookConfig_r19
			}{}
			choiceDec := d.NewChoiceDecoder(lTM_EarlyCSI_ReportConfig_r19LtmCodebookConfigR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ltm_CodebookConfig_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LTM_EarlyCSI_ReportConfig_r19_Ltm_CodebookConfig_r19_TwoToThirtyTwoPorts:
				(*ie.Ltm_CodebookConfig_r19).TwoToThirtyTwoPorts = new(CodebookConfig)
				if err := (*ie.Ltm_CodebookConfig_r19).TwoToThirtyTwoPorts.Decode(d); err != nil {
					return err
				}
			case LTM_EarlyCSI_ReportConfig_r19_Ltm_CodebookConfig_r19_MoreThanThirtyTwoPorts:
				(*ie.Ltm_CodebookConfig_r19).MoreThanThirtyTwoPorts = new(CodebookConfig_r19)
				if err := (*ie.Ltm_CodebookConfig_r19).MoreThanThirtyTwoPorts.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. reportQuantity-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(lTMEarlyCSIReportConfigR19ReportQuantityR19Constraints)
			if err != nil {
				return err
			}
			ie.ReportQuantity_r19 = &idx
		}
	}

	// 6. ltm-CQI-Table-r19: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Ltm_CQI_Table_r19 = new(CQI_Table)
			if err := ie.Ltm_CQI_Table_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
