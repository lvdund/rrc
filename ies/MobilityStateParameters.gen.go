// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MobilityStateParameters (line 10096).

var mobilityStateParametersConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "t-Evaluation"},
		{Name: "t-HystNormal"},
		{Name: "n-CellChangeMedium"},
		{Name: "n-CellChangeHigh"},
	},
}

const (
	MobilityStateParameters_T_Evaluation_S30    = 0
	MobilityStateParameters_T_Evaluation_S60    = 1
	MobilityStateParameters_T_Evaluation_S120   = 2
	MobilityStateParameters_T_Evaluation_S180   = 3
	MobilityStateParameters_T_Evaluation_S240   = 4
	MobilityStateParameters_T_Evaluation_Spare3 = 5
	MobilityStateParameters_T_Evaluation_Spare2 = 6
	MobilityStateParameters_T_Evaluation_Spare1 = 7
)

var mobilityStateParametersTEvaluationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MobilityStateParameters_T_Evaluation_S30, MobilityStateParameters_T_Evaluation_S60, MobilityStateParameters_T_Evaluation_S120, MobilityStateParameters_T_Evaluation_S180, MobilityStateParameters_T_Evaluation_S240, MobilityStateParameters_T_Evaluation_Spare3, MobilityStateParameters_T_Evaluation_Spare2, MobilityStateParameters_T_Evaluation_Spare1},
}

const (
	MobilityStateParameters_T_HystNormal_S30    = 0
	MobilityStateParameters_T_HystNormal_S60    = 1
	MobilityStateParameters_T_HystNormal_S120   = 2
	MobilityStateParameters_T_HystNormal_S180   = 3
	MobilityStateParameters_T_HystNormal_S240   = 4
	MobilityStateParameters_T_HystNormal_Spare3 = 5
	MobilityStateParameters_T_HystNormal_Spare2 = 6
	MobilityStateParameters_T_HystNormal_Spare1 = 7
)

var mobilityStateParametersTHystNormalConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MobilityStateParameters_T_HystNormal_S30, MobilityStateParameters_T_HystNormal_S60, MobilityStateParameters_T_HystNormal_S120, MobilityStateParameters_T_HystNormal_S180, MobilityStateParameters_T_HystNormal_S240, MobilityStateParameters_T_HystNormal_Spare3, MobilityStateParameters_T_HystNormal_Spare2, MobilityStateParameters_T_HystNormal_Spare1},
}

var mobilityStateParametersNCellChangeMediumConstraints = per.Constrained(1, 16)

var mobilityStateParametersNCellChangeHighConstraints = per.Constrained(1, 16)

type MobilityStateParameters struct {
	T_Evaluation       int64
	T_HystNormal       int64
	N_CellChangeMedium int64
	N_CellChangeHigh   int64
}

func (ie *MobilityStateParameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mobilityStateParametersConstraints)
	_ = seq

	// 1. t-Evaluation: enumerated
	{
		if err := e.EncodeEnumerated(ie.T_Evaluation, mobilityStateParametersTEvaluationConstraints); err != nil {
			return err
		}
	}

	// 2. t-HystNormal: enumerated
	{
		if err := e.EncodeEnumerated(ie.T_HystNormal, mobilityStateParametersTHystNormalConstraints); err != nil {
			return err
		}
	}

	// 3. n-CellChangeMedium: integer
	{
		if err := e.EncodeInteger(ie.N_CellChangeMedium, mobilityStateParametersNCellChangeMediumConstraints); err != nil {
			return err
		}
	}

	// 4. n-CellChangeHigh: integer
	{
		if err := e.EncodeInteger(ie.N_CellChangeHigh, mobilityStateParametersNCellChangeHighConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MobilityStateParameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mobilityStateParametersConstraints)
	_ = seq

	// 1. t-Evaluation: enumerated
	{
		v0, err := d.DecodeEnumerated(mobilityStateParametersTEvaluationConstraints)
		if err != nil {
			return err
		}
		ie.T_Evaluation = v0
	}

	// 2. t-HystNormal: enumerated
	{
		v1, err := d.DecodeEnumerated(mobilityStateParametersTHystNormalConstraints)
		if err != nil {
			return err
		}
		ie.T_HystNormal = v1
	}

	// 3. n-CellChangeMedium: integer
	{
		v2, err := d.DecodeInteger(mobilityStateParametersNCellChangeMediumConstraints)
		if err != nil {
			return err
		}
		ie.N_CellChangeMedium = v2
	}

	// 4. n-CellChangeHigh: integer
	{
		v3, err := d.DecodeInteger(mobilityStateParametersNCellChangeHighConstraints)
		if err != nil {
			return err
		}
		ie.N_CellChangeHigh = v3
	}

	return nil
}
