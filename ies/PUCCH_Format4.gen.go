// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-format4 (line 12149).

var pUCCHFormat4Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofSymbols"},
		{Name: "occ-Length"},
		{Name: "occ-Index"},
		{Name: "startingSymbolIndex"},
	},
}

var pUCCHFormat4NrofSymbolsConstraints = per.Constrained(4, 14)

const (
	PUCCH_Format4_Occ_Length_N2 = 0
	PUCCH_Format4_Occ_Length_N4 = 1
)

var pUCCHFormat4OccLengthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Format4_Occ_Length_N2, PUCCH_Format4_Occ_Length_N4},
}

const (
	PUCCH_Format4_Occ_Index_N0 = 0
	PUCCH_Format4_Occ_Index_N1 = 1
	PUCCH_Format4_Occ_Index_N2 = 2
	PUCCH_Format4_Occ_Index_N3 = 3
)

var pUCCHFormat4OccIndexConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Format4_Occ_Index_N0, PUCCH_Format4_Occ_Index_N1, PUCCH_Format4_Occ_Index_N2, PUCCH_Format4_Occ_Index_N3},
}

var pUCCHFormat4StartingSymbolIndexConstraints = per.Constrained(0, 10)

type PUCCH_Format4 struct {
	NrofSymbols         int64
	Occ_Length          int64
	Occ_Index           int64
	StartingSymbolIndex int64
}

func (ie *PUCCH_Format4) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHFormat4Constraints)
	_ = seq

	// 1. nrofSymbols: integer
	{
		if err := e.EncodeInteger(ie.NrofSymbols, pUCCHFormat4NrofSymbolsConstraints); err != nil {
			return err
		}
	}

	// 2. occ-Length: enumerated
	{
		if err := e.EncodeEnumerated(ie.Occ_Length, pUCCHFormat4OccLengthConstraints); err != nil {
			return err
		}
	}

	// 3. occ-Index: enumerated
	{
		if err := e.EncodeEnumerated(ie.Occ_Index, pUCCHFormat4OccIndexConstraints); err != nil {
			return err
		}
	}

	// 4. startingSymbolIndex: integer
	{
		if err := e.EncodeInteger(ie.StartingSymbolIndex, pUCCHFormat4StartingSymbolIndexConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PUCCH_Format4) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHFormat4Constraints)
	_ = seq

	// 1. nrofSymbols: integer
	{
		v0, err := d.DecodeInteger(pUCCHFormat4NrofSymbolsConstraints)
		if err != nil {
			return err
		}
		ie.NrofSymbols = v0
	}

	// 2. occ-Length: enumerated
	{
		v1, err := d.DecodeEnumerated(pUCCHFormat4OccLengthConstraints)
		if err != nil {
			return err
		}
		ie.Occ_Length = v1
	}

	// 3. occ-Index: enumerated
	{
		v2, err := d.DecodeEnumerated(pUCCHFormat4OccIndexConstraints)
		if err != nil {
			return err
		}
		ie.Occ_Index = v2
	}

	// 4. startingSymbolIndex: integer
	{
		v3, err := d.DecodeInteger(pUCCHFormat4StartingSymbolIndexConstraints)
		if err != nil {
			return err
		}
		ie.StartingSymbolIndex = v3
	}

	return nil
}
