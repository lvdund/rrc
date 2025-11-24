package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IntraFreqCAG_CellListPerPLMN_r16 struct {
	Plmn_IdentityIndex_r16 int64       `lb:1,ub:maxPLMN,madatory`
	Cag_CellList_r16       []PCI_Range `lb:1,ub:maxCAG_Cell_r16,madatory`
}

func (ie *IntraFreqCAG_CellListPerPLMN_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Plmn_IdentityIndex_r16, &uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
		return utils.WrapError("WriteInteger Plmn_IdentityIndex_r16", err)
	}
	tmp_Cag_CellList_r16 := utils.NewSequence[*PCI_Range]([]*PCI_Range{}, uper.Constraint{Lb: 1, Ub: maxCAG_Cell_r16}, false)
	for _, i := range ie.Cag_CellList_r16 {
		tmp_Cag_CellList_r16.Value = append(tmp_Cag_CellList_r16.Value, &i)
	}
	if err = tmp_Cag_CellList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Cag_CellList_r16", err)
	}
	return nil
}

func (ie *IntraFreqCAG_CellListPerPLMN_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_Plmn_IdentityIndex_r16 int64
	if tmp_int_Plmn_IdentityIndex_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
		return utils.WrapError("ReadInteger Plmn_IdentityIndex_r16", err)
	}
	ie.Plmn_IdentityIndex_r16 = tmp_int_Plmn_IdentityIndex_r16
	tmp_Cag_CellList_r16 := utils.NewSequence[*PCI_Range]([]*PCI_Range{}, uper.Constraint{Lb: 1, Ub: maxCAG_Cell_r16}, false)
	fn_Cag_CellList_r16 := func() *PCI_Range {
		return new(PCI_Range)
	}
	if err = tmp_Cag_CellList_r16.Decode(r, fn_Cag_CellList_r16); err != nil {
		return utils.WrapError("Decode Cag_CellList_r16", err)
	}
	ie.Cag_CellList_r16 = []PCI_Range{}
	for _, i := range tmp_Cag_CellList_r16.Value {
		ie.Cag_CellList_r16 = append(ie.Cag_CellList_r16, *i)
	}
	return nil
}
