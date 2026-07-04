// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-FrequencyOccupation (line 6931).

var cSIFrequencyOccupationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "startingRB"},
		{Name: "nrofRBs"},
	},
}

var cSIFrequencyOccupationStartingRBConstraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

var cSIFrequencyOccupationNrofRBsConstraints = per.Constrained(24, common.MaxNrofPhysicalResourceBlocksPlus1)

type CSI_FrequencyOccupation struct {
	StartingRB int64
	NrofRBs    int64
}

func (ie *CSI_FrequencyOccupation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIFrequencyOccupationConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. startingRB: integer
	{
		if err := e.EncodeInteger(ie.StartingRB, cSIFrequencyOccupationStartingRBConstraints); err != nil {
			return err
		}
	}

	// 3. nrofRBs: integer
	{
		if err := e.EncodeInteger(ie.NrofRBs, cSIFrequencyOccupationNrofRBsConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CSI_FrequencyOccupation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIFrequencyOccupationConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. startingRB: integer
	{
		v0, err := d.DecodeInteger(cSIFrequencyOccupationStartingRBConstraints)
		if err != nil {
			return err
		}
		ie.StartingRB = v0
	}

	// 3. nrofRBs: integer
	{
		v1, err := d.DecodeInteger(cSIFrequencyOccupationNrofRBsConstraints)
		if err != nil {
			return err
		}
		ie.NrofRBs = v1
	}

	return nil
}
