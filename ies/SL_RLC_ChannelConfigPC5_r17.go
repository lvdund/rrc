package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RLC_ChannelConfigPC5_r17 struct {
	Sl_RLC_ChannelID_PC5_r17           SL_RLC_ChannelID_r17            `madatory`
	Sl_RLC_ConfigPC5_r17               *SL_RLC_ConfigPC5_r16           `optional`
	Sl_MAC_LogicalChannelConfigPC5_r17 *SL_LogicalChannelConfigPC5_r16 `optional`
}

func (ie *SL_RLC_ChannelConfigPC5_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_RLC_ConfigPC5_r17 != nil, ie.Sl_MAC_LogicalChannelConfigPC5_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sl_RLC_ChannelID_PC5_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_RLC_ChannelID_PC5_r17", err)
	}
	if ie.Sl_RLC_ConfigPC5_r17 != nil {
		if err = ie.Sl_RLC_ConfigPC5_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RLC_ConfigPC5_r17", err)
		}
	}
	if ie.Sl_MAC_LogicalChannelConfigPC5_r17 != nil {
		if err = ie.Sl_MAC_LogicalChannelConfigPC5_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MAC_LogicalChannelConfigPC5_r17", err)
		}
	}
	return nil
}

func (ie *SL_RLC_ChannelConfigPC5_r17) Decode(r *uper.UperReader) error {
	var err error
	var Sl_RLC_ConfigPC5_r17Present bool
	if Sl_RLC_ConfigPC5_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MAC_LogicalChannelConfigPC5_r17Present bool
	if Sl_MAC_LogicalChannelConfigPC5_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sl_RLC_ChannelID_PC5_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_RLC_ChannelID_PC5_r17", err)
	}
	if Sl_RLC_ConfigPC5_r17Present {
		ie.Sl_RLC_ConfigPC5_r17 = new(SL_RLC_ConfigPC5_r16)
		if err = ie.Sl_RLC_ConfigPC5_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_RLC_ConfigPC5_r17", err)
		}
	}
	if Sl_MAC_LogicalChannelConfigPC5_r17Present {
		ie.Sl_MAC_LogicalChannelConfigPC5_r17 = new(SL_LogicalChannelConfigPC5_r16)
		if err = ie.Sl_MAC_LogicalChannelConfigPC5_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MAC_LogicalChannelConfigPC5_r17", err)
		}
	}
	return nil
}
