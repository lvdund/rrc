// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1750 (line 25825).

var uENRCapabilityV1750Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "crossCarrierSchedulingConfigurationRelease-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_NR_Capability_v1750_CrossCarrierSchedulingConfigurationRelease_r17_Supported = 0
)

var uENRCapabilityV1750CrossCarrierSchedulingConfigurationReleaseR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1750_CrossCarrierSchedulingConfigurationRelease_r17_Supported},
}

type UE_NR_Capability_v1750 struct {
	CrossCarrierSchedulingConfigurationRelease_r17 *int64
	NonCriticalExtension                           *UE_NR_Capability_v1800
}

func (ie *UE_NR_Capability_v1750) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1750Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CrossCarrierSchedulingConfigurationRelease_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. crossCarrierSchedulingConfigurationRelease-r17: enumerated
	{
		if ie.CrossCarrierSchedulingConfigurationRelease_r17 != nil {
			if err := e.EncodeEnumerated(*ie.CrossCarrierSchedulingConfigurationRelease_r17, uENRCapabilityV1750CrossCarrierSchedulingConfigurationReleaseR17Constraints); err != nil {
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

func (ie *UE_NR_Capability_v1750) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1750Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. crossCarrierSchedulingConfigurationRelease-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1750CrossCarrierSchedulingConfigurationReleaseR17Constraints)
			if err != nil {
				return err
			}
			ie.CrossCarrierSchedulingConfigurationRelease_r17 = &idx
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1800)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
