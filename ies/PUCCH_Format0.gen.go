// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-format0 (line 12124).

var pUCCHFormat0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "initialCyclicShift"},
		{Name: "nrofSymbols"},
		{Name: "startingSymbolIndex"},
	},
}

var pUCCHFormat0InitialCyclicShiftConstraints = per.Constrained(0, 11)

var pUCCHFormat0NrofSymbolsConstraints = per.Constrained(1, 2)

var pUCCHFormat0StartingSymbolIndexConstraints = per.Constrained(0, 13)

type PUCCH_Format0 struct {
	InitialCyclicShift  int64
	NrofSymbols         int64
	StartingSymbolIndex int64
}

func (ie *PUCCH_Format0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHFormat0Constraints)
	_ = seq

	// 1. initialCyclicShift: integer
	{
		if err := e.EncodeInteger(ie.InitialCyclicShift, pUCCHFormat0InitialCyclicShiftConstraints); err != nil {
			return err
		}
	}

	// 2. nrofSymbols: integer
	{
		if err := e.EncodeInteger(ie.NrofSymbols, pUCCHFormat0NrofSymbolsConstraints); err != nil {
			return err
		}
	}

	// 3. startingSymbolIndex: integer
	{
		if err := e.EncodeInteger(ie.StartingSymbolIndex, pUCCHFormat0StartingSymbolIndexConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PUCCH_Format0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHFormat0Constraints)
	_ = seq

	// 1. initialCyclicShift: integer
	{
		v0, err := d.DecodeInteger(pUCCHFormat0InitialCyclicShiftConstraints)
		if err != nil {
			return err
		}
		ie.InitialCyclicShift = v0
	}

	// 2. nrofSymbols: integer
	{
		v1, err := d.DecodeInteger(pUCCHFormat0NrofSymbolsConstraints)
		if err != nil {
			return err
		}
		ie.NrofSymbols = v1
	}

	// 3. startingSymbolIndex: integer
	{
		v2, err := d.DecodeInteger(pUCCHFormat0StartingSymbolIndexConstraints)
		if err != nil {
			return err
		}
		ie.StartingSymbolIndex = v2
	}

	return nil
}
