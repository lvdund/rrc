// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-MappingConfig-U2U-r18 (line 28299).

var sLMappingConfigU2UR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RemoteUE-SLRB-Identity-r18"},
		{Name: "sl-EgressRLC-ChannelPC5-r18"},
	},
}

type SL_MappingConfig_U2U_r18 struct {
	Sl_RemoteUE_SLRB_Identity_r18 SLRB_Uu_ConfigIndex_r16
	Sl_EgressRLC_ChannelPC5_r18   SL_RLC_ChannelID_r17
}

func (ie *SL_MappingConfig_U2U_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLMappingConfigU2UR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sl-RemoteUE-SLRB-Identity-r18: ref
	{
		if err := ie.Sl_RemoteUE_SLRB_Identity_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-EgressRLC-ChannelPC5-r18: ref
	{
		if err := ie.Sl_EgressRLC_ChannelPC5_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_MappingConfig_U2U_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLMappingConfigU2UR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sl-RemoteUE-SLRB-Identity-r18: ref
	{
		if err := ie.Sl_RemoteUE_SLRB_Identity_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-EgressRLC-ChannelPC5-r18: ref
	{
		if err := ie.Sl_EgressRLC_ChannelPC5_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
