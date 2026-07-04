// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PathlossReferenceRS-Config (line 15388).

var pathlossReferenceRSConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-Index"},
		{Name: "csi-RS-Index"},
	},
}

const (
	PathlossReferenceRS_Config_Ssb_Index    = 0
	PathlossReferenceRS_Config_Csi_RS_Index = 1
)

type PathlossReferenceRS_Config struct {
	Choice       int
	Ssb_Index    *SSB_Index
	Csi_RS_Index *NZP_CSI_RS_ResourceId
}

func (ie *PathlossReferenceRS_Config) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(pathlossReferenceRSConfigConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case PathlossReferenceRS_Config_Ssb_Index:
		if err := ie.Ssb_Index.Encode(e); err != nil {
			return err
		}
	case PathlossReferenceRS_Config_Csi_RS_Index:
		if err := ie.Csi_RS_Index.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "PathlossReferenceRS-Config",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *PathlossReferenceRS_Config) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(pathlossReferenceRSConfigConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "PathlossReferenceRS-Config", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case PathlossReferenceRS_Config_Ssb_Index:
		ie.Ssb_Index = new(SSB_Index)
		if err := ie.Ssb_Index.Decode(d); err != nil {
			return err
		}
	case PathlossReferenceRS_Config_Csi_RS_Index:
		ie.Csi_RS_Index = new(NZP_CSI_RS_ResourceId)
		if err := ie.Csi_RS_Index.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "PathlossReferenceRS-Config", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
