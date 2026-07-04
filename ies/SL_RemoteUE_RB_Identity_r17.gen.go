// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RemoteUE-RB-Identity-r17 (line 28280).

var sLRemoteUERBIdentityR17Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "srb-Identity-r17"},
		{Name: "drb-Identity-r17"},
	},
}

const (
	SL_RemoteUE_RB_Identity_r17_Srb_Identity_r17 = 0
	SL_RemoteUE_RB_Identity_r17_Drb_Identity_r17 = 1
)

type SL_RemoteUE_RB_Identity_r17 struct {
	Choice           int
	Srb_Identity_r17 *int64
	Drb_Identity_r17 *DRB_Identity
}

func (ie *SL_RemoteUE_RB_Identity_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sLRemoteUERBIdentityR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_RemoteUE_RB_Identity_r17_Srb_Identity_r17:
		if err := e.EncodeInteger((*ie.Srb_Identity_r17), per.Constrained(0, 3)); err != nil {
			return err
		}
	case SL_RemoteUE_RB_Identity_r17_Drb_Identity_r17:
		if err := ie.Drb_Identity_r17.Encode(e); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "SL-RemoteUE-RB-Identity-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SL_RemoteUE_RB_Identity_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sLRemoteUERBIdentityR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SL-RemoteUE-RB-Identity-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SL_RemoteUE_RB_Identity_r17_Srb_Identity_r17:
		ie.Srb_Identity_r17 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 3))
		if err != nil {
			return err
		}
		(*ie.Srb_Identity_r17) = v
	case SL_RemoteUE_RB_Identity_r17_Drb_Identity_r17:
		ie.Drb_Identity_r17 = new(DRB_Identity)
		if err := ie.Drb_Identity_r17.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "SL-RemoteUE-RB-Identity-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
