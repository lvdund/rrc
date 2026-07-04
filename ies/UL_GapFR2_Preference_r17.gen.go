// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-GapFR2-Preference-r17 (line 2709).

var uLGapFR2PreferenceR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-GapFR2-PatternPreference-r17", Optional: true},
	},
}

var uLGapFR2PreferenceR17UlGapFR2PatternPreferenceR17Constraints = per.Constrained(0, 3)

type UL_GapFR2_Preference_r17 struct {
	Ul_GapFR2_PatternPreference_r17 *int64
}

func (ie *UL_GapFR2_Preference_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLGapFR2PreferenceR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ul_GapFR2_PatternPreference_r17 != nil}); err != nil {
		return err
	}

	// 2. ul-GapFR2-PatternPreference-r17: integer
	{
		if ie.Ul_GapFR2_PatternPreference_r17 != nil {
			if err := e.EncodeInteger(*ie.Ul_GapFR2_PatternPreference_r17, uLGapFR2PreferenceR17UlGapFR2PatternPreferenceR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UL_GapFR2_Preference_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLGapFR2PreferenceR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ul-GapFR2-PatternPreference-r17: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(uLGapFR2PreferenceR17UlGapFR2PatternPreferenceR17Constraints)
			if err != nil {
				return err
			}
			ie.Ul_GapFR2_PatternPreference_r17 = &v
		}
	}

	return nil
}
