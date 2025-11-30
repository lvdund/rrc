package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Uu_RelayRLC_ChannelConfig_r17 struct {
	Uu_LogicalChannelIdentity_r17 *LogicalChannelIdentity                           `optional`
	Uu_RelayRLC_ChannelID_r17     Uu_RelayRLC_ChannelID_r17                         `madatory`
	ReestablishRLC_r17            *Uu_RelayRLC_ChannelConfig_r17_reestablishRLC_r17 `optional`
	Rlc_Config_r17                *RLC_Config                                       `optional`
	Mac_LogicalChannelConfig_r17  *LogicalChannelConfig                             `optional`
}

func (ie *Uu_RelayRLC_ChannelConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Uu_LogicalChannelIdentity_r17 != nil, ie.ReestablishRLC_r17 != nil, ie.Rlc_Config_r17 != nil, ie.Mac_LogicalChannelConfig_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Uu_LogicalChannelIdentity_r17 != nil {
		if err = ie.Uu_LogicalChannelIdentity_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Uu_LogicalChannelIdentity_r17", err)
		}
	}
	if err = ie.Uu_RelayRLC_ChannelID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Uu_RelayRLC_ChannelID_r17", err)
	}
	if ie.ReestablishRLC_r17 != nil {
		if err = ie.ReestablishRLC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ReestablishRLC_r17", err)
		}
	}
	if ie.Rlc_Config_r17 != nil {
		if err = ie.Rlc_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Rlc_Config_r17", err)
		}
	}
	if ie.Mac_LogicalChannelConfig_r17 != nil {
		if err = ie.Mac_LogicalChannelConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_LogicalChannelConfig_r17", err)
		}
	}
	return nil
}

func (ie *Uu_RelayRLC_ChannelConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	var Uu_LogicalChannelIdentity_r17Present bool
	if Uu_LogicalChannelIdentity_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReestablishRLC_r17Present bool
	if ReestablishRLC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rlc_Config_r17Present bool
	if Rlc_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mac_LogicalChannelConfig_r17Present bool
	if Mac_LogicalChannelConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Uu_LogicalChannelIdentity_r17Present {
		ie.Uu_LogicalChannelIdentity_r17 = new(LogicalChannelIdentity)
		if err = ie.Uu_LogicalChannelIdentity_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Uu_LogicalChannelIdentity_r17", err)
		}
	}
	if err = ie.Uu_RelayRLC_ChannelID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Uu_RelayRLC_ChannelID_r17", err)
	}
	if ReestablishRLC_r17Present {
		ie.ReestablishRLC_r17 = new(Uu_RelayRLC_ChannelConfig_r17_reestablishRLC_r17)
		if err = ie.ReestablishRLC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ReestablishRLC_r17", err)
		}
	}
	if Rlc_Config_r17Present {
		ie.Rlc_Config_r17 = new(RLC_Config)
		if err = ie.Rlc_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Rlc_Config_r17", err)
		}
	}
	if Mac_LogicalChannelConfig_r17Present {
		ie.Mac_LogicalChannelConfig_r17 = new(LogicalChannelConfig)
		if err = ie.Mac_LogicalChannelConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_LogicalChannelConfig_r17", err)
		}
	}
	return nil
}
