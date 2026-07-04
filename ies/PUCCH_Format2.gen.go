// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-format2 (line 12137).

var pUCCHFormat2Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofPRBs"},
		{Name: "nrofSymbols"},
		{Name: "startingSymbolIndex"},
	},
}

var pUCCHFormat2NrofPRBsConstraints = per.Constrained(1, 16)

var pUCCHFormat2NrofSymbolsConstraints = per.Constrained(1, 2)

var pUCCHFormat2StartingSymbolIndexConstraints = per.Constrained(0, 13)

type PUCCH_Format2 struct {
	NrofPRBs            int64
	NrofSymbols         int64
	StartingSymbolIndex int64
}

func (ie *PUCCH_Format2) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHFormat2Constraints)
	_ = seq

	// 1. nrofPRBs: integer
	{
		if err := e.EncodeInteger(ie.NrofPRBs, pUCCHFormat2NrofPRBsConstraints); err != nil {
			return err
		}
	}

	// 2. nrofSymbols: integer
	{
		if err := e.EncodeInteger(ie.NrofSymbols, pUCCHFormat2NrofSymbolsConstraints); err != nil {
			return err
		}
	}

	// 3. startingSymbolIndex: integer
	{
		if err := e.EncodeInteger(ie.StartingSymbolIndex, pUCCHFormat2StartingSymbolIndexConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PUCCH_Format2) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHFormat2Constraints)
	_ = seq

	// 1. nrofPRBs: integer
	{
		v0, err := d.DecodeInteger(pUCCHFormat2NrofPRBsConstraints)
		if err != nil {
			return err
		}
		ie.NrofPRBs = v0
	}

	// 2. nrofSymbols: integer
	{
		v1, err := d.DecodeInteger(pUCCHFormat2NrofSymbolsConstraints)
		if err != nil {
			return err
		}
		ie.NrofSymbols = v1
	}

	// 3. startingSymbolIndex: integer
	{
		v2, err := d.DecodeInteger(pUCCHFormat2StartingSymbolIndexConstraints)
		if err != nil {
			return err
		}
		ie.StartingSymbolIndex = v2
	}

	return nil
}
