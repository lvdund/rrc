package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17 struct {
	MaxNumNZP_CSI_RS_r17        int64                                                                       `lb:2,ub:8,madatory`
	CSI_Report_mode_r17         MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17_cSI_Report_mode_r17  `madatory`
	SupportedComboAcrossCCs_r17 []CSI_MultiTRP_SupportedCombinations_r17                                    `lb:1,ub:16,madatory`
	CodebookModeNCJT_r17        MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17_codebookModeNCJT_r17 `madatory`
}

func (ie *MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.MaxNumNZP_CSI_RS_r17, &uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumNZP_CSI_RS_r17", err)
	}
	if err = ie.CSI_Report_mode_r17.Encode(w); err != nil {
		return utils.WrapError("Encode CSI_Report_mode_r17", err)
	}
	tmp_SupportedComboAcrossCCs_r17 := utils.NewSequence[*CSI_MultiTRP_SupportedCombinations_r17]([]*CSI_MultiTRP_SupportedCombinations_r17{}, uper.Constraint{Lb: 1, Ub: 16}, false)
	for _, i := range ie.SupportedComboAcrossCCs_r17 {
		tmp_SupportedComboAcrossCCs_r17.Value = append(tmp_SupportedComboAcrossCCs_r17.Value, &i)
	}
	if err = tmp_SupportedComboAcrossCCs_r17.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedComboAcrossCCs_r17", err)
	}
	if err = ie.CodebookModeNCJT_r17.Encode(w); err != nil {
		return utils.WrapError("Encode CodebookModeNCJT_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_MaxNumNZP_CSI_RS_r17 int64
	if tmp_int_MaxNumNZP_CSI_RS_r17, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumNZP_CSI_RS_r17", err)
	}
	ie.MaxNumNZP_CSI_RS_r17 = tmp_int_MaxNumNZP_CSI_RS_r17
	if err = ie.CSI_Report_mode_r17.Decode(r); err != nil {
		return utils.WrapError("Decode CSI_Report_mode_r17", err)
	}
	tmp_SupportedComboAcrossCCs_r17 := utils.NewSequence[*CSI_MultiTRP_SupportedCombinations_r17]([]*CSI_MultiTRP_SupportedCombinations_r17{}, uper.Constraint{Lb: 1, Ub: 16}, false)
	fn_SupportedComboAcrossCCs_r17 := func() *CSI_MultiTRP_SupportedCombinations_r17 {
		return new(CSI_MultiTRP_SupportedCombinations_r17)
	}
	if err = tmp_SupportedComboAcrossCCs_r17.Decode(r, fn_SupportedComboAcrossCCs_r17); err != nil {
		return utils.WrapError("Decode SupportedComboAcrossCCs_r17", err)
	}
	ie.SupportedComboAcrossCCs_r17 = []CSI_MultiTRP_SupportedCombinations_r17{}
	for _, i := range tmp_SupportedComboAcrossCCs_r17.Value {
		ie.SupportedComboAcrossCCs_r17 = append(ie.SupportedComboAcrossCCs_r17, *i)
	}
	if err = ie.CodebookModeNCJT_r17.Decode(r); err != nil {
		return utils.WrapError("Decode CodebookModeNCJT_r17", err)
	}
	return nil
}
