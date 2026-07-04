// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PRACH-ResourceDedicatedBFR (line 5135).

var pRACHResourceDedicatedBFRConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb"},
		{Name: "csi-RS"},
	},
}

const (
	PRACH_ResourceDedicatedBFR_Ssb    = 0
	PRACH_ResourceDedicatedBFR_Csi_RS = 1
)

type PRACH_ResourceDedicatedBFR struct {
	Choice int
	Ssb    *BFR_SSB_Resource
	Csi_RS *BFR_CSIRS_Resource
}

func (ie *PRACH_ResourceDedicatedBFR) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(pRACHResourceDedicatedBFRConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case PRACH_ResourceDedicatedBFR_Ssb:
		if err := ie.Ssb.Encode(e); err != nil {
			return err
		}
	case PRACH_ResourceDedicatedBFR_Csi_RS:
		if err := ie.Csi_RS.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "PRACH-ResourceDedicatedBFR",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *PRACH_ResourceDedicatedBFR) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(pRACHResourceDedicatedBFRConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "PRACH-ResourceDedicatedBFR", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case PRACH_ResourceDedicatedBFR_Ssb:
		ie.Ssb = new(BFR_SSB_Resource)
		if err := ie.Ssb.Decode(d); err != nil {
			return err
		}
	case PRACH_ResourceDedicatedBFR_Csi_RS:
		ie.Csi_RS = new(BFR_CSIRS_Resource)
		if err := ie.Csi_RS.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "PRACH-ResourceDedicatedBFR", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
