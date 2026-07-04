// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PerRAAttemptInfo-r16 (line 3199).

var perRAAttemptInfoR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "contentionDetected-r16", Optional: true},
		{Name: "dlRSRPAboveThreshold-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	PerRAAttemptInfo_r16_Ext_FallbackToFourStepRA_r17_True = 0
)

var perRAAttemptInfoR16ExtFallbackToFourStepRAR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PerRAAttemptInfo_r16_Ext_FallbackToFourStepRA_r17_True},
}

type PerRAAttemptInfo_r16 struct {
	ContentionDetected_r16   *bool
	DlRSRPAboveThreshold_r16 *bool
	FallbackToFourStepRA_r17 *int64
}

func (ie *PerRAAttemptInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(perRAAttemptInfoR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.FallbackToFourStepRA_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ContentionDetected_r16 != nil, ie.DlRSRPAboveThreshold_r16 != nil}); err != nil {
		return err
	}

	// 3. contentionDetected-r16: boolean
	{
		if ie.ContentionDetected_r16 != nil {
			if err := e.EncodeBoolean(*ie.ContentionDetected_r16); err != nil {
				return err
			}
		}
	}

	// 4. dlRSRPAboveThreshold-r16: boolean
	{
		if ie.DlRSRPAboveThreshold_r16 != nil {
			if err := e.EncodeBoolean(*ie.DlRSRPAboveThreshold_r16); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "fallbackToFourStepRA-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FallbackToFourStepRA_r17 != nil}); err != nil {
				return err
			}

			if ie.FallbackToFourStepRA_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.FallbackToFourStepRA_r17, perRAAttemptInfoR16ExtFallbackToFourStepRAR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PerRAAttemptInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(perRAAttemptInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. contentionDetected-r16: boolean
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.ContentionDetected_r16 = &v
		}
	}

	// 4. dlRSRPAboveThreshold-r16: boolean
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.DlRSRPAboveThreshold_r16 = &v
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "fallbackToFourStepRA-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(perRAAttemptInfoR16ExtFallbackToFourStepRAR17Constraints)
			if err != nil {
				return err
			}
			ie.FallbackToFourStepRA_r17 = &v
		}
	}

	return nil
}
