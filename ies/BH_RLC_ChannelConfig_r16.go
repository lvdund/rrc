package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BH_RLC_ChannelConfig_r16 struct {
	Bh_LogicalChannelIdentity_r16 *BH_LogicalChannelIdentity_r16               `optional`
	Bh_RLC_ChannelID_r16          BH_RLC_ChannelID_r16                         `madatory`
	ReestablishRLC_r16            *BH_RLC_ChannelConfig_r16_reestablishRLC_r16 `optional`
	Rlc_Config_r16                *RLC_Config                                  `optional`
	Mac_LogicalChannelConfig_r16  *LogicalChannelConfig                        `optional`
}

func (ie *BH_RLC_ChannelConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Bh_LogicalChannelIdentity_r16 != nil, ie.ReestablishRLC_r16 != nil, ie.Rlc_Config_r16 != nil, ie.Mac_LogicalChannelConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Bh_LogicalChannelIdentity_r16 != nil {
		if err = ie.Bh_LogicalChannelIdentity_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Bh_LogicalChannelIdentity_r16", err)
		}
	}
	if err = ie.Bh_RLC_ChannelID_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Bh_RLC_ChannelID_r16", err)
	}
	if ie.ReestablishRLC_r16 != nil {
		if err = ie.ReestablishRLC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ReestablishRLC_r16", err)
		}
	}
	if ie.Rlc_Config_r16 != nil {
		if err = ie.Rlc_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rlc_Config_r16", err)
		}
	}
	if ie.Mac_LogicalChannelConfig_r16 != nil {
		if err = ie.Mac_LogicalChannelConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_LogicalChannelConfig_r16", err)
		}
	}
	return nil
}

func (ie *BH_RLC_ChannelConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var Bh_LogicalChannelIdentity_r16Present bool
	if Bh_LogicalChannelIdentity_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReestablishRLC_r16Present bool
	if ReestablishRLC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rlc_Config_r16Present bool
	if Rlc_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mac_LogicalChannelConfig_r16Present bool
	if Mac_LogicalChannelConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Bh_LogicalChannelIdentity_r16Present {
		ie.Bh_LogicalChannelIdentity_r16 = new(BH_LogicalChannelIdentity_r16)
		if err = ie.Bh_LogicalChannelIdentity_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Bh_LogicalChannelIdentity_r16", err)
		}
	}
	if err = ie.Bh_RLC_ChannelID_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Bh_RLC_ChannelID_r16", err)
	}
	if ReestablishRLC_r16Present {
		ie.ReestablishRLC_r16 = new(BH_RLC_ChannelConfig_r16_reestablishRLC_r16)
		if err = ie.ReestablishRLC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReestablishRLC_r16", err)
		}
	}
	if Rlc_Config_r16Present {
		ie.Rlc_Config_r16 = new(RLC_Config)
		if err = ie.Rlc_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rlc_Config_r16", err)
		}
	}
	if Mac_LogicalChannelConfig_r16Present {
		ie.Mac_LogicalChannelConfig_r16 = new(LogicalChannelConfig)
		if err = ie.Mac_LogicalChannelConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_LogicalChannelConfig_r16", err)
		}
	}
	return nil
}
