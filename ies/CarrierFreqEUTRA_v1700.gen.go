// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CarrierFreqEUTRA-v1700 (line 4157).

var carrierFreqEUTRAV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "eutra-FreqNeighHSDN-CellList-r17", Optional: true},
	},
}

type CarrierFreqEUTRA_v1700 struct {
	Eutra_FreqNeighHSDN_CellList_r17 *EUTRA_FreqNeighHSDN_CellList_r17
}

func (ie *CarrierFreqEUTRA_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(carrierFreqEUTRAV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Eutra_FreqNeighHSDN_CellList_r17 != nil}); err != nil {
		return err
	}

	// 2. eutra-FreqNeighHSDN-CellList-r17: ref
	{
		if ie.Eutra_FreqNeighHSDN_CellList_r17 != nil {
			if err := ie.Eutra_FreqNeighHSDN_CellList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CarrierFreqEUTRA_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(carrierFreqEUTRAV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. eutra-FreqNeighHSDN-CellList-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Eutra_FreqNeighHSDN_CellList_r17 = new(EUTRA_FreqNeighHSDN_CellList_r17)
			if err := ie.Eutra_FreqNeighHSDN_CellList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
