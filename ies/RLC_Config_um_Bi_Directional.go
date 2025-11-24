package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLC_Config_um_Bi_Directional struct {
	Ul_UM_RLC UL_UM_RLC `madatory`
	Dl_UM_RLC DL_UM_RLC `madatory`
}

func (ie *RLC_Config_um_Bi_Directional) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Ul_UM_RLC.Encode(w); err != nil {
		return utils.WrapError("Encode Ul_UM_RLC", err)
	}
	if err = ie.Dl_UM_RLC.Encode(w); err != nil {
		return utils.WrapError("Encode Dl_UM_RLC", err)
	}
	return nil
}

func (ie *RLC_Config_um_Bi_Directional) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Ul_UM_RLC.Decode(r); err != nil {
		return utils.WrapError("Decode Ul_UM_RLC", err)
	}
	if err = ie.Dl_UM_RLC.Decode(r); err != nil {
		return utils.WrapError("Decode Dl_UM_RLC", err)
	}
	return nil
}
