// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RLC-ChannelConfig-r17 (line 28157).

var sLRLCChannelConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RLC-ChannelID-r17"},
		{Name: "sl-RLC-Config-r17", Optional: true},
		{Name: "sl-MAC-LogicalChannelConfig-r17", Optional: true},
		{Name: "sl-PacketDelayBudget-r17", Optional: true},
	},
}

var sLRLCChannelConfigR17SlPacketDelayBudgetR17Constraints = per.Constrained(0, 1023)

type SL_RLC_ChannelConfig_r17 struct {
	Sl_RLC_ChannelID_r17            SL_RLC_ChannelID_r17
	Sl_RLC_Config_r17               *SL_RLC_Config_r16
	Sl_MAC_LogicalChannelConfig_r17 *SL_LogicalChannelConfig_r16
	Sl_PacketDelayBudget_r17        *int64
}

func (ie *SL_RLC_ChannelConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLRLCChannelConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_RLC_Config_r17 != nil, ie.Sl_MAC_LogicalChannelConfig_r17 != nil, ie.Sl_PacketDelayBudget_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-RLC-ChannelID-r17: ref
	{
		if err := ie.Sl_RLC_ChannelID_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-RLC-Config-r17: ref
	{
		if ie.Sl_RLC_Config_r17 != nil {
			if err := ie.Sl_RLC_Config_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-MAC-LogicalChannelConfig-r17: ref
	{
		if ie.Sl_MAC_LogicalChannelConfig_r17 != nil {
			if err := ie.Sl_MAC_LogicalChannelConfig_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sl-PacketDelayBudget-r17: integer
	{
		if ie.Sl_PacketDelayBudget_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_PacketDelayBudget_r17, sLRLCChannelConfigR17SlPacketDelayBudgetR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_RLC_ChannelConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLRLCChannelConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-RLC-ChannelID-r17: ref
	{
		if err := ie.Sl_RLC_ChannelID_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-RLC-Config-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_RLC_Config_r17 = new(SL_RLC_Config_r16)
			if err := ie.Sl_RLC_Config_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-MAC-LogicalChannelConfig-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_MAC_LogicalChannelConfig_r17 = new(SL_LogicalChannelConfig_r16)
			if err := ie.Sl_MAC_LogicalChannelConfig_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sl-PacketDelayBudget-r17: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(sLRLCChannelConfigR17SlPacketDelayBudgetR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PacketDelayBudget_r17 = &v
		}
	}

	return nil
}
