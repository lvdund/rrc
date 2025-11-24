package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_CandidateList_r17_IEs struct {
	Cg_CandidateToAddModList_r17  []CG_CandidateInfo_r17   `lb:1,ub:maxNrofCondCells_r16,optional`
	Cg_CandidateToReleaseList_r17 []CG_CandidateInfoId_r17 `lb:1,ub:maxNrofCondCells_r16,optional`
	NonCriticalExtension          interface{}              `optional`
}

func (ie *CG_CandidateList_r17_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Cg_CandidateToAddModList_r17) > 0, len(ie.Cg_CandidateToReleaseList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Cg_CandidateToAddModList_r17) > 0 {
		tmp_Cg_CandidateToAddModList_r17 := utils.NewSequence[*CG_CandidateInfo_r17]([]*CG_CandidateInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
		for _, i := range ie.Cg_CandidateToAddModList_r17 {
			tmp_Cg_CandidateToAddModList_r17.Value = append(tmp_Cg_CandidateToAddModList_r17.Value, &i)
		}
		if err = tmp_Cg_CandidateToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Cg_CandidateToAddModList_r17", err)
		}
	}
	if len(ie.Cg_CandidateToReleaseList_r17) > 0 {
		tmp_Cg_CandidateToReleaseList_r17 := utils.NewSequence[*CG_CandidateInfoId_r17]([]*CG_CandidateInfoId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
		for _, i := range ie.Cg_CandidateToReleaseList_r17 {
			tmp_Cg_CandidateToReleaseList_r17.Value = append(tmp_Cg_CandidateToReleaseList_r17.Value, &i)
		}
		if err = tmp_Cg_CandidateToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Cg_CandidateToReleaseList_r17", err)
		}
	}
	return nil
}

func (ie *CG_CandidateList_r17_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Cg_CandidateToAddModList_r17Present bool
	if Cg_CandidateToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Cg_CandidateToReleaseList_r17Present bool
	if Cg_CandidateToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Cg_CandidateToAddModList_r17Present {
		tmp_Cg_CandidateToAddModList_r17 := utils.NewSequence[*CG_CandidateInfo_r17]([]*CG_CandidateInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
		fn_Cg_CandidateToAddModList_r17 := func() *CG_CandidateInfo_r17 {
			return new(CG_CandidateInfo_r17)
		}
		if err = tmp_Cg_CandidateToAddModList_r17.Decode(r, fn_Cg_CandidateToAddModList_r17); err != nil {
			return utils.WrapError("Decode Cg_CandidateToAddModList_r17", err)
		}
		ie.Cg_CandidateToAddModList_r17 = []CG_CandidateInfo_r17{}
		for _, i := range tmp_Cg_CandidateToAddModList_r17.Value {
			ie.Cg_CandidateToAddModList_r17 = append(ie.Cg_CandidateToAddModList_r17, *i)
		}
	}
	if Cg_CandidateToReleaseList_r17Present {
		tmp_Cg_CandidateToReleaseList_r17 := utils.NewSequence[*CG_CandidateInfoId_r17]([]*CG_CandidateInfoId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
		fn_Cg_CandidateToReleaseList_r17 := func() *CG_CandidateInfoId_r17 {
			return new(CG_CandidateInfoId_r17)
		}
		if err = tmp_Cg_CandidateToReleaseList_r17.Decode(r, fn_Cg_CandidateToReleaseList_r17); err != nil {
			return utils.WrapError("Decode Cg_CandidateToReleaseList_r17", err)
		}
		ie.Cg_CandidateToReleaseList_r17 = []CG_CandidateInfoId_r17{}
		for _, i := range tmp_Cg_CandidateToReleaseList_r17.Value {
			ie.Cg_CandidateToReleaseList_r17 = append(ie.Cg_CandidateToReleaseList_r17, *i)
		}
	}
	return nil
}
