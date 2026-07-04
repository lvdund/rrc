// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MBS-NeighbourCell-r17 (line 28456).

var mBSNeighbourCellR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId-r17"},
		{Name: "carrierFreq-r17", Optional: true},
	},
}

type MBS_NeighbourCell_r17 struct {
	PhysCellId_r17  PhysCellId
	CarrierFreq_r17 *ARFCN_ValueNR
}

func (ie *MBS_NeighbourCell_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSNeighbourCellR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CarrierFreq_r17 != nil}); err != nil {
		return err
	}

	// 2. physCellId-r17: ref
	{
		if err := ie.PhysCellId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. carrierFreq-r17: ref
	{
		if ie.CarrierFreq_r17 != nil {
			if err := ie.CarrierFreq_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MBS_NeighbourCell_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSNeighbourCellR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. physCellId-r17: ref
	{
		if err := ie.PhysCellId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. carrierFreq-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.CarrierFreq_r17 = new(ARFCN_ValueNR)
			if err := ie.CarrierFreq_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
