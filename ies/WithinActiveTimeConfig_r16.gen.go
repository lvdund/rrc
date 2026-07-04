// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: WithinActiveTimeConfig-r16 (line 14790).

var withinActiveTimeConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "firstWithinActiveTimeBWP-Id-r16", Optional: true},
		{Name: "dormancyGroupWithinActiveTime-r16", Optional: true},
	},
}

type WithinActiveTimeConfig_r16 struct {
	FirstWithinActiveTimeBWP_Id_r16   *BWP_Id
	DormancyGroupWithinActiveTime_r16 *DormancyGroupID_r16
}

func (ie *WithinActiveTimeConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(withinActiveTimeConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FirstWithinActiveTimeBWP_Id_r16 != nil, ie.DormancyGroupWithinActiveTime_r16 != nil}); err != nil {
		return err
	}

	// 2. firstWithinActiveTimeBWP-Id-r16: ref
	{
		if ie.FirstWithinActiveTimeBWP_Id_r16 != nil {
			if err := ie.FirstWithinActiveTimeBWP_Id_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. dormancyGroupWithinActiveTime-r16: ref
	{
		if ie.DormancyGroupWithinActiveTime_r16 != nil {
			if err := ie.DormancyGroupWithinActiveTime_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *WithinActiveTimeConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(withinActiveTimeConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. firstWithinActiveTimeBWP-Id-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FirstWithinActiveTimeBWP_Id_r16 = new(BWP_Id)
			if err := ie.FirstWithinActiveTimeBWP_Id_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. dormancyGroupWithinActiveTime-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.DormancyGroupWithinActiveTime_r16 = new(DormancyGroupID_r16)
			if err := ie.DormancyGroupWithinActiveTime_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
