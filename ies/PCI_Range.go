package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PCI_Range struct {
	Start     PhysCellId           `madatory`
	Range_pci *PCI_Range_range_pci `optional`
}

func (ie *PCI_Range) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Range_pci != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Start.Encode(w); err != nil {
		return utils.WrapError("Encode Start", err)
	}
	if ie.Range_pci != nil {
		if err = ie.Range_pci.Encode(w); err != nil {
			return utils.WrapError("Encode Range_pci", err)
		}
	}
	return nil
}

func (ie *PCI_Range) Decode(r *aper.AperReader) error {
	var err error
	var Range_pciPresent bool
	if Range_pciPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Start.Decode(r); err != nil {
		return utils.WrapError("Decode Start", err)
	}
	if Range_pciPresent {
		ie.Range_pci = new(PCI_Range_range_pci)
		if err = ie.Range_pci.Decode(r); err != nil {
			return utils.WrapError("Decode Range_pci", err)
		}
	}
	return nil
}
