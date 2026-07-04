// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MaxMIMO-LayerPreference-r16 (line 2564).

var maxMIMOLayerPreferenceR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reducedMaxMIMO-LayersFR1-r16", Optional: true},
		{Name: "reducedMaxMIMO-LayersFR2-r16", Optional: true},
	},
}

type MaxMIMO_LayerPreference_r16 struct {
	ReducedMaxMIMO_LayersFR1_r16 *struct {
		ReducedMIMO_LayersFR1_DL_r16 int64
		ReducedMIMO_LayersFR1_UL_r16 int64
	}
	ReducedMaxMIMO_LayersFR2_r16 *struct {
		ReducedMIMO_LayersFR2_DL_r16 int64
		ReducedMIMO_LayersFR2_UL_r16 int64
	}
}

func (ie *MaxMIMO_LayerPreference_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(maxMIMOLayerPreferenceR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReducedMaxMIMO_LayersFR1_r16 != nil, ie.ReducedMaxMIMO_LayersFR2_r16 != nil}); err != nil {
		return err
	}

	// 2. reducedMaxMIMO-LayersFR1-r16: sequence
	{
		if ie.ReducedMaxMIMO_LayersFR1_r16 != nil {
			c := ie.ReducedMaxMIMO_LayersFR1_r16
			if err := e.EncodeInteger(c.ReducedMIMO_LayersFR1_DL_r16, per.Constrained(1, 8)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.ReducedMIMO_LayersFR1_UL_r16, per.Constrained(1, 4)); err != nil {
				return err
			}
		}
	}

	// 3. reducedMaxMIMO-LayersFR2-r16: sequence
	{
		if ie.ReducedMaxMIMO_LayersFR2_r16 != nil {
			c := ie.ReducedMaxMIMO_LayersFR2_r16
			if err := e.EncodeInteger(c.ReducedMIMO_LayersFR2_DL_r16, per.Constrained(1, 8)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.ReducedMIMO_LayersFR2_UL_r16, per.Constrained(1, 4)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MaxMIMO_LayerPreference_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(maxMIMOLayerPreferenceR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. reducedMaxMIMO-LayersFR1-r16: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.ReducedMaxMIMO_LayersFR1_r16 = &struct {
				ReducedMIMO_LayersFR1_DL_r16 int64
				ReducedMIMO_LayersFR1_UL_r16 int64
			}{}
			c := ie.ReducedMaxMIMO_LayersFR1_r16
			{
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.ReducedMIMO_LayersFR1_DL_r16 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.ReducedMIMO_LayersFR1_UL_r16 = v
			}
		}
	}

	// 3. reducedMaxMIMO-LayersFR2-r16: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.ReducedMaxMIMO_LayersFR2_r16 = &struct {
				ReducedMIMO_LayersFR2_DL_r16 int64
				ReducedMIMO_LayersFR2_UL_r16 int64
			}{}
			c := ie.ReducedMaxMIMO_LayersFR2_r16
			{
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.ReducedMIMO_LayersFR2_DL_r16 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.ReducedMIMO_LayersFR2_UL_r16 = v
			}
		}
	}

	return nil
}
