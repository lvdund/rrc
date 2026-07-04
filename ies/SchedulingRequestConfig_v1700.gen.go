// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SchedulingRequestConfig-v1700 (line 14263).

var schedulingRequestConfigV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "schedulingRequestToAddModListExt-v1700", Optional: true},
	},
}

var schedulingRequestConfigV1700SchedulingRequestToAddModListExtV1700Constraints = per.SizeRange(1, common.MaxNrofSR_ConfigPerCellGroup)

type SchedulingRequestConfig_v1700 struct {
	SchedulingRequestToAddModListExt_v1700 []SchedulingRequestToAddModExt_v1700
}

func (ie *SchedulingRequestConfig_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(schedulingRequestConfigV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SchedulingRequestToAddModListExt_v1700 != nil}); err != nil {
		return err
	}

	// 2. schedulingRequestToAddModListExt-v1700: sequence-of
	{
		if ie.SchedulingRequestToAddModListExt_v1700 != nil {
			seqOf := e.NewSequenceOfEncoder(schedulingRequestConfigV1700SchedulingRequestToAddModListExtV1700Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SchedulingRequestToAddModListExt_v1700))); err != nil {
				return err
			}
			for i := range ie.SchedulingRequestToAddModListExt_v1700 {
				if err := ie.SchedulingRequestToAddModListExt_v1700[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SchedulingRequestConfig_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(schedulingRequestConfigV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. schedulingRequestToAddModListExt-v1700: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(schedulingRequestConfigV1700SchedulingRequestToAddModListExtV1700Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SchedulingRequestToAddModListExt_v1700 = make([]SchedulingRequestToAddModExt_v1700, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SchedulingRequestToAddModListExt_v1700[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
