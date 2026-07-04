// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParameters-v1700 (line 21108).

var measAndMobParametersV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersFR2-2-r17", Optional: true},
	},
}

type MeasAndMobParameters_v1700 struct {
	MeasAndMobParametersFR2_2_r17 *MeasAndMobParametersFR2_2_r17
}

func (ie *MeasAndMobParameters_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersFR2_2_r17 != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersFR2-2-r17: ref
	{
		if ie.MeasAndMobParametersFR2_2_r17 != nil {
			if err := ie.MeasAndMobParametersFR2_2_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParameters_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersFR2-2-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersFR2_2_r17 = new(MeasAndMobParametersFR2_2_r17)
			if err := ie.MeasAndMobParametersFR2_2_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
