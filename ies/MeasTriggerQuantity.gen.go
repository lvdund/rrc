// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasTriggerQuantity (line 10045).

var measTriggerQuantityConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rsrp"},
		{Name: "rsrq"},
		{Name: "sinr"},
	},
}

const (
	MeasTriggerQuantity_Rsrp = 0
	MeasTriggerQuantity_Rsrq = 1
	MeasTriggerQuantity_Sinr = 2
)

type MeasTriggerQuantity struct {
	Choice int
	Rsrp   *RSRP_Range
	Rsrq   *RSRQ_Range
	Sinr   *SINR_Range
}

func (ie *MeasTriggerQuantity) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(measTriggerQuantityConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantity_Rsrp:
		if err := ie.Rsrp.Encode(e); err != nil {
			return err
		}
	case MeasTriggerQuantity_Rsrq:
		if err := ie.Rsrq.Encode(e); err != nil {
			return err
		}
	case MeasTriggerQuantity_Sinr:
		if err := ie.Sinr.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "MeasTriggerQuantity",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *MeasTriggerQuantity) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(measTriggerQuantityConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "MeasTriggerQuantity", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case MeasTriggerQuantity_Rsrp:
		ie.Rsrp = new(RSRP_Range)
		if err := ie.Rsrp.Decode(d); err != nil {
			return err
		}
	case MeasTriggerQuantity_Rsrq:
		ie.Rsrq = new(RSRQ_Range)
		if err := ie.Rsrq.Decode(d); err != nil {
			return err
		}
	case MeasTriggerQuantity_Sinr:
		ie.Sinr = new(SINR_Range)
		if err := ie.Sinr.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "MeasTriggerQuantity", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
