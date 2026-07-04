// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: HighSpeedParameters-r16 (line 20838).

var highSpeedParametersR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measurementEnhancement-r16", Optional: true},
		{Name: "demodulationEnhancement-r16", Optional: true},
	},
}

const (
	HighSpeedParameters_r16_MeasurementEnhancement_r16_Supported = 0
)

var highSpeedParametersR16MeasurementEnhancementR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{HighSpeedParameters_r16_MeasurementEnhancement_r16_Supported},
}

const (
	HighSpeedParameters_r16_DemodulationEnhancement_r16_Supported = 0
)

var highSpeedParametersR16DemodulationEnhancementR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{HighSpeedParameters_r16_DemodulationEnhancement_r16_Supported},
}

type HighSpeedParameters_r16 struct {
	MeasurementEnhancement_r16  *int64
	DemodulationEnhancement_r16 *int64
}

func (ie *HighSpeedParameters_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(highSpeedParametersR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasurementEnhancement_r16 != nil, ie.DemodulationEnhancement_r16 != nil}); err != nil {
		return err
	}

	// 2. measurementEnhancement-r16: enumerated
	{
		if ie.MeasurementEnhancement_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MeasurementEnhancement_r16, highSpeedParametersR16MeasurementEnhancementR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. demodulationEnhancement-r16: enumerated
	{
		if ie.DemodulationEnhancement_r16 != nil {
			if err := e.EncodeEnumerated(*ie.DemodulationEnhancement_r16, highSpeedParametersR16DemodulationEnhancementR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *HighSpeedParameters_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(highSpeedParametersR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measurementEnhancement-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(highSpeedParametersR16MeasurementEnhancementR16Constraints)
			if err != nil {
				return err
			}
			ie.MeasurementEnhancement_r16 = &idx
		}
	}

	// 3. demodulationEnhancement-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(highSpeedParametersR16DemodulationEnhancementR16Constraints)
			if err != nil {
				return err
			}
			ie.DemodulationEnhancement_r16 = &idx
		}
	}

	return nil
}
