// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-format3 (line 12143).

var pUCCHFormat3Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofPRBs"},
		{Name: "nrofSymbols"},
		{Name: "startingSymbolIndex"},
	},
}

var pUCCHFormat3NrofPRBsConstraints = per.Constrained(1, 16)

var pUCCHFormat3NrofSymbolsConstraints = per.Constrained(4, 14)

var pUCCHFormat3StartingSymbolIndexConstraints = per.Constrained(0, 10)

type PUCCH_Format3 struct {
	NrofPRBs            int64
	NrofSymbols         int64
	StartingSymbolIndex int64
}

func (ie *PUCCH_Format3) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHFormat3Constraints)
	_ = seq

	// 1. nrofPRBs: integer
	{
		if err := e.EncodeInteger(ie.NrofPRBs, pUCCHFormat3NrofPRBsConstraints); err != nil {
			return err
		}
	}

	// 2. nrofSymbols: integer
	{
		if err := e.EncodeInteger(ie.NrofSymbols, pUCCHFormat3NrofSymbolsConstraints); err != nil {
			return err
		}
	}

	// 3. startingSymbolIndex: integer
	{
		if err := e.EncodeInteger(ie.StartingSymbolIndex, pUCCHFormat3StartingSymbolIndexConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PUCCH_Format3) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHFormat3Constraints)
	_ = seq

	// 1. nrofPRBs: integer
	{
		v0, err := d.DecodeInteger(pUCCHFormat3NrofPRBsConstraints)
		if err != nil {
			return err
		}
		ie.NrofPRBs = v0
	}

	// 2. nrofSymbols: integer
	{
		v1, err := d.DecodeInteger(pUCCHFormat3NrofSymbolsConstraints)
		if err != nil {
			return err
		}
		ie.NrofSymbols = v1
	}

	// 3. startingSymbolIndex: integer
	{
		v2, err := d.DecodeInteger(pUCCHFormat3StartingSymbolIndexConstraints)
		if err != nil {
			return err
		}
		ie.StartingSymbolIndex = v2
	}

	return nil
}
