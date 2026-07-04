// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v16d0 (line 25774).

var uENRCapabilityV16d0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "featureSets-v16d0", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_NR_Capability_V16d0 struct {
	FeatureSets_V16d0    *FeatureSets_V16d0
	NonCriticalExtension *UE_NR_Capability_V16j0
}

func (ie *UE_NR_Capability_V16d0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV16d0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FeatureSets_V16d0 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. featureSets-v16d0: ref
	{
		if ie.FeatureSets_V16d0 != nil {
			if err := ie.FeatureSets_V16d0.Encode(e); err != nil {
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

func (ie *UE_NR_Capability_V16d0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV16d0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. featureSets-v16d0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FeatureSets_V16d0 = new(FeatureSets_V16d0)
			if err := ie.FeatureSets_V16d0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UE_NR_Capability_V16j0)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
