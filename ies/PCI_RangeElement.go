package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PCI_RangeElement struct {
	Pci_RangeIndex PCI_RangeIndex `madatory`
	Pci_Range      PCI_Range      `madatory`
}

func (ie *PCI_RangeElement) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Pci_RangeIndex.Encode(w); err != nil {
		return utils.WrapError("Encode Pci_RangeIndex", err)
	}
	if err = ie.Pci_Range.Encode(w); err != nil {
		return utils.WrapError("Encode Pci_Range", err)
	}
	return nil
}

func (ie *PCI_RangeElement) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Pci_RangeIndex.Decode(r); err != nil {
		return utils.WrapError("Decode Pci_RangeIndex", err)
	}
	if err = ie.Pci_Range.Decode(r); err != nil {
		return utils.WrapError("Decode Pci_Range", err)
	}
	return nil
}
