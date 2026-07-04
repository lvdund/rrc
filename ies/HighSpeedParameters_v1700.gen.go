// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: HighSpeedParameters-v1700 (line 20848).

var highSpeedParametersV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measurementEnhancementCA-r17", Optional: true},
		{Name: "measurementEnhancementInterFreq-r17", Optional: true},
	},
}

const (
	HighSpeedParameters_v1700_MeasurementEnhancementCA_r17_Supported = 0
)

var highSpeedParametersV1700MeasurementEnhancementCAR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{HighSpeedParameters_v1700_MeasurementEnhancementCA_r17_Supported},
}

const (
	HighSpeedParameters_v1700_MeasurementEnhancementInterFreq_r17_Supported = 0
)

var highSpeedParametersV1700MeasurementEnhancementInterFreqR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{HighSpeedParameters_v1700_MeasurementEnhancementInterFreq_r17_Supported},
}

type HighSpeedParameters_v1700 struct {
	MeasurementEnhancementCA_r17        *int64
	MeasurementEnhancementInterFreq_r17 *int64
}

func (ie *HighSpeedParameters_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(highSpeedParametersV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasurementEnhancementCA_r17 != nil, ie.MeasurementEnhancementInterFreq_r17 != nil}); err != nil {
		return err
	}

	// 2. measurementEnhancementCA-r17: enumerated
	{
		if ie.MeasurementEnhancementCA_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MeasurementEnhancementCA_r17, highSpeedParametersV1700MeasurementEnhancementCAR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. measurementEnhancementInterFreq-r17: enumerated
	{
		if ie.MeasurementEnhancementInterFreq_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MeasurementEnhancementInterFreq_r17, highSpeedParametersV1700MeasurementEnhancementInterFreqR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *HighSpeedParameters_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(highSpeedParametersV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measurementEnhancementCA-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(highSpeedParametersV1700MeasurementEnhancementCAR17Constraints)
			if err != nil {
				return err
			}
			ie.MeasurementEnhancementCA_r17 = &idx
		}
	}

	// 3. measurementEnhancementInterFreq-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(highSpeedParametersV1700MeasurementEnhancementInterFreqR17Constraints)
			if err != nil {
				return err
			}
			ie.MeasurementEnhancementInterFreq_r17 = &idx
		}
	}

	return nil
}
