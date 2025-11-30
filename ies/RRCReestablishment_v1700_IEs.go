package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCReestablishment_v1700_IEs struct {
	Sl_L2RemoteUE_Config_r17 *SL_L2RemoteUE_Config_r17 `optional,setuprelease`
	NonCriticalExtension     interface{}               `optional`
}

func (ie *RRCReestablishment_v1700_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_L2RemoteUE_Config_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_L2RemoteUE_Config_r17 != nil {
		tmp_Sl_L2RemoteUE_Config_r17 := utils.SetupRelease[*SL_L2RemoteUE_Config_r17]{
			Setup: ie.Sl_L2RemoteUE_Config_r17,
		}
		if err = tmp_Sl_L2RemoteUE_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_L2RemoteUE_Config_r17", err)
		}
	}
	return nil
}

func (ie *RRCReestablishment_v1700_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Sl_L2RemoteUE_Config_r17Present bool
	if Sl_L2RemoteUE_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_L2RemoteUE_Config_r17Present {
		tmp_Sl_L2RemoteUE_Config_r17 := utils.SetupRelease[*SL_L2RemoteUE_Config_r17]{}
		if err = tmp_Sl_L2RemoteUE_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_L2RemoteUE_Config_r17", err)
		}
		ie.Sl_L2RemoteUE_Config_r17 = tmp_Sl_L2RemoteUE_Config_r17.Setup
	}
	return nil
}
