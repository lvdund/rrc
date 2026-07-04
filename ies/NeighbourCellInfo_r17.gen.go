// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NeighbourCellInfo-r17 (line 26497).

var neighbourCellInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "epochTime-r17"},
		{Name: "ephemerisInfo-r17"},
	},
}

type NeighbourCellInfo_r17 struct {
	EpochTime_r17     EpochTime_r17
	EphemerisInfo_r17 EphemerisInfo_r17
}

func (ie *NeighbourCellInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(neighbourCellInfoR17Constraints)
	_ = seq

	// 1. epochTime-r17: ref
	{
		if err := ie.EpochTime_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. ephemerisInfo-r17: ref
	{
		if err := ie.EphemerisInfo_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NeighbourCellInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(neighbourCellInfoR17Constraints)
	_ = seq

	// 1. epochTime-r17: ref
	{
		if err := ie.EpochTime_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. ephemerisInfo-r17: ref
	{
		if err := ie.EphemerisInfo_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
