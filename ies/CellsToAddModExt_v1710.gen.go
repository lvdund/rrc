// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CellsToAddModExt-v1710 (line 9552).

var cellsToAddModExtV1710Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ntn-PolarizationDL-r17", Optional: true},
		{Name: "ntn-PolarizationUL-r17", Optional: true},
	},
}

const (
	CellsToAddModExt_v1710_Ntn_PolarizationDL_r17_Rhcp   = 0
	CellsToAddModExt_v1710_Ntn_PolarizationDL_r17_Lhcp   = 1
	CellsToAddModExt_v1710_Ntn_PolarizationDL_r17_Linear = 2
)

var cellsToAddModExtV1710NtnPolarizationDLR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellsToAddModExt_v1710_Ntn_PolarizationDL_r17_Rhcp, CellsToAddModExt_v1710_Ntn_PolarizationDL_r17_Lhcp, CellsToAddModExt_v1710_Ntn_PolarizationDL_r17_Linear},
}

const (
	CellsToAddModExt_v1710_Ntn_PolarizationUL_r17_Rhcp   = 0
	CellsToAddModExt_v1710_Ntn_PolarizationUL_r17_Lhcp   = 1
	CellsToAddModExt_v1710_Ntn_PolarizationUL_r17_Linear = 2
)

var cellsToAddModExtV1710NtnPolarizationULR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellsToAddModExt_v1710_Ntn_PolarizationUL_r17_Rhcp, CellsToAddModExt_v1710_Ntn_PolarizationUL_r17_Lhcp, CellsToAddModExt_v1710_Ntn_PolarizationUL_r17_Linear},
}

type CellsToAddModExt_v1710 struct {
	Ntn_PolarizationDL_r17 *int64
	Ntn_PolarizationUL_r17 *int64
}

func (ie *CellsToAddModExt_v1710) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellsToAddModExtV1710Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ntn_PolarizationDL_r17 != nil, ie.Ntn_PolarizationUL_r17 != nil}); err != nil {
		return err
	}

	// 2. ntn-PolarizationDL-r17: enumerated
	{
		if ie.Ntn_PolarizationDL_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ntn_PolarizationDL_r17, cellsToAddModExtV1710NtnPolarizationDLR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. ntn-PolarizationUL-r17: enumerated
	{
		if ie.Ntn_PolarizationUL_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ntn_PolarizationUL_r17, cellsToAddModExtV1710NtnPolarizationULR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CellsToAddModExt_v1710) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellsToAddModExtV1710Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ntn-PolarizationDL-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cellsToAddModExtV1710NtnPolarizationDLR17Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_PolarizationDL_r17 = &idx
		}
	}

	// 3. ntn-PolarizationUL-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cellsToAddModExtV1710NtnPolarizationULR17Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_PolarizationUL_r17 = &idx
		}
	}

	return nil
}
