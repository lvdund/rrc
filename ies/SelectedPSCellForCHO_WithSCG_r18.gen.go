// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SelectedPSCellForCHO-WithSCG-r18 (line 14571).

var selectedPSCellForCHOWithSCGR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssbFrequency-r18"},
		{Name: "physCellId-r18"},
	},
}

type SelectedPSCellForCHO_WithSCG_r18 struct {
	SsbFrequency_r18 ARFCN_ValueNR
	PhysCellId_r18   PhysCellId
}

func (ie *SelectedPSCellForCHO_WithSCG_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(selectedPSCellForCHOWithSCGR18Constraints)
	_ = seq

	// 1. ssbFrequency-r18: ref
	{
		if err := ie.SsbFrequency_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. physCellId-r18: ref
	{
		if err := ie.PhysCellId_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SelectedPSCellForCHO_WithSCG_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(selectedPSCellForCHOWithSCGR18Constraints)
	_ = seq

	// 1. ssbFrequency-r18: ref
	{
		if err := ie.SsbFrequency_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. physCellId-r18: ref
	{
		if err := ie.PhysCellId_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
