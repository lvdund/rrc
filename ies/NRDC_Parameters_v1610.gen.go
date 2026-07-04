// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NRDC-Parameters-v1610 (line 22665).

var nRDCParametersV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersNRDC-v1610", Optional: true},
	},
}

type NRDC_Parameters_v1610 struct {
	MeasAndMobParametersNRDC_v1610 *MeasAndMobParametersMRDC_v1610
}

func (ie *NRDC_Parameters_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRDCParametersV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersNRDC_v1610 != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersNRDC-v1610: ref
	{
		if ie.MeasAndMobParametersNRDC_v1610 != nil {
			if err := ie.MeasAndMobParametersNRDC_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NRDC_Parameters_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRDCParametersV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersNRDC-v1610: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersNRDC_v1610 = new(MeasAndMobParametersMRDC_v1610)
			if err := ie.MeasAndMobParametersNRDC_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
