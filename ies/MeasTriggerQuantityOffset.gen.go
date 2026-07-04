// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasTriggerQuantityOffset (line 10069).

var measTriggerQuantityOffsetConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rsrp"},
		{Name: "rsrq"},
		{Name: "sinr"},
	},
}

const (
	MeasTriggerQuantityOffset_Rsrp = 0
	MeasTriggerQuantityOffset_Rsrq = 1
	MeasTriggerQuantityOffset_Sinr = 2
)

type MeasTriggerQuantityOffset struct {
	Choice int
	Rsrp   *int64
	Rsrq   *int64
	Sinr   *int64
}

func (ie *MeasTriggerQuantityOffset) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(measTriggerQuantityOffsetConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantityOffset_Rsrp:
		if err := e.EncodeInteger((*ie.Rsrp), per.Constrained(-30, 30)); err != nil {
			return err
		}
	case MeasTriggerQuantityOffset_Rsrq:
		if err := e.EncodeInteger((*ie.Rsrq), per.Constrained(-30, 30)); err != nil {
			return err
		}
	case MeasTriggerQuantityOffset_Sinr:
		if err := e.EncodeInteger((*ie.Sinr), per.Constrained(-30, 30)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "MeasTriggerQuantityOffset",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *MeasTriggerQuantityOffset) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(measTriggerQuantityOffsetConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "MeasTriggerQuantityOffset", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case MeasTriggerQuantityOffset_Rsrp:
		ie.Rsrp = new(int64)
		v, err := d.DecodeInteger(per.Constrained(-30, 30))
		if err != nil {
			return err
		}
		(*ie.Rsrp) = v
	case MeasTriggerQuantityOffset_Rsrq:
		ie.Rsrq = new(int64)
		v, err := d.DecodeInteger(per.Constrained(-30, 30))
		if err != nil {
			return err
		}
		(*ie.Rsrq) = v
	case MeasTriggerQuantityOffset_Sinr:
		ie.Sinr = new(int64)
		v, err := d.DecodeInteger(per.Constrained(-30, 30))
		if err != nil {
			return err
		}
		(*ie.Sinr) = v
	default:
		return &per.DecodeError{TypeName: "MeasTriggerQuantityOffset", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
