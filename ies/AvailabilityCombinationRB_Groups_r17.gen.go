// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AvailabilityCombinationRB-Groups-r17 (line 5073).

var availabilityCombinationRBGroupsR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "availabilityCombinationId-r17"},
		{Name: "rb-SetGroups-r17", Optional: true},
		{Name: "resourceAvailability-r17", Optional: true},
	},
}

var availabilityCombinationRBGroupsR17RbSetGroupsR17Constraints = per.SizeRange(1, common.MaxNrofRB_SetGroups_r17)

var availabilityCombinationRBGroupsR17ResourceAvailabilityR17Constraints = per.SizeRange(1, common.MaxNrofResourceAvailabilityPerCombination_r16)

type AvailabilityCombinationRB_Groups_r17 struct {
	AvailabilityCombinationId_r17 AvailabilityCombinationId_r16
	Rb_SetGroups_r17              []RB_SetGroup_r17
	ResourceAvailability_r17      []int64
}

func (ie *AvailabilityCombinationRB_Groups_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(availabilityCombinationRBGroupsR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rb_SetGroups_r17 != nil, ie.ResourceAvailability_r17 != nil}); err != nil {
		return err
	}

	// 2. availabilityCombinationId-r17: ref
	{
		if err := ie.AvailabilityCombinationId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. rb-SetGroups-r17: sequence-of
	{
		if ie.Rb_SetGroups_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(availabilityCombinationRBGroupsR17RbSetGroupsR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Rb_SetGroups_r17))); err != nil {
				return err
			}
			for i := range ie.Rb_SetGroups_r17 {
				if err := ie.Rb_SetGroups_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. resourceAvailability-r17: sequence-of
	{
		if ie.ResourceAvailability_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(availabilityCombinationRBGroupsR17ResourceAvailabilityR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.ResourceAvailability_r17))); err != nil {
				return err
			}
			for i := range ie.ResourceAvailability_r17 {
				if err := e.EncodeInteger(ie.ResourceAvailability_r17[i], per.Constrained(0, 7)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *AvailabilityCombinationRB_Groups_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(availabilityCombinationRBGroupsR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. availabilityCombinationId-r17: ref
	{
		if err := ie.AvailabilityCombinationId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. rb-SetGroups-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(availabilityCombinationRBGroupsR17RbSetGroupsR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Rb_SetGroups_r17 = make([]RB_SetGroup_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Rb_SetGroups_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. resourceAvailability-r17: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(availabilityCombinationRBGroupsR17ResourceAvailabilityR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ResourceAvailability_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, 7))
				if err != nil {
					return err
				}
				ie.ResourceAvailability_r17[i] = v
			}
		}
	}

	return nil
}
