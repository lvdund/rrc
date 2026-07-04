// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SchedulingRequestResourceConfigExt-v1610 (line 14304).

var schedulingRequestResourceConfigExtV1610Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "phy-PriorityIndex-r16", Optional: true},
	},
}

const (
	SchedulingRequestResourceConfigExt_v1610_Phy_PriorityIndex_r16_P0 = 0
	SchedulingRequestResourceConfigExt_v1610_Phy_PriorityIndex_r16_P1 = 1
)

var schedulingRequestResourceConfigExtV1610PhyPriorityIndexR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SchedulingRequestResourceConfigExt_v1610_Phy_PriorityIndex_r16_P0, SchedulingRequestResourceConfigExt_v1610_Phy_PriorityIndex_r16_P1},
}

type SchedulingRequestResourceConfigExt_v1610 struct {
	Phy_PriorityIndex_r16 *int64
}

func (ie *SchedulingRequestResourceConfigExt_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(schedulingRequestResourceConfigExtV1610Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Phy_PriorityIndex_r16 != nil}); err != nil {
		return err
	}

	// 3. phy-PriorityIndex-r16: enumerated
	{
		if ie.Phy_PriorityIndex_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Phy_PriorityIndex_r16, schedulingRequestResourceConfigExtV1610PhyPriorityIndexR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SchedulingRequestResourceConfigExt_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(schedulingRequestResourceConfigExtV1610Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. phy-PriorityIndex-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(schedulingRequestResourceConfigExtV1610PhyPriorityIndexR16Constraints)
			if err != nil {
				return err
			}
			ie.Phy_PriorityIndex_r16 = &idx
		}
	}

	return nil
}
