// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-PositionQCL-CellsToAddMod-r16 (line 9584).

var sSBPositionQCLCellsToAddModR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId-r16"},
		{Name: "ssb-PositionQCL-r16"},
	},
}

type SSB_PositionQCL_CellsToAddMod_r16 struct {
	PhysCellId_r16      PhysCellId
	Ssb_PositionQCL_r16 SSB_PositionQCL_Relation_r16
}

func (ie *SSB_PositionQCL_CellsToAddMod_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBPositionQCLCellsToAddModR16Constraints)
	_ = seq

	// 1. physCellId-r16: ref
	{
		if err := ie.PhysCellId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. ssb-PositionQCL-r16: ref
	{
		if err := ie.Ssb_PositionQCL_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SSB_PositionQCL_CellsToAddMod_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBPositionQCLCellsToAddModR16Constraints)
	_ = seq

	// 1. physCellId-r16: ref
	{
		if err := ie.PhysCellId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. ssb-PositionQCL-r16: ref
	{
		if err := ie.Ssb_PositionQCL_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
