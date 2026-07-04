// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IntraFreqNeighCellInfo-v1610 (line 3909).

var intraFreqNeighCellInfoV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-PositionQCL-r16", Optional: true},
	},
}

type IntraFreqNeighCellInfo_v1610 struct {
	Ssb_PositionQCL_r16 *SSB_PositionQCL_Relation_r16
}

func (ie *IntraFreqNeighCellInfo_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(intraFreqNeighCellInfoV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_PositionQCL_r16 != nil}); err != nil {
		return err
	}

	// 2. ssb-PositionQCL-r16: ref
	{
		if ie.Ssb_PositionQCL_r16 != nil {
			if err := ie.Ssb_PositionQCL_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IntraFreqNeighCellInfo_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(intraFreqNeighCellInfoV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ssb-PositionQCL-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ssb_PositionQCL_r16 = new(SSB_PositionQCL_Relation_r16)
			if err := ie.Ssb_PositionQCL_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
