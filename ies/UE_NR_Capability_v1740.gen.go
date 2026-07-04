// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1740 (line 25820).

var uENRCapabilityV1740Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "redCapParameters-v1740"},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_NR_Capability_v1740 struct {
	RedCapParameters_v1740 RedCapParameters_v1740
	NonCriticalExtension   *UE_NR_Capability_v1750
}

func (ie *UE_NR_Capability_v1740) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1740Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. redCapParameters-v1740: ref
	{
		if err := ie.RedCapParameters_v1740.Encode(e); err != nil {
			return err
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

func (ie *UE_NR_Capability_v1740) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1740Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. redCapParameters-v1740: ref
	{
		if err := ie.RedCapParameters_v1740.Decode(d); err != nil {
			return err
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1750)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
