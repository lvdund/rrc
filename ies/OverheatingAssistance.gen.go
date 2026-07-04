// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OverheatingAssistance (line 2415).

var overheatingAssistanceConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reducedMaxCCs", Optional: true},
		{Name: "reducedMaxBW-FR1", Optional: true},
		{Name: "reducedMaxBW-FR2", Optional: true},
		{Name: "reducedMaxMIMO-LayersFR1", Optional: true},
		{Name: "reducedMaxMIMO-LayersFR2", Optional: true},
	},
}

type OverheatingAssistance struct {
	ReducedMaxCCs            *ReducedMaxCCs_r16
	ReducedMaxBW_FR1         *ReducedMaxBW_FRx_r16
	ReducedMaxBW_FR2         *ReducedMaxBW_FRx_r16
	ReducedMaxMIMO_LayersFR1 *struct {
		ReducedMIMO_LayersFR1_DL MIMO_LayersDL
		ReducedMIMO_LayersFR1_UL MIMO_LayersUL
	}
	ReducedMaxMIMO_LayersFR2 *struct {
		ReducedMIMO_LayersFR2_DL MIMO_LayersDL
		ReducedMIMO_LayersFR2_UL MIMO_LayersUL
	}
}

func (ie *OverheatingAssistance) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(overheatingAssistanceConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReducedMaxCCs != nil, ie.ReducedMaxBW_FR1 != nil, ie.ReducedMaxBW_FR2 != nil, ie.ReducedMaxMIMO_LayersFR1 != nil, ie.ReducedMaxMIMO_LayersFR2 != nil}); err != nil {
		return err
	}

	// 2. reducedMaxCCs: ref
	{
		if ie.ReducedMaxCCs != nil {
			if err := ie.ReducedMaxCCs.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. reducedMaxBW-FR1: ref
	{
		if ie.ReducedMaxBW_FR1 != nil {
			if err := ie.ReducedMaxBW_FR1.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. reducedMaxBW-FR2: ref
	{
		if ie.ReducedMaxBW_FR2 != nil {
			if err := ie.ReducedMaxBW_FR2.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. reducedMaxMIMO-LayersFR1: sequence
	{
		if ie.ReducedMaxMIMO_LayersFR1 != nil {
			c := ie.ReducedMaxMIMO_LayersFR1
			if err := c.ReducedMIMO_LayersFR1_DL.Encode(e); err != nil {
				return err
			}
			if err := c.ReducedMIMO_LayersFR1_UL.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. reducedMaxMIMO-LayersFR2: sequence
	{
		if ie.ReducedMaxMIMO_LayersFR2 != nil {
			c := ie.ReducedMaxMIMO_LayersFR2
			if err := c.ReducedMIMO_LayersFR2_DL.Encode(e); err != nil {
				return err
			}
			if err := c.ReducedMIMO_LayersFR2_UL.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *OverheatingAssistance) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(overheatingAssistanceConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. reducedMaxCCs: ref
	{
		if seq.IsComponentPresent(0) {
			ie.ReducedMaxCCs = new(ReducedMaxCCs_r16)
			if err := ie.ReducedMaxCCs.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. reducedMaxBW-FR1: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ReducedMaxBW_FR1 = new(ReducedMaxBW_FRx_r16)
			if err := ie.ReducedMaxBW_FR1.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. reducedMaxBW-FR2: ref
	{
		if seq.IsComponentPresent(2) {
			ie.ReducedMaxBW_FR2 = new(ReducedMaxBW_FRx_r16)
			if err := ie.ReducedMaxBW_FR2.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. reducedMaxMIMO-LayersFR1: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.ReducedMaxMIMO_LayersFR1 = &struct {
				ReducedMIMO_LayersFR1_DL MIMO_LayersDL
				ReducedMIMO_LayersFR1_UL MIMO_LayersUL
			}{}
			c := ie.ReducedMaxMIMO_LayersFR1
			{
				if err := c.ReducedMIMO_LayersFR1_DL.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.ReducedMIMO_LayersFR1_UL.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. reducedMaxMIMO-LayersFR2: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.ReducedMaxMIMO_LayersFR2 = &struct {
				ReducedMIMO_LayersFR2_DL MIMO_LayersDL
				ReducedMIMO_LayersFR2_UL MIMO_LayersUL
			}{}
			c := ie.ReducedMaxMIMO_LayersFR2
			{
				if err := c.ReducedMIMO_LayersFR2_DL.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.ReducedMIMO_LayersFR2_UL.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
