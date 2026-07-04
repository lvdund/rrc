// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParameters-v15t0 (line 21104).

var measAndMobParametersV15t0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersCommon-v15t0", Optional: true},
	},
}

type MeasAndMobParameters_V15t0 struct {
	MeasAndMobParametersCommon_V15t0 *MeasAndMobParametersCommon_V15t0
}

func (ie *MeasAndMobParameters_V15t0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersV15t0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersCommon_V15t0 != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersCommon-v15t0: ref
	{
		if ie.MeasAndMobParametersCommon_V15t0 != nil {
			if err := ie.MeasAndMobParametersCommon_V15t0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParameters_V15t0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersV15t0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersCommon-v15t0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersCommon_V15t0 = new(MeasAndMobParametersCommon_V15t0)
			if err := ie.MeasAndMobParametersCommon_V15t0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
