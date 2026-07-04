// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MaxBW-Preference-r16 (line 2548).

var maxBWPreferenceR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reducedMaxBW-FR1-r16", Optional: true},
		{Name: "reducedMaxBW-FR2-r16", Optional: true},
	},
}

type MaxBW_Preference_r16 struct {
	ReducedMaxBW_FR1_r16 *ReducedMaxBW_FRx_r16
	ReducedMaxBW_FR2_r16 *ReducedMaxBW_FRx_r16
}

func (ie *MaxBW_Preference_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(maxBWPreferenceR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReducedMaxBW_FR1_r16 != nil, ie.ReducedMaxBW_FR2_r16 != nil}); err != nil {
		return err
	}

	// 2. reducedMaxBW-FR1-r16: ref
	{
		if ie.ReducedMaxBW_FR1_r16 != nil {
			if err := ie.ReducedMaxBW_FR1_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. reducedMaxBW-FR2-r16: ref
	{
		if ie.ReducedMaxBW_FR2_r16 != nil {
			if err := ie.ReducedMaxBW_FR2_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MaxBW_Preference_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(maxBWPreferenceR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. reducedMaxBW-FR1-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.ReducedMaxBW_FR1_r16 = new(ReducedMaxBW_FRx_r16)
			if err := ie.ReducedMaxBW_FR1_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. reducedMaxBW-FR2-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ReducedMaxBW_FR2_r16 = new(ReducedMaxBW_FRx_r16)
			if err := ie.ReducedMaxBW_FR2_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
