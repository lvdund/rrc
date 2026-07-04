// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: L1-MeasResultPerCell-r19 (line 3540).

var l1MeasResultPerCellR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId-r19"},
		{Name: "resultsSSB-Indexes-r19", Optional: true},
	},
}

type L1_MeasResultPerCell_r19 struct {
	PhysCellId_r19         PhysCellId
	ResultsSSB_Indexes_r19 *ResultsPerSSB_IndexList
}

func (ie *L1_MeasResultPerCell_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(l1MeasResultPerCellR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ResultsSSB_Indexes_r19 != nil}); err != nil {
		return err
	}

	// 3. physCellId-r19: ref
	{
		if err := ie.PhysCellId_r19.Encode(e); err != nil {
			return err
		}
	}

	// 4. resultsSSB-Indexes-r19: ref
	{
		if ie.ResultsSSB_Indexes_r19 != nil {
			if err := ie.ResultsSSB_Indexes_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *L1_MeasResultPerCell_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(l1MeasResultPerCellR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. physCellId-r19: ref
	{
		if err := ie.PhysCellId_r19.Decode(d); err != nil {
			return err
		}
	}

	// 4. resultsSSB-Indexes-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ResultsSSB_Indexes_r19 = new(ResultsPerSSB_IndexList)
			if err := ie.ResultsSSB_Indexes_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
