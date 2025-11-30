package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookComboParameterMultiTRP_PerBC_r17 struct {
	NCJT_null_null                       []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_null_null                    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT_Type2_null_r16                  []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT_Type2PS_null_r16                []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT_eType2R1_null_r16               []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT_eType2R2_null_r16               []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT_eType2R1PS_null_r16             []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT_eType2R2PS_null_r16             []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT_Type2_Type2PS_r16               []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_Type2_null_r16               []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_Type2PS_null_r16             []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_eType2R1_null_r16            []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_eType2R2_null_r16            []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_eType2R1PS_null_r16          []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_eType2R2PS_null_r16          []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_Type2_Type2PS_r16            []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT_feType2PS_null_r17              []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT_feType2PS_M2R1_null_r17         []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT_feType2PS_M2R2_null_r17         []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT_Type2_feType2_PS_M1_r17         []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT_Type2_feType2_PS_M2R1_r17       []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT_eType2R1_feType2_PS_M1_r17      []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT_eType2R1_feType2_PS_M2R1_r17    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_feType2PS_null_r17           []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_feType2PS_M2R1_null_r17      []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_feType2PS_M2R2_null_r1       []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_Type2_feType2_PS_M1_r17      []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_Type2_feType2_PS_M2R1_r17    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_eType2R1_feType2_PS_M1_r17   []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	NCJT1SP_eType2R1_feType2_PS_M2R1_r17 []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
}

