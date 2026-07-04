// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RA-Prioritization (line 13083).

var rAPrioritizationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "powerRampingStepHighPriority"},
		{Name: "scalingFactorBI", Optional: true},
	},
}

const (
	RA_Prioritization_PowerRampingStepHighPriority_DB0 = 0
	RA_Prioritization_PowerRampingStepHighPriority_DB2 = 1
	RA_Prioritization_PowerRampingStepHighPriority_DB4 = 2
	RA_Prioritization_PowerRampingStepHighPriority_DB6 = 3
)

var rAPrioritizationPowerRampingStepHighPriorityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_Prioritization_PowerRampingStepHighPriority_DB0, RA_Prioritization_PowerRampingStepHighPriority_DB2, RA_Prioritization_PowerRampingStepHighPriority_DB4, RA_Prioritization_PowerRampingStepHighPriority_DB6},
}

const (
	RA_Prioritization_ScalingFactorBI_Zero  = 0
	RA_Prioritization_ScalingFactorBI_Dot25 = 1
	RA_Prioritization_ScalingFactorBI_Dot5  = 2
	RA_Prioritization_ScalingFactorBI_Dot75 = 3
)

var rAPrioritizationScalingFactorBIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_Prioritization_ScalingFactorBI_Zero, RA_Prioritization_ScalingFactorBI_Dot25, RA_Prioritization_ScalingFactorBI_Dot5, RA_Prioritization_ScalingFactorBI_Dot75},
}

type RA_Prioritization struct {
	PowerRampingStepHighPriority int64
	ScalingFactorBI              *int64
}

func (ie *RA_Prioritization) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rAPrioritizationConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ScalingFactorBI != nil}); err != nil {
		return err
	}

	// 3. powerRampingStepHighPriority: enumerated
	{
		if err := e.EncodeEnumerated(ie.PowerRampingStepHighPriority, rAPrioritizationPowerRampingStepHighPriorityConstraints); err != nil {
			return err
		}
	}

	// 4. scalingFactorBI: enumerated
	{
		if ie.ScalingFactorBI != nil {
			if err := e.EncodeEnumerated(*ie.ScalingFactorBI, rAPrioritizationScalingFactorBIConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RA_Prioritization) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rAPrioritizationConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. powerRampingStepHighPriority: enumerated
	{
		v0, err := d.DecodeEnumerated(rAPrioritizationPowerRampingStepHighPriorityConstraints)
		if err != nil {
			return err
		}
		ie.PowerRampingStepHighPriority = v0
	}

	// 4. scalingFactorBI: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(rAPrioritizationScalingFactorBIConstraints)
			if err != nil {
				return err
			}
			ie.ScalingFactorBI = &idx
		}
	}

	return nil
}
