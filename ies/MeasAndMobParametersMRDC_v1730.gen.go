// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC-v1730 (line 21390).

var measAndMobParametersMRDCV1730Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersMRDC-Common-v1730", Optional: true},
	},
}

type MeasAndMobParametersMRDC_v1730 struct {
	MeasAndMobParametersMRDC_Common_v1730 *MeasAndMobParametersMRDC_Common_v1730
}

func (ie *MeasAndMobParametersMRDC_v1730) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCV1730Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersMRDC_Common_v1730 != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-Common-v1730: ref
	{
		if ie.MeasAndMobParametersMRDC_Common_v1730 != nil {
			if err := ie.MeasAndMobParametersMRDC_Common_v1730.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC_v1730) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCV1730Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-Common-v1730: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersMRDC_Common_v1730 = new(MeasAndMobParametersMRDC_Common_v1730)
			if err := ie.MeasAndMobParametersMRDC_Common_v1730.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
