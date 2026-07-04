// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v15g0 (line 25703).

var uENRCapabilityV15g0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rf-Parameters-v15g0", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_NR_Capability_V15g0 struct {
	Rf_Parameters_V15g0  *RF_Parameters_V15g0
	NonCriticalExtension *UE_NR_Capability_V15j0
}

func (ie *UE_NR_Capability_V15g0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV15g0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rf_Parameters_V15g0 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. rf-Parameters-v15g0: ref
	{
		if ie.Rf_Parameters_V15g0 != nil {
			if err := ie.Rf_Parameters_V15g0.Encode(e); err != nil {
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

func (ie *UE_NR_Capability_V15g0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV15g0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. rf-Parameters-v15g0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Rf_Parameters_V15g0 = new(RF_Parameters_V15g0)
			if err := ie.Rf_Parameters_V15g0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UE_NR_Capability_V15j0)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
