// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-MeasTriggerQuantity-r16 (line 27929).

var sLMeasTriggerQuantityR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-RSRP-r16"},
	},
}

const (
	SL_MeasTriggerQuantity_r16_Sl_RSRP_r16 = 0
)

type SL_MeasTriggerQuantity_r16 struct {
	Choice      int
	Sl_RSRP_r16 *RSRP_Range
}

func (ie *SL_MeasTriggerQuantity_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sLMeasTriggerQuantityR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_MeasTriggerQuantity_r16_Sl_RSRP_r16:
		if err := ie.Sl_RSRP_r16.Encode(e); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "SL-MeasTriggerQuantity-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SL_MeasTriggerQuantity_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sLMeasTriggerQuantityR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SL-MeasTriggerQuantity-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SL_MeasTriggerQuantity_r16_Sl_RSRP_r16:
		ie.Sl_RSRP_r16 = new(RSRP_Range)
		if err := ie.Sl_RSRP_r16.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "SL-MeasTriggerQuantity-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
