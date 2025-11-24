package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TwoPUCCH_Grp_Configurations_r17 struct {
	PrimaryPUCCH_GroupConfig_r17   PUCCH_Group_Config_r17 `madatory`
	SecondaryPUCCH_GroupConfig_r17 PUCCH_Group_Config_r17 `madatory`
}

func (ie *TwoPUCCH_Grp_Configurations_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.PrimaryPUCCH_GroupConfig_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PrimaryPUCCH_GroupConfig_r17", err)
	}
	if err = ie.SecondaryPUCCH_GroupConfig_r17.Encode(w); err != nil {
		return utils.WrapError("Encode SecondaryPUCCH_GroupConfig_r17", err)
	}
	return nil
}

func (ie *TwoPUCCH_Grp_Configurations_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.PrimaryPUCCH_GroupConfig_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PrimaryPUCCH_GroupConfig_r17", err)
	}
	if err = ie.SecondaryPUCCH_GroupConfig_r17.Decode(r); err != nil {
		return utils.WrapError("Decode SecondaryPUCCH_GroupConfig_r17", err)
	}
	return nil
}
