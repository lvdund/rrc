// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC-v1560 (line 21377).

var measAndMobParametersMRDCV1560Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersMRDC-XDD-Diff-v1560", Optional: true},
	},
}

type MeasAndMobParametersMRDC_v1560 struct {
	MeasAndMobParametersMRDC_XDD_Diff_v1560 *MeasAndMobParametersMRDC_XDD_Diff_v1560
}

func (ie *MeasAndMobParametersMRDC_v1560) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCV1560Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersMRDC_XDD_Diff_v1560 != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-XDD-Diff-v1560: ref
	{
		if ie.MeasAndMobParametersMRDC_XDD_Diff_v1560 != nil {
			if err := ie.MeasAndMobParametersMRDC_XDD_Diff_v1560.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC_v1560) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCV1560Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-XDD-Diff-v1560: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersMRDC_XDD_Diff_v1560 = new(MeasAndMobParametersMRDC_XDD_Diff_v1560)
			if err := ie.MeasAndMobParametersMRDC_XDD_Diff_v1560.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
