// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v15t0 (line 25714).

var uENRCapabilityV15t0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "featureSets-v15t0", Optional: true},
		{Name: "measAndMobParameters-v15t0", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_NR_Capability_V15t0 struct {
	FeatureSets_V15t0          *FeatureSets_V15t0
	MeasAndMobParameters_V15t0 *MeasAndMobParameters_V15t0
	NonCriticalExtension       *struct{}
}

func (ie *UE_NR_Capability_V15t0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV15t0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FeatureSets_V15t0 != nil, ie.MeasAndMobParameters_V15t0 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. featureSets-v15t0: ref
	{
		if ie.FeatureSets_V15t0 != nil {
			if err := ie.FeatureSets_V15t0.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. measAndMobParameters-v15t0: ref
	{
		if ie.MeasAndMobParameters_V15t0 != nil {
			if err := ie.MeasAndMobParameters_V15t0.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *UE_NR_Capability_V15t0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV15t0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. featureSets-v15t0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FeatureSets_V15t0 = new(FeatureSets_V15t0)
			if err := ie.FeatureSets_V15t0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. measAndMobParameters-v15t0: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MeasAndMobParameters_V15t0 = new(MeasAndMobParameters_V15t0)
			if err := ie.MeasAndMobParameters_V15t0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
