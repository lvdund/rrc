package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParametersfetype2PerBC_r17 struct {
	Fetype2basic_r17 []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,madatory`
	Fetype2R1_r17    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r17,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Fetype2R2_r17    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r17,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
}

func (ie *CodebookParametersfetype2PerBC_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Fetype2R1_r17) > 0, len(ie.Fetype2R2_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_Fetype2basic_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
	for _, i := range ie.Fetype2basic_r17 {
		tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
		tmp_Fetype2basic_r17.Value = append(tmp_Fetype2basic_r17.Value, &tmp_ie)
	}
	if err = tmp_Fetype2basic_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Fetype2basic_r17", err)
	}
	if len(ie.Fetype2R1_r17) > 0 {
		tmp_Fetype2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r17}, false)
		for _, i := range ie.Fetype2R1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Fetype2R1_r17.Value = append(tmp_Fetype2R1_r17.Value, &tmp_ie)
		}
		if err = tmp_Fetype2R1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Fetype2R1_r17", err)
		}
	}
	if len(ie.Fetype2R2_r17) > 0 {
		tmp_Fetype2R2_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r17}, false)
		for _, i := range ie.Fetype2R2_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Fetype2R2_r17.Value = append(tmp_Fetype2R2_r17.Value, &tmp_ie)
		}
		if err = tmp_Fetype2R2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Fetype2R2_r17", err)
		}
	}
	return nil
}

func (ie *CodebookParametersfetype2PerBC_r17) Decode(r *aper.AperReader) error {
	var err error
	var Fetype2R1_r17Present bool
	if Fetype2R1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fetype2R2_r17Present bool
	if Fetype2R2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_Fetype2basic_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
	fn_Fetype2basic_r17 := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
		return &ie
	}
	if err = tmp_Fetype2basic_r17.Decode(r, fn_Fetype2basic_r17); err != nil {
		return utils.WrapError("Decode Fetype2basic_r17", err)
	}
	ie.Fetype2basic_r17 = []int64{}
	for _, i := range tmp_Fetype2basic_r17.Value {
		ie.Fetype2basic_r17 = append(ie.Fetype2basic_r17, int64(i.Value))
	}
	if Fetype2R1_r17Present {
		tmp_Fetype2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r17}, false)
		fn_Fetype2R1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Fetype2R1_r17.Decode(r, fn_Fetype2R1_r17); err != nil {
			return utils.WrapError("Decode Fetype2R1_r17", err)
		}
		ie.Fetype2R1_r17 = []int64{}
		for _, i := range tmp_Fetype2R1_r17.Value {
			ie.Fetype2R1_r17 = append(ie.Fetype2R1_r17, int64(i.Value))
		}
	}
	if Fetype2R2_r17Present {
		tmp_Fetype2R2_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r17}, false)
		fn_Fetype2R2_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Fetype2R2_r17.Decode(r, fn_Fetype2R2_r17); err != nil {
			return utils.WrapError("Decode Fetype2R2_r17", err)
		}
		ie.Fetype2R2_r17 = []int64{}
		for _, i := range tmp_Fetype2R2_r17.Value {
			ie.Fetype2R2_r17 = append(ie.Fetype2R2_r17, int64(i.Value))
		}
	}
	return nil
}
