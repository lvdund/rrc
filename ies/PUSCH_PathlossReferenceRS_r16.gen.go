// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUSCH-PathlossReferenceRS-r16 (line 12600).

var pUSCHPathlossReferenceRSR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pusch-PathlossReferenceRS-Id-r16"},
		{Name: "referenceSignal-r16"},
	},
}

var pUSCH_PathlossReferenceRS_r16ReferenceSignalR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-Index-r16"},
		{Name: "csi-RS-Index-r16"},
	},
}

const (
	PUSCH_PathlossReferenceRS_r16_ReferenceSignal_r16_Ssb_Index_r16    = 0
	PUSCH_PathlossReferenceRS_r16_ReferenceSignal_r16_Csi_RS_Index_r16 = 1
)

type PUSCH_PathlossReferenceRS_r16 struct {
	Pusch_PathlossReferenceRS_Id_r16 PUSCH_PathlossReferenceRS_Id_v1610
	ReferenceSignal_r16              struct {
		Choice           int
		Ssb_Index_r16    *SSB_Index
		Csi_RS_Index_r16 *NZP_CSI_RS_ResourceId
	}
}

func (ie *PUSCH_PathlossReferenceRS_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHPathlossReferenceRSR16Constraints)
	_ = seq

	// 1. pusch-PathlossReferenceRS-Id-r16: ref
	{
		if err := ie.Pusch_PathlossReferenceRS_Id_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. referenceSignal-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(pUSCH_PathlossReferenceRS_r16ReferenceSignalR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReferenceSignal_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReferenceSignal_r16.Choice {
		case PUSCH_PathlossReferenceRS_r16_ReferenceSignal_r16_Ssb_Index_r16:
			if err := ie.ReferenceSignal_r16.Ssb_Index_r16.Encode(e); err != nil {
				return err
			}
		case PUSCH_PathlossReferenceRS_r16_ReferenceSignal_r16_Csi_RS_Index_r16:
			if err := ie.ReferenceSignal_r16.Csi_RS_Index_r16.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReferenceSignal_r16.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *PUSCH_PathlossReferenceRS_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHPathlossReferenceRSR16Constraints)
	_ = seq

	// 1. pusch-PathlossReferenceRS-Id-r16: ref
	{
		if err := ie.Pusch_PathlossReferenceRS_Id_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. referenceSignal-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(pUSCH_PathlossReferenceRS_r16ReferenceSignalR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReferenceSignal_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case PUSCH_PathlossReferenceRS_r16_ReferenceSignal_r16_Ssb_Index_r16:
			ie.ReferenceSignal_r16.Ssb_Index_r16 = new(SSB_Index)
			if err := ie.ReferenceSignal_r16.Ssb_Index_r16.Decode(d); err != nil {
				return err
			}
		case PUSCH_PathlossReferenceRS_r16_ReferenceSignal_r16_Csi_RS_Index_r16:
			ie.ReferenceSignal_r16.Csi_RS_Index_r16 = new(NZP_CSI_RS_ResourceId)
			if err := ie.ReferenceSignal_r16.Csi_RS_Index_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
