// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRB-InfoBroadcast-r17 (line 28524).

var mRBInfoBroadcastR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcp-Config-r17"},
		{Name: "rlc-Config-r17"},
	},
}

type MRB_InfoBroadcast_r17 struct {
	Pdcp_Config_r17 MRB_PDCP_ConfigBroadcast_r17
	Rlc_Config_r17  MRB_RLC_ConfigBroadcast_r17
}

func (ie *MRB_InfoBroadcast_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRBInfoBroadcastR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. pdcp-Config-r17: ref
	{
		if err := ie.Pdcp_Config_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. rlc-Config-r17: ref
	{
		if err := ie.Rlc_Config_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MRB_InfoBroadcast_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRBInfoBroadcastR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. pdcp-Config-r17: ref
	{
		if err := ie.Pdcp_Config_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. rlc-Config-r17: ref
	{
		if err := ie.Rlc_Config_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
