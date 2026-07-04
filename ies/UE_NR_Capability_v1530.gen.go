// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1530 (line 25660).

var uENRCapabilityV1530Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fdd-Add-UE-NR-Capabilities-v1530", Optional: true},
		{Name: "tdd-Add-UE-NR-Capabilities-v1530", Optional: true},
		{Name: "dummy", Optional: true},
		{Name: "interRAT-Parameters", Optional: true},
		{Name: "inactiveState", Optional: true},
		{Name: "delayBudgetReporting", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_NR_Capability_v1530_Dummy_Supported = 0
)

var uENRCapabilityV1530DummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1530_Dummy_Supported},
}

const (
	UE_NR_Capability_v1530_InactiveState_Supported = 0
)

var uENRCapabilityV1530InactiveStateConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1530_InactiveState_Supported},
}

const (
	UE_NR_Capability_v1530_DelayBudgetReporting_Supported = 0
)

var uENRCapabilityV1530DelayBudgetReportingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1530_DelayBudgetReporting_Supported},
}

type UE_NR_Capability_v1530 struct {
	Fdd_Add_UE_NR_Capabilities_v1530 *UE_NR_CapabilityAddXDD_Mode_v1530
	Tdd_Add_UE_NR_Capabilities_v1530 *UE_NR_CapabilityAddXDD_Mode_v1530
	Dummy                            *int64
	InterRAT_Parameters              *InterRAT_Parameters
	InactiveState                    *int64
	DelayBudgetReporting             *int64
	NonCriticalExtension             *UE_NR_Capability_v1540
}

func (ie *UE_NR_Capability_v1530) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1530Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Fdd_Add_UE_NR_Capabilities_v1530 != nil, ie.Tdd_Add_UE_NR_Capabilities_v1530 != nil, ie.Dummy != nil, ie.InterRAT_Parameters != nil, ie.InactiveState != nil, ie.DelayBudgetReporting != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. fdd-Add-UE-NR-Capabilities-v1530: ref
	{
		if ie.Fdd_Add_UE_NR_Capabilities_v1530 != nil {
			if err := ie.Fdd_Add_UE_NR_Capabilities_v1530.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. tdd-Add-UE-NR-Capabilities-v1530: ref
	{
		if ie.Tdd_Add_UE_NR_Capabilities_v1530 != nil {
			if err := ie.Tdd_Add_UE_NR_Capabilities_v1530.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. dummy: enumerated
	{
		if ie.Dummy != nil {
			if err := e.EncodeEnumerated(*ie.Dummy, uENRCapabilityV1530DummyConstraints); err != nil {
				return err
			}
		}
	}

	// 5. interRAT-Parameters: ref
	{
		if ie.InterRAT_Parameters != nil {
			if err := ie.InterRAT_Parameters.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. inactiveState: enumerated
	{
		if ie.InactiveState != nil {
			if err := e.EncodeEnumerated(*ie.InactiveState, uENRCapabilityV1530InactiveStateConstraints); err != nil {
				return err
			}
		}
	}

	// 7. delayBudgetReporting: enumerated
	{
		if ie.DelayBudgetReporting != nil {
			if err := e.EncodeEnumerated(*ie.DelayBudgetReporting, uENRCapabilityV1530DelayBudgetReportingConstraints); err != nil {
				return err
			}
		}
	}

	// 8. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_Capability_v1530) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1530Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. fdd-Add-UE-NR-Capabilities-v1530: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Fdd_Add_UE_NR_Capabilities_v1530 = new(UE_NR_CapabilityAddXDD_Mode_v1530)
			if err := ie.Fdd_Add_UE_NR_Capabilities_v1530.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. tdd-Add-UE-NR-Capabilities-v1530: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Tdd_Add_UE_NR_Capabilities_v1530 = new(UE_NR_CapabilityAddXDD_Mode_v1530)
			if err := ie.Tdd_Add_UE_NR_Capabilities_v1530.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. dummy: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1530DummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &idx
		}
	}

	// 5. interRAT-Parameters: ref
	{
		if seq.IsComponentPresent(3) {
			ie.InterRAT_Parameters = new(InterRAT_Parameters)
			if err := ie.InterRAT_Parameters.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. inactiveState: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1530InactiveStateConstraints)
			if err != nil {
				return err
			}
			ie.InactiveState = &idx
		}
	}

	// 7. delayBudgetReporting: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1530DelayBudgetReportingConstraints)
			if err != nil {
				return err
			}
			ie.DelayBudgetReporting = &idx
		}
	}

	// 8. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(6) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1540)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
