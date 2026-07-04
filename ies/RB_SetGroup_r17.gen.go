// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RB-SetGroup-r17 (line 5079).

var rBSetGroupR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resourceAvailability-r17", Optional: true},
		{Name: "rb-Sets-r17", Optional: true},
	},
}

var rBSetGroupR17ResourceAvailabilityR17Constraints = per.SizeRange(1, common.MaxNrofResourceAvailabilityPerCombination_r16)

var rBSetGroupR17RbSetsR17Constraints = per.SizeRange(1, common.MaxNrofRB_Sets_r17)

type RB_SetGroup_r17 struct {
	ResourceAvailability_r17 []int64
	Rb_Sets_r17              []int64
}

func (ie *RB_SetGroup_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rBSetGroupR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ResourceAvailability_r17 != nil, ie.Rb_Sets_r17 != nil}); err != nil {
		return err
	}

	// 2. resourceAvailability-r17: sequence-of
	{
		if ie.ResourceAvailability_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(rBSetGroupR17ResourceAvailabilityR17Constraints)
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

	// 3. rb-Sets-r17: sequence-of
	{
		if ie.Rb_Sets_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(rBSetGroupR17RbSetsR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Rb_Sets_r17))); err != nil {
				return err
			}
			for i := range ie.Rb_Sets_r17 {
				if err := e.EncodeInteger(ie.Rb_Sets_r17[i], per.Constrained(0, 7)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *RB_SetGroup_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rBSetGroupR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. resourceAvailability-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(rBSetGroupR17ResourceAvailabilityR17Constraints)
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

	// 3. rb-Sets-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(rBSetGroupR17RbSetsR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Rb_Sets_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, 7))
				if err != nil {
					return err
				}
				ie.Rb_Sets_r17[i] = v
			}
		}
	}

	return nil
}
