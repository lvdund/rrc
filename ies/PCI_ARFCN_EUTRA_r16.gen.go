// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PCI-ARFCN-EUTRA-r16 (line 10941).

var pCIARFCNEUTRAR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId-r16"},
		{Name: "carrierFreq-r16"},
	},
}

type PCI_ARFCN_EUTRA_r16 struct {
	PhysCellId_r16  EUTRA_PhysCellId
	CarrierFreq_r16 ARFCN_ValueEUTRA
}

func (ie *PCI_ARFCN_EUTRA_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pCIARFCNEUTRAR16Constraints)
	_ = seq

	// 1. physCellId-r16: ref
	{
		if err := ie.PhysCellId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PCI_ARFCN_EUTRA_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pCIARFCNEUTRAR16Constraints)
	_ = seq

	// 1. physCellId-r16: ref
	{
		if err := ie.PhysCellId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
