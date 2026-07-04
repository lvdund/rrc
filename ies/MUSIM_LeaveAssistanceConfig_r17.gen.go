// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-LeaveAssistanceConfig-r17 (line 26393).

var mUSIMLeaveAssistanceConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "musim-LeaveWithoutResponseTimer-r17"},
	},
}

const (
	MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Ms10   = 0
	MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Ms20   = 1
	MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Ms40   = 2
	MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Ms60   = 3
	MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Ms80   = 4
	MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Ms100  = 5
	MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Spare2 = 6
	MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Spare1 = 7
)

var mUSIMLeaveAssistanceConfigR17MusimLeaveWithoutResponseTimerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Ms10, MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Ms20, MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Ms40, MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Ms60, MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Ms80, MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Ms100, MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Spare2, MUSIM_LeaveAssistanceConfig_r17_Musim_LeaveWithoutResponseTimer_r17_Spare1},
}

type MUSIM_LeaveAssistanceConfig_r17 struct {
	Musim_LeaveWithoutResponseTimer_r17 int64
}

func (ie *MUSIM_LeaveAssistanceConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMLeaveAssistanceConfigR17Constraints)
	_ = seq

	// 1. musim-LeaveWithoutResponseTimer-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Musim_LeaveWithoutResponseTimer_r17, mUSIMLeaveAssistanceConfigR17MusimLeaveWithoutResponseTimerR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MUSIM_LeaveAssistanceConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMLeaveAssistanceConfigR17Constraints)
	_ = seq

	// 1. musim-LeaveWithoutResponseTimer-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(mUSIMLeaveAssistanceConfigR17MusimLeaveWithoutResponseTimerR17Constraints)
		if err != nil {
			return err
		}
		ie.Musim_LeaveWithoutResponseTimer_r17 = v0
	}

	return nil
}
