// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OverheatingAssistance-r17 (line 2428).

var overheatingAssistanceR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reducedMaxBW-FR2-2-r17", Optional: true},
		{Name: "reducedMaxMIMO-LayersFR2-2", Optional: true},
	},
}

type OverheatingAssistance_r17 struct {
	ReducedMaxBW_FR2_2_r17 *struct {
		ReducedBW_FR2_2_DL_r17 ReducedAggregatedBandwidth_r17
		ReducedBW_FR2_2_UL_r17 ReducedAggregatedBandwidth_r17
	}
	ReducedMaxMIMO_LayersFR2_2 *struct {
		ReducedMIMO_LayersFR2_2_DL MIMO_LayersDL
		ReducedMIMO_LayersFR2_2_UL MIMO_LayersUL
	}
}

func (ie *OverheatingAssistance_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(overheatingAssistanceR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReducedMaxBW_FR2_2_r17 != nil, ie.ReducedMaxMIMO_LayersFR2_2 != nil}); err != nil {
		return err
	}

	// 2. reducedMaxBW-FR2-2-r17: sequence
	{
		if ie.ReducedMaxBW_FR2_2_r17 != nil {
			c := ie.ReducedMaxBW_FR2_2_r17
			if err := c.ReducedBW_FR2_2_DL_r17.Encode(e); err != nil {
				return err
			}
			if err := c.ReducedBW_FR2_2_UL_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. reducedMaxMIMO-LayersFR2-2: sequence
	{
		if ie.ReducedMaxMIMO_LayersFR2_2 != nil {
			c := ie.ReducedMaxMIMO_LayersFR2_2
			if err := c.ReducedMIMO_LayersFR2_2_DL.Encode(e); err != nil {
				return err
			}
			if err := c.ReducedMIMO_LayersFR2_2_UL.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *OverheatingAssistance_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(overheatingAssistanceR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. reducedMaxBW-FR2-2-r17: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.ReducedMaxBW_FR2_2_r17 = &struct {
				ReducedBW_FR2_2_DL_r17 ReducedAggregatedBandwidth_r17
				ReducedBW_FR2_2_UL_r17 ReducedAggregatedBandwidth_r17
			}{}
			c := ie.ReducedMaxBW_FR2_2_r17
			{
				if err := c.ReducedBW_FR2_2_DL_r17.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.ReducedBW_FR2_2_UL_r17.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. reducedMaxMIMO-LayersFR2-2: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.ReducedMaxMIMO_LayersFR2_2 = &struct {
				ReducedMIMO_LayersFR2_2_DL MIMO_LayersDL
				ReducedMIMO_LayersFR2_2_UL MIMO_LayersUL
			}{}
			c := ie.ReducedMaxMIMO_LayersFR2_2
			{
				if err := c.ReducedMIMO_LayersFR2_2_DL.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.ReducedMIMO_LayersFR2_2_UL.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
