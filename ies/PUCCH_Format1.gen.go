// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-format1 (line 12130).

var pUCCHFormat1Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "initialCyclicShift"},
		{Name: "nrofSymbols"},
		{Name: "startingSymbolIndex"},
		{Name: "timeDomainOCC"},
	},
}

var pUCCHFormat1InitialCyclicShiftConstraints = per.Constrained(0, 11)

var pUCCHFormat1NrofSymbolsConstraints = per.Constrained(4, 14)

var pUCCHFormat1StartingSymbolIndexConstraints = per.Constrained(0, 10)

var pUCCHFormat1TimeDomainOCCConstraints = per.Constrained(0, 6)

type PUCCH_Format1 struct {
	InitialCyclicShift  int64
	NrofSymbols         int64
	StartingSymbolIndex int64
	TimeDomainOCC       int64
}

func (ie *PUCCH_Format1) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHFormat1Constraints)
	_ = seq

	// 1. initialCyclicShift: integer
	{
		if err := e.EncodeInteger(ie.InitialCyclicShift, pUCCHFormat1InitialCyclicShiftConstraints); err != nil {
			return err
		}
	}

	// 2. nrofSymbols: integer
	{
		if err := e.EncodeInteger(ie.NrofSymbols, pUCCHFormat1NrofSymbolsConstraints); err != nil {
			return err
		}
	}

	// 3. startingSymbolIndex: integer
	{
		if err := e.EncodeInteger(ie.StartingSymbolIndex, pUCCHFormat1StartingSymbolIndexConstraints); err != nil {
			return err
		}
	}

	// 4. timeDomainOCC: integer
	{
		if err := e.EncodeInteger(ie.TimeDomainOCC, pUCCHFormat1TimeDomainOCCConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PUCCH_Format1) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHFormat1Constraints)
	_ = seq

	// 1. initialCyclicShift: integer
	{
		v0, err := d.DecodeInteger(pUCCHFormat1InitialCyclicShiftConstraints)
		if err != nil {
			return err
		}
		ie.InitialCyclicShift = v0
	}

	// 2. nrofSymbols: integer
	{
		v1, err := d.DecodeInteger(pUCCHFormat1NrofSymbolsConstraints)
		if err != nil {
			return err
		}
		ie.NrofSymbols = v1
	}

	// 3. startingSymbolIndex: integer
	{
		v2, err := d.DecodeInteger(pUCCHFormat1StartingSymbolIndexConstraints)
		if err != nil {
			return err
		}
		ie.StartingSymbolIndex = v2
	}

	// 4. timeDomainOCC: integer
	{
		v3, err := d.DecodeInteger(pUCCHFormat1TimeDomainOCCConstraints)
		if err != nil {
			return err
		}
		ie.TimeDomainOCC = v3
	}

	return nil
}
