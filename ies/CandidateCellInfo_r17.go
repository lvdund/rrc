package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CandidateCellInfo_r17 struct {
	SsbFrequency_r17  ARFCN_ValueNR       `madatory`
	CandidateList_r17 []CandidateCell_r17 `lb:1,ub:maxNrofCondCells_r16,madatory`
}

func (ie *CandidateCellInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.SsbFrequency_r17.Encode(w); err != nil {
		return utils.WrapError("Encode SsbFrequency_r17", err)
	}
	tmp_CandidateList_r17 := utils.NewSequence[*CandidateCell_r17]([]*CandidateCell_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
	for _, i := range ie.CandidateList_r17 {
		tmp_CandidateList_r17.Value = append(tmp_CandidateList_r17.Value, &i)
	}
	if err = tmp_CandidateList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode CandidateList_r17", err)
	}
	return nil
}

func (ie *CandidateCellInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.SsbFrequency_r17.Decode(r); err != nil {
		return utils.WrapError("Decode SsbFrequency_r17", err)
	}
	tmp_CandidateList_r17 := utils.NewSequence[*CandidateCell_r17]([]*CandidateCell_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
	fn_CandidateList_r17 := func() *CandidateCell_r17 {
		return new(CandidateCell_r17)
	}
	if err = tmp_CandidateList_r17.Decode(r, fn_CandidateList_r17); err != nil {
		return utils.WrapError("Decode CandidateList_r17", err)
	}
	ie.CandidateList_r17 = []CandidateCell_r17{}
	for _, i := range tmp_CandidateList_r17.Value {
		ie.CandidateList_r17 = append(ie.CandidateList_r17, *i)
	}
	return nil
}
