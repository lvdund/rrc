// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CellsToAddModExt-v1800 (line 9557).

var cellsToAddModExtV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ntn-NeighbourCellInfo-r18", Optional: true},
	},
}

type CellsToAddModExt_v1800 struct {
	Ntn_NeighbourCellInfo_r18 *NTN_NeighbourCellInfo_r18
}

func (ie *CellsToAddModExt_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellsToAddModExtV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ntn_NeighbourCellInfo_r18 != nil}); err != nil {
		return err
	}

	// 2. ntn-NeighbourCellInfo-r18: ref
	{
		if ie.Ntn_NeighbourCellInfo_r18 != nil {
			if err := ie.Ntn_NeighbourCellInfo_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CellsToAddModExt_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellsToAddModExtV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ntn-NeighbourCellInfo-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ntn_NeighbourCellInfo_r18 = new(NTN_NeighbourCellInfo_r18)
			if err := ie.Ntn_NeighbourCellInfo_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
