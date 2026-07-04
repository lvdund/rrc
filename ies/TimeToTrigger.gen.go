// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TimeToTrigger (line 16155).

const (
	TimeToTrigger_Ms0    = 0
	TimeToTrigger_Ms40   = 1
	TimeToTrigger_Ms64   = 2
	TimeToTrigger_Ms80   = 3
	TimeToTrigger_Ms100  = 4
	TimeToTrigger_Ms128  = 5
	TimeToTrigger_Ms160  = 6
	TimeToTrigger_Ms256  = 7
	TimeToTrigger_Ms320  = 8
	TimeToTrigger_Ms480  = 9
	TimeToTrigger_Ms512  = 10
	TimeToTrigger_Ms640  = 11
	TimeToTrigger_Ms1024 = 12
	TimeToTrigger_Ms1280 = 13
	TimeToTrigger_Ms2560 = 14
	TimeToTrigger_Ms5120 = 15
)

var timeToTriggerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TimeToTrigger_Ms0, TimeToTrigger_Ms40, TimeToTrigger_Ms64, TimeToTrigger_Ms80, TimeToTrigger_Ms100, TimeToTrigger_Ms128, TimeToTrigger_Ms160, TimeToTrigger_Ms256, TimeToTrigger_Ms320, TimeToTrigger_Ms480, TimeToTrigger_Ms512, TimeToTrigger_Ms640, TimeToTrigger_Ms1024, TimeToTrigger_Ms1280, TimeToTrigger_Ms2560, TimeToTrigger_Ms5120},
}

type TimeToTrigger struct {
	Value int64
}

func (v *TimeToTrigger) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, timeToTriggerConstraints)
}

func (v *TimeToTrigger) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(timeToTriggerConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
