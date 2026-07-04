// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PCI-RangeElement (line 10971).

var pCIRangeElementConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pci-RangeIndex"},
		{Name: "pci-Range"},
	},
}

type PCI_RangeElement struct {
	Pci_RangeIndex PCI_RangeIndex
	Pci_Range      PCI_Range
}

func (ie *PCI_RangeElement) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pCIRangeElementConstraints)
	_ = seq

	// 1. pci-RangeIndex: ref
	{
		if err := ie.Pci_RangeIndex.Encode(e); err != nil {
			return err
		}
	}

	// 2. pci-Range: ref
	{
		if err := ie.Pci_Range.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PCI_RangeElement) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pCIRangeElementConstraints)
	_ = seq

	// 1. pci-RangeIndex: ref
	{
		if err := ie.Pci_RangeIndex.Decode(d); err != nil {
			return err
		}
	}

	// 2. pci-Range: ref
	{
		if err := ie.Pci_Range.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
