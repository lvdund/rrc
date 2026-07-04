// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResult2EUTRA (line 9893).

var measResult2EUTRAConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq"},
		{Name: "measResultServingCell", Optional: true},
		{Name: "measResultBestNeighCell", Optional: true},
	},
}

type MeasResult2EUTRA struct {
	CarrierFreq             ARFCN_ValueEUTRA
	MeasResultServingCell   *MeasResultEUTRA
	MeasResultBestNeighCell *MeasResultEUTRA
}

func (ie *MeasResult2EUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResult2EUTRAConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasResultServingCell != nil, ie.MeasResultBestNeighCell != nil}); err != nil {
		return err
	}

	// 3. carrierFreq: ref
	{
		if err := ie.CarrierFreq.Encode(e); err != nil {
			return err
		}
	}

	// 4. measResultServingCell: ref
	{
		if ie.MeasResultServingCell != nil {
			if err := ie.MeasResultServingCell.Encode(e); err != nil {
				return err
			}
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

func (ie *MeasResult2EUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResult2EUTRAConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. carrierFreq: ref
	{
		if err := ie.CarrierFreq.Decode(d); err != nil {
			return err
		}
	}

	// 4. measResultServingCell: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MeasResultServingCell = new(MeasResultEUTRA)
			if err := ie.MeasResultServingCell.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. measResultBestNeighCell: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MeasResultBestNeighCell = new(MeasResultEUTRA)
			if err := ie.MeasResultBestNeighCell.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
