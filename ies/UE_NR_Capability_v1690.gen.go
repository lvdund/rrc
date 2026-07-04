// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1690 (line 25757).

var uENRCapabilityV1690Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-RRC-Segmentation-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_NR_Capability_v1690_Ul_RRC_Segmentation_r16_Supported = 0
)

var uENRCapabilityV1690UlRRCSegmentationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1690_Ul_RRC_Segmentation_r16_Supported},
}

type UE_NR_Capability_v1690 struct {
	Ul_RRC_Segmentation_r16 *int64
	NonCriticalExtension    *UE_NR_Capability_v1700
}

func (ie *UE_NR_Capability_v1690) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1690Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ul_RRC_Segmentation_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. ul-RRC-Segmentation-r16: enumerated
	{
		if ie.Ul_RRC_Segmentation_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_RRC_Segmentation_r16, uENRCapabilityV1690UlRRCSegmentationR16Constraints); err != nil {
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

func (ie *UE_NR_Capability_v1690) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1690Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ul-RRC-Segmentation-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1690UlRRCSegmentationR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_RRC_Segmentation_r16 = &idx
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1700)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
