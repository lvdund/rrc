package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParametersAddition_r16 struct {
	Etype2_r16    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Etype2_PS_r16 []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
}

func (ie *CodebookParametersAddition_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Etype2_r16) > 0, len(ie.Etype2_PS_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Etype2_r16) > 0 {
		tmp_Etype2_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Etype2_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Etype2_r16.Value = append(tmp_Etype2_r16.Value, &tmp_ie)
		}
		if err = tmp_Etype2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Etype2_r16", err)
		}
	}
	if len(ie.Etype2_PS_r16) > 0 {
		tmp_Etype2_PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Etype2_PS_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Etype2_PS_r16.Value = append(tmp_Etype2_PS_r16.Value, &tmp_ie)
		}
		if err = tmp_Etype2_PS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Etype2_PS_r16", err)
		}
	}
	return nil
}

func (ie *CodebookParametersAddition_r16) Decode(r *aper.AperReader) error {
	var err error
	var Etype2_r16Present bool
	if Etype2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Etype2_PS_r16Present bool
	if Etype2_PS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Etype2_r16Present {
		tmp_Etype2_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Etype2_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Etype2_r16.Decode(r, fn_Etype2_r16); err != nil {
			return utils.WrapError("Decode Etype2_r16", err)
		}
		ie.Etype2_r16 = []int64{}
		for _, i := range tmp_Etype2_r16.Value {
			ie.Etype2_r16 = append(ie.Etype2_r16, int64(i.Value))
		}
	}
	if Etype2_PS_r16Present {
		tmp_Etype2_PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Etype2_PS_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Etype2_PS_r16.Decode(r, fn_Etype2_PS_r16); err != nil {
			return utils.WrapError("Decode Etype2_PS_r16", err)
		}
		ie.Etype2_PS_r16 = []int64{}
		for _, i := range tmp_Etype2_PS_r16.Value {
			ie.Etype2_PS_r16 = append(ie.Etype2_PS_r16, int64(i.Value))
		}
	}
	return nil
}
