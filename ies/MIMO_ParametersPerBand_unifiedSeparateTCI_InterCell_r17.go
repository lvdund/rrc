package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17 struct {
	K_DL_PerCC_r17    MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17_k_DL_PerCC_r17    `madatory`
	K_UL_PerCC_r17    MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17_k_UL_PerCC_r17    `madatory`
	K_DL_AcrossCC_r17 MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17_k_DL_AcrossCC_r17 `madatory`
	K_UL_AcrossCC_r17 MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17_k_UL_AcrossCC_r17 `madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.K_DL_PerCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode K_DL_PerCC_r17", err)
	}
	if err = ie.K_UL_PerCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode K_UL_PerCC_r17", err)
	}
	if err = ie.K_DL_AcrossCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode K_DL_AcrossCC_r17", err)
	}
	if err = ie.K_UL_AcrossCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode K_UL_AcrossCC_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.K_DL_PerCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode K_DL_PerCC_r17", err)
	}
	if err = ie.K_UL_PerCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode K_UL_PerCC_r17", err)
	}
	if err = ie.K_DL_AcrossCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode K_DL_AcrossCC_r17", err)
	}
	if err = ie.K_UL_AcrossCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode K_UL_AcrossCC_r17", err)
	}
	return nil
}
