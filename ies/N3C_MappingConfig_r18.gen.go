// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: N3C-MappingConfig-r18 (line 10252).

var n3CMappingConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "n3c-RemoteUE-RB-Identity-r18"},
		{Name: "n3c-RLC-ChannelUu-r18"},
	},
}

type N3C_MappingConfig_r18 struct {
	N3c_RemoteUE_RB_Identity_r18 SL_RemoteUE_RB_Identity_r17
	N3c_RLC_ChannelUu_r18        Uu_RelayRLC_ChannelID_r17
}

func (ie *N3C_MappingConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(n3CMappingConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. n3c-RemoteUE-RB-Identity-r18: ref
	{
		if err := ie.N3c_RemoteUE_RB_Identity_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. n3c-RLC-ChannelUu-r18: ref
	{
		if err := ie.N3c_RLC_ChannelUu_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *N3C_MappingConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(n3CMappingConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. n3c-RemoteUE-RB-Identity-r18: ref
	{
		if err := ie.N3c_RemoteUE_RB_Identity_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. n3c-RLC-ChannelUu-r18: ref
	{
		if err := ie.N3c_RLC_ChannelUu_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
