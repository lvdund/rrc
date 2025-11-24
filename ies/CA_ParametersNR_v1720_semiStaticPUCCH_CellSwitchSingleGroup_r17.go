package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1720_semiStaticPUCCH_CellSwitchSingleGroup_r17 struct {
	Pucch_Group_r17        CA_ParametersNR_v1720_semiStaticPUCCH_CellSwitchSingleGroup_r17_pucch_Group_r17 `madatory`
	Pucch_Group_Config_r17 PUCCH_Group_Config_r17                                                          `madatory`
}

func (ie *CA_ParametersNR_v1720_semiStaticPUCCH_CellSwitchSingleGroup_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Pucch_Group_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_Group_r17", err)
	}
	if err = ie.Pucch_Group_Config_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_Group_Config_r17", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v1720_semiStaticPUCCH_CellSwitchSingleGroup_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Pucch_Group_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_Group_r17", err)
	}
	if err = ie.Pucch_Group_Config_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_Group_Config_r17", err)
	}
	return nil
}
