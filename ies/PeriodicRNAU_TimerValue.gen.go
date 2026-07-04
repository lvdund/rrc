// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PeriodicRNAU-TimerValue (line 1314).
// PeriodicRNAU-TimerValue ::=         ENUMERATED { min5, min10, min20, min30, min60, min120, min360, min720}

const (
	PeriodicRNAU_TimerValue_Min5   = 0
	PeriodicRNAU_TimerValue_Min10  = 1
	PeriodicRNAU_TimerValue_Min20  = 2
	PeriodicRNAU_TimerValue_Min30  = 3
	PeriodicRNAU_TimerValue_Min60  = 4
	PeriodicRNAU_TimerValue_Min120 = 5
	PeriodicRNAU_TimerValue_Min360 = 6
	PeriodicRNAU_TimerValue_Min720 = 7
)

var periodicRNAUTimerValueConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PeriodicRNAU_TimerValue_Min5, PeriodicRNAU_TimerValue_Min10, PeriodicRNAU_TimerValue_Min20, PeriodicRNAU_TimerValue_Min30, PeriodicRNAU_TimerValue_Min60, PeriodicRNAU_TimerValue_Min120, PeriodicRNAU_TimerValue_Min360, PeriodicRNAU_TimerValue_Min720},
}

type PeriodicRNAU_TimerValue struct {
	Value int64
}

func (v *PeriodicRNAU_TimerValue) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, periodicRNAUTimerValueConstraints)
}

func (v *PeriodicRNAU_TimerValue) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(periodicRNAUTimerValueConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
