// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: HighSpeedConfig-v1700 (line 8483).

var highSpeedConfigV1700Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "highSpeedMeasCA-Scell-r17", Optional: true},
		{Name: "highSpeedMeasInterFreq-r17", Optional: true},
		{Name: "highSpeedDemodCA-Scell-r17", Optional: true},
	},
}

const (
	HighSpeedConfig_v1700_HighSpeedMeasCA_Scell_r17_True = 0
)

var highSpeedConfigV1700HighSpeedMeasCAScellR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{HighSpeedConfig_v1700_HighSpeedMeasCA_Scell_r17_True},
}

const (
	HighSpeedConfig_v1700_HighSpeedMeasInterFreq_r17_True = 0
)

var highSpeedConfigV1700HighSpeedMeasInterFreqR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{HighSpeedConfig_v1700_HighSpeedMeasInterFreq_r17_True},
}

const (
	HighSpeedConfig_v1700_HighSpeedDemodCA_Scell_r17_True = 0
)

var highSpeedConfigV1700HighSpeedDemodCAScellR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{HighSpeedConfig_v1700_HighSpeedDemodCA_Scell_r17_True},
}

type HighSpeedConfig_v1700 struct {
	HighSpeedMeasCA_Scell_r17  *int64
	HighSpeedMeasInterFreq_r17 *int64
	HighSpeedDemodCA_Scell_r17 *int64
}

func (ie *HighSpeedConfig_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(highSpeedConfigV1700Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.HighSpeedMeasCA_Scell_r17 != nil, ie.HighSpeedMeasInterFreq_r17 != nil, ie.HighSpeedDemodCA_Scell_r17 != nil}); err != nil {
		return err
	}

	// 3. highSpeedMeasCA-Scell-r17: enumerated
	{
		if ie.HighSpeedMeasCA_Scell_r17 != nil {
			if err := e.EncodeEnumerated(*ie.HighSpeedMeasCA_Scell_r17, highSpeedConfigV1700HighSpeedMeasCAScellR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. highSpeedMeasInterFreq-r17: enumerated
	{
		if ie.HighSpeedMeasInterFreq_r17 != nil {
			if err := e.EncodeEnumerated(*ie.HighSpeedMeasInterFreq_r17, highSpeedConfigV1700HighSpeedMeasInterFreqR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. highSpeedDemodCA-Scell-r17: enumerated
	{
		if ie.HighSpeedDemodCA_Scell_r17 != nil {
			if err := e.EncodeEnumerated(*ie.HighSpeedDemodCA_Scell_r17, highSpeedConfigV1700HighSpeedDemodCAScellR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *HighSpeedConfig_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(highSpeedConfigV1700Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. highSpeedMeasCA-Scell-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(highSpeedConfigV1700HighSpeedMeasCAScellR17Constraints)
			if err != nil {
				return err
			}
			ie.HighSpeedMeasCA_Scell_r17 = &idx
		}
	}

	// 4. highSpeedMeasInterFreq-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(highSpeedConfigV1700HighSpeedMeasInterFreqR17Constraints)
			if err != nil {
				return err
			}
			ie.HighSpeedMeasInterFreq_r17 = &idx
		}
	}

	// 5. highSpeedDemodCA-Scell-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(highSpeedConfigV1700HighSpeedDemodCAScellR17Constraints)
			if err != nil {
				return err
			}
			ie.HighSpeedDemodCA_Scell_r17 = &idx
		}
	}

	return nil
}
