// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: HighSpeedConfigFR2-r17 (line 8490).

var highSpeedConfigFR2R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "highSpeedMeasFlagFR2-r17", Optional: true},
		{Name: "highSpeedDeploymentTypeFR2-r17", Optional: true},
		{Name: "highSpeedLargeOneStepUL-TimingFR2-r17", Optional: true},
	},
}

const (
	HighSpeedConfigFR2_r17_HighSpeedMeasFlagFR2_r17_Set1 = 0
	HighSpeedConfigFR2_r17_HighSpeedMeasFlagFR2_r17_Set2 = 1
)

var highSpeedConfigFR2R17HighSpeedMeasFlagFR2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{HighSpeedConfigFR2_r17_HighSpeedMeasFlagFR2_r17_Set1, HighSpeedConfigFR2_r17_HighSpeedMeasFlagFR2_r17_Set2},
}

const (
	HighSpeedConfigFR2_r17_HighSpeedDeploymentTypeFR2_r17_Unidirectional = 0
	HighSpeedConfigFR2_r17_HighSpeedDeploymentTypeFR2_r17_Bidirectional  = 1
)

var highSpeedConfigFR2R17HighSpeedDeploymentTypeFR2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{HighSpeedConfigFR2_r17_HighSpeedDeploymentTypeFR2_r17_Unidirectional, HighSpeedConfigFR2_r17_HighSpeedDeploymentTypeFR2_r17_Bidirectional},
}

const (
	HighSpeedConfigFR2_r17_HighSpeedLargeOneStepUL_TimingFR2_r17_True = 0
)

var highSpeedConfigFR2R17HighSpeedLargeOneStepULTimingFR2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{HighSpeedConfigFR2_r17_HighSpeedLargeOneStepUL_TimingFR2_r17_True},
}

type HighSpeedConfigFR2_r17 struct {
	HighSpeedMeasFlagFR2_r17              *int64
	HighSpeedDeploymentTypeFR2_r17        *int64
	HighSpeedLargeOneStepUL_TimingFR2_r17 *int64
}

func (ie *HighSpeedConfigFR2_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(highSpeedConfigFR2R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.HighSpeedMeasFlagFR2_r17 != nil, ie.HighSpeedDeploymentTypeFR2_r17 != nil, ie.HighSpeedLargeOneStepUL_TimingFR2_r17 != nil}); err != nil {
		return err
	}

	// 3. highSpeedMeasFlagFR2-r17: enumerated
	{
		if ie.HighSpeedMeasFlagFR2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.HighSpeedMeasFlagFR2_r17, highSpeedConfigFR2R17HighSpeedMeasFlagFR2R17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. highSpeedDeploymentTypeFR2-r17: enumerated
	{
		if ie.HighSpeedDeploymentTypeFR2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.HighSpeedDeploymentTypeFR2_r17, highSpeedConfigFR2R17HighSpeedDeploymentTypeFR2R17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. highSpeedLargeOneStepUL-TimingFR2-r17: enumerated
	{
		if ie.HighSpeedLargeOneStepUL_TimingFR2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.HighSpeedLargeOneStepUL_TimingFR2_r17, highSpeedConfigFR2R17HighSpeedLargeOneStepULTimingFR2R17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *HighSpeedConfigFR2_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(highSpeedConfigFR2R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. highSpeedMeasFlagFR2-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(highSpeedConfigFR2R17HighSpeedMeasFlagFR2R17Constraints)
			if err != nil {
				return err
			}
			ie.HighSpeedMeasFlagFR2_r17 = &idx
		}
	}

	// 4. highSpeedDeploymentTypeFR2-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(highSpeedConfigFR2R17HighSpeedDeploymentTypeFR2R17Constraints)
			if err != nil {
				return err
			}
			ie.HighSpeedDeploymentTypeFR2_r17 = &idx
		}
	}

	// 5. highSpeedLargeOneStepUL-TimingFR2-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(highSpeedConfigFR2R17HighSpeedLargeOneStepULTimingFR2R17Constraints)
			if err != nil {
				return err
			}
			ie.HighSpeedLargeOneStepUL_TimingFR2_r17 = &idx
		}
	}

	return nil
}
