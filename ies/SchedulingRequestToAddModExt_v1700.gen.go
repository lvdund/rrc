// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SchedulingRequestToAddModExt-v1700 (line 14268).

var schedulingRequestToAddModExtV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sr-ProhibitTimer-v1700", Optional: true},
	},
}

const (
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms192  = 0
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms256  = 1
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms320  = 2
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms384  = 3
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms448  = 4
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms512  = 5
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms576  = 6
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms640  = 7
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms1082 = 8
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Spare7 = 9
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Spare6 = 10
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Spare5 = 11
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Spare4 = 12
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Spare3 = 13
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Spare2 = 14
	SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Spare1 = 15
)

var schedulingRequestToAddModExtV1700SrProhibitTimerV1700Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms192, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms256, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms320, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms384, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms448, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms512, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms576, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms640, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Ms1082, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Spare7, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Spare6, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Spare5, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Spare4, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Spare3, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Spare2, SchedulingRequestToAddModExt_v1700_Sr_ProhibitTimer_v1700_Spare1},
}

type SchedulingRequestToAddModExt_v1700 struct {
	Sr_ProhibitTimer_v1700 *int64
}

func (ie *SchedulingRequestToAddModExt_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(schedulingRequestToAddModExtV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sr_ProhibitTimer_v1700 != nil}); err != nil {
		return err
	}

	// 2. sr-ProhibitTimer-v1700: enumerated
	{
		if ie.Sr_ProhibitTimer_v1700 != nil {
			if err := e.EncodeEnumerated(*ie.Sr_ProhibitTimer_v1700, schedulingRequestToAddModExtV1700SrProhibitTimerV1700Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SchedulingRequestToAddModExt_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(schedulingRequestToAddModExtV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sr-ProhibitTimer-v1700: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(schedulingRequestToAddModExtV1700SrProhibitTimerV1700Constraints)
			if err != nil {
				return err
			}
			ie.Sr_ProhibitTimer_v1700 = &idx
		}
	}

	return nil
}
