package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17_codebookModeNCJT_r17_Enum_mode1     aper.Enumerated = 0
	MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17_codebookModeNCJT_r17_Enum_mode1And2 aper.Enumerated = 1
)

type MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17_codebookModeNCJT_r17 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17_codebookModeNCJT_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17_codebookModeNCJT_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17_codebookModeNCJT_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17_codebookModeNCJT_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
