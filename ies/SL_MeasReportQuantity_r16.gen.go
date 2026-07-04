// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-MeasReportQuantity-r16 (line 27924).

var sLMeasReportQuantityR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-RSRP-r16"},
	},
}

const (
	SL_MeasReportQuantity_r16_Sl_RSRP_r16 = 0
)

type SL_MeasReportQuantity_r16 struct {
	Choice      int
	Sl_RSRP_r16 *bool
}

func (ie *SL_MeasReportQuantity_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sLMeasReportQuantityR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_MeasReportQuantity_r16_Sl_RSRP_r16:
		if err := e.EncodeBoolean((*ie.Sl_RSRP_r16)); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "SL-MeasReportQuantity-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SL_MeasReportQuantity_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sLMeasReportQuantityR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SL-MeasReportQuantity-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SL_MeasReportQuantity_r16_Sl_RSRP_r16:
		ie.Sl_RSRP_r16 = new(bool)
		v, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		(*ie.Sl_RSRP_r16) = v
	default:
		return &per.DecodeError{TypeName: "SL-MeasReportQuantity-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
