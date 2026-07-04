// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: HighSpeedParameters-v1650 (line 20843).

var highSpeedParametersV1650Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "intraNR-MeasurementEnhancement-r16"},
		{Name: "interRAT-MeasurementEnhancement-r16"},
	},
}

const (
	HighSpeedParameters_v1650_IntraNR_MeasurementEnhancement_r16  = 0
	HighSpeedParameters_v1650_InterRAT_MeasurementEnhancement_r16 = 1
)

const (
	HighSpeedParameters_v1650_IntraNR_MeasurementEnhancement_r16_Supported = 0
)

var highSpeedParametersV1650IntraNRMeasurementEnhancementR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{HighSpeedParameters_v1650_IntraNR_MeasurementEnhancement_r16_Supported},
}

const (
	HighSpeedParameters_v1650_InterRAT_MeasurementEnhancement_r16_Supported = 0
)

var highSpeedParametersV1650InterRATMeasurementEnhancementR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{HighSpeedParameters_v1650_InterRAT_MeasurementEnhancement_r16_Supported},
}

type HighSpeedParameters_v1650 struct {
	Choice                              int
	IntraNR_MeasurementEnhancement_r16  *int64
	InterRAT_MeasurementEnhancement_r16 *int64
}

func (ie *HighSpeedParameters_v1650) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(highSpeedParametersV1650Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case HighSpeedParameters_v1650_IntraNR_MeasurementEnhancement_r16:
		if err := e.EncodeEnumerated((*ie.IntraNR_MeasurementEnhancement_r16), highSpeedParametersV1650IntraNRMeasurementEnhancementR16Constraints); err != nil {
			return err
		}
	case HighSpeedParameters_v1650_InterRAT_MeasurementEnhancement_r16:
		if err := e.EncodeEnumerated((*ie.InterRAT_MeasurementEnhancement_r16), highSpeedParametersV1650InterRATMeasurementEnhancementR16Constraints); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "HighSpeedParameters-v1650",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *HighSpeedParameters_v1650) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(highSpeedParametersV1650Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "HighSpeedParameters-v1650", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case HighSpeedParameters_v1650_IntraNR_MeasurementEnhancement_r16:
		ie.IntraNR_MeasurementEnhancement_r16 = new(int64)
		v, err := d.DecodeEnumerated(highSpeedParametersV1650IntraNRMeasurementEnhancementR16Constraints)
		if err != nil {
			return err
		}
		(*ie.IntraNR_MeasurementEnhancement_r16) = v
	case HighSpeedParameters_v1650_InterRAT_MeasurementEnhancement_r16:
		ie.InterRAT_MeasurementEnhancement_r16 = new(int64)
		v, err := d.DecodeEnumerated(highSpeedParametersV1650InterRATMeasurementEnhancementR16Constraints)
		if err != nil {
			return err
		}
		(*ie.InterRAT_MeasurementEnhancement_r16) = v
	default:
		return &per.DecodeError{TypeName: "HighSpeedParameters-v1650", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
