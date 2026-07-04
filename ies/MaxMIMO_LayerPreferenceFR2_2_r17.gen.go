// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MaxMIMO-LayerPreferenceFR2-2-r17 (line 2575).

var maxMIMOLayerPreferenceFR22R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reducedMaxMIMO-LayersFR2-2-r17", Optional: true},
	},
}

type MaxMIMO_LayerPreferenceFR2_2_r17 struct {
	ReducedMaxMIMO_LayersFR2_2_r17 *struct {
		ReducedMIMO_LayersFR2_2_DL_r17 int64
		ReducedMIMO_LayersFR2_2_UL_r17 int64
	}
}

func (ie *MaxMIMO_LayerPreferenceFR2_2_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(maxMIMOLayerPreferenceFR22R17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReducedMaxMIMO_LayersFR2_2_r17 != nil}); err != nil {
		return err
	}

	// 2. reducedMaxMIMO-LayersFR2-2-r17: sequence
	{
		if ie.ReducedMaxMIMO_LayersFR2_2_r17 != nil {
			c := ie.ReducedMaxMIMO_LayersFR2_2_r17
			if err := e.EncodeInteger(c.ReducedMIMO_LayersFR2_2_DL_r17, per.Constrained(1, 8)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.ReducedMIMO_LayersFR2_2_UL_r17, per.Constrained(1, 4)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MaxMIMO_LayerPreferenceFR2_2_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(maxMIMOLayerPreferenceFR22R17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. reducedMaxMIMO-LayersFR2-2-r17: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.ReducedMaxMIMO_LayersFR2_2_r17 = &struct {
				ReducedMIMO_LayersFR2_2_DL_r17 int64
				ReducedMIMO_LayersFR2_2_UL_r17 int64
			}{}
			c := ie.ReducedMaxMIMO_LayersFR2_2_r17
			{
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.ReducedMIMO_LayersFR2_2_DL_r17 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.ReducedMIMO_LayersFR2_2_UL_r17 = v
			}
		}
	}

	return nil
}
