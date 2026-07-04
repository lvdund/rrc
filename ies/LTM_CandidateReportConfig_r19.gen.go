// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LTM-CandidateReportConfig-r19 (line 8883).

var lTMCandidateReportConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ltm-CandidateReportConfigId-r19"},
		{Name: "candidateSpecificOffset-r19", Optional: true},
	},
}

type LTM_CandidateReportConfig_r19 struct {
	Ltm_CandidateReportConfigId_r19 LTM_CandidateId_r18
	CandidateSpecificOffset_r19     *MeasTriggerQuantityOffset
}

func (ie *LTM_CandidateReportConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMCandidateReportConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CandidateSpecificOffset_r19 != nil}); err != nil {
		return err
	}

	// 3. ltm-CandidateReportConfigId-r19: ref
	{
		if err := ie.Ltm_CandidateReportConfigId_r19.Encode(e); err != nil {
			return err
		}
	}

	// 4. candidateSpecificOffset-r19: ref
	{
		if ie.CandidateSpecificOffset_r19 != nil {
			if err := ie.CandidateSpecificOffset_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LTM_CandidateReportConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMCandidateReportConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ltm-CandidateReportConfigId-r19: ref
	{
		if err := ie.Ltm_CandidateReportConfigId_r19.Decode(d); err != nil {
			return err
		}
	}

	// 4. candidateSpecificOffset-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.CandidateSpecificOffset_r19 = new(MeasTriggerQuantityOffset)
			if err := ie.CandidateSpecificOffset_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
