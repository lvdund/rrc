// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Uu-RelayRLC-ChannelConfig-r17 (line 16353).

var uuRelayRLCChannelConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "uu-LogicalChannelIdentity-r17", Optional: true},
		{Name: "uu-RelayRLC-ChannelID-r17"},
		{Name: "reestablishRLC-r17", Optional: true},
		{Name: "rlc-Config-r17", Optional: true},
		{Name: "mac-LogicalChannelConfig-r17", Optional: true},
	},
}

const (
	Uu_RelayRLC_ChannelConfig_r17_ReestablishRLC_r17_True = 0
)

var uuRelayRLCChannelConfigR17ReestablishRLCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Uu_RelayRLC_ChannelConfig_r17_ReestablishRLC_r17_True},
}

type Uu_RelayRLC_ChannelConfig_r17 struct {
	Uu_LogicalChannelIdentity_r17 *LogicalChannelIdentity
	Uu_RelayRLC_ChannelID_r17     Uu_RelayRLC_ChannelID_r17
	ReestablishRLC_r17            *int64
	Rlc_Config_r17                *RLC_Config
	Mac_LogicalChannelConfig_r17  *LogicalChannelConfig
}

func (ie *Uu_RelayRLC_ChannelConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uuRelayRLCChannelConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Uu_LogicalChannelIdentity_r17 != nil, ie.ReestablishRLC_r17 != nil, ie.Rlc_Config_r17 != nil, ie.Mac_LogicalChannelConfig_r17 != nil}); err != nil {
		return err
	}

	// 3. uu-LogicalChannelIdentity-r17: ref
	{
		if ie.Uu_LogicalChannelIdentity_r17 != nil {
			if err := ie.Uu_LogicalChannelIdentity_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. uu-RelayRLC-ChannelID-r17: ref
	{
		if err := ie.Uu_RelayRLC_ChannelID_r17.Encode(e); err != nil {
			return err
		}
	}

	// 5. reestablishRLC-r17: enumerated
	{
		if ie.ReestablishRLC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ReestablishRLC_r17, uuRelayRLCChannelConfigR17ReestablishRLCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. rlc-Config-r17: ref
	{
		if ie.Rlc_Config_r17 != nil {
			if err := ie.Rlc_Config_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. mac-LogicalChannelConfig-r17: ref
	{
		if ie.Mac_LogicalChannelConfig_r17 != nil {
			if err := ie.Mac_LogicalChannelConfig_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Uu_RelayRLC_ChannelConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uuRelayRLCChannelConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. uu-LogicalChannelIdentity-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Uu_LogicalChannelIdentity_r17 = new(LogicalChannelIdentity)
			if err := ie.Uu_LogicalChannelIdentity_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. uu-RelayRLC-ChannelID-r17: ref
	{
		if err := ie.Uu_RelayRLC_ChannelID_r17.Decode(d); err != nil {
			return err
		}
	}

	// 5. reestablishRLC-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(uuRelayRLCChannelConfigR17ReestablishRLCR17Constraints)
			if err != nil {
				return err
			}
			ie.ReestablishRLC_r17 = &idx
		}
	}

	// 6. rlc-Config-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Rlc_Config_r17 = new(RLC_Config)
			if err := ie.Rlc_Config_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. mac-LogicalChannelConfig-r17: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Mac_LogicalChannelConfig_r17 = new(LogicalChannelConfig)
			if err := ie.Mac_LogicalChannelConfig_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
