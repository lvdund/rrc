// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultServMO (line 9753).

var measResultServMOConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "servCellId"},
		{Name: "measResultServingCell"},
		{Name: "measResultBestNeighCell", Optional: true},
	},
}

type MeasResultServMO struct {
	ServCellId              ServCellIndex
	MeasResultServingCell   MeasResultNR
	MeasResultBestNeighCell *MeasResultNR
}

func (ie *MeasResultServMO) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultServMOConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasResultBestNeighCell != nil}); err != nil {
		return err
	}

	// 3. servCellId: ref
	{
		if err := ie.ServCellId.Encode(e); err != nil {
			return err
		}
	}

	// 4. measResultServingCell: ref
	{
		if err := ie.MeasResultServingCell.Encode(e); err != nil {
			return err
		}
	}

	// 5. measResultBestNeighCell: ref
	{
		if ie.MeasResultBestNeighCell != nil {
			if err := ie.MeasResultBestNeighCell.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResultServMO) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultServMOConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. servCellId: ref
	{
		if err := ie.ServCellId.Decode(d); err != nil {
			return err
		}
	}

	// 4. measResultServingCell: ref
	{
		if err := ie.MeasResultServingCell.Decode(d); err != nil {
			return err
		}
	}

	// 5. measResultBestNeighCell: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MeasResultBestNeighCell = new(MeasResultNR)
			if err := ie.MeasResultBestNeighCell.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
