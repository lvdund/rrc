package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookComboParameterMixedType_r17 struct {
	Type1SP_feType2PS_null_r17           []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Type1SP_feType2PS_M2R1_null_r17      []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Type1SP_feType2PS_M2R2_null_r1       []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Type1SP_Type2_feType2_PS_M1_r17      []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Type1SP_Type2_feType2_PS_M2R1_r17    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Type1SP_eType2R1_feType2_PS_M1_r17   []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Type1SP_eType2R1_feType2_PS_M2R1_r17 []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Type1MP_feType2PS_null_r17           []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Type1MP_feType2PS_M2R1_null_r17      []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Type1MP_feType2PS_M2R2_null_r17      []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Type1MP_Type2_feType2_PS_M1_r17      []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Type1MP_Type2_feType2_PS_M2R1_r17    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Type1MP_eType2R1_feType2_PS_M1_r17   []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	Type1MP_eType2R1_feType2_PS_M2R1_r17 []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
}

func (ie *CodebookComboParameterMixedType_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Type1SP_feType2PS_null_r17) > 0, len(ie.Type1SP_feType2PS_M2R1_null_r17) > 0, len(ie.Type1SP_feType2PS_M2R2_null_r1) > 0, len(ie.Type1SP_Type2_feType2_PS_M1_r17) > 0, len(ie.Type1SP_Type2_feType2_PS_M2R1_r17) > 0, len(ie.Type1SP_eType2R1_feType2_PS_M1_r17) > 0, len(ie.Type1SP_eType2R1_feType2_PS_M2R1_r17) > 0, len(ie.Type1MP_feType2PS_null_r17) > 0, len(ie.Type1MP_feType2PS_M2R1_null_r17) > 0, len(ie.Type1MP_feType2PS_M2R2_null_r17) > 0, len(ie.Type1MP_Type2_feType2_PS_M1_r17) > 0, len(ie.Type1MP_Type2_feType2_PS_M2R1_r17) > 0, len(ie.Type1MP_eType2R1_feType2_PS_M1_r17) > 0, len(ie.Type1MP_eType2R1_feType2_PS_M2R1_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Type1SP_feType2PS_null_r17) > 0 {
		tmp_Type1SP_feType2PS_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Type1SP_feType2PS_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Type1SP_feType2PS_null_r17.Value = append(tmp_Type1SP_feType2PS_null_r17.Value, &tmp_ie)
		}
		if err = tmp_Type1SP_feType2PS_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1SP_feType2PS_null_r17", err)
		}
	}
	if len(ie.Type1SP_feType2PS_M2R1_null_r17) > 0 {
		tmp_Type1SP_feType2PS_M2R1_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Type1SP_feType2PS_M2R1_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Type1SP_feType2PS_M2R1_null_r17.Value = append(tmp_Type1SP_feType2PS_M2R1_null_r17.Value, &tmp_ie)
		}
		if err = tmp_Type1SP_feType2PS_M2R1_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1SP_feType2PS_M2R1_null_r17", err)
		}
	}
	if len(ie.Type1SP_feType2PS_M2R2_null_r1) > 0 {
		tmp_Type1SP_feType2PS_M2R2_null_r1 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Type1SP_feType2PS_M2R2_null_r1 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Type1SP_feType2PS_M2R2_null_r1.Value = append(tmp_Type1SP_feType2PS_M2R2_null_r1.Value, &tmp_ie)
		}
		if err = tmp_Type1SP_feType2PS_M2R2_null_r1.Encode(w); err != nil {
			return utils.WrapError("Encode Type1SP_feType2PS_M2R2_null_r1", err)
		}
	}
	if len(ie.Type1SP_Type2_feType2_PS_M1_r17) > 0 {
		tmp_Type1SP_Type2_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Type1SP_Type2_feType2_PS_M1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Type1SP_Type2_feType2_PS_M1_r17.Value = append(tmp_Type1SP_Type2_feType2_PS_M1_r17.Value, &tmp_ie)
		}
		if err = tmp_Type1SP_Type2_feType2_PS_M1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1SP_Type2_feType2_PS_M1_r17", err)
		}
	}
	if len(ie.Type1SP_Type2_feType2_PS_M2R1_r17) > 0 {
		tmp_Type1SP_Type2_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Type1SP_Type2_feType2_PS_M2R1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Type1SP_Type2_feType2_PS_M2R1_r17.Value = append(tmp_Type1SP_Type2_feType2_PS_M2R1_r17.Value, &tmp_ie)
		}
		if err = tmp_Type1SP_Type2_feType2_PS_M2R1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1SP_Type2_feType2_PS_M2R1_r17", err)
		}
	}
	if len(ie.Type1SP_eType2R1_feType2_PS_M1_r17) > 0 {
		tmp_Type1SP_eType2R1_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Type1SP_eType2R1_feType2_PS_M1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Type1SP_eType2R1_feType2_PS_M1_r17.Value = append(tmp_Type1SP_eType2R1_feType2_PS_M1_r17.Value, &tmp_ie)
		}
		if err = tmp_Type1SP_eType2R1_feType2_PS_M1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1SP_eType2R1_feType2_PS_M1_r17", err)
		}
	}
	if len(ie.Type1SP_eType2R1_feType2_PS_M2R1_r17) > 0 {
		tmp_Type1SP_eType2R1_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Type1SP_eType2R1_feType2_PS_M2R1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Type1SP_eType2R1_feType2_PS_M2R1_r17.Value = append(tmp_Type1SP_eType2R1_feType2_PS_M2R1_r17.Value, &tmp_ie)
		}
		if err = tmp_Type1SP_eType2R1_feType2_PS_M2R1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1SP_eType2R1_feType2_PS_M2R1_r17", err)
		}
	}
	if len(ie.Type1MP_feType2PS_null_r17) > 0 {
		tmp_Type1MP_feType2PS_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Type1MP_feType2PS_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Type1MP_feType2PS_null_r17.Value = append(tmp_Type1MP_feType2PS_null_r17.Value, &tmp_ie)
		}
		if err = tmp_Type1MP_feType2PS_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1MP_feType2PS_null_r17", err)
		}
	}
	if len(ie.Type1MP_feType2PS_M2R1_null_r17) > 0 {
		tmp_Type1MP_feType2PS_M2R1_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Type1MP_feType2PS_M2R1_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Type1MP_feType2PS_M2R1_null_r17.Value = append(tmp_Type1MP_feType2PS_M2R1_null_r17.Value, &tmp_ie)
		}
		if err = tmp_Type1MP_feType2PS_M2R1_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1MP_feType2PS_M2R1_null_r17", err)
		}
	}
	if len(ie.Type1MP_feType2PS_M2R2_null_r17) > 0 {
		tmp_Type1MP_feType2PS_M2R2_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Type1MP_feType2PS_M2R2_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Type1MP_feType2PS_M2R2_null_r17.Value = append(tmp_Type1MP_feType2PS_M2R2_null_r17.Value, &tmp_ie)
		}
		if err = tmp_Type1MP_feType2PS_M2R2_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1MP_feType2PS_M2R2_null_r17", err)
		}
	}
	if len(ie.Type1MP_Type2_feType2_PS_M1_r17) > 0 {
		tmp_Type1MP_Type2_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Type1MP_Type2_feType2_PS_M1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Type1MP_Type2_feType2_PS_M1_r17.Value = append(tmp_Type1MP_Type2_feType2_PS_M1_r17.Value, &tmp_ie)
		}
		if err = tmp_Type1MP_Type2_feType2_PS_M1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1MP_Type2_feType2_PS_M1_r17", err)
		}
	}
	if len(ie.Type1MP_Type2_feType2_PS_M2R1_r17) > 0 {
		tmp_Type1MP_Type2_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Type1MP_Type2_feType2_PS_M2R1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Type1MP_Type2_feType2_PS_M2R1_r17.Value = append(tmp_Type1MP_Type2_feType2_PS_M2R1_r17.Value, &tmp_ie)
		}
		if err = tmp_Type1MP_Type2_feType2_PS_M2R1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1MP_Type2_feType2_PS_M2R1_r17", err)
		}
	}
	if len(ie.Type1MP_eType2R1_feType2_PS_M1_r17) > 0 {
		tmp_Type1MP_eType2R1_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Type1MP_eType2R1_feType2_PS_M1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Type1MP_eType2R1_feType2_PS_M1_r17.Value = append(tmp_Type1MP_eType2R1_feType2_PS_M1_r17.Value, &tmp_ie)
		}
		if err = tmp_Type1MP_eType2R1_feType2_PS_M1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1MP_eType2R1_feType2_PS_M1_r17", err)
		}
	}
	if len(ie.Type1MP_eType2R1_feType2_PS_M2R1_r17) > 0 {
		tmp_Type1MP_eType2R1_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.Type1MP_eType2R1_feType2_PS_M2R1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_Type1MP_eType2R1_feType2_PS_M2R1_r17.Value = append(tmp_Type1MP_eType2R1_feType2_PS_M2R1_r17.Value, &tmp_ie)
		}
		if err = tmp_Type1MP_eType2R1_feType2_PS_M2R1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type1MP_eType2R1_feType2_PS_M2R1_r17", err)
		}
	}
	return nil
}

