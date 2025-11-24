package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedJointTCI_r17 struct {
	MaxConfiguredJointTCI_r17   MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17   `madatory`
	MaxActivatedTCIAcrossCC_r17 MIMO_ParametersPerBand_unifiedJointTCI_r17_maxActivatedTCIAcrossCC_r17 `madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MaxConfiguredJointTCI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxConfiguredJointTCI_r17", err)
	}
	if err = ie.MaxActivatedTCIAcrossCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxActivatedTCIAcrossCC_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MaxConfiguredJointTCI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxConfiguredJointTCI_r17", err)
	}
	if err = ie.MaxActivatedTCIAcrossCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxActivatedTCIAcrossCC_r17", err)
	}
	return nil
}
