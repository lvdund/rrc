// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MaxBW-PreferenceFR2-2-r17 (line 2553).

var maxBWPreferenceFR22R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reducedMaxBW-FR2-2-r17", Optional: true},
	},
}

var maxBWPreferenceFR22R17ReducedMaxBWFR22R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reducedBW-FR2-2-DL-r17", Optional: true},
		{Name: "reducedBW-FR2-2-UL-r17", Optional: true},
	},
}

type MaxBW_PreferenceFR2_2_r17 struct {
	ReducedMaxBW_FR2_2_r17 *struct {
		ReducedBW_FR2_2_DL_r17 *ReducedAggregatedBandwidth_r17
		ReducedBW_FR2_2_UL_r17 *ReducedAggregatedBandwidth_r17
	}
}

func (ie *MaxBW_PreferenceFR2_2_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(maxBWPreferenceFR22R17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReducedMaxBW_FR2_2_r17 != nil}); err != nil {
		return err
	}

	// 2. reducedMaxBW-FR2-2-r17: sequence
	{
		if ie.ReducedMaxBW_FR2_2_r17 != nil {
			c := ie.ReducedMaxBW_FR2_2_r17
			maxBWPreferenceFR22R17ReducedMaxBWFR22R17Seq := e.NewSequenceEncoder(maxBWPreferenceFR22R17ReducedMaxBWFR22R17Constraints)
			if err := maxBWPreferenceFR22R17ReducedMaxBWFR22R17Seq.EncodePreamble([]bool{c.ReducedBW_FR2_2_DL_r17 != nil, c.ReducedBW_FR2_2_UL_r17 != nil}); err != nil {
				return err
			}
			if c.ReducedBW_FR2_2_DL_r17 != nil {
				if err := c.ReducedBW_FR2_2_DL_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.ReducedBW_FR2_2_UL_r17 != nil {
				if err := c.ReducedBW_FR2_2_UL_r17.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MaxBW_PreferenceFR2_2_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(maxBWPreferenceFR22R17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. reducedMaxBW-FR2-2-r17: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.ReducedMaxBW_FR2_2_r17 = &struct {
				ReducedBW_FR2_2_DL_r17 *ReducedAggregatedBandwidth_r17
				ReducedBW_FR2_2_UL_r17 *ReducedAggregatedBandwidth_r17
			}{}
			c := ie.ReducedMaxBW_FR2_2_r17
			maxBWPreferenceFR22R17ReducedMaxBWFR22R17Seq := d.NewSequenceDecoder(maxBWPreferenceFR22R17ReducedMaxBWFR22R17Constraints)
			if err := maxBWPreferenceFR22R17ReducedMaxBWFR22R17Seq.DecodePreamble(); err != nil {
				return err
			}
			if maxBWPreferenceFR22R17ReducedMaxBWFR22R17Seq.IsComponentPresent(0) {
				c.ReducedBW_FR2_2_DL_r17 = new(ReducedAggregatedBandwidth_r17)
				if err := (*c.ReducedBW_FR2_2_DL_r17).Decode(d); err != nil {
					return err
				}
			}
			if maxBWPreferenceFR22R17ReducedMaxBWFR22R17Seq.IsComponentPresent(1) {
				c.ReducedBW_FR2_2_UL_r17 = new(ReducedAggregatedBandwidth_r17)
				if err := (*c.ReducedBW_FR2_2_UL_r17).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
