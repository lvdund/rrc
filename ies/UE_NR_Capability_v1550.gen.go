// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1550 (line 25680).

var uENRCapabilityV1550Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reducedCP-Latency", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_NR_Capability_v1550_ReducedCP_Latency_Supported = 0
)

var uENRCapabilityV1550ReducedCPLatencyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1550_ReducedCP_Latency_Supported},
}

type UE_NR_Capability_v1550 struct {
	ReducedCP_Latency    *int64
	NonCriticalExtension *UE_NR_Capability_v1560
}

func (ie *UE_NR_Capability_v1550) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1550Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReducedCP_Latency != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. reducedCP-Latency: enumerated
	{
		if ie.ReducedCP_Latency != nil {
			if err := e.EncodeEnumerated(*ie.ReducedCP_Latency, uENRCapabilityV1550ReducedCPLatencyConstraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_Capability_v1550) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1550Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. reducedCP-Latency: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1550ReducedCPLatencyConstraints)
			if err != nil {
				return err
			}
			ie.ReducedCP_Latency = &idx
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1560)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
