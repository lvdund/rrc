// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CellsToAddModUTRA-FDD-r16 (line 9675).

var cellsToAddModUTRAFDDR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellIndexUTRA-FDD-r16"},
		{Name: "physCellId-r16"},
	},
}

type CellsToAddModUTRA_FDD_r16 struct {
	CellIndexUTRA_FDD_r16 UTRA_FDD_CellIndex_r16
	PhysCellId_r16        PhysCellIdUTRA_FDD_r16
}

func (ie *CellsToAddModUTRA_FDD_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellsToAddModUTRAFDDR16Constraints)
	_ = seq

	// 1. cellIndexUTRA-FDD-r16: ref
	{
		if err := ie.CellIndexUTRA_FDD_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. physCellId-r16: ref
	{
		if err := ie.PhysCellId_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CellsToAddModUTRA_FDD_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellsToAddModUTRAFDDR16Constraints)
	_ = seq

	// 1. cellIndexUTRA-FDD-r16: ref
	{
		if err := ie.CellIndexUTRA_FDD_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. physCellId-r16: ref
	{
		if err := ie.PhysCellId_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
