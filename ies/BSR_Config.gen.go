// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BSR-Config (line 5218).

var bSRConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicBSR-Timer"},
		{Name: "retxBSR-Timer"},
		{Name: "logicalChannelSR-DelayTimer", Optional: true},
	},
}

const (
	BSR_Config_PeriodicBSR_Timer_Sf1      = 0
	BSR_Config_PeriodicBSR_Timer_Sf5      = 1
	BSR_Config_PeriodicBSR_Timer_Sf10     = 2
	BSR_Config_PeriodicBSR_Timer_Sf16     = 3
	BSR_Config_PeriodicBSR_Timer_Sf20     = 4
	BSR_Config_PeriodicBSR_Timer_Sf32     = 5
	BSR_Config_PeriodicBSR_Timer_Sf40     = 6
	BSR_Config_PeriodicBSR_Timer_Sf64     = 7
	BSR_Config_PeriodicBSR_Timer_Sf80     = 8
	BSR_Config_PeriodicBSR_Timer_Sf128    = 9
	BSR_Config_PeriodicBSR_Timer_Sf160    = 10
	BSR_Config_PeriodicBSR_Timer_Sf320    = 11
	BSR_Config_PeriodicBSR_Timer_Sf640    = 12
	BSR_Config_PeriodicBSR_Timer_Sf1280   = 13
	BSR_Config_PeriodicBSR_Timer_Sf2560   = 14
	BSR_Config_PeriodicBSR_Timer_Infinity = 15
)

var bSRConfigPeriodicBSRTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BSR_Config_PeriodicBSR_Timer_Sf1, BSR_Config_PeriodicBSR_Timer_Sf5, BSR_Config_PeriodicBSR_Timer_Sf10, BSR_Config_PeriodicBSR_Timer_Sf16, BSR_Config_PeriodicBSR_Timer_Sf20, BSR_Config_PeriodicBSR_Timer_Sf32, BSR_Config_PeriodicBSR_Timer_Sf40, BSR_Config_PeriodicBSR_Timer_Sf64, BSR_Config_PeriodicBSR_Timer_Sf80, BSR_Config_PeriodicBSR_Timer_Sf128, BSR_Config_PeriodicBSR_Timer_Sf160, BSR_Config_PeriodicBSR_Timer_Sf320, BSR_Config_PeriodicBSR_Timer_Sf640, BSR_Config_PeriodicBSR_Timer_Sf1280, BSR_Config_PeriodicBSR_Timer_Sf2560, BSR_Config_PeriodicBSR_Timer_Infinity},
}

const (
	BSR_Config_RetxBSR_Timer_Sf10    = 0
	BSR_Config_RetxBSR_Timer_Sf20    = 1
	BSR_Config_RetxBSR_Timer_Sf40    = 2
	BSR_Config_RetxBSR_Timer_Sf80    = 3
	BSR_Config_RetxBSR_Timer_Sf160   = 4
	BSR_Config_RetxBSR_Timer_Sf320   = 5
	BSR_Config_RetxBSR_Timer_Sf640   = 6
	BSR_Config_RetxBSR_Timer_Sf1280  = 7
	BSR_Config_RetxBSR_Timer_Sf2560  = 8
	BSR_Config_RetxBSR_Timer_Sf5120  = 9
	BSR_Config_RetxBSR_Timer_Sf10240 = 10
	BSR_Config_RetxBSR_Timer_Spare5  = 11
	BSR_Config_RetxBSR_Timer_Spare4  = 12
	BSR_Config_RetxBSR_Timer_Spare3  = 13
	BSR_Config_RetxBSR_Timer_Spare2  = 14
	BSR_Config_RetxBSR_Timer_Spare1  = 15
)

var bSRConfigRetxBSRTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BSR_Config_RetxBSR_Timer_Sf10, BSR_Config_RetxBSR_Timer_Sf20, BSR_Config_RetxBSR_Timer_Sf40, BSR_Config_RetxBSR_Timer_Sf80, BSR_Config_RetxBSR_Timer_Sf160, BSR_Config_RetxBSR_Timer_Sf320, BSR_Config_RetxBSR_Timer_Sf640, BSR_Config_RetxBSR_Timer_Sf1280, BSR_Config_RetxBSR_Timer_Sf2560, BSR_Config_RetxBSR_Timer_Sf5120, BSR_Config_RetxBSR_Timer_Sf10240, BSR_Config_RetxBSR_Timer_Spare5, BSR_Config_RetxBSR_Timer_Spare4, BSR_Config_RetxBSR_Timer_Spare3, BSR_Config_RetxBSR_Timer_Spare2, BSR_Config_RetxBSR_Timer_Spare1},
}

const (
	BSR_Config_LogicalChannelSR_DelayTimer_Sf20   = 0
	BSR_Config_LogicalChannelSR_DelayTimer_Sf40   = 1
	BSR_Config_LogicalChannelSR_DelayTimer_Sf64   = 2
	BSR_Config_LogicalChannelSR_DelayTimer_Sf128  = 3
	BSR_Config_LogicalChannelSR_DelayTimer_Sf512  = 4
	BSR_Config_LogicalChannelSR_DelayTimer_Sf1024 = 5
	BSR_Config_LogicalChannelSR_DelayTimer_Sf2560 = 6
	BSR_Config_LogicalChannelSR_DelayTimer_Spare1 = 7
)

var bSRConfigLogicalChannelSRDelayTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BSR_Config_LogicalChannelSR_DelayTimer_Sf20, BSR_Config_LogicalChannelSR_DelayTimer_Sf40, BSR_Config_LogicalChannelSR_DelayTimer_Sf64, BSR_Config_LogicalChannelSR_DelayTimer_Sf128, BSR_Config_LogicalChannelSR_DelayTimer_Sf512, BSR_Config_LogicalChannelSR_DelayTimer_Sf1024, BSR_Config_LogicalChannelSR_DelayTimer_Sf2560, BSR_Config_LogicalChannelSR_DelayTimer_Spare1},
}

type BSR_Config struct {
	PeriodicBSR_Timer           int64
	RetxBSR_Timer               int64
	LogicalChannelSR_DelayTimer *int64
}

func (ie *BSR_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bSRConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LogicalChannelSR_DelayTimer != nil}); err != nil {
		return err
	}

	// 3. periodicBSR-Timer: enumerated
	{
		if err := e.EncodeEnumerated(ie.PeriodicBSR_Timer, bSRConfigPeriodicBSRTimerConstraints); err != nil {
			return err
		}
	}

	// 4. retxBSR-Timer: enumerated
	{
		if err := e.EncodeEnumerated(ie.RetxBSR_Timer, bSRConfigRetxBSRTimerConstraints); err != nil {
			return err
		}
	}

	// 5. logicalChannelSR-DelayTimer: enumerated
	{
		if ie.LogicalChannelSR_DelayTimer != nil {
			if err := e.EncodeEnumerated(*ie.LogicalChannelSR_DelayTimer, bSRConfigLogicalChannelSRDelayTimerConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BSR_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bSRConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. periodicBSR-Timer: enumerated
	{
		v0, err := d.DecodeEnumerated(bSRConfigPeriodicBSRTimerConstraints)
		if err != nil {
			return err
		}
		ie.PeriodicBSR_Timer = v0
	}

	// 4. retxBSR-Timer: enumerated
	{
		v1, err := d.DecodeEnumerated(bSRConfigRetxBSRTimerConstraints)
		if err != nil {
			return err
		}
		ie.RetxBSR_Timer = v1
	}

	// 5. logicalChannelSR-DelayTimer: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(bSRConfigLogicalChannelSRDelayTimerConstraints)
			if err != nil {
				return err
			}
			ie.LogicalChannelSR_DelayTimer = &idx
		}
	}

	return nil
}
