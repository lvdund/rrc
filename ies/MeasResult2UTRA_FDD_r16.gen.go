// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResult2UTRA-FDD-r16 (line 719).

var measResult2UTRAFDDR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq-r16"},
		{Name: "measResultNeighCellList-r16"},
	},
}

type MeasResult2UTRA_FDD_r16 struct {
	CarrierFreq_r16             ARFCN_ValueUTRA_FDD_r16
	MeasResultNeighCellList_r16 MeasResultListUTRA_FDD_r16
}

func (ie *MeasResult2UTRA_FDD_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResult2UTRAFDDR16Constraints)
	_ = seq

	// 1. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. measResultNeighCellList-r16: ref
	{
		if err := ie.MeasResultNeighCellList_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasResult2UTRA_FDD_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResult2UTRAFDDR16Constraints)
	_ = seq

	// 1. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. measResultNeighCellList-r16: ref
	{
		if err := ie.MeasResultNeighCellList_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
