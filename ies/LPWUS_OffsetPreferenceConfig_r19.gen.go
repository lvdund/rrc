// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LPWUS-OffsetPreferenceConfig-r19 (line 26532).

var lPWUSOffsetPreferenceConfigR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lpwus-OffsetPreferenceProhibitTimer-r19"},
	},
}

const (
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S0     = 0
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S0dot5 = 1
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S1     = 2
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S2     = 3
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S5     = 4
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S10    = 5
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S20    = 6
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S30    = 7
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S60    = 8
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S90    = 9
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S120   = 10
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S300   = 11
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S600   = 12
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_Spare3 = 13
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_Spare2 = 14
	LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_Spare1 = 15
)

var lPWUSOffsetPreferenceConfigR19LpwusOffsetPreferenceProhibitTimerR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S0, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S0dot5, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S1, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S2, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S5, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S10, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S20, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S30, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S60, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S90, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S120, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S300, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_S600, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_Spare3, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_Spare2, LPWUS_OffsetPreferenceConfig_r19_Lpwus_OffsetPreferenceProhibitTimer_r19_Spare1},
}

type LPWUS_OffsetPreferenceConfig_r19 struct {
	Lpwus_OffsetPreferenceProhibitTimer_r19 int64
}

func (ie *LPWUS_OffsetPreferenceConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lPWUSOffsetPreferenceConfigR19Constraints)
	_ = seq

	// 1. lpwus-OffsetPreferenceProhibitTimer-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.Lpwus_OffsetPreferenceProhibitTimer_r19, lPWUSOffsetPreferenceConfigR19LpwusOffsetPreferenceProhibitTimerR19Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *LPWUS_OffsetPreferenceConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lPWUSOffsetPreferenceConfigR19Constraints)
	_ = seq

	// 1. lpwus-OffsetPreferenceProhibitTimer-r19: enumerated
	{
		v0, err := d.DecodeEnumerated(lPWUSOffsetPreferenceConfigR19LpwusOffsetPreferenceProhibitTimerR19Constraints)
		if err != nil {
			return err
		}
		ie.Lpwus_OffsetPreferenceProhibitTimer_r19 = v0
	}

	return nil
}
