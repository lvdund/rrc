// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CandidateBeamRS-r16 (line 5453).

var candidateBeamRSR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "candidateBeamConfig-r16"},
		{Name: "servingCellId", Optional: true},
	},
}

var candidateBeamRS_r16CandidateBeamConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-r16"},
		{Name: "csi-RS-r16"},
	},
}

const (
	CandidateBeamRS_r16_CandidateBeamConfig_r16_Ssb_r16    = 0
	CandidateBeamRS_r16_CandidateBeamConfig_r16_Csi_RS_r16 = 1
)

type CandidateBeamRS_r16 struct {
	CandidateBeamConfig_r16 struct {
		Choice     int
		Ssb_r16    *SSB_Index
		Csi_RS_r16 *NZP_CSI_RS_ResourceId
	}
	ServingCellId *ServCellIndex
}

func (ie *CandidateBeamRS_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(candidateBeamRSR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ServingCellId != nil}); err != nil {
		return err
	}

	// 2. candidateBeamConfig-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(candidateBeamRS_r16CandidateBeamConfigR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CandidateBeamConfig_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CandidateBeamConfig_r16.Choice {
		case CandidateBeamRS_r16_CandidateBeamConfig_r16_Ssb_r16:
			if err := ie.CandidateBeamConfig_r16.Ssb_r16.Encode(e); err != nil {
				return err
			}
		case CandidateBeamRS_r16_CandidateBeamConfig_r16_Csi_RS_r16:
			if err := ie.CandidateBeamConfig_r16.Csi_RS_r16.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CandidateBeamConfig_r16.Choice), Constraint: "index out of range"}
		}
	}

	// 3. servingCellId: ref
	{
		if ie.ServingCellId != nil {
			if err := ie.ServingCellId.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CandidateBeamRS_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(candidateBeamRSR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. candidateBeamConfig-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(candidateBeamRS_r16CandidateBeamConfigR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CandidateBeamConfig_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CandidateBeamRS_r16_CandidateBeamConfig_r16_Ssb_r16:
			ie.CandidateBeamConfig_r16.Ssb_r16 = new(SSB_Index)
			if err := ie.CandidateBeamConfig_r16.Ssb_r16.Decode(d); err != nil {
				return err
			}
		case CandidateBeamRS_r16_CandidateBeamConfig_r16_Csi_RS_r16:
			ie.CandidateBeamConfig_r16.Csi_RS_r16 = new(NZP_CSI_RS_ResourceId)
			if err := ie.CandidateBeamConfig_r16.Csi_RS_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. servingCellId: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ServingCellId = new(ServCellIndex)
			if err := ie.ServingCellId.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
