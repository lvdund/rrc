// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasTriggerQuantityUTRA-FDD-r16 (line 13509).

var measTriggerQuantityUTRAFDDR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "utra-FDD-RSCP-r16"},
		{Name: "utra-FDD-EcN0-r16"},
	},
}

const (
	MeasTriggerQuantityUTRA_FDD_r16_Utra_FDD_RSCP_r16 = 0
	MeasTriggerQuantityUTRA_FDD_r16_Utra_FDD_EcN0_r16 = 1
)

type MeasTriggerQuantityUTRA_FDD_r16 struct {
	Choice            int
	Utra_FDD_RSCP_r16 *int64
	Utra_FDD_EcN0_r16 *int64
}

func (ie *MeasTriggerQuantityUTRA_FDD_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(measTriggerQuantityUTRAFDDR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantityUTRA_FDD_r16_Utra_FDD_RSCP_r16:
		if err := e.EncodeInteger((*ie.Utra_FDD_RSCP_r16), per.Constrained(-5, 91)); err != nil {
			return err
		}
	case MeasTriggerQuantityUTRA_FDD_r16_Utra_FDD_EcN0_r16:
		if err := e.EncodeInteger((*ie.Utra_FDD_EcN0_r16), per.Constrained(0, 49)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "MeasTriggerQuantityUTRA-FDD-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *MeasTriggerQuantityUTRA_FDD_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(measTriggerQuantityUTRAFDDR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "MeasTriggerQuantityUTRA-FDD-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case MeasTriggerQuantityUTRA_FDD_r16_Utra_FDD_RSCP_r16:
		ie.Utra_FDD_RSCP_r16 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(-5, 91))
		if err != nil {
			return err
		}
		(*ie.Utra_FDD_RSCP_r16) = v
	case MeasTriggerQuantityUTRA_FDD_r16_Utra_FDD_EcN0_r16:
		ie.Utra_FDD_EcN0_r16 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 49))
		if err != nil {
			return err
		}
		(*ie.Utra_FDD_EcN0_r16) = v
	default:
		return &per.DecodeError{TypeName: "MeasTriggerQuantityUTRA-FDD-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
