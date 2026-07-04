// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BH-RLC-ChannelConfig-r16 (line 5201).

var bHRLCChannelConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "bh-LogicalChannelIdentity-r16", Optional: true},
		{Name: "bh-RLC-ChannelID-r16"},
		{Name: "reestablishRLC-r16", Optional: true},
		{Name: "rlc-Config-r16", Optional: true},
		{Name: "mac-LogicalChannelConfig-r16", Optional: true},
	},
}

const (
	BH_RLC_ChannelConfig_r16_ReestablishRLC_r16_True = 0
)

var bHRLCChannelConfigR16ReestablishRLCR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BH_RLC_ChannelConfig_r16_ReestablishRLC_r16_True},
}

type BH_RLC_ChannelConfig_r16 struct {
	Bh_LogicalChannelIdentity_r16 *BH_LogicalChannelIdentity_r16
	Bh_RLC_ChannelID_r16          BH_RLC_ChannelID_r16
	ReestablishRLC_r16            *int64
	Rlc_Config_r16                *RLC_Config
	Mac_LogicalChannelConfig_r16  *LogicalChannelConfig
}

func (ie *BH_RLC_ChannelConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bHRLCChannelConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Bh_LogicalChannelIdentity_r16 != nil, ie.ReestablishRLC_r16 != nil, ie.Rlc_Config_r16 != nil, ie.Mac_LogicalChannelConfig_r16 != nil}); err != nil {
		return err
	}

	// 3. bh-LogicalChannelIdentity-r16: ref
	{
		if ie.Bh_LogicalChannelIdentity_r16 != nil {
			if err := ie.Bh_LogicalChannelIdentity_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. bh-RLC-ChannelID-r16: ref
	{
		if err := ie.Bh_RLC_ChannelID_r16.Encode(e); err != nil {
			return err
		}
	}

	// 5. reestablishRLC-r16: enumerated
	{
		if ie.ReestablishRLC_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ReestablishRLC_r16, bHRLCChannelConfigR16ReestablishRLCR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. rlc-Config-r16: ref
	{
		if ie.Rlc_Config_r16 != nil {
			if err := ie.Rlc_Config_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. mac-LogicalChannelConfig-r16: ref
	{
		if ie.Mac_LogicalChannelConfig_r16 != nil {
			if err := ie.Mac_LogicalChannelConfig_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BH_RLC_ChannelConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bHRLCChannelConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. bh-LogicalChannelIdentity-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Bh_LogicalChannelIdentity_r16 = new(BH_LogicalChannelIdentity_r16)
			if err := ie.Bh_LogicalChannelIdentity_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. bh-RLC-ChannelID-r16: ref
	{
		if err := ie.Bh_RLC_ChannelID_r16.Decode(d); err != nil {
			return err
		}
	}

	// 5. reestablishRLC-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(bHRLCChannelConfigR16ReestablishRLCR16Constraints)
			if err != nil {
				return err
			}
			ie.ReestablishRLC_r16 = &idx
		}
	}

	// 6. rlc-Config-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Rlc_Config_r16 = new(RLC_Config)
			if err := ie.Rlc_Config_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. mac-LogicalChannelConfig-r16: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Mac_LogicalChannelConfig_r16 = new(LogicalChannelConfig)
			if err := ie.Mac_LogicalChannelConfig_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
