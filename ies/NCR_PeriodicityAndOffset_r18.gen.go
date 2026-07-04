// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NCR-PeriodicityAndOffset-r18 (line 10320).

var nCRPeriodicityAndOffsetR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "slot"},
		{Name: "ms"},
	},
}

const (
	NCR_PeriodicityAndOffset_r18_Slot = 0
	NCR_PeriodicityAndOffset_r18_Ms   = 1
)

type NCR_PeriodicityAndOffset_r18 struct {
	Choice int
	Slot   *NCR_SlotPeriodicityAndSlotOffset_r18
	Ms     *NCR_MsPeriodicityAndSlotOffset_r18
}

func (ie *NCR_PeriodicityAndOffset_r18) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(nCRPeriodicityAndOffsetR18Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case NCR_PeriodicityAndOffset_r18_Slot:
		if err := ie.Slot.Encode(e); err != nil {
			return err
		}
	case NCR_PeriodicityAndOffset_r18_Ms:
		if err := ie.Ms.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "NCR-PeriodicityAndOffset-r18",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *NCR_PeriodicityAndOffset_r18) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(nCRPeriodicityAndOffsetR18Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "NCR-PeriodicityAndOffset-r18", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case NCR_PeriodicityAndOffset_r18_Slot:
		ie.Slot = new(NCR_SlotPeriodicityAndSlotOffset_r18)
		if err := ie.Slot.Decode(d); err != nil {
			return err
		}
	case NCR_PeriodicityAndOffset_r18_Ms:
		ie.Ms = new(NCR_MsPeriodicityAndSlotOffset_r18)
		if err := ie.Ms.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "NCR-PeriodicityAndOffset-r18", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
