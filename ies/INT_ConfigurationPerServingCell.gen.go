// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: INT-ConfigurationPerServingCell (line 8097).

var iNTConfigurationPerServingCellConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "servingCellId"},
		{Name: "positionInDCI"},
	},
}

var iNTConfigurationPerServingCellPositionInDCIConstraints = per.Constrained(0, common.MaxINT_DCI_PayloadSize_1)

type INT_ConfigurationPerServingCell struct {
	ServingCellId ServCellIndex
	PositionInDCI int64
}

func (ie *INT_ConfigurationPerServingCell) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iNTConfigurationPerServingCellConstraints)
	_ = seq

	// 1. servingCellId: ref
	{
		if err := ie.ServingCellId.Encode(e); err != nil {
			return err
		}
	}

	// 2. positionInDCI: integer
	{
		if err := e.EncodeInteger(ie.PositionInDCI, iNTConfigurationPerServingCellPositionInDCIConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *INT_ConfigurationPerServingCell) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iNTConfigurationPerServingCellConstraints)
	_ = seq

	// 1. servingCellId: ref
	{
		if err := ie.ServingCellId.Decode(d); err != nil {
			return err
		}
	}

	// 2. positionInDCI: integer
	{
		v1, err := d.DecodeInteger(iNTConfigurationPerServingCellPositionInDCIConstraints)
		if err != nil {
			return err
		}
		ie.PositionInDCI = v1
	}

	return nil
}
