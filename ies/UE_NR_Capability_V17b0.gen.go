// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v17b0 (line 25831).

var uENRCapabilityV17b0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mac-Parameters-v17b0", Optional: true},
		{Name: "rf-Parameters-v17b0", Optional: true},
		{Name: "ul-RRC-MaxCapaSegments-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_NR_Capability_V17b0_Ul_RRC_MaxCapaSegments_r17_Supported = 0
)

var uENRCapabilityV17b0UlRRCMaxCapaSegmentsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_V17b0_Ul_RRC_MaxCapaSegments_r17_Supported},
}

type UE_NR_Capability_V17b0 struct {
	Mac_Parameters_V17b0       *MAC_Parameters_V17b0
	Rf_Parameters_V17b0        *RF_Parameters_V17b0
	Ul_RRC_MaxCapaSegments_r17 *int64
	NonCriticalExtension       *UE_NR_Capability_V17c0
}

func (ie *UE_NR_Capability_V17b0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV17b0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mac_Parameters_V17b0 != nil, ie.Rf_Parameters_V17b0 != nil, ie.Ul_RRC_MaxCapaSegments_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. mac-Parameters-v17b0: ref
	{
		if ie.Mac_Parameters_V17b0 != nil {
			if err := ie.Mac_Parameters_V17b0.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. rf-Parameters-v17b0: ref
	{
		if ie.Rf_Parameters_V17b0 != nil {
			if err := ie.Rf_Parameters_V17b0.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ul-RRC-MaxCapaSegments-r17: enumerated
	{
		if ie.Ul_RRC_MaxCapaSegments_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_RRC_MaxCapaSegments_r17, uENRCapabilityV17b0UlRRCMaxCapaSegmentsR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_Capability_V17b0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV17b0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mac-Parameters-v17b0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mac_Parameters_V17b0 = new(MAC_Parameters_V17b0)
			if err := ie.Mac_Parameters_V17b0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. rf-Parameters-v17b0: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Rf_Parameters_V17b0 = new(RF_Parameters_V17b0)
			if err := ie.Rf_Parameters_V17b0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ul-RRC-MaxCapaSegments-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV17b0UlRRCMaxCapaSegmentsR17Constraints)
			if err != nil {
				return err
			}
			ie.Ul_RRC_MaxCapaSegments_r17 = &idx
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(3) {
			ie.NonCriticalExtension = new(UE_NR_Capability_V17c0)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
