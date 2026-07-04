// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OverheatingAssistanceConfig (line 26422).

var overheatingAssistanceConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "overheatingIndicationProhibitTimer"},
	},
}

const (
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S0     = 0
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S0dot5 = 1
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S1     = 2
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S2     = 3
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S5     = 4
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S10    = 5
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S20    = 6
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S30    = 7
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S60    = 8
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S90    = 9
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S120   = 10
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S300   = 11
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S600   = 12
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_Spare3 = 13
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_Spare2 = 14
	OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_Spare1 = 15
)

var overheatingAssistanceConfigOverheatingIndicationProhibitTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S0, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S0dot5, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S1, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S2, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S5, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S10, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S20, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S30, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S60, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S90, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S120, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S300, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_S600, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_Spare3, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_Spare2, OverheatingAssistanceConfig_OverheatingIndicationProhibitTimer_Spare1},
}

type OverheatingAssistanceConfig struct {
	OverheatingIndicationProhibitTimer int64
}

func (ie *OverheatingAssistanceConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(overheatingAssistanceConfigConstraints)
	_ = seq

	// 1. overheatingIndicationProhibitTimer: enumerated
	{
		if err := e.EncodeEnumerated(ie.OverheatingIndicationProhibitTimer, overheatingAssistanceConfigOverheatingIndicationProhibitTimerConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *OverheatingAssistanceConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(overheatingAssistanceConfigConstraints)
	_ = seq

	// 1. overheatingIndicationProhibitTimer: enumerated
	{
		v0, err := d.DecodeEnumerated(overheatingAssistanceConfigOverheatingIndicationProhibitTimerConstraints)
		if err != nil {
			return err
		}
		ie.OverheatingIndicationProhibitTimer = v0
	}

	return nil
}
