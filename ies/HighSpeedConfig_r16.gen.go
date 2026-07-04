// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: HighSpeedConfig-r16 (line 8477).

var highSpeedConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "highSpeedMeasFlag-r16", Optional: true},
		{Name: "highSpeedDemodFlag-r16", Optional: true},
	},
}

const (
	HighSpeedConfig_r16_HighSpeedMeasFlag_r16_True = 0
)

var highSpeedConfigR16HighSpeedMeasFlagR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{HighSpeedConfig_r16_HighSpeedMeasFlag_r16_True},
}

const (
	HighSpeedConfig_r16_HighSpeedDemodFlag_r16_True = 0
)

var highSpeedConfigR16HighSpeedDemodFlagR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{HighSpeedConfig_r16_HighSpeedDemodFlag_r16_True},
}

type HighSpeedConfig_r16 struct {
	HighSpeedMeasFlag_r16  *int64
	HighSpeedDemodFlag_r16 *int64
}

func (ie *HighSpeedConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(highSpeedConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.HighSpeedMeasFlag_r16 != nil, ie.HighSpeedDemodFlag_r16 != nil}); err != nil {
		return err
	}

	// 3. highSpeedMeasFlag-r16: enumerated
	{
		if ie.HighSpeedMeasFlag_r16 != nil {
			if err := e.EncodeEnumerated(*ie.HighSpeedMeasFlag_r16, highSpeedConfigR16HighSpeedMeasFlagR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. highSpeedDemodFlag-r16: enumerated
	{
		if ie.HighSpeedDemodFlag_r16 != nil {
			if err := e.EncodeEnumerated(*ie.HighSpeedDemodFlag_r16, highSpeedConfigR16HighSpeedDemodFlagR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *HighSpeedConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(highSpeedConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. highSpeedMeasFlag-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(highSpeedConfigR16HighSpeedMeasFlagR16Constraints)
			if err != nil {
				return err
			}
			ie.HighSpeedMeasFlag_r16 = &idx
		}
	}

	// 4. highSpeedDemodFlag-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(highSpeedConfigR16HighSpeedDemodFlagR16Constraints)
			if err != nil {
				return err
			}
			ie.HighSpeedDemodFlag_r16 = &idx
		}
	}

	return nil
}
