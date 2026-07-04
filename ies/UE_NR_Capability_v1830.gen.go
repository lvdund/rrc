// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1830 (line 25875).

var uENRCapabilityV1830Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sib19-Support-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_NR_Capability_v1830_Sib19_Support_r18_Supported = 0
)

var uENRCapabilityV1830Sib19SupportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1830_Sib19_Support_r18_Supported},
}

type UE_NR_Capability_v1830 struct {
	Sib19_Support_r18    *int64
	NonCriticalExtension *UE_NR_Capability_v1860
}

func (ie *UE_NR_Capability_v1830) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1830Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sib19_Support_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. sib19-Support-r18: enumerated
	{
		if ie.Sib19_Support_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sib19_Support_r18, uENRCapabilityV1830Sib19SupportR18Constraints); err != nil {
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

func (ie *UE_NR_Capability_v1830) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1830Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sib19-Support-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1830Sib19SupportR18Constraints)
			if err != nil {
				return err
			}
			ie.Sib19_Support_r18 = &idx
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1860)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
