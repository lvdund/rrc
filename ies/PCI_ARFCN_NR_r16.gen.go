// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PCI-ARFCN-NR-r16 (line 10949).

var pCIARFCNNRR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId-r16"},
		{Name: "carrierFreq-r16"},
	},
}

type PCI_ARFCN_NR_r16 struct {
	PhysCellId_r16  PhysCellId
	CarrierFreq_r16 ARFCN_ValueNR
}

func (ie *PCI_ARFCN_NR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pCIARFCNNRR16Constraints)
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

func (ie *PCI_ARFCN_NR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pCIARFCNNRR16Constraints)
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