func (ie *CodebookComboParameterMultiTRP_PerBC_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.NCJT_null_null) > 0, len(ie.NCJT1SP_null_null) > 0, len(ie.NCJT_Type2_null_r16) > 0, len(ie.NCJT_Type2PS_null_r16) > 0, len(ie.NCJT_eType2R1_null_r16) > 0, len(ie.NCJT_eType2R2_null_r16) > 0, len(ie.NCJT_eType2R1PS_null_r16) > 0, len(ie.NCJT_eType2R2PS_null_r16) > 0, len(ie.NCJT_Type2_Type2PS_r16) > 0, len(ie.NCJT1SP_Type2_null_r16) > 0, len(ie.NCJT1SP_Type2PS_null_r16) > 0, len(ie.NCJT1SP_eType2R1_null_r16) > 0, len(ie.NCJT1SP_eType2R2_null_r16) > 0, len(ie.NCJT1SP_eType2R1PS_null_r16) > 0, len(ie.NCJT1SP_eType2R2PS_null_r16) > 0, len(ie.NCJT1SP_Type2_Type2PS_r16) > 0, len(ie.NCJT_feType2PS_null_r17) > 0, len(ie.NCJT_feType2PS_M2R1_null_r17) > 0, len(ie.NCJT_feType2PS_M2R2_null_r17) > 0, len(ie.NCJT_Type2_feType2_PS_M1_r17) > 0, len(ie.NCJT_Type2_feType2_PS_M2R1_r17) > 0, len(ie.NCJT_eType2R1_feType2_PS_M1_r17) > 0, len(ie.NCJT_eType2R1_feType2_PS_M2R1_r17) > 0, len(ie.NCJT1SP_feType2PS_null_r17) > 0, len(ie.NCJT1SP_feType2PS_M2R1_null_r17) > 0, len(ie.NCJT1SP_feType2PS_M2R2_null_r1) > 0, len(ie.NCJT1SP_Type2_feType2_PS_M1_r17) > 0, len(ie.NCJT1SP_Type2_feType2_PS_M2R1_r17) > 0, len(ie.NCJT1SP_eType2R1_feType2_PS_M1_r17) > 0, len(ie.NCJT1SP_eType2R1_feType2_PS_M2R1_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.NCJT_null_null) > 0 {
		tmp_NCJT_null_null := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_null_null {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_null_null.Value = append(tmp_NCJT_null_null.Value, &tmp_ie)
		}
		if err = tmp_NCJT_null_null.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_null_null", err)
		}
	}
	if len(ie.NCJT1SP_null_null) > 0 {
		tmp_NCJT1SP_null_null := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_null_null {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_null_null.Value = append(tmp_NCJT1SP_null_null.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_null_null.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_null_null", err)
		}
	}
	if len(ie.NCJT_Type2_null_r16) > 0 {
		tmp_NCJT_Type2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_Type2_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_Type2_null_r16.Value = append(tmp_NCJT_Type2_null_r16.Value, &tmp_ie)
		}
		if err = tmp_NCJT_Type2_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_Type2_null_r16", err)
		}
	}
	if len(ie.NCJT_Type2PS_null_r16) > 0 {
		tmp_NCJT_Type2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_Type2PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_Type2PS_null_r16.Value = append(tmp_NCJT_Type2PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_NCJT_Type2PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_Type2PS_null_r16", err)
		}
	}
	if len(ie.NCJT_eType2R1_null_r16) > 0 {
		tmp_NCJT_eType2R1_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_eType2R1_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_eType2R1_null_r16.Value = append(tmp_NCJT_eType2R1_null_r16.Value, &tmp_ie)
		}
		if err = tmp_NCJT_eType2R1_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_eType2R1_null_r16", err)
		}
	}
	if len(ie.NCJT_eType2R2_null_r16) > 0 {
		tmp_NCJT_eType2R2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_eType2R2_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_eType2R2_null_r16.Value = append(tmp_NCJT_eType2R2_null_r16.Value, &tmp_ie)
		}
		if err = tmp_NCJT_eType2R2_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_eType2R2_null_r16", err)
		}
	}
	if len(ie.NCJT_eType2R1PS_null_r16) > 0 {
		tmp_NCJT_eType2R1PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_eType2R1PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_eType2R1PS_null_r16.Value = append(tmp_NCJT_eType2R1PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_NCJT_eType2R1PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_eType2R1PS_null_r16", err)
		}
	}
	if len(ie.NCJT_eType2R2PS_null_r16) > 0 {
		tmp_NCJT_eType2R2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_eType2R2PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_eType2R2PS_null_r16.Value = append(tmp_NCJT_eType2R2PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_NCJT_eType2R2PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_eType2R2PS_null_r16", err)
		}
	}
	if len(ie.NCJT_Type2_Type2PS_r16) > 0 {
		tmp_NCJT_Type2_Type2PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_Type2_Type2PS_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_Type2_Type2PS_r16.Value = append(tmp_NCJT_Type2_Type2PS_r16.Value, &tmp_ie)
		}
		if err = tmp_NCJT_Type2_Type2PS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_Type2_Type2PS_r16", err)
		}
	}
	if len(ie.NCJT1SP_Type2_null_r16) > 0 {
		tmp_NCJT1SP_Type2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_Type2_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_Type2_null_r16.Value = append(tmp_NCJT1SP_Type2_null_r16.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_Type2_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_Type2_null_r16", err)
		}
	}
	if len(ie.NCJT1SP_Type2PS_null_r16) > 0 {
		tmp_NCJT1SP_Type2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_Type2PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_Type2PS_null_r16.Value = append(tmp_NCJT1SP_Type2PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_Type2PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_Type2PS_null_r16", err)
		}
	}
	if len(ie.NCJT1SP_eType2R1_null_r16) > 0 {
		tmp_NCJT1SP_eType2R1_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_eType2R1_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_eType2R1_null_r16.Value = append(tmp_NCJT1SP_eType2R1_null_r16.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_eType2R1_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_eType2R1_null_r16", err)
		}
	}
	if len(ie.NCJT1SP_eType2R2_null_r16) > 0 {
		tmp_NCJT1SP_eType2R2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_eType2R2_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_eType2R2_null_r16.Value = append(tmp_NCJT1SP_eType2R2_null_r16.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_eType2R2_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_eType2R2_null_r16", err)
		}
	}
	if len(ie.NCJT1SP_eType2R1PS_null_r16) > 0 {
		tmp_NCJT1SP_eType2R1PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_eType2R1PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_eType2R1PS_null_r16.Value = append(tmp_NCJT1SP_eType2R1PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_eType2R1PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_eType2R1PS_null_r16", err)
		}
	}
	if len(ie.NCJT1SP_eType2R2PS_null_r16) > 0 {
		tmp_NCJT1SP_eType2R2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_eType2R2PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_eType2R2PS_null_r16.Value = append(tmp_NCJT1SP_eType2R2PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_eType2R2PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_eType2R2PS_null_r16", err)
		}
	}
	if len(ie.NCJT1SP_Type2_Type2PS_r16) > 0 {
		tmp_NCJT1SP_Type2_Type2PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_Type2_Type2PS_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_Type2_Type2PS_r16.Value = append(tmp_NCJT1SP_Type2_Type2PS_r16.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_Type2_Type2PS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_Type2_Type2PS_r16", err)
		}
	}
	if len(ie.NCJT_feType2PS_null_r17) > 0 {
		tmp_NCJT_feType2PS_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_feType2PS_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_feType2PS_null_r17.Value = append(tmp_NCJT_feType2PS_null_r17.Value, &tmp_ie)
		}
		if err = tmp_NCJT_feType2PS_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_feType2PS_null_r17", err)
		}
	}
	if len(ie.NCJT_feType2PS_M2R1_null_r17) > 0 {
		tmp_NCJT_feType2PS_M2R1_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_feType2PS_M2R1_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_feType2PS_M2R1_null_r17.Value = append(tmp_NCJT_feType2PS_M2R1_null_r17.Value, &tmp_ie)
		}
		if err = tmp_NCJT_feType2PS_M2R1_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_feType2PS_M2R1_null_r17", err)
		}
	}
	if len(ie.NCJT_feType2PS_M2R2_null_r17) > 0 {
		tmp_NCJT_feType2PS_M2R2_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_feType2PS_M2R2_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_feType2PS_M2R2_null_r17.Value = append(tmp_NCJT_feType2PS_M2R2_null_r17.Value, &tmp_ie)
		}
		if err = tmp_NCJT_feType2PS_M2R2_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_feType2PS_M2R2_null_r17", err)
		}
	}
	if len(ie.NCJT_Type2_feType2_PS_M1_r17) > 0 {
		tmp_NCJT_Type2_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_Type2_feType2_PS_M1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_Type2_feType2_PS_M1_r17.Value = append(tmp_NCJT_Type2_feType2_PS_M1_r17.Value, &tmp_ie)
		}
		if err = tmp_NCJT_Type2_feType2_PS_M1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_Type2_feType2_PS_M1_r17", err)
		}
	}
	if len(ie.NCJT_Type2_feType2_PS_M2R1_r17) > 0 {
		tmp_NCJT_Type2_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_Type2_feType2_PS_M2R1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_Type2_feType2_PS_M2R1_r17.Value = append(tmp_NCJT_Type2_feType2_PS_M2R1_r17.Value, &tmp_ie)
		}
		if err = tmp_NCJT_Type2_feType2_PS_M2R1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_Type2_feType2_PS_M2R1_r17", err)
		}
	}
	if len(ie.NCJT_eType2R1_feType2_PS_M1_r17) > 0 {
		tmp_NCJT_eType2R1_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_eType2R1_feType2_PS_M1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_eType2R1_feType2_PS_M1_r17.Value = append(tmp_NCJT_eType2R1_feType2_PS_M1_r17.Value, &tmp_ie)
		}
		if err = tmp_NCJT_eType2R1_feType2_PS_M1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_eType2R1_feType2_PS_M1_r17", err)
		}
	}
	if len(ie.NCJT_eType2R1_feType2_PS_M2R1_r17) > 0 {
		tmp_NCJT_eType2R1_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT_eType2R1_feType2_PS_M2R1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT_eType2R1_feType2_PS_M2R1_r17.Value = append(tmp_NCJT_eType2R1_feType2_PS_M2R1_r17.Value, &tmp_ie)
		}
		if err = tmp_NCJT_eType2R1_feType2_PS_M2R1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT_eType2R1_feType2_PS_M2R1_r17", err)
		}
	}
	if len(ie.NCJT1SP_feType2PS_null_r17) > 0 {
		tmp_NCJT1SP_feType2PS_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_feType2PS_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_feType2PS_null_r17.Value = append(tmp_NCJT1SP_feType2PS_null_r17.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_feType2PS_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_feType2PS_null_r17", err)
		}
	}
	if len(ie.NCJT1SP_feType2PS_M2R1_null_r17) > 0 {
		tmp_NCJT1SP_feType2PS_M2R1_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_feType2PS_M2R1_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_feType2PS_M2R1_null_r17.Value = append(tmp_NCJT1SP_feType2PS_M2R1_null_r17.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_feType2PS_M2R1_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_feType2PS_M2R1_null_r17", err)
		}
	}
	if len(ie.NCJT1SP_feType2PS_M2R2_null_r1) > 0 {
		tmp_NCJT1SP_feType2PS_M2R2_null_r1 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_feType2PS_M2R2_null_r1 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_feType2PS_M2R2_null_r1.Value = append(tmp_NCJT1SP_feType2PS_M2R2_null_r1.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_feType2PS_M2R2_null_r1.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_feType2PS_M2R2_null_r1", err)
		}
	}
	if len(ie.NCJT1SP_Type2_feType2_PS_M1_r17) > 0 {
		tmp_NCJT1SP_Type2_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_Type2_feType2_PS_M1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_Type2_feType2_PS_M1_r17.Value = append(tmp_NCJT1SP_Type2_feType2_PS_M1_r17.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_Type2_feType2_PS_M1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_Type2_feType2_PS_M1_r17", err)
		}
	}
	if len(ie.NCJT1SP_Type2_feType2_PS_M2R1_r17) > 0 {
		tmp_NCJT1SP_Type2_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_Type2_feType2_PS_M2R1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_Type2_feType2_PS_M2R1_r17.Value = append(tmp_NCJT1SP_Type2_feType2_PS_M2R1_r17.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_Type2_feType2_PS_M2R1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_Type2_feType2_PS_M2R1_r17", err)
		}
	}
	if len(ie.NCJT1SP_eType2R1_feType2_PS_M1_r17) > 0 {
		tmp_NCJT1SP_eType2R1_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_eType2R1_feType2_PS_M1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_eType2R1_feType2_PS_M1_r17.Value = append(tmp_NCJT1SP_eType2R1_feType2_PS_M1_r17.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_eType2R1_feType2_PS_M1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_eType2R1_feType2_PS_M1_r17", err)
		}
	}
	if len(ie.NCJT1SP_eType2R1_feType2_PS_M2R1_r17) > 0 {
		tmp_NCJT1SP_eType2R1_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.NCJT1SP_eType2R1_feType2_PS_M2R1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_NCJT1SP_eType2R1_feType2_PS_M2R1_r17.Value = append(tmp_NCJT1SP_eType2R1_feType2_PS_M2R1_r17.Value, &tmp_ie)
		}
		if err = tmp_NCJT1SP_eType2R1_feType2_PS_M2R1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NCJT1SP_eType2R1_feType2_PS_M2R1_r17", err)
		}
	}
	return nil
}

