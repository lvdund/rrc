// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TimeAlignmentTimer (line 16150).
// TimeAlignmentTimer ::=              ENUMERATED {ms500, ms750, ms1280, ms1920, ms2560, ms5120, ms10240, infinity}

const (
	TimeAlignmentTimer_Ms500    = 0
	TimeAlignmentTimer_Ms750    = 1
	TimeAlignmentTimer_Ms1280   = 2
	TimeAlignmentTimer_Ms1920   = 3
	TimeAlignmentTimer_Ms2560   = 4
	TimeAlignmentTimer_Ms5120   = 5
	TimeAlignmentTimer_Ms10240  = 6
	TimeAlignmentTimer_Infinity = 7
)

var timeAlignmentTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TimeAlignmentTimer_Ms500, TimeAlignmentTimer_Ms750, TimeAlignmentTimer_Ms1280, TimeAlignmentTimer_Ms1920, TimeAlignmentTimer_Ms2560, TimeAlignmentTimer_Ms5120, TimeAlignmentTimer_Ms10240, TimeAlignmentTimer_Infinity},
}

type TimeAlignmentTimer struct {
	Value int64
}

func (v *TimeAlignmentTimer) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, timeAlignmentTimerConstraints)
}

func (v *TimeAlignmentTimer) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(timeAlignmentTimerConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
