// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MaxCC-Preference-r16 (line 2560).

var maxCCPreferenceR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reducedMaxCCs-r16", Optional: true},
	},
}

type MaxCC_Preference_r16 struct {
	ReducedMaxCCs_r16 *ReducedMaxCCs_r16
}

func (ie *MaxCC_Preference_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(maxCCPreferenceR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReducedMaxCCs_r16 != nil}); err != nil {
		return err
	}

	// 2. reducedMaxCCs-r16: ref
	{
		if ie.ReducedMaxCCs_r16 != nil {
			if err := ie.ReducedMaxCCs_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MaxCC_Preference_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(maxCCPreferenceR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. reducedMaxCCs-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.ReducedMaxCCs_r16 = new(ReducedMaxCCs_r16)
			if err := ie.ReducedMaxCCs_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
