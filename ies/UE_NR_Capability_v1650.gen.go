// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1650 (line 25751).

var uENRCapabilityV1650Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mpsPriorityIndication-r16", Optional: true},
		{Name: "highSpeedParameters-v1650", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_NR_Capability_v1650_MpsPriorityIndication_r16_Supported = 0
)

var uENRCapabilityV1650MpsPriorityIndicationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1650_MpsPriorityIndication_r16_Supported},
}

type UE_NR_Capability_v1650 struct {
	MpsPriorityIndication_r16 *int64
	HighSpeedParameters_v1650 *HighSpeedParameters_v1650
	NonCriticalExtension      *UE_NR_Capability_v1690
}

func (ie *UE_NR_Capability_v1650) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1650Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MpsPriorityIndication_r16 != nil, ie.HighSpeedParameters_v1650 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. mpsPriorityIndication-r16: enumerated
	{
		if ie.MpsPriorityIndication_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MpsPriorityIndication_r16, uENRCapabilityV1650MpsPriorityIndicationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. highSpeedParameters-v1650: ref
	{
		if ie.HighSpeedParameters_v1650 != nil {
			if err := ie.HighSpeedParameters_v1650.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_Capability_v1650) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1650Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mpsPriorityIndication-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1650MpsPriorityIndicationR16Constraints)
			if err != nil {
				return err
			}
			ie.MpsPriorityIndication_r16 = &idx
		}
	}

	// 3. highSpeedParameters-v1650: ref
	{
		if seq.IsComponentPresent(1) {
			ie.HighSpeedParameters_v1650 = new(HighSpeedParameters_v1650)
			if err := ie.HighSpeedParameters_v1650.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1690)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
