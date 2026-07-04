// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OnDemandSIB-Request-r16 (line 1094).

var onDemandSIBRequestR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "onDemandSIB-RequestProhibitTimer-r16"},
	},
}

const (
	OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S0     = 0
	OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S0dot5 = 1
	OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S1     = 2
	OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S2     = 3
	OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S5     = 4
	OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S10    = 5
	OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S20    = 6
	OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S30    = 7
)

var onDemandSIBRequestR16OnDemandSIBRequestProhibitTimerR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S0, OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S0dot5, OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S1, OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S2, OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S5, OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S10, OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S20, OnDemandSIB_Request_r16_OnDemandSIB_RequestProhibitTimer_r16_S30},
}

type OnDemandSIB_Request_r16 struct {
	OnDemandSIB_RequestProhibitTimer_r16 int64
}

func (ie *OnDemandSIB_Request_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(onDemandSIBRequestR16Constraints)
	_ = seq

	// 1. onDemandSIB-RequestProhibitTimer-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.OnDemandSIB_RequestProhibitTimer_r16, onDemandSIBRequestR16OnDemandSIBRequestProhibitTimerR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *OnDemandSIB_Request_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(onDemandSIBRequestR16Constraints)
	_ = seq

	// 1. onDemandSIB-RequestProhibitTimer-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(onDemandSIBRequestR16OnDemandSIBRequestProhibitTimerR16Constraints)
		if err != nil {
			return err
		}
		ie.OnDemandSIB_RequestProhibitTimer_r16 = v0
	}

	return nil
}
