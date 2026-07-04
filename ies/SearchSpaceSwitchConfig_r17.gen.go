// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SearchSpaceSwitchConfig-r17 (line 11028).

var searchSpaceSwitchConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "searchSpaceSwitchTimer-r17", Optional: true},
		{Name: "searchSpaceSwitchDelay-r17", Optional: true},
	},
}

var searchSpaceSwitchConfigR17SearchSpaceSwitchDelayR17Constraints = per.Constrained(10, 52)

type SearchSpaceSwitchConfig_r17 struct {
	SearchSpaceSwitchTimer_r17 *SCS_SpecificDuration_r17
	SearchSpaceSwitchDelay_r17 *int64
}

func (ie *SearchSpaceSwitchConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(searchSpaceSwitchConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SearchSpaceSwitchTimer_r17 != nil, ie.SearchSpaceSwitchDelay_r17 != nil}); err != nil {
		return err
	}

	// 2. searchSpaceSwitchTimer-r17: ref
	{
		if ie.SearchSpaceSwitchTimer_r17 != nil {
			if err := ie.SearchSpaceSwitchTimer_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. searchSpaceSwitchDelay-r17: integer
	{
		if ie.SearchSpaceSwitchDelay_r17 != nil {
			if err := e.EncodeInteger(*ie.SearchSpaceSwitchDelay_r17, searchSpaceSwitchConfigR17SearchSpaceSwitchDelayR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SearchSpaceSwitchConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(searchSpaceSwitchConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. searchSpaceSwitchTimer-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SearchSpaceSwitchTimer_r17 = new(SCS_SpecificDuration_r17)
			if err := ie.SearchSpaceSwitchTimer_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. searchSpaceSwitchDelay-r17: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(searchSpaceSwitchConfigR17SearchSpaceSwitchDelayR17Constraints)
			if err != nil {
				return err
			}
			ie.SearchSpaceSwitchDelay_r17 = &v
		}
	}

	return nil
}
