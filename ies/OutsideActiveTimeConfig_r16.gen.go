// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OutsideActiveTimeConfig-r16 (line 14795).

var outsideActiveTimeConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "firstOutsideActiveTimeBWP-Id-r16", Optional: true},
		{Name: "dormancyGroupOutsideActiveTime-r16", Optional: true},
	},
}

type OutsideActiveTimeConfig_r16 struct {
	FirstOutsideActiveTimeBWP_Id_r16   *BWP_Id
	DormancyGroupOutsideActiveTime_r16 *DormancyGroupID_r16
}

func (ie *OutsideActiveTimeConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(outsideActiveTimeConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FirstOutsideActiveTimeBWP_Id_r16 != nil, ie.DormancyGroupOutsideActiveTime_r16 != nil}); err != nil {
		return err
	}

	// 2. firstOutsideActiveTimeBWP-Id-r16: ref
	{
		if ie.FirstOutsideActiveTimeBWP_Id_r16 != nil {
			if err := ie.FirstOutsideActiveTimeBWP_Id_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. dormancyGroupOutsideActiveTime-r16: ref
	{
		if ie.DormancyGroupOutsideActiveTime_r16 != nil {
			if err := ie.DormancyGroupOutsideActiveTime_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *OutsideActiveTimeConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(outsideActiveTimeConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. firstOutsideActiveTimeBWP-Id-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FirstOutsideActiveTimeBWP_Id_r16 = new(BWP_Id)
			if err := ie.FirstOutsideActiveTimeBWP_Id_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. dormancyGroupOutsideActiveTime-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.DormancyGroupOutsideActiveTime_r16 = new(DormancyGroupID_r16)
			if err := ie.DormancyGroupOutsideActiveTime_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
