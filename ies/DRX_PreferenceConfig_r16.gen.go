// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DRX-PreferenceConfig-r16 (line 26432).

var dRXPreferenceConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "drx-PreferenceProhibitTimer-r16"},
	},
}

const (
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S0     = 0
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S0dot5 = 1
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S1     = 2
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S2     = 3
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S3     = 4
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S4     = 5
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S5     = 6
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S6     = 7
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S7     = 8
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S8     = 9
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S9     = 10
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S10    = 11
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S20    = 12
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S30    = 13
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_Spare2 = 14
	DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_Spare1 = 15
)

var dRXPreferenceConfigR16DrxPreferenceProhibitTimerR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S0, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S0dot5, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S1, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S2, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S3, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S4, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S5, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S6, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S7, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S8, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S9, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S10, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S20, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_S30, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_Spare2, DRX_PreferenceConfig_r16_Drx_PreferenceProhibitTimer_r16_Spare1},
}

type DRX_PreferenceConfig_r16 struct {
	Drx_PreferenceProhibitTimer_r16 int64
}

func (ie *DRX_PreferenceConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRXPreferenceConfigR16Constraints)
	_ = seq

	// 1. drx-PreferenceProhibitTimer-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Drx_PreferenceProhibitTimer_r16, dRXPreferenceConfigR16DrxPreferenceProhibitTimerR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DRX_PreferenceConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRXPreferenceConfigR16Constraints)
	_ = seq

	// 1. drx-PreferenceProhibitTimer-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(dRXPreferenceConfigR16DrxPreferenceProhibitTimerR16Constraints)
		if err != nil {
			return err
		}
		ie.Drx_PreferenceProhibitTimer_r16 = v0
	}

	return nil
}
