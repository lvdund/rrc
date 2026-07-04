// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1920 (line 25903).

var uENRCapabilityV1920Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ntn-RedirectionWithSatelliteInfo-r19", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_NR_Capability_v1920_Ntn_RedirectionWithSatelliteInfo_r19_Supported = 0
)

var uENRCapabilityV1920NtnRedirectionWithSatelliteInfoR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1920_Ntn_RedirectionWithSatelliteInfo_r19_Supported},
}

type UE_NR_Capability_v1920 struct {
	Ntn_RedirectionWithSatelliteInfo_r19 *int64
	NonCriticalExtension                 *struct{}
}

func (ie *UE_NR_Capability_v1920) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1920Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ntn_RedirectionWithSatelliteInfo_r19 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. ntn-RedirectionWithSatelliteInfo-r19: enumerated
	{
		if ie.Ntn_RedirectionWithSatelliteInfo_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Ntn_RedirectionWithSatelliteInfo_r19, uENRCapabilityV1920NtnRedirectionWithSatelliteInfoR19Constraints); err != nil {
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

func (ie *UE_NR_Capability_v1920) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1920Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ntn-RedirectionWithSatelliteInfo-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1920NtnRedirectionWithSatelliteInfoR19Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_RedirectionWithSatelliteInfo_r19 = &idx
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
