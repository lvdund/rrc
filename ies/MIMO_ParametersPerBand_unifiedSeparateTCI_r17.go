package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedSeparateTCI_r17 struct {
	MaxConfiguredDL_TCI_r17        MIMO_ParametersPerBand_unifiedSeparateTCI_r17_maxConfiguredDL_TCI_r17        `madatory`
	MaxConfiguredUL_TCI_r17        MIMO_ParametersPerBand_unifiedSeparateTCI_r17_maxConfiguredUL_TCI_r17        `madatory`
	MaxActivatedDL_TCIAcrossCC_r17 MIMO_ParametersPerBand_unifiedSeparateTCI_r17_maxActivatedDL_TCIAcrossCC_r17 `madatory`
	MaxActivatedUL_TCIAcrossCC_r17 MIMO_ParametersPerBand_unifiedSeparateTCI_r17_maxActivatedUL_TCIAcrossCC_r17 `madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MaxConfiguredDL_TCI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxConfiguredDL_TCI_r17", err)
	}
	if err = ie.MaxConfiguredUL_TCI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxConfiguredUL_TCI_r17", err)
	}
	if err = ie.MaxActivatedDL_TCIAcrossCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxActivatedDL_TCIAcrossCC_r17", err)
	}
	if err = ie.MaxActivatedUL_TCIAcrossCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxActivatedUL_TCIAcrossCC_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MaxConfiguredDL_TCI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxConfiguredDL_TCI_r17", err)
	}
	if err = ie.MaxConfiguredUL_TCI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxConfiguredUL_TCI_r17", err)
	}
	if err = ie.MaxActivatedDL_TCIAcrossCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxActivatedDL_TCIAcrossCC_r17", err)
	}
	if err = ie.MaxActivatedUL_TCIAcrossCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxActivatedUL_TCIAcrossCC_r17", err)
	}
	return nil
}
