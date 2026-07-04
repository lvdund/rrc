// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RepFactorAndTimeGap-r17 (line 10736).

var repFactorAndTimeGapR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "repetitionFactor-r17"},
		{Name: "timeGap-r17"},
	},
}

const (
	RepFactorAndTimeGap_r17_RepetitionFactor_r17_N2     = 0
	RepFactorAndTimeGap_r17_RepetitionFactor_r17_N4     = 1
	RepFactorAndTimeGap_r17_RepetitionFactor_r17_N6     = 2
	RepFactorAndTimeGap_r17_RepetitionFactor_r17_N8     = 3
	RepFactorAndTimeGap_r17_RepetitionFactor_r17_N16    = 4
	RepFactorAndTimeGap_r17_RepetitionFactor_r17_N32    = 5
	RepFactorAndTimeGap_r17_RepetitionFactor_r17_Spare2 = 6
	RepFactorAndTimeGap_r17_RepetitionFactor_r17_Spare1 = 7
)

var repFactorAndTimeGapR17RepetitionFactorR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RepFactorAndTimeGap_r17_RepetitionFactor_r17_N2, RepFactorAndTimeGap_r17_RepetitionFactor_r17_N4, RepFactorAndTimeGap_r17_RepetitionFactor_r17_N6, RepFactorAndTimeGap_r17_RepetitionFactor_r17_N8, RepFactorAndTimeGap_r17_RepetitionFactor_r17_N16, RepFactorAndTimeGap_r17_RepetitionFactor_r17_N32, RepFactorAndTimeGap_r17_RepetitionFactor_r17_Spare2, RepFactorAndTimeGap_r17_RepetitionFactor_r17_Spare1},
}

const (
	RepFactorAndTimeGap_r17_TimeGap_r17_S1     = 0
	RepFactorAndTimeGap_r17_TimeGap_r17_S2     = 1
	RepFactorAndTimeGap_r17_TimeGap_r17_S4     = 2
	RepFactorAndTimeGap_r17_TimeGap_r17_S8     = 3
	RepFactorAndTimeGap_r17_TimeGap_r17_S16    = 4
	RepFactorAndTimeGap_r17_TimeGap_r17_S32    = 5
	RepFactorAndTimeGap_r17_TimeGap_r17_Spare2 = 6
	RepFactorAndTimeGap_r17_TimeGap_r17_Spare1 = 7
)

var repFactorAndTimeGapR17TimeGapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RepFactorAndTimeGap_r17_TimeGap_r17_S1, RepFactorAndTimeGap_r17_TimeGap_r17_S2, RepFactorAndTimeGap_r17_TimeGap_r17_S4, RepFactorAndTimeGap_r17_TimeGap_r17_S8, RepFactorAndTimeGap_r17_TimeGap_r17_S16, RepFactorAndTimeGap_r17_TimeGap_r17_S32, RepFactorAndTimeGap_r17_TimeGap_r17_Spare2, RepFactorAndTimeGap_r17_TimeGap_r17_Spare1},
}

type RepFactorAndTimeGap_r17 struct {
	RepetitionFactor_r17 int64
	TimeGap_r17          int64
}

func (ie *RepFactorAndTimeGap_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(repFactorAndTimeGapR17Constraints)
	_ = seq

	// 1. repetitionFactor-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.RepetitionFactor_r17, repFactorAndTimeGapR17RepetitionFactorR17Constraints); err != nil {
			return err
		}
	}

	// 2. timeGap-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.TimeGap_r17, repFactorAndTimeGapR17TimeGapR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RepFactorAndTimeGap_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(repFactorAndTimeGapR17Constraints)
	_ = seq

	// 1. repetitionFactor-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(repFactorAndTimeGapR17RepetitionFactorR17Constraints)
		if err != nil {
			return err
		}
		ie.RepetitionFactor_r17 = v0
	}

	// 2. timeGap-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(repFactorAndTimeGapR17TimeGapR17Constraints)
		if err != nil {
			return err
		}
		ie.TimeGap_r17 = v1
	}

	return nil
}
