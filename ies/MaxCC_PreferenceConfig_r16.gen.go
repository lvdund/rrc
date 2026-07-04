// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MaxCC-PreferenceConfig-r16 (line 26444).

var maxCCPreferenceConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxCC-PreferenceProhibitTimer-r16"},
	},
}

const (
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S0     = 0
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S0dot5 = 1
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S1     = 2
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S2     = 3
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S3     = 4
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S4     = 5
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S5     = 6
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S6     = 7
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S7     = 8
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S8     = 9
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S9     = 10
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S10    = 11
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S20    = 12
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S30    = 13
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_Spare2 = 14
	MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_Spare1 = 15
)

var maxCCPreferenceConfigR16MaxCCPreferenceProhibitTimerR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S0, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S0dot5, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S1, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S2, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S3, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S4, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S5, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S6, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S7, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S8, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S9, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S10, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S20, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_S30, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_Spare2, MaxCC_PreferenceConfig_r16_MaxCC_PreferenceProhibitTimer_r16_Spare1},
}

type MaxCC_PreferenceConfig_r16 struct {
	MaxCC_PreferenceProhibitTimer_r16 int64
}

func (ie *MaxCC_PreferenceConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(maxCCPreferenceConfigR16Constraints)
	_ = seq

	// 1. maxCC-PreferenceProhibitTimer-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxCC_PreferenceProhibitTimer_r16, maxCCPreferenceConfigR16MaxCCPreferenceProhibitTimerR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MaxCC_PreferenceConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(maxCCPreferenceConfigR16Constraints)
	_ = seq

	// 1. maxCC-PreferenceProhibitTimer-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(maxCCPreferenceConfigR16MaxCCPreferenceProhibitTimerR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxCC_PreferenceProhibitTimer_r16 = v0
	}

	return nil
}
