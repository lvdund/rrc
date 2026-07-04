// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CarrierState-r17 (line 5797).

var carrierStateR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "deActivated-r17"},
		{Name: "activeBWP-r17"},
	},
}

const (
	CarrierState_r17_DeActivated_r17 = 0
	CarrierState_r17_ActiveBWP_r17   = 1
)

type CarrierState_r17 struct {
	Choice          int
	DeActivated_r17 *struct{}
	ActiveBWP_r17   *int64
}

func (ie *CarrierState_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(carrierStateR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case CarrierState_r17_DeActivated_r17:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case CarrierState_r17_ActiveBWP_r17:
		if err := e.EncodeInteger((*ie.ActiveBWP_r17), per.Constrained(0, common.MaxNrofBWPs)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "CarrierState-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *CarrierState_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(carrierStateR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "CarrierState-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case CarrierState_r17_DeActivated_r17:
		ie.DeActivated_r17 = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case CarrierState_r17_ActiveBWP_r17:
		ie.ActiveBWP_r17 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofBWPs))
		if err != nil {
			return err
		}
		(*ie.ActiveBWP_r17) = v
	default:
		return &per.DecodeError{TypeName: "CarrierState-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
