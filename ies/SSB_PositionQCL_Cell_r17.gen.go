// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-PositionQCL-Cell-r17 (line 9591).

var sSBPositionQCLCellR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId-r17"},
		{Name: "ssb-PositionQCL-r17"},
	},
}

type SSB_PositionQCL_Cell_r17 struct {
	PhysCellId_r17      PhysCellId
	Ssb_PositionQCL_r17 SSB_PositionQCL_Relation_r17
}

func (ie *SSB_PositionQCL_Cell_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBPositionQCLCellR17Constraints)
	_ = seq

	// 1. physCellId-r17: ref
	{
		if err := ie.PhysCellId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. ssb-PositionQCL-r17: ref
	{
		if err := ie.Ssb_PositionQCL_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SSB_PositionQCL_Cell_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBPositionQCLCellR17Constraints)
	_ = seq

	// 1. physCellId-r17: ref
	{
		if err := ie.PhysCellId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. ssb-PositionQCL-r17: ref
	{
		if err := ie.Ssb_PositionQCL_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
