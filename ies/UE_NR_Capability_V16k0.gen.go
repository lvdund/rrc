// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v16k0 (line 25786).

var uENRCapabilityV16k0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "featureSets-v16k0", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_NR_Capability_V16k0 struct {
	FeatureSets_V16k0    *FeatureSets_V16k0
	NonCriticalExtension *struct{}
}

func (ie *UE_NR_Capability_V16k0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV16k0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FeatureSets_V16k0 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. featureSets-v16k0: ref
	{
		if ie.FeatureSets_V16k0 != nil {
			if err := ie.FeatureSets_V16k0.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *UE_NR_Capability_V16k0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV16k0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. featureSets-v16k0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FeatureSets_V16k0 = new(FeatureSets_V16k0)
			if err := ie.FeatureSets_V16k0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
