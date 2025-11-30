package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PCI_Range_range_pci_Enum_n4     aper.Enumerated = 0
	PCI_Range_range_pci_Enum_n8     aper.Enumerated = 1
	PCI_Range_range_pci_Enum_n12    aper.Enumerated = 2
	PCI_Range_range_pci_Enum_n16    aper.Enumerated = 3
	PCI_Range_range_pci_Enum_n24    aper.Enumerated = 4
	PCI_Range_range_pci_Enum_n32    aper.Enumerated = 5
	PCI_Range_range_pci_Enum_n48    aper.Enumerated = 6
	PCI_Range_range_pci_Enum_n64    aper.Enumerated = 7
	PCI_Range_range_pci_Enum_n84    aper.Enumerated = 8
	PCI_Range_range_pci_Enum_n96    aper.Enumerated = 9
	PCI_Range_range_pci_Enum_n128   aper.Enumerated = 10
	PCI_Range_range_pci_Enum_n168   aper.Enumerated = 11
	PCI_Range_range_pci_Enum_n252   aper.Enumerated = 12
	PCI_Range_range_pci_Enum_n504   aper.Enumerated = 13
	PCI_Range_range_pci_Enum_n1008  aper.Enumerated = 14
	PCI_Range_range_pci_Enum_spare1 aper.Enumerated = 15
)

type PCI_Range_range_pci struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *PCI_Range_range_pci) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode PCI_Range_range_pci", err)
	}
	return nil
}

func (ie *PCI_Range_range_pci) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode PCI_Range_range_pci", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
