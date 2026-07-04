// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MaxBW-PreferenceConfig-r16 (line 26438).

var maxBWPreferenceConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxBW-PreferenceProhibitTimer-r16"},
	},
}

const (
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S0     = 0
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S0dot5 = 1
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S1     = 2
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S2     = 3
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S3     = 4
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S4     = 5
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S5     = 6
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S6     = 7
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S7     = 8
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S8     = 9
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S9     = 10
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S10    = 11
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S20    = 12
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S30    = 13
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_Spare2 = 14
	MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_Spare1 = 15
)

var maxBWPreferenceConfigR16MaxBWPreferenceProhibitTimerR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S0, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S0dot5, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S1, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S2, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S3, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S4, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S5, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S6, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S7, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S8, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S9, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S10, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S20, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_S30, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_Spare2, MaxBW_PreferenceConfig_r16_MaxBW_PreferenceProhibitTimer_r16_Spare1},
}

type MaxBW_PreferenceConfig_r16 struct {
	MaxBW_PreferenceProhibitTimer_r16 int64
}

func (ie *MaxBW_PreferenceConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(maxBWPreferenceConfigR16Constraints)
	_ = seq

	// 1. maxBW-PreferenceProhibitTimer-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxBW_PreferenceProhibitTimer_r16, maxBWPreferenceConfigR16MaxBWPreferenceProhibitTimerR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MaxBW_PreferenceConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(maxBWPreferenceConfigR16Constraints)
	_ = seq

	// 1. maxBW-PreferenceProhibitTimer-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(maxBWPreferenceConfigR16MaxBWPreferenceProhibitTimerR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxBW_PreferenceProhibitTimer_r16 = v0
	}

	return nil
}
