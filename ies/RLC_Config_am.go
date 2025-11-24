package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLC_Config_am struct {
	Ul_AM_RLC UL_AM_RLC `madatory`
	Dl_AM_RLC DL_AM_RLC `madatory`
}

func (ie *RLC_Config_am) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Ul_AM_RLC.Encode(w); err != nil {
		return utils.WrapError("Encode Ul_AM_RLC", err)
	}
	if err = ie.Dl_AM_RLC.Encode(w); err != nil {
		return utils.WrapError("Encode Dl_AM_RLC", err)
	}
	return nil
}

func (ie *RLC_Config_am) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Ul_AM_RLC.Decode(r); err != nil {
		return utils.WrapError("Decode Ul_AM_RLC", err)
	}
	if err = ie.Dl_AM_RLC.Decode(r); err != nil {
		return utils.WrapError("Decode Dl_AM_RLC", err)
	}
	return nil
}
