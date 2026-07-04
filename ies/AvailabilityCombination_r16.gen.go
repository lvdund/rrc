// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AvailabilityCombination-r16 (line 5066).

var availabilityCombinationR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "availabilityCombinationId-r16"},
		{Name: "resourceAvailability-r16"},
	},
}

var availabilityCombinationR16ResourceAvailabilityR16Constraints = per.SizeRange(1, common.MaxNrofResourceAvailabilityPerCombination_r16)

type AvailabilityCombination_r16 struct {
	AvailabilityCombinationId_r16 AvailabilityCombinationId_r16
	ResourceAvailability_r16      []int64
}

func (ie *AvailabilityCombination_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(availabilityCombinationR16Constraints)
	_ = seq

	// 1. availabilityCombinationId-r16: ref
	{
		if err := ie.AvailabilityCombinationId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. resourceAvailability-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(availabilityCombinationR16ResourceAvailabilityR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.ResourceAvailability_r16))); err != nil {
			return err
		}
		for i := range ie.ResourceAvailability_r16 {
			if err := e.EncodeInteger(ie.ResourceAvailability_r16[i], per.Constrained(0, 7)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *AvailabilityCombination_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(availabilityCombinationR16Constraints)
	_ = seq

	// 1. availabilityCombinationId-r16: ref
	{
		if err := ie.AvailabilityCombinationId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. resourceAvailability-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(availabilityCombinationR16ResourceAvailabilityR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.ResourceAvailability_r16 = make([]int64, n)
		for i := int64(0); i < n; i++ {
			v, err := d.DecodeInteger(per.Constrained(0, 7))
			if err != nil {
				return err
			}
			ie.ResourceAvailability_r16[i] = v
		}
	}

	return nil
}
