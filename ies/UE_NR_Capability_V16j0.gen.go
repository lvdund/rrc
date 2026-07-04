// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v16j0 (line 25779).

var uENRCapabilityV16j0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rf-Parameters-v16j0", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uENRCapabilityV16j0LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type UE_NR_Capability_V16j0 struct {
	Rf_Parameters_V16j0      *RF_Parameters_V16j0
	LateNonCriticalExtension []byte
	NonCriticalExtension     *UE_NR_Capability_V17b0
}

func (ie *UE_NR_Capability_V16j0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV16j0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rf_Parameters_V16j0 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. rf-Parameters-v16j0: ref
	{
		if ie.Rf_Parameters_V16j0 != nil {
			if err := ie.Rf_Parameters_V16j0.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uENRCapabilityV16j0LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_Capability_V16j0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV16j0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. rf-Parameters-v16j0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Rf_Parameters_V16j0 = new(RF_Parameters_V16j0)
			if err := ie.Rf_Parameters_V16j0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(uENRCapabilityV16j0LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(UE_NR_Capability_V17b0)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
