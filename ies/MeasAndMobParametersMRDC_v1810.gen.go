// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC-v1810 (line 21394).

var measAndMobParametersMRDCV1810Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersMRDC-Common-v1810", Optional: true},
	},
}

type MeasAndMobParametersMRDC_v1810 struct {
	MeasAndMobParametersMRDC_Common_v1810 *MeasAndMobParametersMRDC_Common_v1810
}

func (ie *MeasAndMobParametersMRDC_v1810) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCV1810Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersMRDC_Common_v1810 != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-Common-v1810: ref
	{
		if ie.MeasAndMobParametersMRDC_Common_v1810 != nil {
			if err := ie.MeasAndMobParametersMRDC_Common_v1810.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC_v1810) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCV1810Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-Common-v1810: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersMRDC_Common_v1810 = new(MeasAndMobParametersMRDC_Common_v1810)
			if err := ie.MeasAndMobParametersMRDC_Common_v1810.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