func (ie *CodebookComboParameterMixedType_r17) Decode(r *aper.AperReader) error {
	var err error
	var Type1SP_feType2PS_null_r17Present bool
	if Type1SP_feType2PS_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1SP_feType2PS_M2R1_null_r17Present bool
	if Type1SP_feType2PS_M2R1_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1SP_feType2PS_M2R2_null_r1Present bool
	if Type1SP_feType2PS_M2R2_null_r1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1SP_Type2_feType2_PS_M1_r17Present bool
	if Type1SP_Type2_feType2_PS_M1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1SP_Type2_feType2_PS_M2R1_r17Present bool
	if Type1SP_Type2_feType2_PS_M2R1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1SP_eType2R1_feType2_PS_M1_r17Present bool
	if Type1SP_eType2R1_feType2_PS_M1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1SP_eType2R1_feType2_PS_M2R1_r17Present bool
	if Type1SP_eType2R1_feType2_PS_M2R1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1MP_feType2PS_null_r17Present bool
	if Type1MP_feType2PS_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1MP_feType2PS_M2R1_null_r17Present bool
	if Type1MP_feType2PS_M2R1_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1MP_feType2PS_M2R2_null_r17Present bool
	if Type1MP_feType2PS_M2R2_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1MP_Type2_feType2_PS_M1_r17Present bool
	if Type1MP_Type2_feType2_PS_M1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1MP_Type2_feType2_PS_M2R1_r17Present bool
	if Type1MP_Type2_feType2_PS_M2R1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1MP_eType2R1_feType2_PS_M1_r17Present bool
	if Type1MP_eType2R1_feType2_PS_M1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1MP_eType2R1_feType2_PS_M2R1_r17Present bool
	if Type1MP_eType2R1_feType2_PS_M2R1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Type1SP_feType2PS_null_r17Present {
		tmp_Type1SP_feType2PS_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Type1SP_feType2PS_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Type1SP_feType2PS_null_r17.Decode(r, fn_Type1SP_feType2PS_null_r17); err != nil {
			return utils.WrapError("Decode Type1SP_feType2PS_null_r17", err)
		}
		ie.Type1SP_feType2PS_null_r17 = []int64{}
		for _, i := range tmp_Type1SP_feType2PS_null_r17.Value {
			ie.Type1SP_feType2PS_null_r17 = append(ie.Type1SP_feType2PS_null_r17, int64(i.Value))
		}
	}
	if Type1SP_feType2PS_M2R1_null_r17Present {
		tmp_Type1SP_feType2PS_M2R1_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Type1SP_feType2PS_M2R1_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Type1SP_feType2PS_M2R1_null_r17.Decode(r, fn_Type1SP_feType2PS_M2R1_null_r17); err != nil {
			return utils.WrapError("Decode Type1SP_feType2PS_M2R1_null_r17", err)
		}
		ie.Type1SP_feType2PS_M2R1_null_r17 = []int64{}
		for _, i := range tmp_Type1SP_feType2PS_M2R1_null_r17.Value {
			ie.Type1SP_feType2PS_M2R1_null_r17 = append(ie.Type1SP_feType2PS_M2R1_null_r17, int64(i.Value))
		}
	}
	if Type1SP_feType2PS_M2R2_null_r1Present {
		tmp_Type1SP_feType2PS_M2R2_null_r1 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Type1SP_feType2PS_M2R2_null_r1 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Type1SP_feType2PS_M2R2_null_r1.Decode(r, fn_Type1SP_feType2PS_M2R2_null_r1); err != nil {
			return utils.WrapError("Decode Type1SP_feType2PS_M2R2_null_r1", err)
		}
		ie.Type1SP_feType2PS_M2R2_null_r1 = []int64{}
		for _, i := range tmp_Type1SP_feType2PS_M2R2_null_r1.Value {
			ie.Type1SP_feType2PS_M2R2_null_r1 = append(ie.Type1SP_feType2PS_M2R2_null_r1, int64(i.Value))
		}
	}
	if Type1SP_Type2_feType2_PS_M1_r17Present {
		tmp_Type1SP_Type2_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Type1SP_Type2_feType2_PS_M1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Type1SP_Type2_feType2_PS_M1_r17.Decode(r, fn_Type1SP_Type2_feType2_PS_M1_r17); err != nil {
			return utils.WrapError("Decode Type1SP_Type2_feType2_PS_M1_r17", err)
		}
		ie.Type1SP_Type2_feType2_PS_M1_r17 = []int64{}
		for _, i := range tmp_Type1SP_Type2_feType2_PS_M1_r17.Value {
			ie.Type1SP_Type2_feType2_PS_M1_r17 = append(ie.Type1SP_Type2_feType2_PS_M1_r17, int64(i.Value))
		}
	}
	if Type1SP_Type2_feType2_PS_M2R1_r17Present {
		tmp_Type1SP_Type2_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Type1SP_Type2_feType2_PS_M2R1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Type1SP_Type2_feType2_PS_M2R1_r17.Decode(r, fn_Type1SP_Type2_feType2_PS_M2R1_r17); err != nil {
			return utils.WrapError("Decode Type1SP_Type2_feType2_PS_M2R1_r17", err)
		}
		ie.Type1SP_Type2_feType2_PS_M2R1_r17 = []int64{}
		for _, i := range tmp_Type1SP_Type2_feType2_PS_M2R1_r17.Value {
			ie.Type1SP_Type2_feType2_PS_M2R1_r17 = append(ie.Type1SP_Type2_feType2_PS_M2R1_r17, int64(i.Value))
		}
	}
	if Type1SP_eType2R1_feType2_PS_M1_r17Present {
		tmp_Type1SP_eType2R1_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Type1SP_eType2R1_feType2_PS_M1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Type1SP_eType2R1_feType2_PS_M1_r17.Decode(r, fn_Type1SP_eType2R1_feType2_PS_M1_r17); err != nil {
			return utils.WrapError("Decode Type1SP_eType2R1_feType2_PS_M1_r17", err)
		}
		ie.Type1SP_eType2R1_feType2_PS_M1_r17 = []int64{}
		for _, i := range tmp_Type1SP_eType2R1_feType2_PS_M1_r17.Value {
			ie.Type1SP_eType2R1_feType2_PS_M1_r17 = append(ie.Type1SP_eType2R1_feType2_PS_M1_r17, int64(i.Value))
		}
	}
	if Type1SP_eType2R1_feType2_PS_M2R1_r17Present {
		tmp_Type1SP_eType2R1_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Type1SP_eType2R1_feType2_PS_M2R1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Type1SP_eType2R1_feType2_PS_M2R1_r17.Decode(r, fn_Type1SP_eType2R1_feType2_PS_M2R1_r17); err != nil {
			return utils.WrapError("Decode Type1SP_eType2R1_feType2_PS_M2R1_r17", err)
		}
		ie.Type1SP_eType2R1_feType2_PS_M2R1_r17 = []int64{}
		for _, i := range tmp_Type1SP_eType2R1_feType2_PS_M2R1_r17.Value {
			ie.Type1SP_eType2R1_feType2_PS_M2R1_r17 = append(ie.Type1SP_eType2R1_feType2_PS_M2R1_r17, int64(i.Value))
		}
	}
	if Type1MP_feType2PS_null_r17Present {
		tmp_Type1MP_feType2PS_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Type1MP_feType2PS_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Type1MP_feType2PS_null_r17.Decode(r, fn_Type1MP_feType2PS_null_r17); err != nil {
			return utils.WrapError("Decode Type1MP_feType2PS_null_r17", err)
		}
		ie.Type1MP_feType2PS_null_r17 = []int64{}
		for _, i := range tmp_Type1MP_feType2PS_null_r17.Value {
			ie.Type1MP_feType2PS_null_r17 = append(ie.Type1MP_feType2PS_null_r17, int64(i.Value))
		}
	}
	if Type1MP_feType2PS_M2R1_null_r17Present {
		tmp_Type1MP_feType2PS_M2R1_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Type1MP_feType2PS_M2R1_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Type1MP_feType2PS_M2R1_null_r17.Decode(r, fn_Type1MP_feType2PS_M2R1_null_r17); err != nil {
			return utils.WrapError("Decode Type1MP_feType2PS_M2R1_null_r17", err)
		}
		ie.Type1MP_feType2PS_M2R1_null_r17 = []int64{}
		for _, i := range tmp_Type1MP_feType2PS_M2R1_null_r17.Value {
			ie.Type1MP_feType2PS_M2R1_null_r17 = append(ie.Type1MP_feType2PS_M2R1_null_r17, int64(i.Value))
		}
	}
	if Type1MP_feType2PS_M2R2_null_r17Present {
		tmp_Type1MP_feType2PS_M2R2_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Type1MP_feType2PS_M2R2_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Type1MP_feType2PS_M2R2_null_r17.Decode(r, fn_Type1MP_feType2PS_M2R2_null_r17); err != nil {
			return utils.WrapError("Decode Type1MP_feType2PS_M2R2_null_r17", err)
		}
		ie.Type1MP_feType2PS_M2R2_null_r17 = []int64{}
		for _, i := range tmp_Type1MP_feType2PS_M2R2_null_r17.Value {
			ie.Type1MP_feType2PS_M2R2_null_r17 = append(ie.Type1MP_feType2PS_M2R2_null_r17, int64(i.Value))
		}
	}
	if Type1MP_Type2_feType2_PS_M1_r17Present {
		tmp_Type1MP_Type2_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Type1MP_Type2_feType2_PS_M1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Type1MP_Type2_feType2_PS_M1_r17.Decode(r, fn_Type1MP_Type2_feType2_PS_M1_r17); err != nil {
			return utils.WrapError("Decode Type1MP_Type2_feType2_PS_M1_r17", err)
		}
		ie.Type1MP_Type2_feType2_PS_M1_r17 = []int64{}
		for _, i := range tmp_Type1MP_Type2_feType2_PS_M1_r17.Value {
			ie.Type1MP_Type2_feType2_PS_M1_r17 = append(ie.Type1MP_Type2_feType2_PS_M1_r17, int64(i.Value))
		}
	}
	if Type1MP_Type2_feType2_PS_M2R1_r17Present {
		tmp_Type1MP_Type2_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Type1MP_Type2_feType2_PS_M2R1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Type1MP_Type2_feType2_PS_M2R1_r17.Decode(r, fn_Type1MP_Type2_feType2_PS_M2R1_r17); err != nil {
			return utils.WrapError("Decode Type1MP_Type2_feType2_PS_M2R1_r17", err)
		}
		ie.Type1MP_Type2_feType2_PS_M2R1_r17 = []int64{}
		for _, i := range tmp_Type1MP_Type2_feType2_PS_M2R1_r17.Value {
			ie.Type1MP_Type2_feType2_PS_M2R1_r17 = append(ie.Type1MP_Type2_feType2_PS_M2R1_r17, int64(i.Value))
		}
	}
	if Type1MP_eType2R1_feType2_PS_M1_r17Present {
		tmp_Type1MP_eType2R1_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Type1MP_eType2R1_feType2_PS_M1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Type1MP_eType2R1_feType2_PS_M1_r17.Decode(r, fn_Type1MP_eType2R1_feType2_PS_M1_r17); err != nil {
			return utils.WrapError("Decode Type1MP_eType2R1_feType2_PS_M1_r17", err)
		}
		ie.Type1MP_eType2R1_feType2_PS_M1_r17 = []int64{}
		for _, i := range tmp_Type1MP_eType2R1_feType2_PS_M1_r17.Value {
			ie.Type1MP_eType2R1_feType2_PS_M1_r17 = append(ie.Type1MP_eType2R1_feType2_PS_M1_r17, int64(i.Value))
		}
	}
	if Type1MP_eType2R1_feType2_PS_M2R1_r17Present {
		tmp_Type1MP_eType2R1_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_Type1MP_eType2R1_feType2_PS_M2R1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_Type1MP_eType2R1_feType2_PS_M2R1_r17.Decode(r, fn_Type1MP_eType2R1_feType2_PS_M2R1_r17); err != nil {
			return utils.WrapError("Decode Type1MP_eType2R1_feType2_PS_M2R1_r17", err)
		}
		ie.Type1MP_eType2R1_feType2_PS_M2R1_r17 = []int64{}
		for _, i := range tmp_Type1MP_eType2R1_feType2_PS_M2R1_r17.Value {
			ie.Type1MP_eType2R1_feType2_PS_M2R1_r17 = append(ie.Type1MP_eType2R1_feType2_PS_M2R1_r17, int64(i.Value))
		}
	}
	return nil
}
