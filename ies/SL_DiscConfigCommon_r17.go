package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_DiscConfigCommon_r17 struct {
	Sl_RelayUE_ConfigCommon_r17  SL_RelayUE_Config_r17  `madatory`
	Sl_RemoteUE_ConfigCommon_r17 SL_RemoteUE_Config_r17 `madatory`
}

func (ie *SL_DiscConfigCommon_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Sl_RelayUE_ConfigCommon_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_RelayUE_ConfigCommon_r17", err)
	}
	if err = ie.Sl_RemoteUE_ConfigCommon_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_RemoteUE_ConfigCommon_r17", err)
	}
	return nil
}

func (ie *SL_DiscConfigCommon_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Sl_RelayUE_ConfigCommon_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_RelayUE_ConfigCommon_r17", err)
	}
	if err = ie.Sl_RemoteUE_ConfigCommon_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_RemoteUE_ConfigCommon_r17", err)
	}
	return nil
}
