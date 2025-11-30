package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RLC_Config_um_Uni_Directional_UL struct {
	Ul_UM_RLC UL_UM_RLC `madatory`
}

func (ie *RLC_Config_um_Uni_Directional_UL) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Ul_UM_RLC.Encode(w); err != nil {
		return utils.WrapError("Encode Ul_UM_RLC", err)
	}
	return nil
}

func (ie *RLC_Config_um_Uni_Directional_UL) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Ul_UM_RLC.Decode(r); err != nil {
		return utils.WrapError("Decode Ul_UM_RLC", err)
	}
	return nil
}
