// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-Cell (line 9411).

var eUTRACellConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellIndexEUTRA"},
		{Name: "physCellId"},
		{Name: "cellIndividualOffset"},
	},
}

type EUTRA_Cell struct {
	CellIndexEUTRA       EUTRA_CellIndex
	PhysCellId           EUTRA_PhysCellId
	CellIndividualOffset EUTRA_Q_OffsetRange
}

func (ie *EUTRA_Cell) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRACellConstraints)
	_ = seq

	// 1. cellIndexEUTRA: ref
	{
		if err := ie.CellIndexEUTRA.Encode(e); err != nil {
			return err
		}
	}

	// 2. physCellId: ref
	{
		if err := ie.PhysCellId.Encode(e); err != nil {
			return err
		}
	}

	// 3. cellIndividualOffset: ref
	{
		if err := ie.CellIndividualOffset.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *EUTRA_Cell) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRACellConstraints)
	_ = seq

	// 1. cellIndexEUTRA: ref
	{
		if err := ie.CellIndexEUTRA.Decode(d); err != nil {
			return err
		}
	}

	// 2. physCellId: ref
	{
		if err := ie.PhysCellId.Decode(d); err != nil {
			return err
		}
	}

	// 3. cellIndividualOffset: ref
	{
		if err := ie.CellIndividualOffset.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
