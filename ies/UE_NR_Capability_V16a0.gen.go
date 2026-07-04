// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v16a0 (line 25763).

var uENRCapabilityV16a0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "phy-Parameters-v16a0", Optional: true},
		{Name: "rf-Parameters-v16a0", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UE_NR_Capability_V16a0 struct {
	Phy_Parameters_V16a0 *Phy_Parameters_V16a0
	Rf_Parameters_V16a0  *RF_Parameters_V16a0
	NonCriticalExtension *UE_NR_Capability_V16c0
}

func (ie *UE_NR_Capability_V16a0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV16a0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Phy_Parameters_V16a0 != nil, ie.Rf_Parameters_V16a0 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. phy-Parameters-v16a0: ref
	{
		if ie.Phy_Parameters_V16a0 != nil {
			if err := ie.Phy_Parameters_V16a0.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. rf-Parameters-v16a0: ref
	{
		if ie.Rf_Parameters_V16a0 != nil {
			if err := ie.Rf_Parameters_V16a0.Encode(e); err != nil {
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

func (ie *UE_NR_Capability_V16a0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV16a0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. phy-Parameters-v16a0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Phy_Parameters_V16a0 = new(Phy_Parameters_V16a0)
			if err := ie.Phy_Parameters_V16a0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. rf-Parameters-v16a0: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Rf_Parameters_V16a0 = new(RF_Parameters_V16a0)
			if err := ie.Rf_Parameters_V16a0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(UE_NR_Capability_V16c0)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
