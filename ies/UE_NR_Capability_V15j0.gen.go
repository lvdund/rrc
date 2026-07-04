// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v15j0 (line 25708).

var uENRCapabilityV15j0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uENRCapabilityV15j0LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type UE_NR_Capability_V15j0 struct {
	LateNonCriticalExtension []byte
	NonCriticalExtension     *UE_NR_Capability_V16a0
}

func (ie *UE_NR_Capability_V15j0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV15j0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uENRCapabilityV15j0LateNonCriticalExtensionConstraints); err != nil {
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

func (ie *UE_NR_Capability_V15j0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV15j0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(uENRCapabilityV15j0LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UE_NR_Capability_V16a0)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
