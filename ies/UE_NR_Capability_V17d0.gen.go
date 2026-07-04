// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v17d0 (line 25843).

var uENRCapabilityV17d0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "featureSets-v17d0", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_NR_Capability_V17d0 struct {
	FeatureSets_V17d0    *FeatureSets_V17d0
	NonCriticalExtension *struct{}
}

func (ie *UE_NR_Capability_V17d0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV17d0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FeatureSets_V17d0 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. featureSets-v17d0: ref
	{
		if ie.FeatureSets_V17d0 != nil {
			if err := ie.FeatureSets_V17d0.Encode(e); err != nil {
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

func (ie *UE_NR_Capability_V17d0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV17d0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. featureSets-v17d0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FeatureSets_V17d0 = new(FeatureSets_V17d0)
			if err := ie.FeatureSets_V17d0.Decode(d); err != nil {
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
