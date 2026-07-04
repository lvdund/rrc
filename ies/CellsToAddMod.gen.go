// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CellsToAddMod (line 9547).

var cellsToAddModConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId"},
		{Name: "cellIndividualOffset"},
	},
}

type CellsToAddMod struct {
	PhysCellId           PhysCellId
	CellIndividualOffset Q_OffsetRangeList
}

func (ie *CellsToAddMod) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellsToAddModConstraints)
	_ = seq

	// 1. physCellId: ref
	{
		if err := ie.PhysCellId.Encode(e); err != nil {
			return err
		}
	}

	// 2. cellIndividualOffset: ref
	{
		if err := ie.CellIndividualOffset.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CellsToAddMod) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellsToAddModConstraints)
	_ = seq

	// 1. physCellId: ref
	{
		if err := ie.PhysCellId.Decode(d); err != nil {
			return err
		}
	}

	// 2. cellIndividualOffset: ref
	{
		if err := ie.CellIndividualOffset.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
