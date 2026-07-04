// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IDC-AssistanceConfig-r16 (line 26427).

var iDCAssistanceConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "candidateServingFreqListNR-r16", Optional: true},
	},
}

type IDC_AssistanceConfig_r16 struct {
	CandidateServingFreqListNR_r16 *CandidateServingFreqListNR_r16
}

func (ie *IDC_AssistanceConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iDCAssistanceConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CandidateServingFreqListNR_r16 != nil}); err != nil {
		return err
	}

	// 3. candidateServingFreqListNR-r16: ref
	{
		if ie.CandidateServingFreqListNR_r16 != nil {
			if err := ie.CandidateServingFreqListNR_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IDC_AssistanceConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iDCAssistanceConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. candidateServingFreqListNR-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.CandidateServingFreqListNR_r16 = new(CandidateServingFreqListNR_r16)
			if err := ie.CandidateServingFreqListNR_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
