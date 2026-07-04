// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IntraFreqNeighCellInfo-v1710 (line 3913).

var intraFreqNeighCellInfoV1710Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-PositionQCL-r17", Optional: true},
	},
}

type IntraFreqNeighCellInfo_v1710 struct {
	Ssb_PositionQCL_r17 *SSB_PositionQCL_Relation_r17
}

func (ie *IntraFreqNeighCellInfo_v1710) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(intraFreqNeighCellInfoV1710Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_PositionQCL_r17 != nil}); err != nil {
		return err
	}

	// 2. ssb-PositionQCL-r17: ref
	{
		if ie.Ssb_PositionQCL_r17 != nil {
			if err := ie.Ssb_PositionQCL_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IntraFreqNeighCellInfo_v1710) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(intraFreqNeighCellInfoV1710Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ssb-PositionQCL-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ssb_PositionQCL_r17 = new(SSB_PositionQCL_Relation_r17)
			if err := ie.Ssb_PositionQCL_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
