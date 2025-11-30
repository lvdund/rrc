package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_L2RemoteUE_Config_r17 struct {
	Sl_SRAP_ConfigRemote_r17 *SL_SRAP_Config_r17 `optional`
	Sl_UEIdentityRemote_r17  *RNTI_Value         `optional`
}

func (ie *SL_L2RemoteUE_Config_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_SRAP_ConfigRemote_r17 != nil, ie.Sl_UEIdentityRemote_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_SRAP_ConfigRemote_r17 != nil {
		if err = ie.Sl_SRAP_ConfigRemote_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SRAP_ConfigRemote_r17", err)
		}
	}
	if ie.Sl_UEIdentityRemote_r17 != nil {
		if err = ie.Sl_UEIdentityRemote_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_UEIdentityRemote_r17", err)
		}
	}
	return nil
}

func (ie *SL_L2RemoteUE_Config_r17) Decode(r *aper.AperReader) error {
	var err error
	var Sl_SRAP_ConfigRemote_r17Present bool
	if Sl_SRAP_ConfigRemote_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_UEIdentityRemote_r17Present bool
	if Sl_UEIdentityRemote_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_SRAP_ConfigRemote_r17Present {
		ie.Sl_SRAP_ConfigRemote_r17 = new(SL_SRAP_Config_r17)
		if err = ie.Sl_SRAP_ConfigRemote_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SRAP_ConfigRemote_r17", err)
		}
	}
	if Sl_UEIdentityRemote_r17Present {
		ie.Sl_UEIdentityRemote_r17 = new(RNTI_Value)
		if err = ie.Sl_UEIdentityRemote_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_UEIdentityRemote_r17", err)
		}
	}
	return nil
}
