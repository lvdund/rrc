package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SSB_MTC2 struct {
	Pci_List    []PhysCellId         `lb:1,ub:maxNrofPCIsPerSMTC,optional`
	Periodicity SSB_MTC2_periodicity `madatory`
}

func (ie *SSB_MTC2) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Pci_List) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Pci_List) > 0 {
		tmp_Pci_List := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, aper.Constraint{Lb: 1, Ub: maxNrofPCIsPerSMTC}, false)
		for _, i := range ie.Pci_List {
			tmp_Pci_List.Value = append(tmp_Pci_List.Value, &i)
		}
		if err = tmp_Pci_List.Encode(w); err != nil {
			return utils.WrapError("Encode Pci_List", err)
		}
	}
	if err = ie.Periodicity.Encode(w); err != nil {
		return utils.WrapError("Encode Periodicity", err)
	}
	return nil
}

func (ie *SSB_MTC2) Decode(r *aper.AperReader) error {
	var err error
	var Pci_ListPresent bool
	if Pci_ListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Pci_ListPresent {
		tmp_Pci_List := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, aper.Constraint{Lb: 1, Ub: maxNrofPCIsPerSMTC}, false)
		fn_Pci_List := func() *PhysCellId {
			return new(PhysCellId)
		}
		if err = tmp_Pci_List.Decode(r, fn_Pci_List); err != nil {
			return utils.WrapError("Decode Pci_List", err)
		}
		ie.Pci_List = []PhysCellId{}
		for _, i := range tmp_Pci_List.Value {
			ie.Pci_List = append(ie.Pci_List, *i)
		}
	}
	if err = ie.Periodicity.Decode(r); err != nil {
		return utils.WrapError("Decode Periodicity", err)
	}
	return nil
}
