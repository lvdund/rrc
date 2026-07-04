// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IDC-FDM-AssistanceConfig-r18 (line 26502).

var iDCFDMAssistanceConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "candidateServingFreqRangeListNR-r18", Optional: true},
	},
}

type IDC_FDM_AssistanceConfig_r18 struct {
	CandidateServingFreqRangeListNR_r18 *CandidateServingFreqRangeListNR_r18
}

func (ie *IDC_FDM_AssistanceConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iDCFDMAssistanceConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CandidateServingFreqRangeListNR_r18 != nil}); err != nil {
		return err
	}

	// 3. candidateServingFreqRangeListNR-r18: ref
	{
		if ie.CandidateServingFreqRangeListNR_r18 != nil {
			if err := ie.CandidateServingFreqRangeListNR_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IDC_FDM_AssistanceConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iDCFDMAssistanceConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. candidateServingFreqRangeListNR-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.CandidateServingFreqRangeListNR_r18 = new(CandidateServingFreqRangeListNR_r18)
			if err := ie.CandidateServingFreqRangeListNR_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
