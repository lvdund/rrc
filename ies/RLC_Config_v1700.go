package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLC_Config_v1700 struct {
	Dl_AM_RLC_v1700 DL_AM_RLC_v1700 `madatory`
	Dl_UM_RLC_v1700 DL_UM_RLC_v1700 `madatory`
}

func (ie *RLC_Config_v1700) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Dl_AM_RLC_v1700.Encode(w); err != nil {
		return utils.WrapError("Encode Dl_AM_RLC_v1700", err)
	}
	if err = ie.Dl_UM_RLC_v1700.Encode(w); err != nil {
		return utils.WrapError("Encode Dl_UM_RLC_v1700", err)
	}
	return nil
}

func (ie *RLC_Config_v1700) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Dl_AM_RLC_v1700.Decode(r); err != nil {
		return utils.WrapError("Decode Dl_AM_RLC_v1700", err)
	}
	if err = ie.Dl_UM_RLC_v1700.Decode(r); err != nil {
		return utils.WrapError("Decode Dl_UM_RLC_v1700", err)
	}
	return nil
}
