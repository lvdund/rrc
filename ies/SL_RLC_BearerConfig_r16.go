package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RLC_BearerConfig_r16 struct {
	Sl_RLC_BearerConfigIndex_r16    SL_RLC_BearerConfigIndex_r16 `madatory`
	Sl_ServedRadioBearer_r16        *SLRB_Uu_ConfigIndex_r16     `optional`
	Sl_RLC_Config_r16               *SL_RLC_Config_r16           `optional`
	Sl_MAC_LogicalChannelConfig_r16 *SL_LogicalChannelConfig_r16 `optional`
}

func (ie *SL_RLC_BearerConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_ServedRadioBearer_r16 != nil, ie.Sl_RLC_Config_r16 != nil, ie.Sl_MAC_LogicalChannelConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sl_RLC_BearerConfigIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_RLC_BearerConfigIndex_r16", err)
	}
	if ie.Sl_ServedRadioBearer_r16 != nil {
		if err = ie.Sl_ServedRadioBearer_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ServedRadioBearer_r16", err)
		}
	}
	if ie.Sl_RLC_Config_r16 != nil {
		if err = ie.Sl_RLC_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RLC_Config_r16", err)
		}
	}
	if ie.Sl_MAC_LogicalChannelConfig_r16 != nil {
		if err = ie.Sl_MAC_LogicalChannelConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MAC_LogicalChannelConfig_r16", err)
		}
	}
	return nil
}

func (ie *SL_RLC_BearerConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_ServedRadioBearer_r16Present bool
	if Sl_ServedRadioBearer_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RLC_Config_r16Present bool
	if Sl_RLC_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MAC_LogicalChannelConfig_r16Present bool
	if Sl_MAC_LogicalChannelConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sl_RLC_BearerConfigIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_RLC_BearerConfigIndex_r16", err)
	}
	if Sl_ServedRadioBearer_r16Present {
		ie.Sl_ServedRadioBearer_r16 = new(SLRB_Uu_ConfigIndex_r16)
		if err = ie.Sl_ServedRadioBearer_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ServedRadioBearer_r16", err)
		}
	}
	if Sl_RLC_Config_r16Present {
		ie.Sl_RLC_Config_r16 = new(SL_RLC_Config_r16)
		if err = ie.Sl_RLC_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_RLC_Config_r16", err)
		}
	}
	if Sl_MAC_LogicalChannelConfig_r16Present {
		ie.Sl_MAC_LogicalChannelConfig_r16 = new(SL_LogicalChannelConfig_r16)
		if err = ie.Sl_MAC_LogicalChannelConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MAC_LogicalChannelConfig_r16", err)
		}
	}
	return nil
}
