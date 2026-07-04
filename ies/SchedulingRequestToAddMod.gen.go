// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SchedulingRequestToAddMod (line 14257).

var schedulingRequestToAddModConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "schedulingRequestId"},
		{Name: "sr-ProhibitTimer", Optional: true},
		{Name: "sr-TransMax"},
	},
}

const (
	SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms1   = 0
	SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms2   = 1
	SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms4   = 2
	SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms8   = 3
	SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms16  = 4
	SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms32  = 5
	SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms64  = 6
	SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms128 = 7
)

var schedulingRequestToAddModSrProhibitTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms1, SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms2, SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms4, SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms8, SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms16, SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms32, SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms64, SchedulingRequestToAddMod_Sr_ProhibitTimer_Ms128},
}

const (
	SchedulingRequestToAddMod_Sr_TransMax_N4     = 0
	SchedulingRequestToAddMod_Sr_TransMax_N8     = 1
	SchedulingRequestToAddMod_Sr_TransMax_N16    = 2
	SchedulingRequestToAddMod_Sr_TransMax_N32    = 3
	SchedulingRequestToAddMod_Sr_TransMax_N64    = 4
	SchedulingRequestToAddMod_Sr_TransMax_Spare3 = 5
	SchedulingRequestToAddMod_Sr_TransMax_Spare2 = 6
	SchedulingRequestToAddMod_Sr_TransMax_Spare1 = 7
)

var schedulingRequestToAddModSrTransMaxConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SchedulingRequestToAddMod_Sr_TransMax_N4, SchedulingRequestToAddMod_Sr_TransMax_N8, SchedulingRequestToAddMod_Sr_TransMax_N16, SchedulingRequestToAddMod_Sr_TransMax_N32, SchedulingRequestToAddMod_Sr_TransMax_N64, SchedulingRequestToAddMod_Sr_TransMax_Spare3, SchedulingRequestToAddMod_Sr_TransMax_Spare2, SchedulingRequestToAddMod_Sr_TransMax_Spare1},
}

type SchedulingRequestToAddMod struct {
	SchedulingRequestId SchedulingRequestId
	Sr_ProhibitTimer    *int64
	Sr_TransMax         int64
}

func (ie *SchedulingRequestToAddMod) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(schedulingRequestToAddModConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sr_ProhibitTimer != nil}); err != nil {
		return err
	}

	// 2. schedulingRequestId: ref
	{
		if err := ie.SchedulingRequestId.Encode(e); err != nil {
			return err
		}
	}

	// 3. sr-ProhibitTimer: enumerated
	{
		if ie.Sr_ProhibitTimer != nil {
			if err := e.EncodeEnumerated(*ie.Sr_ProhibitTimer, schedulingRequestToAddModSrProhibitTimerConstraints); err != nil {
				return err
			}
		}
	}

	// 4. sr-TransMax: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sr_TransMax, schedulingRequestToAddModSrTransMaxConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SchedulingRequestToAddMod) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(schedulingRequestToAddModConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. schedulingRequestId: ref
	{
		if err := ie.SchedulingRequestId.Decode(d); err != nil {
			return err
		}
	}

	// 3. sr-ProhibitTimer: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(schedulingRequestToAddModSrProhibitTimerConstraints)
			if err != nil {
				return err
			}
			ie.Sr_ProhibitTimer = &idx
		}
	}

	// 4. sr-TransMax: enumerated
	{
		v2, err := d.DecodeEnumerated(schedulingRequestToAddModSrTransMaxConstraints)
		if err != nil {
			return err
		}
		ie.Sr_TransMax = v2
	}

	return nil
}
