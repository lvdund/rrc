// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AvailableRB-SetsPerCell-r16 (line 15204).

var availableRBSetsPerCellR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "servingCellId-r16"},
		{Name: "positionInDCI-r16"},
	},
}

var availableRBSetsPerCellR16PositionInDCIR16Constraints = per.Constrained(0, common.MaxSFI_DCI_PayloadSize_1)

type AvailableRB_SetsPerCell_r16 struct {
	ServingCellId_r16 ServCellIndex
	PositionInDCI_r16 int64
}

func (ie *AvailableRB_SetsPerCell_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(availableRBSetsPerCellR16Constraints)
	_ = seq

	// 1. servingCellId-r16: ref
	{
		if err := ie.ServingCellId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. positionInDCI-r16: integer
	{
		if err := e.EncodeInteger(ie.PositionInDCI_r16, availableRBSetsPerCellR16PositionInDCIR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *AvailableRB_SetsPerCell_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(availableRBSetsPerCellR16Constraints)
	_ = seq

	// 1. servingCellId-r16: ref
	{
		if err := ie.ServingCellId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. positionInDCI-r16: integer
	{
		v1, err := d.DecodeInteger(availableRBSetsPerCellR16PositionInDCIR16Constraints)
		if err != nil {
			return err
		}
		ie.PositionInDCI_r16 = v1
	}

	return nil
}
