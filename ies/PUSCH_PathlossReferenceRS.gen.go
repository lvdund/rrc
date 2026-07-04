// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUSCH-PathlossReferenceRS (line 12592).

var pUSCHPathlossReferenceRSConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pusch-PathlossReferenceRS-Id"},
		{Name: "referenceSignal"},
	},
}

var pUSCH_PathlossReferenceRSReferenceSignalConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-Index"},
		{Name: "csi-RS-Index"},
	},
}

const (
	PUSCH_PathlossReferenceRS_ReferenceSignal_Ssb_Index    = 0
	PUSCH_PathlossReferenceRS_ReferenceSignal_Csi_RS_Index = 1
)

type PUSCH_PathlossReferenceRS struct {
	Pusch_PathlossReferenceRS_Id PUSCH_PathlossReferenceRS_Id
	ReferenceSignal              struct {
		Choice       int
		Ssb_Index    *SSB_Index
		Csi_RS_Index *NZP_CSI_RS_ResourceId
	}
}

func (ie *PUSCH_PathlossReferenceRS) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHPathlossReferenceRSConstraints)
	_ = seq

	// 1. pusch-PathlossReferenceRS-Id: ref
	{
		if err := ie.Pusch_PathlossReferenceRS_Id.Encode(e); err != nil {
			return err
		}
	}

	// 2. referenceSignal: choice
	{
		choiceEnc := e.NewChoiceEncoder(pUSCH_PathlossReferenceRSReferenceSignalConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReferenceSignal.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReferenceSignal.Choice {
		case PUSCH_PathlossReferenceRS_ReferenceSignal_Ssb_Index:
			if err := ie.ReferenceSignal.Ssb_Index.Encode(e); err != nil {
				return err
			}
		case PUSCH_PathlossReferenceRS_ReferenceSignal_Csi_RS_Index:
			if err := ie.ReferenceSignal.Csi_RS_Index.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReferenceSignal.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *PUSCH_PathlossReferenceRS) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHPathlossReferenceRSConstraints)
	_ = seq

	// 1. pusch-PathlossReferenceRS-Id: ref
	{
		if err := ie.Pusch_PathlossReferenceRS_Id.Decode(d); err != nil {
			return err
		}
	}

	// 2. referenceSignal: choice
	{
		choiceDec := d.NewChoiceDecoder(pUSCH_PathlossReferenceRSReferenceSignalConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReferenceSignal.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case PUSCH_PathlossReferenceRS_ReferenceSignal_Ssb_Index:
			ie.ReferenceSignal.Ssb_Index = new(SSB_Index)
			if err := ie.ReferenceSignal.Ssb_Index.Decode(d); err != nil {
				return err
			}
		case PUSCH_PathlossReferenceRS_ReferenceSignal_Csi_RS_Index:
			ie.ReferenceSignal.Csi_RS_Index = new(NZP_CSI_RS_ResourceId)
			if err := ie.ReferenceSignal.Csi_RS_Index.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
