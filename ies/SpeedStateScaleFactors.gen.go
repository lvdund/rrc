// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SpeedStateScaleFactors (line 15225).

var speedStateScaleFactorsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sf-Medium"},
		{Name: "sf-High"},
	},
}

const (
	SpeedStateScaleFactors_Sf_Medium_ODot25 = 0
	SpeedStateScaleFactors_Sf_Medium_ODot5  = 1
	SpeedStateScaleFactors_Sf_Medium_ODot75 = 2
	SpeedStateScaleFactors_Sf_Medium_LDot0  = 3
)

var speedStateScaleFactorsSfMediumConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpeedStateScaleFactors_Sf_Medium_ODot25, SpeedStateScaleFactors_Sf_Medium_ODot5, SpeedStateScaleFactors_Sf_Medium_ODot75, SpeedStateScaleFactors_Sf_Medium_LDot0},
}

const (
	SpeedStateScaleFactors_Sf_High_ODot25 = 0
	SpeedStateScaleFactors_Sf_High_ODot5  = 1
	SpeedStateScaleFactors_Sf_High_ODot75 = 2
	SpeedStateScaleFactors_Sf_High_LDot0  = 3
)

var speedStateScaleFactorsSfHighConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpeedStateScaleFactors_Sf_High_ODot25, SpeedStateScaleFactors_Sf_High_ODot5, SpeedStateScaleFactors_Sf_High_ODot75, SpeedStateScaleFactors_Sf_High_LDot0},
}

type SpeedStateScaleFactors struct {
	Sf_Medium int64
	Sf_High   int64
}

func (ie *SpeedStateScaleFactors) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(speedStateScaleFactorsConstraints)
	_ = seq

	// 1. sf-Medium: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sf_Medium, speedStateScaleFactorsSfMediumConstraints); err != nil {
			return err
		}
	}

	// 2. sf-High: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sf_High, speedStateScaleFactorsSfHighConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SpeedStateScaleFactors) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(speedStateScaleFactorsConstraints)
	_ = seq

	// 1. sf-Medium: enumerated
	{
		v0, err := d.DecodeEnumerated(speedStateScaleFactorsSfMediumConstraints)
		if err != nil {
			return err
		}
		ie.Sf_Medium = v0
	}

	// 2. sf-High: enumerated
	{
		v1, err := d.DecodeEnumerated(speedStateScaleFactorsSfHighConstraints)
		if err != nil {
			return err
		}
		ie.Sf_High = v1
	}

	return nil
}
