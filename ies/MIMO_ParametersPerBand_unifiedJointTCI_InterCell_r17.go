package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedJointTCI_InterCell_r17 struct {
	AdditionalMAC_CE_PerCC_r17    MIMO_ParametersPerBand_unifiedJointTCI_InterCell_r17_additionalMAC_CE_PerCC_r17    `madatory`
	AdditionalMAC_CE_AcrossCC_r17 MIMO_ParametersPerBand_unifiedJointTCI_InterCell_r17_additionalMAC_CE_AcrossCC_r17 `madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_InterCell_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.AdditionalMAC_CE_PerCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode AdditionalMAC_CE_PerCC_r17", err)
	}
	if err = ie.AdditionalMAC_CE_AcrossCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode AdditionalMAC_CE_AcrossCC_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_InterCell_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.AdditionalMAC_CE_PerCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode AdditionalMAC_CE_PerCC_r17", err)
	}
	if err = ie.AdditionalMAC_CE_AcrossCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode AdditionalMAC_CE_AcrossCC_r17", err)
	}
	return nil
}