func (ie *CodebookComboParameterMultiTRP_PerBC_r17) Decode(r *aper.AperReader) error {
	var err error
	var NCJT_null_nullPresent bool
	if NCJT_null_nullPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_null_nullPresent bool
	if NCJT1SP_null_nullPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT_Type2_null_r16Present bool
	if NCJT_Type2_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT_Type2PS_null_r16Present bool
	if NCJT_Type2PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT_eType2R1_null_r16Present bool
	if NCJT_eType2R1_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT_eType2R2_null_r16Present bool
	if NCJT_eType2R2_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT_eType2R1PS_null_r16Present bool
	if NCJT_eType2R1PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT_eType2R2PS_null_r16Present bool
	if NCJT_eType2R2PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT_Type2_Type2PS_r16Present bool
	if NCJT_Type2_Type2PS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_Type2_null_r16Present bool
	if NCJT1SP_Type2_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_Type2PS_null_r16Present bool
	if NCJT1SP_Type2PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_eType2R1_null_r16Present bool
	if NCJT1SP_eType2R1_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_eType2R2_null_r16Present bool
	if NCJT1SP_eType2R2_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_eType2R1PS_null_r16Present bool
	if NCJT1SP_eType2R1PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_eType2R2PS_null_r16Present bool
	if NCJT1SP_eType2R2PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_Type2_Type2PS_r16Present bool
	if NCJT1SP_Type2_Type2PS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT_feType2PS_null_r17Present bool
	if NCJT_feType2PS_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT_feType2PS_M2R1_null_r17Present bool
	if NCJT_feType2PS_M2R1_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT_feType2PS_M2R2_null_r17Present bool
	if NCJT_feType2PS_M2R2_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT_Type2_feType2_PS_M1_r17Present bool
	if NCJT_Type2_feType2_PS_M1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT_Type2_feType2_PS_M2R1_r17Present bool
	if NCJT_Type2_feType2_PS_M2R1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT_eType2R1_feType2_PS_M1_r17Present bool
	if NCJT_eType2R1_feType2_PS_M1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT_eType2R1_feType2_PS_M2R1_r17Present bool
	if NCJT_eType2R1_feType2_PS_M2R1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_feType2PS_null_r17Present bool
	if NCJT1SP_feType2PS_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_feType2PS_M2R1_null_r17Present bool
	if NCJT1SP_feType2PS_M2R1_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_feType2PS_M2R2_null_r1Present bool
	if NCJT1SP_feType2PS_M2R2_null_r1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_Type2_feType2_PS_M1_r17Present bool
	if NCJT1SP_Type2_feType2_PS_M1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_Type2_feType2_PS_M2R1_r17Present bool
	if NCJT1SP_Type2_feType2_PS_M2R1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_eType2R1_feType2_PS_M1_r17Present bool
	if NCJT1SP_eType2R1_feType2_PS_M1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NCJT1SP_eType2R1_feType2_PS_M2R1_r17Present bool
	if NCJT1SP_eType2R1_feType2_PS_M2R1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if NCJT_null_nullPresent {
		tmp_NCJT_null_null := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_null_null := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_null_null.Decode(r, fn_NCJT_null_null); err != nil {
			return utils.WrapError("Decode NCJT_null_null", err)
		}
		ie.NCJT_null_null = []int64{}
		for _, i := range tmp_NCJT_null_null.Value {
			ie.NCJT_null_null = append(ie.NCJT_null_null, int64(i.Value))
		}
	}
	if NCJT1SP_null_nullPresent {
		tmp_NCJT1SP_null_null := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_null_null := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_null_null.Decode(r, fn_NCJT1SP_null_null); err != nil {
			return utils.WrapError("Decode NCJT1SP_null_null", err)
		}
		ie.NCJT1SP_null_null = []int64{}
		for _, i := range tmp_NCJT1SP_null_null.Value {
			ie.NCJT1SP_null_null = append(ie.NCJT1SP_null_null, int64(i.Value))
		}
	}
	if NCJT_Type2_null_r16Present {
		tmp_NCJT_Type2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_Type2_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_Type2_null_r16.Decode(r, fn_NCJT_Type2_null_r16); err != nil {
			return utils.WrapError("Decode NCJT_Type2_null_r16", err)
		}
		ie.NCJT_Type2_null_r16 = []int64{}
		for _, i := range tmp_NCJT_Type2_null_r16.Value {
			ie.NCJT_Type2_null_r16 = append(ie.NCJT_Type2_null_r16, int64(i.Value))
		}
	}
	if NCJT_Type2PS_null_r16Present {
		tmp_NCJT_Type2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_Type2PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_Type2PS_null_r16.Decode(r, fn_NCJT_Type2PS_null_r16); err != nil {
			return utils.WrapError("Decode NCJT_Type2PS_null_r16", err)
		}
		ie.NCJT_Type2PS_null_r16 = []int64{}
		for _, i := range tmp_NCJT_Type2PS_null_r16.Value {
			ie.NCJT_Type2PS_null_r16 = append(ie.NCJT_Type2PS_null_r16, int64(i.Value))
		}
	}
	if NCJT_eType2R1_null_r16Present {
		tmp_NCJT_eType2R1_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_eType2R1_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_eType2R1_null_r16.Decode(r, fn_NCJT_eType2R1_null_r16); err != nil {
			return utils.WrapError("Decode NCJT_eType2R1_null_r16", err)
		}
		ie.NCJT_eType2R1_null_r16 = []int64{}
		for _, i := range tmp_NCJT_eType2R1_null_r16.Value {
			ie.NCJT_eType2R1_null_r16 = append(ie.NCJT_eType2R1_null_r16, int64(i.Value))
		}
	}
	if NCJT_eType2R2_null_r16Present {
		tmp_NCJT_eType2R2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_eType2R2_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_eType2R2_null_r16.Decode(r, fn_NCJT_eType2R2_null_r16); err != nil {
			return utils.WrapError("Decode NCJT_eType2R2_null_r16", err)
		}
		ie.NCJT_eType2R2_null_r16 = []int64{}
		for _, i := range tmp_NCJT_eType2R2_null_r16.Value {
			ie.NCJT_eType2R2_null_r16 = append(ie.NCJT_eType2R2_null_r16, int64(i.Value))
		}
	}
	if NCJT_eType2R1PS_null_r16Present {
		tmp_NCJT_eType2R1PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_eType2R1PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_eType2R1PS_null_r16.Decode(r, fn_NCJT_eType2R1PS_null_r16); err != nil {
			return utils.WrapError("Decode NCJT_eType2R1PS_null_r16", err)
		}
		ie.NCJT_eType2R1PS_null_r16 = []int64{}
		for _, i := range tmp_NCJT_eType2R1PS_null_r16.Value {
			ie.NCJT_eType2R1PS_null_r16 = append(ie.NCJT_eType2R1PS_null_r16, int64(i.Value))
		}
	}
	if NCJT_eType2R2PS_null_r16Present {
		tmp_NCJT_eType2R2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_eType2R2PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_eType2R2PS_null_r16.Decode(r, fn_NCJT_eType2R2PS_null_r16); err != nil {
			return utils.WrapError("Decode NCJT_eType2R2PS_null_r16", err)
		}
		ie.NCJT_eType2R2PS_null_r16 = []int64{}
		for _, i := range tmp_NCJT_eType2R2PS_null_r16.Value {
			ie.NCJT_eType2R2PS_null_r16 = append(ie.NCJT_eType2R2PS_null_r16, int64(i.Value))
		}
	}
	if NCJT_Type2_Type2PS_r16Present {
		tmp_NCJT_Type2_Type2PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_Type2_Type2PS_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_Type2_Type2PS_r16.Decode(r, fn_NCJT_Type2_Type2PS_r16); err != nil {
			return utils.WrapError("Decode NCJT_Type2_Type2PS_r16", err)
		}
		ie.NCJT_Type2_Type2PS_r16 = []int64{}
		for _, i := range tmp_NCJT_Type2_Type2PS_r16.Value {
			ie.NCJT_Type2_Type2PS_r16 = append(ie.NCJT_Type2_Type2PS_r16, int64(i.Value))
		}
	}
	if NCJT1SP_Type2_null_r16Present {
		tmp_NCJT1SP_Type2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_Type2_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_Type2_null_r16.Decode(r, fn_NCJT1SP_Type2_null_r16); err != nil {
			return utils.WrapError("Decode NCJT1SP_Type2_null_r16", err)
		}
		ie.NCJT1SP_Type2_null_r16 = []int64{}
		for _, i := range tmp_NCJT1SP_Type2_null_r16.Value {
			ie.NCJT1SP_Type2_null_r16 = append(ie.NCJT1SP_Type2_null_r16, int64(i.Value))
		}
	}
	if NCJT1SP_Type2PS_null_r16Present {
		tmp_NCJT1SP_Type2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_Type2PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_Type2PS_null_r16.Decode(r, fn_NCJT1SP_Type2PS_null_r16); err != nil {
			return utils.WrapError("Decode NCJT1SP_Type2PS_null_r16", err)
		}
		ie.NCJT1SP_Type2PS_null_r16 = []int64{}
		for _, i := range tmp_NCJT1SP_Type2PS_null_r16.Value {
			ie.NCJT1SP_Type2PS_null_r16 = append(ie.NCJT1SP_Type2PS_null_r16, int64(i.Value))
		}
	}
	if NCJT1SP_eType2R1_null_r16Present {
		tmp_NCJT1SP_eType2R1_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_eType2R1_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_eType2R1_null_r16.Decode(r, fn_NCJT1SP_eType2R1_null_r16); err != nil {
			return utils.WrapError("Decode NCJT1SP_eType2R1_null_r16", err)
		}
		ie.NCJT1SP_eType2R1_null_r16 = []int64{}
		for _, i := range tmp_NCJT1SP_eType2R1_null_r16.Value {
			ie.NCJT1SP_eType2R1_null_r16 = append(ie.NCJT1SP_eType2R1_null_r16, int64(i.Value))
		}
	}
	if NCJT1SP_eType2R2_null_r16Present {
		tmp_NCJT1SP_eType2R2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_eType2R2_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_eType2R2_null_r16.Decode(r, fn_NCJT1SP_eType2R2_null_r16); err != nil {
			return utils.WrapError("Decode NCJT1SP_eType2R2_null_r16", err)
		}
		ie.NCJT1SP_eType2R2_null_r16 = []int64{}
		for _, i := range tmp_NCJT1SP_eType2R2_null_r16.Value {
			ie.NCJT1SP_eType2R2_null_r16 = append(ie.NCJT1SP_eType2R2_null_r16, int64(i.Value))
		}
	}
	if NCJT1SP_eType2R1PS_null_r16Present {
		tmp_NCJT1SP_eType2R1PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_eType2R1PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_eType2R1PS_null_r16.Decode(r, fn_NCJT1SP_eType2R1PS_null_r16); err != nil {
			return utils.WrapError("Decode NCJT1SP_eType2R1PS_null_r16", err)
		}
		ie.NCJT1SP_eType2R1PS_null_r16 = []int64{}
		for _, i := range tmp_NCJT1SP_eType2R1PS_null_r16.Value {
			ie.NCJT1SP_eType2R1PS_null_r16 = append(ie.NCJT1SP_eType2R1PS_null_r16, int64(i.Value))
		}
	}
	if NCJT1SP_eType2R2PS_null_r16Present {
		tmp_NCJT1SP_eType2R2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_eType2R2PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_eType2R2PS_null_r16.Decode(r, fn_NCJT1SP_eType2R2PS_null_r16); err != nil {
			return utils.WrapError("Decode NCJT1SP_eType2R2PS_null_r16", err)
		}
		ie.NCJT1SP_eType2R2PS_null_r16 = []int64{}
		for _, i := range tmp_NCJT1SP_eType2R2PS_null_r16.Value {
			ie.NCJT1SP_eType2R2PS_null_r16 = append(ie.NCJT1SP_eType2R2PS_null_r16, int64(i.Value))
		}
	}
	if NCJT1SP_Type2_Type2PS_r16Present {
		tmp_NCJT1SP_Type2_Type2PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_Type2_Type2PS_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_Type2_Type2PS_r16.Decode(r, fn_NCJT1SP_Type2_Type2PS_r16); err != nil {
			return utils.WrapError("Decode NCJT1SP_Type2_Type2PS_r16", err)
		}
		ie.NCJT1SP_Type2_Type2PS_r16 = []int64{}
		for _, i := range tmp_NCJT1SP_Type2_Type2PS_r16.Value {
			ie.NCJT1SP_Type2_Type2PS_r16 = append(ie.NCJT1SP_Type2_Type2PS_r16, int64(i.Value))
		}
	}
	if NCJT_feType2PS_null_r17Present {
		tmp_NCJT_feType2PS_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_feType2PS_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_feType2PS_null_r17.Decode(r, fn_NCJT_feType2PS_null_r17); err != nil {
			return utils.WrapError("Decode NCJT_feType2PS_null_r17", err)
		}
		ie.NCJT_feType2PS_null_r17 = []int64{}
		for _, i := range tmp_NCJT_feType2PS_null_r17.Value {
			ie.NCJT_feType2PS_null_r17 = append(ie.NCJT_feType2PS_null_r17, int64(i.Value))
		}
	}
	if NCJT_feType2PS_M2R1_null_r17Present {
		tmp_NCJT_feType2PS_M2R1_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_feType2PS_M2R1_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_feType2PS_M2R1_null_r17.Decode(r, fn_NCJT_feType2PS_M2R1_null_r17); err != nil {
			return utils.WrapError("Decode NCJT_feType2PS_M2R1_null_r17", err)
		}
		ie.NCJT_feType2PS_M2R1_null_r17 = []int64{}
		for _, i := range tmp_NCJT_feType2PS_M2R1_null_r17.Value {
			ie.NCJT_feType2PS_M2R1_null_r17 = append(ie.NCJT_feType2PS_M2R1_null_r17, int64(i.Value))
		}
	}
	if NCJT_feType2PS_M2R2_null_r17Present {
		tmp_NCJT_feType2PS_M2R2_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_feType2PS_M2R2_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_feType2PS_M2R2_null_r17.Decode(r, fn_NCJT_feType2PS_M2R2_null_r17); err != nil {
			return utils.WrapError("Decode NCJT_feType2PS_M2R2_null_r17", err)
		}
		ie.NCJT_feType2PS_M2R2_null_r17 = []int64{}
		for _, i := range tmp_NCJT_feType2PS_M2R2_null_r17.Value {
			ie.NCJT_feType2PS_M2R2_null_r17 = append(ie.NCJT_feType2PS_M2R2_null_r17, int64(i.Value))
		}
	}
	if NCJT_Type2_feType2_PS_M1_r17Present {
		tmp_NCJT_Type2_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_Type2_feType2_PS_M1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_Type2_feType2_PS_M1_r17.Decode(r, fn_NCJT_Type2_feType2_PS_M1_r17); err != nil {
			return utils.WrapError("Decode NCJT_Type2_feType2_PS_M1_r17", err)
		}
		ie.NCJT_Type2_feType2_PS_M1_r17 = []int64{}
		for _, i := range tmp_NCJT_Type2_feType2_PS_M1_r17.Value {
			ie.NCJT_Type2_feType2_PS_M1_r17 = append(ie.NCJT_Type2_feType2_PS_M1_r17, int64(i.Value))
		}
	}
	if NCJT_Type2_feType2_PS_M2R1_r17Present {
		tmp_NCJT_Type2_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_Type2_feType2_PS_M2R1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_Type2_feType2_PS_M2R1_r17.Decode(r, fn_NCJT_Type2_feType2_PS_M2R1_r17); err != nil {
			return utils.WrapError("Decode NCJT_Type2_feType2_PS_M2R1_r17", err)
		}
		ie.NCJT_Type2_feType2_PS_M2R1_r17 = []int64{}
		for _, i := range tmp_NCJT_Type2_feType2_PS_M2R1_r17.Value {
			ie.NCJT_Type2_feType2_PS_M2R1_r17 = append(ie.NCJT_Type2_feType2_PS_M2R1_r17, int64(i.Value))
		}
	}
	if NCJT_eType2R1_feType2_PS_M1_r17Present {
		tmp_NCJT_eType2R1_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_eType2R1_feType2_PS_M1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_eType2R1_feType2_PS_M1_r17.Decode(r, fn_NCJT_eType2R1_feType2_PS_M1_r17); err != nil {
			return utils.WrapError("Decode NCJT_eType2R1_feType2_PS_M1_r17", err)
		}
		ie.NCJT_eType2R1_feType2_PS_M1_r17 = []int64{}
		for _, i := range tmp_NCJT_eType2R1_feType2_PS_M1_r17.Value {
			ie.NCJT_eType2R1_feType2_PS_M1_r17 = append(ie.NCJT_eType2R1_feType2_PS_M1_r17, int64(i.Value))
		}
	}
	if NCJT_eType2R1_feType2_PS_M2R1_r17Present {
		tmp_NCJT_eType2R1_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT_eType2R1_feType2_PS_M2R1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT_eType2R1_feType2_PS_M2R1_r17.Decode(r, fn_NCJT_eType2R1_feType2_PS_M2R1_r17); err != nil {
			return utils.WrapError("Decode NCJT_eType2R1_feType2_PS_M2R1_r17", err)
		}
		ie.NCJT_eType2R1_feType2_PS_M2R1_r17 = []int64{}
		for _, i := range tmp_NCJT_eType2R1_feType2_PS_M2R1_r17.Value {
			ie.NCJT_eType2R1_feType2_PS_M2R1_r17 = append(ie.NCJT_eType2R1_feType2_PS_M2R1_r17, int64(i.Value))
		}
	}
	if NCJT1SP_feType2PS_null_r17Present {
		tmp_NCJT1SP_feType2PS_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_feType2PS_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_feType2PS_null_r17.Decode(r, fn_NCJT1SP_feType2PS_null_r17); err != nil {
			return utils.WrapError("Decode NCJT1SP_feType2PS_null_r17", err)
		}
		ie.NCJT1SP_feType2PS_null_r17 = []int64{}
		for _, i := range tmp_NCJT1SP_feType2PS_null_r17.Value {
			ie.NCJT1SP_feType2PS_null_r17 = append(ie.NCJT1SP_feType2PS_null_r17, int64(i.Value))
		}
	}
	if NCJT1SP_feType2PS_M2R1_null_r17Present {
		tmp_NCJT1SP_feType2PS_M2R1_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_feType2PS_M2R1_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_feType2PS_M2R1_null_r17.Decode(r, fn_NCJT1SP_feType2PS_M2R1_null_r17); err != nil {
			return utils.WrapError("Decode NCJT1SP_feType2PS_M2R1_null_r17", err)
		}
		ie.NCJT1SP_feType2PS_M2R1_null_r17 = []int64{}
		for _, i := range tmp_NCJT1SP_feType2PS_M2R1_null_r17.Value {
			ie.NCJT1SP_feType2PS_M2R1_null_r17 = append(ie.NCJT1SP_feType2PS_M2R1_null_r17, int64(i.Value))
		}
	}
	if NCJT1SP_feType2PS_M2R2_null_r1Present {
		tmp_NCJT1SP_feType2PS_M2R2_null_r1 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_feType2PS_M2R2_null_r1 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_feType2PS_M2R2_null_r1.Decode(r, fn_NCJT1SP_feType2PS_M2R2_null_r1); err != nil {
			return utils.WrapError("Decode NCJT1SP_feType2PS_M2R2_null_r1", err)
		}
		ie.NCJT1SP_feType2PS_M2R2_null_r1 = []int64{}
		for _, i := range tmp_NCJT1SP_feType2PS_M2R2_null_r1.Value {
			ie.NCJT1SP_feType2PS_M2R2_null_r1 = append(ie.NCJT1SP_feType2PS_M2R2_null_r1, int64(i.Value))
		}
	}
	if NCJT1SP_Type2_feType2_PS_M1_r17Present {
		tmp_NCJT1SP_Type2_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_Type2_feType2_PS_M1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_Type2_feType2_PS_M1_r17.Decode(r, fn_NCJT1SP_Type2_feType2_PS_M1_r17); err != nil {
			return utils.WrapError("Decode NCJT1SP_Type2_feType2_PS_M1_r17", err)
		}
		ie.NCJT1SP_Type2_feType2_PS_M1_r17 = []int64{}
		for _, i := range tmp_NCJT1SP_Type2_feType2_PS_M1_r17.Value {
			ie.NCJT1SP_Type2_feType2_PS_M1_r17 = append(ie.NCJT1SP_Type2_feType2_PS_M1_r17, int64(i.Value))
		}
	}
	if NCJT1SP_Type2_feType2_PS_M2R1_r17Present {
		tmp_NCJT1SP_Type2_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_Type2_feType2_PS_M2R1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_Type2_feType2_PS_M2R1_r17.Decode(r, fn_NCJT1SP_Type2_feType2_PS_M2R1_r17); err != nil {
			return utils.WrapError("Decode NCJT1SP_Type2_feType2_PS_M2R1_r17", err)
		}
		ie.NCJT1SP_Type2_feType2_PS_M2R1_r17 = []int64{}
		for _, i := range tmp_NCJT1SP_Type2_feType2_PS_M2R1_r17.Value {
			ie.NCJT1SP_Type2_feType2_PS_M2R1_r17 = append(ie.NCJT1SP_Type2_feType2_PS_M2R1_r17, int64(i.Value))
		}
	}
	if NCJT1SP_eType2R1_feType2_PS_M1_r17Present {
		tmp_NCJT1SP_eType2R1_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_eType2R1_feType2_PS_M1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_eType2R1_feType2_PS_M1_r17.Decode(r, fn_NCJT1SP_eType2R1_feType2_PS_M1_r17); err != nil {
			return utils.WrapError("Decode NCJT1SP_eType2R1_feType2_PS_M1_r17", err)
		}
		ie.NCJT1SP_eType2R1_feType2_PS_M1_r17 = []int64{}
		for _, i := range tmp_NCJT1SP_eType2R1_feType2_PS_M1_r17.Value {
			ie.NCJT1SP_eType2R1_feType2_PS_M1_r17 = append(ie.NCJT1SP_eType2R1_feType2_PS_M1_r17, int64(i.Value))
		}
	}
	if NCJT1SP_eType2R1_feType2_PS_M2R1_r17Present {
		tmp_NCJT1SP_eType2R1_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_NCJT1SP_eType2R1_feType2_PS_M2R1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_NCJT1SP_eType2R1_feType2_PS_M2R1_r17.Decode(r, fn_NCJT1SP_eType2R1_feType2_PS_M2R1_r17); err != nil {
			return utils.WrapError("Decode NCJT1SP_eType2R1_feType2_PS_M2R1_r17", err)
		}
		ie.NCJT1SP_eType2R1_feType2_PS_M2R1_r17 = []int64{}
		for _, i := range tmp_NCJT1SP_eType2R1_feType2_PS_M2R1_r17.Value {
			ie.NCJT1SP_eType2R1_feType2_PS_M2R1_r17 = append(ie.NCJT1SP_eType2R1_feType2_PS_M2R1_r17, int64(i.Value))
		}
	}
	return nil
}
