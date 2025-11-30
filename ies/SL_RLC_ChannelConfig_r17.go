package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_RLC_ChannelConfig_r17 struct {
	Sl_RLC_ChannelID_r17            SL_RLC_ChannelID_r17         `madatory`
	Sl_RLC_Config_r17               *SL_RLC_Config_r16           `optional`
	Sl_MAC_LogicalChannelConfig_r17 *SL_LogicalChannelConfig_r16 `optional`
	Sl_PacketDelayBudget_r17        *int64                       `lb:0,ub:1023,optional`
}

func (ie *SL_RLC_ChannelConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_RLC_Config_r17 != nil, ie.Sl_MAC_LogicalChannelConfig_r17 != nil, ie.Sl_PacketDelayBudget_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sl_RLC_ChannelID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_RLC_ChannelID_r17", err)
	}
	if ie.Sl_RLC_Config_r17 != nil {
		if err = ie.Sl_RLC_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RLC_Config_r17", err)
		}
	}
	if ie.Sl_MAC_LogicalChannelConfig_r17 != nil {
		if err = ie.Sl_MAC_LogicalChannelConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MAC_LogicalChannelConfig_r17", err)
		}
	}
	if ie.Sl_PacketDelayBudget_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_PacketDelayBudget_r17, &aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode Sl_PacketDelayBudget_r17", err)
		}
	}
	return nil
}

func (ie *SL_RLC_ChannelConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	var Sl_RLC_Config_r17Present bool
	if Sl_RLC_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MAC_LogicalChannelConfig_r17Present bool
	if Sl_MAC_LogicalChannelConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PacketDelayBudget_r17Present bool
	if Sl_PacketDelayBudget_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sl_RLC_ChannelID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_RLC_ChannelID_r17", err)
	}
	if Sl_RLC_Config_r17Present {
		ie.Sl_RLC_Config_r17 = new(SL_RLC_Config_r16)
		if err = ie.Sl_RLC_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_RLC_Config_r17", err)
		}
	}
	if Sl_MAC_LogicalChannelConfig_r17Present {
		ie.Sl_MAC_LogicalChannelConfig_r17 = new(SL_LogicalChannelConfig_r16)
		if err = ie.Sl_MAC_LogicalChannelConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MAC_LogicalChannelConfig_r17", err)
		}
	}
	if Sl_PacketDelayBudget_r17Present {
		var tmp_int_Sl_PacketDelayBudget_r17 int64
		if tmp_int_Sl_PacketDelayBudget_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode Sl_PacketDelayBudget_r17", err)
		}
		ie.Sl_PacketDelayBudget_r17 = &tmp_int_Sl_PacketDelayBudget_r17
	}
	return nil
}
