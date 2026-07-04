// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DRX-ConfigExt-v1700 (line 8164).

var dRXConfigExtV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "drx-HARQ-RTT-TimerDL-r17"},
		{Name: "drx-HARQ-RTT-TimerUL-r17"},
	},
}

var dRXConfigExtV1700DrxHARQRTTTimerDLR17Constraints = per.Constrained(0, 448)

var dRXConfigExtV1700DrxHARQRTTTimerULR17Constraints = per.Constrained(0, 448)

type DRX_ConfigExt_v1700 struct {
	Drx_HARQ_RTT_TimerDL_r17 int64
	Drx_HARQ_RTT_TimerUL_r17 int64
}

func (ie *DRX_ConfigExt_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRXConfigExtV1700Constraints)
	_ = seq

	// 1. drx-HARQ-RTT-TimerDL-r17: integer
	{
		if err := e.EncodeInteger(ie.Drx_HARQ_RTT_TimerDL_r17, dRXConfigExtV1700DrxHARQRTTTimerDLR17Constraints); err != nil {
			return err
		}
	}

	// 2. drx-HARQ-RTT-TimerUL-r17: integer
	{
		if err := e.EncodeInteger(ie.Drx_HARQ_RTT_TimerUL_r17, dRXConfigExtV1700DrxHARQRTTTimerULR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DRX_ConfigExt_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRXConfigExtV1700Constraints)
	_ = seq

	// 1. drx-HARQ-RTT-TimerDL-r17: integer
	{
		v0, err := d.DecodeInteger(dRXConfigExtV1700DrxHARQRTTTimerDLR17Constraints)
		if err != nil {
			return err
		}
		ie.Drx_HARQ_RTT_TimerDL_r17 = v0
	}

	// 2. drx-HARQ-RTT-TimerUL-r17: integer
	{
		v1, err := d.DecodeInteger(dRXConfigExtV1700DrxHARQRTTTimerULR17Constraints)
		if err != nil {
			return err
		}
		ie.Drx_HARQ_RTT_TimerUL_r17 = v1
	}

	return nil
}
