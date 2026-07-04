// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC-v1900 (line 21398).

var measAndMobParametersMRDCV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersMRDC-Common-v1900", Optional: true},
	},
}

type MeasAndMobParametersMRDC_v1900 struct {
	MeasAndMobParametersMRDC_Common_v1900 *MeasAndMobParametersMRDC_Common_v1900
}

func (ie *MeasAndMobParametersMRDC_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersMRDC_Common_v1900 != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-Common-v1900: ref
	{
		if ie.MeasAndMobParametersMRDC_Common_v1900 != nil {
			if err := ie.MeasAndMobParametersMRDC_Common_v1900.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-Common-v1900: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersMRDC_Common_v1900 = new(MeasAndMobParametersMRDC_Common_v1900)
			if err := ie.MeasAndMobParametersMRDC_Common_v1900.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
