// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NTN-NeighbourCellInfo-r18 (line 9607).

var nTNNeighbourCellInfoR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "epochTime-r18"},
		{Name: "ephemerisInfo-r18"},
		{Name: "referenceLocation-r18", Optional: true},
	},
}

type NTN_NeighbourCellInfo_r18 struct {
	EpochTime_r18         EpochTime_r17
	EphemerisInfo_r18     EphemerisInfo_r17
	ReferenceLocation_r18 *ReferenceLocation_r17
}

func (ie *NTN_NeighbourCellInfo_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nTNNeighbourCellInfoR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReferenceLocation_r18 != nil}); err != nil {
		return err
	}

	// 2. epochTime-r18: ref
	{
		if err := ie.EpochTime_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. ephemerisInfo-r18: ref
	{
		if err := ie.EphemerisInfo_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. referenceLocation-r18: ref
	{
		if ie.ReferenceLocation_r18 != nil {
			if err := ie.ReferenceLocation_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NTN_NeighbourCellInfo_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nTNNeighbourCellInfoR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. epochTime-r18: ref
	{
		if err := ie.EpochTime_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. ephemerisInfo-r18: ref
	{
		if err := ie.EphemerisInfo_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. referenceLocation-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.ReferenceLocation_r18 = new(ReferenceLocation_r17)
			if err := ie.ReferenceLocation_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
