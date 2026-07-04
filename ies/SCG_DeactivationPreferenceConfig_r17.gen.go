// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SCG-DeactivationPreferenceConfig-r17 (line 26479).

var sCGDeactivationPreferenceConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scg-DeactivationPreferenceProhibitTimer-r17"},
	},
}

const (
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S0    = 0
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S1    = 1
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S2    = 2
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S4    = 3
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S8    = 4
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S10   = 5
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S15   = 6
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S30   = 7
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S60   = 8
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S120  = 9
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S180  = 10
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S240  = 11
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S300  = 12
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S600  = 13
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S900  = 14
	SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S1800 = 15
)

var sCGDeactivationPreferenceConfigR17ScgDeactivationPreferenceProhibitTimerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S0, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S1, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S2, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S4, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S8, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S10, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S15, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S30, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S60, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S120, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S180, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S240, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S300, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S600, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S900, SCG_DeactivationPreferenceConfig_r17_Scg_DeactivationPreferenceProhibitTimer_r17_S1800},
}

type SCG_DeactivationPreferenceConfig_r17 struct {
	Scg_DeactivationPreferenceProhibitTimer_r17 int64
}

func (ie *SCG_DeactivationPreferenceConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sCGDeactivationPreferenceConfigR17Constraints)
	_ = seq

	// 1. scg-DeactivationPreferenceProhibitTimer-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Scg_DeactivationPreferenceProhibitTimer_r17, sCGDeactivationPreferenceConfigR17ScgDeactivationPreferenceProhibitTimerR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SCG_DeactivationPreferenceConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sCGDeactivationPreferenceConfigR17Constraints)
	_ = seq

	// 1. scg-DeactivationPreferenceProhibitTimer-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(sCGDeactivationPreferenceConfigR17ScgDeactivationPreferenceProhibitTimerR17Constraints)
		if err != nil {
			return err
		}
		ie.Scg_DeactivationPreferenceProhibitTimer_r17 = v0
	}

	return nil
}
