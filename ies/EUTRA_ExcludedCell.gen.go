// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-ExcludedCell (line 9418).

var eUTRAExcludedCellConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellIndexEUTRA"},
		{Name: "physCellIdRange"},
	},
}

type EUTRA_ExcludedCell struct {
	CellIndexEUTRA  EUTRA_CellIndex
	PhysCellIdRange EUTRA_PhysCellIdRange
}

func (ie *EUTRA_ExcludedCell) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRAExcludedCellConstraints)
	_ = seq

	// 1. cellIndexEUTRA: ref
	{
		if err := ie.CellIndexEUTRA.Encode(e); err != nil {
			return err
		}
	}

	// 2. physCellIdRange: ref
	{
		if err := ie.PhysCellIdRange.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *EUTRA_ExcludedCell) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRAExcludedCellConstraints)
	_ = seq

	// 1. cellIndexEUTRA: ref
	{
		if err := ie.CellIndexEUTRA.Decode(d); err != nil {
			return err
		}
	}

	// 2. physCellIdRange: ref
	{
		if err := ie.PhysCellIdRange.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
