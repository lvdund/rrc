// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasTriggerQuantityEUTRA (line 10054).

var measTriggerQuantityEUTRAConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rsrp"},
		{Name: "rsrq"},
		{Name: "sinr"},
	},
}

const (
	MeasTriggerQuantityEUTRA_Rsrp = 0
	MeasTriggerQuantityEUTRA_Rsrq = 1
	MeasTriggerQuantityEUTRA_Sinr = 2
)

type MeasTriggerQuantityEUTRA struct {
	Choice int
	Rsrp   *RSRP_RangeEUTRA
	Rsrq   *RSRQ_RangeEUTRA
	Sinr   *SINR_RangeEUTRA
}

func (ie *MeasTriggerQuantityEUTRA) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(measTriggerQuantityEUTRAConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantityEUTRA_Rsrp:
		if err := ie.Rsrp.Encode(e); err != nil {
			return err
		}
	case MeasTriggerQuantityEUTRA_Rsrq:
		if err := ie.Rsrq.Encode(e); err != nil {
			return err
		}
	case MeasTriggerQuantityEUTRA_Sinr:
		if err := ie.Sinr.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "MeasTriggerQuantityEUTRA",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *MeasTriggerQuantityEUTRA) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(measTriggerQuantityEUTRAConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "MeasTriggerQuantityEUTRA", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case MeasTriggerQuantityEUTRA_Rsrp:
		ie.Rsrp = new(RSRP_RangeEUTRA)
		if err := ie.Rsrp.Decode(d); err != nil {
			return err
		}
	case MeasTriggerQuantityEUTRA_Rsrq:
		ie.Rsrq = new(RSRQ_RangeEUTRA)
		if err := ie.Rsrq.Decode(d); err != nil {
			return err
		}
	case MeasTriggerQuantityEUTRA_Sinr:
		ie.Sinr = new(SINR_RangeEUTRA)
		if err := ie.Sinr.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "MeasTriggerQuantityEUTRA", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
