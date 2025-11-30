package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_DiscConfig_r17 struct {
	Sl_RelayUE_Config_r17  *SL_RelayUE_Config_r17  `optional,setuprelease`
	Sl_RemoteUE_Config_r17 *SL_RemoteUE_Config_r17 `optional,setuprelease`
}

func (ie *SL_DiscConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_RelayUE_Config_r17 != nil, ie.Sl_RemoteUE_Config_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_RelayUE_Config_r17 != nil {
		tmp_Sl_RelayUE_Config_r17 := utils.SetupRelease[*SL_RelayUE_Config_r17]{
			Setup: ie.Sl_RelayUE_Config_r17,
		}
		if err = tmp_Sl_RelayUE_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RelayUE_Config_r17", err)
		}
	}
	if ie.Sl_RemoteUE_Config_r17 != nil {
		tmp_Sl_RemoteUE_Config_r17 := utils.SetupRelease[*SL_RemoteUE_Config_r17]{
			Setup: ie.Sl_RemoteUE_Config_r17,
		}
		if err = tmp_Sl_RemoteUE_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RemoteUE_Config_r17", err)
		}
	}
	return nil
}

func (ie *SL_DiscConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	var Sl_RelayUE_Config_r17Present bool
	if Sl_RelayUE_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RemoteUE_Config_r17Present bool
	if Sl_RemoteUE_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_RelayUE_Config_r17Present {
		tmp_Sl_RelayUE_Config_r17 := utils.SetupRelease[*SL_RelayUE_Config_r17]{}
		if err = tmp_Sl_RelayUE_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_RelayUE_Config_r17", err)
		}
		ie.Sl_RelayUE_Config_r17 = tmp_Sl_RelayUE_Config_r17.Setup
	}
	if Sl_RemoteUE_Config_r17Present {
		tmp_Sl_RemoteUE_Config_r17 := utils.SetupRelease[*SL_RemoteUE_Config_r17]{}
		if err = tmp_Sl_RemoteUE_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_RemoteUE_Config_r17", err)
		}
		ie.Sl_RemoteUE_Config_r17 = tmp_Sl_RemoteUE_Config_r17.Setup
	}
	return nil
}
