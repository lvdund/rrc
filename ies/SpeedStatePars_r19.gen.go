// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SpeedStatePars-r19 (line 9138).

var speedStateParsR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mobilityStateParameters-r19"},
		{Name: "timeToTrigger-SF-r19"},
	},
}

type SpeedStatePars_r19 struct {
	MobilityStateParameters_r19 MobilityStateParameters
	TimeToTrigger_SF_r19        SpeedStateScaleFactors
}

func (ie *SpeedStatePars_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(speedStateParsR19Constraints)
	_ = seq

	// 1. mobilityStateParameters-r19: ref
	{
		if err := ie.MobilityStateParameters_r19.Encode(e); err != nil {
			return err
		}
	}

	// 2. timeToTrigger-SF-r19: ref
	{
		if err := ie.TimeToTrigger_SF_r19.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SpeedStatePars_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(speedStateParsR19Constraints)
	_ = seq

	// 1. mobilityStateParameters-r19: ref
	{
		if err := ie.MobilityStateParameters_r19.Decode(d); err != nil {
			return err
		}
	}

	// 2. timeToTrigger-SF-r19: ref
	{
		if err := ie.TimeToTrigger_SF_r19.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
