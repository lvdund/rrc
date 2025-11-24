package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_MTC4_r17 struct {
	Pci_List_r17 []PhysCellId `lb:1,ub:maxNrofPCIsPerSMTC,optional`
	Offset_r17   int64        `lb:0,ub:159,madatory`
}

func (ie *SSB_MTC4_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Pci_List_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Pci_List_r17) > 0 {
		tmp_Pci_List_r17 := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxNrofPCIsPerSMTC}, false)
		for _, i := range ie.Pci_List_r17 {
			tmp_Pci_List_r17.Value = append(tmp_Pci_List_r17.Value, &i)
		}
		if err = tmp_Pci_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pci_List_r17", err)
		}
	}
	if err = w.WriteInteger(ie.Offset_r17, &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("WriteInteger Offset_r17", err)
	}
	return nil
}

func (ie *SSB_MTC4_r17) Decode(r *uper.UperReader) error {
	var err error
	var Pci_List_r17Present bool
	if Pci_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pci_List_r17Present {
		tmp_Pci_List_r17 := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxNrofPCIsPerSMTC}, false)
		fn_Pci_List_r17 := func() *PhysCellId {
			return new(PhysCellId)
		}
		if err = tmp_Pci_List_r17.Decode(r, fn_Pci_List_r17); err != nil {
			return utils.WrapError("Decode Pci_List_r17", err)
		}
		ie.Pci_List_r17 = []PhysCellId{}
		for _, i := range tmp_Pci_List_r17.Value {
			ie.Pci_List_r17 = append(ie.Pci_List_r17, *i)
		}
	}
	var tmp_int_Offset_r17 int64
	if tmp_int_Offset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("ReadInteger Offset_r17", err)
	}
	ie.Offset_r17 = tmp_int_Offset_r17
	return nil
}
