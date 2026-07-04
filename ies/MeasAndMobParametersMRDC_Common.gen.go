// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC-Common (line 21402).

var measAndMobParametersMRDCCommonConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "independentGapConfig", Optional: true},
	},
}

const (
	MeasAndMobParametersMRDC_Common_IndependentGapConfig_Supported = 0
)

var measAndMobParametersMRDCCommonIndependentGapConfigConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_IndependentGapConfig_Supported},
}

type MeasAndMobParametersMRDC_Common struct {
	IndependentGapConfig *int64
}

func (ie *MeasAndMobParametersMRDC_Common) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCCommonConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IndependentGapConfig != nil}); err != nil {
		return err
	}

	// 2. independentGapConfig: enumerated
	{
		if ie.IndependentGapConfig != nil {
			if err := e.EncodeEnumerated(*ie.IndependentGapConfig, measAndMobParametersMRDCCommonIndependentGapConfigConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC_Common) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCCommonConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. independentGapConfig: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonIndependentGapConfigConstraints)
			if err != nil {
				return err
			}
			ie.IndependentGapConfig = &idx
		}
	}

	return nil
}
