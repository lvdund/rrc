// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v17c0 (line 25838).

var uENRCapabilityV17c0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mac-Parameters-v17c0", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_NR_Capability_V17c0 struct {
	Mac_Parameters_V17c0 *MAC_Parameters_V17c0
	NonCriticalExtension *UE_NR_Capability_V17d0
}

func (ie *UE_NR_Capability_V17c0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV17c0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mac_Parameters_V17c0 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. mac-Parameters-v17c0: ref
	{
		if ie.Mac_Parameters_V17c0 != nil {
			if err := ie.Mac_Parameters_V17c0.Encode(e); err != nil {
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

func (ie *UE_NR_Capability_V17c0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV17c0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mac-Parameters-v17c0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mac_Parameters_V17c0 = new(MAC_Parameters_V17c0)
			if err := ie.Mac_Parameters_V17c0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UE_NR_Capability_V17d0)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
