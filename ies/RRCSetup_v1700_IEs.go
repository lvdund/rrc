package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetup_v1700_IEs struct {
	Sl_ConfigDedicatedNR_r17 *SL_ConfigDedicatedNR_r16 `optional`
	Sl_L2RemoteUE_Config_r17 *SL_L2RemoteUE_Config_r17 `optional`
	NonCriticalExtension     interface{}               `optional`
}

func (ie *RRCSetup_v1700_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_ConfigDedicatedNR_r17 != nil, ie.Sl_L2RemoteUE_Config_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_ConfigDedicatedNR_r17 != nil {
		if err = ie.Sl_ConfigDedicatedNR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ConfigDedicatedNR_r17", err)
		}
	}
	if ie.Sl_L2RemoteUE_Config_r17 != nil {
		if err = ie.Sl_L2RemoteUE_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_L2RemoteUE_Config_r17", err)
		}
	}
	return nil
}

func (ie *RRCSetup_v1700_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Sl_ConfigDedicatedNR_r17Present bool
	if Sl_ConfigDedicatedNR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_L2RemoteUE_Config_r17Present bool
	if Sl_L2RemoteUE_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_ConfigDedicatedNR_r17Present {
		ie.Sl_ConfigDedicatedNR_r17 = new(SL_ConfigDedicatedNR_r16)
		if err = ie.Sl_ConfigDedicatedNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ConfigDedicatedNR_r17", err)
		}
	}
	if Sl_L2RemoteUE_Config_r17Present {
		ie.Sl_L2RemoteUE_Config_r17 = new(SL_L2RemoteUE_Config_r17)
		if err = ie.Sl_L2RemoteUE_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_L2RemoteUE_Config_r17", err)
		}
	}
	return nil
}
