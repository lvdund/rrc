// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DRX-ConfigSL-r17 (line 8227).

var dRXConfigSLR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "drx-HARQ-RTT-TimerSL-r17"},
		{Name: "drx-RetransmissionTimerSL-r17"},
	},
}

var dRXConfigSLR17DrxHARQRTTTimerSLR17Constraints = per.Constrained(0, 56)

const (
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl0     = 0
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl1     = 1
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl2     = 2
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl4     = 3
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl6     = 4
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl8     = 5
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl16    = 6
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl24    = 7
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl33    = 8
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl40    = 9
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl64    = 10
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl80    = 11
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl96    = 12
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl112   = 13
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl128   = 14
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl160   = 15
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl320   = 16
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare15 = 17
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare14 = 18
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare13 = 19
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare12 = 20
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare11 = 21
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare10 = 22
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare9  = 23
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare8  = 24
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare7  = 25
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare6  = 26
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare5  = 27
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare4  = 28
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare3  = 29
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare2  = 30
	DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare1  = 31
)

var dRXConfigSLR17DrxRetransmissionTimerSLR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl0, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl1, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl2, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl4, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl6, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl8, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl16, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl24, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl33, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl40, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl64, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl80, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl96, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl112, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl128, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl160, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Sl320, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare15, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare14, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare13, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare12, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare11, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare10, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare9, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare8, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare7, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare6, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare5, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare4, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare3, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare2, DRX_ConfigSL_r17_Drx_RetransmissionTimerSL_r17_Spare1},
}

type DRX_ConfigSL_r17 struct {
	Drx_HARQ_RTT_TimerSL_r17      int64
	Drx_RetransmissionTimerSL_r17 int64
}

func (ie *DRX_ConfigSL_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRXConfigSLR17Constraints)
	_ = seq

	// 1. drx-HARQ-RTT-TimerSL-r17: integer
	{
		if err := e.EncodeInteger(ie.Drx_HARQ_RTT_TimerSL_r17, dRXConfigSLR17DrxHARQRTTTimerSLR17Constraints); err != nil {
			return err
		}
	}

	// 2. drx-RetransmissionTimerSL-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Drx_RetransmissionTimerSL_r17, dRXConfigSLR17DrxRetransmissionTimerSLR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DRX_ConfigSL_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRXConfigSLR17Constraints)
	_ = seq

	// 1. drx-HARQ-RTT-TimerSL-r17: integer
	{
		v0, err := d.DecodeInteger(dRXConfigSLR17DrxHARQRTTTimerSLR17Constraints)
		if err != nil {
			return err
		}
		ie.Drx_HARQ_RTT_TimerSL_r17 = v0
	}

	// 2. drx-RetransmissionTimerSL-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(dRXConfigSLR17DrxRetransmissionTimerSLR17Constraints)
		if err != nil {
			return err
		}
		ie.Drx_RetransmissionTimerSL_r17 = v1
	}

	return nil
}
