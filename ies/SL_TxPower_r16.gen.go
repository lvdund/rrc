// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-TxPower-r16 (line 28345).

var sLTxPowerR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "minusinfinity-r16"},
		{Name: "txPower-r16"},
	},
}

const (
	SL_TxPower_r16_Minusinfinity_r16 = 0
	SL_TxPower_r16_TxPower_r16       = 1
)

type SL_TxPower_r16 struct {
	Choice            int
	Minusinfinity_r16 *struct{}
	TxPower_r16       *int64
}

func (ie *SL_TxPower_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sLTxPowerR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_TxPower_r16_Minusinfinity_r16:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case SL_TxPower_r16_TxPower_r16:
		if err := e.EncodeInteger((*ie.TxPower_r16), per.Constrained(-30, 33)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SL-TxPower-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SL_TxPower_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sLTxPowerR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SL-TxPower-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SL_TxPower_r16_Minusinfinity_r16:
		ie.Minusinfinity_r16 = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case SL_TxPower_r16_TxPower_r16:
		ie.TxPower_r16 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(-30, 33))
		if err != nil {
			return err
		}
		(*ie.TxPower_r16) = v
	default:
		return &per.DecodeError{TypeName: "SL-TxPower-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
